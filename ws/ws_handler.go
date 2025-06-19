package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Struktur pesan yang diterima/dikirim
type Message struct {
	User    string `json:"user"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

// Mengelola client yang terhubung
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var clientsMu sync.Mutex // untuk concurrency-safe access ke map

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // ⚠️ Ubah di production untuk keamanan
	},
}

// Fungsi utama untuk menerima koneksi WebSocket dari client
func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("❌ WebSocket upgrade error:", err)
		return
	}
	defer ws.Close()

	// Register client baru
	clientsMu.Lock()
	clients[ws] = true
	fmt.Printf("🟢 New client connected. Total: %d\n", len(clients))
	clientsMu.Unlock()

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println("🔴 Client disconnected:", err)

			// Hapus client dari list
			clientsMu.Lock()
			delete(clients, ws)
			fmt.Printf("⚠️ Client removed. Total now: %d\n", len(clients))
			clientsMu.Unlock()
			break
		}

		// Kirim pesan ke channel broadcast
		broadcast <- msg
	}
}

// Fungsi untuk menyebarkan pesan ke semua client aktif
func handleMessages() {
	for {
		msg := <-broadcast

		// Konversi ke JSON hanya sekali
		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("❌ Marshal error:", err)
			continue
		}

		// Kirim ke semua client
		clientsMu.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, jsonMsg)
			if err != nil {
				fmt.Println("🔴 Error sending message to client:", err)
				client.Close()
				delete(clients, client)
			}
		}
		clientsMu.Unlock()
	}
}
