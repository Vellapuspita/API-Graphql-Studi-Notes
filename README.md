# 📘 API GraphQL Study Notes

API GraphQL Study Notes adalah backend API yang dikembangkan menggunakan bahasa Go (Golang) dan framework gqlgen. Aplikasi ini bertujuan untuk mengelola catatan belajar (study notes) dengan fitur autentikasi JWT, manajemen topik, dan pencatatan oleh pengguna.

---

### 🚀 Fitur Utama
- Autentikasi menggunakan JWT (JSON Web Token)
- Registrasi dan login user
- CRUD data topik
- CRUD catatan belajar (study notes)
- Middleware untuk proteksi endpoint privat
- Modularisasi folder seperti config, models, graph, utils, dan middleware

### ⚙️ Instalasi
Pastikan kamu sudah menginstal Go dan MySQL.

1. Clone repository ini:
```bash
git clone https://github.com/Vellapuspita/API-Graphql-Studi-Notes.git
cd API-Graphql-Studi-Notes
```

2. Ubah konfigurasi database di config/config.go sesuai dengan kredensial lokalmu.
3. Install dependency:
```bash
go mod tidy
```
4. Generate file GraphQL jika belum ada:

```bash
go run github.com/99designs/gqlgen generate
```

### 🧪 Menjalankan Server
```bash
go run server.go
```

Secara default, server akan berjalan di http://localhost:8080/query.

### 📌 Catatan Tambahan
1. JWT Token harus dikirim di header:
Authorization: Bearer <your-token>
2. GQL playground juga bisa digunakan secara lokal untuk testing

---
### Nama Kelompok

1. Vella Puspitasari Wijayanti      (42230043)
2. Ni Made Ochinana Sephti Pratiwi  (42230027)
3. Ida Ayu Dwi Wirayanti            (42230021)
4. Muhammad Zafran ZUlkifli         (42230047)
