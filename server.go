package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	db "github.com/saikrishna-commits/go-mvc/db"
	"github.com/saikrishna-commits/go-mvc/graph"
	dataloader "github.com/saikrishna-commits/go-mvc/graph/dataloaders"
	"github.com/saikrishna-commits/go-mvc/graph/generated"
)

const defaultPort = "8080"

func init() {
	godotenv.Load(".env") // load env variables from specific file , alterantive we can use viper package
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db.ConnectDatabaseMongo() //connect to mongo
	db.CreatePgConnection() //connect to pg
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{mongo: db.MongoClient,sql: db.PgCon}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dataloader.Middleware(db.PgCon,db.MongoClient,srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
