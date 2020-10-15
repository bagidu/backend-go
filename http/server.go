package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bagiduid/backend/http/graphql/generated"
	"github.com/bagiduid/backend/http/graphql/resolver"
	"github.com/bagiduid/backend/services/mail"
	"github.com/bagiduid/backend/services/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const defaultPort = "8080"

func main() {
	startTime := time.Now()
	// Load env
	godotenv.Load()
	debug := os.Getenv("DEBUG")
	key := os.Getenv("APP_SECRET")

	// JWT
	jwt := jwtauth.New("HS256", []byte(key), nil)

	// Port
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// MONGODB
	mongoURL := os.Getenv("MONGODB_URL")
	if mongoURL == "" {
		panic("MONGODB_URL not set properly")
	}

	// Init Mongo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal("Unable to connect mongo db")
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Unable to ping mongo db")
	}

	db := client.Database("bagidu")

	// Services
	userService := user.NewUserService(db)
	mailService := mail.NewMailgunService()

	// Graphql
	res := &resolver.Resolver{
		UserService: userService,
		MailService: mailService,
		JWT:         jwt,
	}

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: res}))
	srv.AddTransport(transport.Options{})
	// srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	// Http Server
	router := chi.NewRouter()

	// JWT Middleware
	router.Use(jwtauth.Verifier(jwt))

	if debug == "TRUE" {
		srv.Use(extension.Introspection{})
		router.Handle("/playground", playground.Handler("GraphQL playground", "/"))
	}

	router.Handle("/", srv)

	elapsed := time.Since(startTime)
	log.Printf("Startup took %s", elapsed)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
