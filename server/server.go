package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/kotaMiya/meetup"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(meetup.NewExecutableSchema(meetup.Config{Resolvers: &meetup.Resolver{}})))

	log.Printf("connect to http://localhost:#{port}/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":" + port, nil))
}