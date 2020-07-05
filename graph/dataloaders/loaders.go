package dataloader

import (
	"context"
	"database/sql"

	"net/http"
	"strings"
	"time"

	"github.com/saikrishna-commits/go-mvc/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const loadersKey = "dataloaders"

type Loaders struct {
	UserById UserLoader
}
// Middleware Does wrapping of db connections and dl config
func Middleware(sql *sql.DB,mongo *mongo.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			UserById: UserLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []int) ([]*model.User, []error) {
					placeholders := make([]string, len(ids))
					args := make([]interface{}, len(ids))
					for i := 0; i < len(ids); i++ {
						placeholders[i] = "?"
						args[i] = i
					}

					res := sql.LogAndQuery(conn,
						"SELECT id, name from dataloader_example.user WHERE id IN ("+strings.Join(placeholders, ",")+")",
						args...,
					)
					defer res.Close()

					userById := map[int]*model.User{}
					for res.Next() {
						user := model.User{}
						err := res.Scan(&user.ID, &user.Name)
						if err != nil {
							panic(err)
						}
						userById[user.ID] = &user
					}

					users := make([]*model.User, len(ids))
					for i, id := range ids {
						users[i] = userById[id]
						i++
					}

					return users, nil
				},
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}