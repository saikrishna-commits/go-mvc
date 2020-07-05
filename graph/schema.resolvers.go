package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/saikrishna-commits/go-mvc/db"
	"github.com/saikrishna-commits/go-mvc/graph/generated"
	"github.com/saikrishna-commits/go-mvc/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *movieResolver) ID(ctx context.Context, obj *model.Movie) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	randN, _ := rand.Int(rand.Reader, big.NewInt(1000))
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", randN),
		UserID: input.UserID, // fix this line
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	curCtx := context.Background()
	// choose Collection
	collection := db.MongoClient.Database("sample_mflix").Collection("movies")

	//find options
	findOptions := options.Find()
	findOptions.SetLimit(2)
	findOptions.SetSort(bson.D{{"title", -1}})

	// Writing query to fetch the Data from the `movies` collection
	cur, err := collection.Find(curCtx, bson.D{
		{"year", bson.D{
			{"$gt", 2010},
		}}}, findOptions)

	defer cur.Close(curCtx)

	if err = cur.All(ctx, &r.movies); err != nil {
		log.Fatal(err)
	}

	return r.movies, nil
}

func (r *queryResolver) Theaters(ctx context.Context) ([]*model.Theater, error) {
	curCtx := context.Background()
	// choose Collection
	collection := db.MongoClient.Database("sample_mflix").Collection("theaters")

	//find options
	findOptions := options.Find()
	findOptions.SetLimit(2)
	findOptions.SetSort(bson.D{{"theaterId", -1}})

	// Writing query to fetch the Data from the `movies` collection
	cur, err := collection.Find(curCtx, bson.D{
		{"theaterId", bson.D{
			{"$gt", 100},
		}}}, findOptions)

	defer cur.Close(curCtx)

	if err = cur.All(ctx, &r.theaters); err != nil {
		log.Fatal(err)
	}

	return r.theaters, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

func (r *todoResolver) UserLoader(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) UserRaw(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Movie returns generated.MovieResolver implementation.
func (r *Resolver) Movie() generated.MovieResolver { return &movieResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type movieResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *theaterResolver) TheaterID(ctx context.Context, obj *model.Theater) (*int, error) {
	panic("Not implemented")
}

type theaterResolver struct{ *Resolver }
