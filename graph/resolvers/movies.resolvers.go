package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/saikrishna-commits/go-mvc/graph/generated"
	"github.com/saikrishna-commits/go-mvc/graph/model"
)

func (r *movieResolver) ID(ctx context.Context, obj *model.Movie) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Movie returns generated.MovieResolver implementation.
func (r *Resolver) Movie() generated.MovieResolver { return &movieResolver{r} }

type movieResolver struct{ *Resolver }
