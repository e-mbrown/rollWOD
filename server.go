package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/e-mbrown/rollWOD/graph"
	"github.com/e-mbrown/rollWOD/pkg/services"
)

const defaultPort = "8080"
var resolver graph.Config

func init() {
	resolver = graph.NewResolver()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	cache, err := services.NewCache(ctx, "localhost:6379", 12*time.Hour,)
	if err != nil {
		log.Fatalf("cannot create APQ redis cache: %v",err) 
	}

	gqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(resolver))

	gqlHandler.AddTransport(transport.POST{})
	gqlHandler.Use(extension.AutomaticPersistedQuery{Cache: cache})
	//TODO: Remove following line
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", gqlHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}


