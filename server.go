package main

import (
	"log"
	"net/http"
	"os"
	"studynotes/graph"
	"studynotes/config"


	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	// ⏬ Connect to Database
	config.ConnectDatabase()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// ⏬ Initialize GraphQL server
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New )

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New ,
	})

	// ⏬ Route for playground and GraphQL endpoint
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("🚀 Server ready at http://localhost:%s/ 🚀", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
