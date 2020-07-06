package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/saikrishna-commits/go-mvc/graph/generated"
	"github.com/saikrishna-commits/go-mvc/graph/model"
)

func (r *authorResolver) ID(ctx context.Context, obj *model.Author) (*int, error) {
	return &obj.AuthorID, nil
}

func (r *authorResolver) BirthDate(ctx context.Context, obj *model.Author) (*string, error) {
	var s = "Not Yet Implemeneted"
	return &s, nil
}

func (r *authorResolver) Posts(ctx context.Context, obj *model.Author) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) PostID(ctx context.Context, obj *model.Post) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) AuthorID(ctx context.Context, obj *model.Post) (*int, error) {
	panic(fmt.Errorf("not implemented AuthorID"))
}

// Author returns generated.AuthorResolver implementation.
func (r *Resolver) Author() generated.AuthorResolver { return &authorResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type authorResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
