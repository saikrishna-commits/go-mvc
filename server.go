package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	db "github.com/saikrishna-commits/go-mvc/db"
	graph "github.com/saikrishna-commits/go-mvc/graph/Resolvers"
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
	db.CreateSqlConnection()  //connect to mysql
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
