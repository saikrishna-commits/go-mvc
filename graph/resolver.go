package graph
//go:generate go run github.com/99designs/gqlgen
import (
	"database/sql"

	"github.com/saikrishna-commits/go-mvc/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//Resolver = Root resolver
type Resolver struct{
	mongo *mongo.Client
	sql *sql.DB
	movies []*model.Movie
	theaters []*model.Theater
	todos []*model.Todo
}
