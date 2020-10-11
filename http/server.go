package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bagiduid/backend/http/graphql/generated"
	"github.com/bagiduid/backend/http/graphql/resolver"
	"github.com/bagiduid/backend/services/user"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Init Mongo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		log.Fatal("Unable to connect mongo db")
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Unable to ping mongo db")
	}

	db := client.Database("bagidu")

	// Services
	userService := user.NewUserService(db)

	// Graphql
	res := &resolver.Resolver{
		UserService: userService,
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: res}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
