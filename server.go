package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/e-mbrown/rollWOD/pkg/graph"
	"github.com/e-mbrown/rollWOD/pkg/repl"
)

const defaultPort = "8080"

var resolver graph.Config

func init() {
	resolver = graph.NewResolver()
}

func main() {

	//REPL//

	if len(os.Args[1:]) > 0 {
		fmt.Println("This is the WQL REPL")
		fmt.Println("Type some commands")
		repl.Start(os.Stdin, os.Stdout)
	}
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// cache, err := services.NewCache(ctx, "localhost:6379", 12*time.Hour,)
	// if err != nil {
	// 	log.Fatalf("cannot create APQ redis cache: %v",err)
	// }

	gqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(resolver))

	gqlHandler.AddTransport(transport.POST{})
	// gqlHandler.Use(extension.AutomaticPersistedQuery{Cache: cache})
	http.Handle("/", gqlHandler)
	//TODO: Remove following line
	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))

	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
