package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/saikrishna-commits/go-mvc/db"
	"github.com/saikrishna-commits/go-mvc/graph/generated"
	"github.com/saikrishna-commits/go-mvc/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.NewAuthor) (*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
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

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	cur, err := db.SqlCon.Query("SELECT id,first_name as firstName ,last_name as LastName, email FROM authors LIMIT 2")

	defer cur.Close()

	if err != nil {
		panic(err.Error())
	}

	for cur.Next() {
		var id int64
		var firstName, lastName, email string
		err = cur.Scan(&id, &firstName, &lastName, &email)
		if err != nil {
			log.Println(err.Error()) // proper error handling instead of panic in your app
		}
		r.authors = append(r.authors, &model.Author{AuthorID: int(id), FirstName: firstName, LastName: lastName, Email: email})
	}
	return r.authors, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Post(ctx context.Context, id int) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Author(ctx context.Context, id int) (*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
