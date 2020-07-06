package graph

import "github.com/saikrishna-commits/go-mvc/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	movies   []*model.Movie
	authors  []*model.Author
	theaters []*model.Theater
	posts    []*model.Post
}
