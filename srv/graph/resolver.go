package graph

import (
	_ "fmt"
	"net/http"
	"context"
	"database/sql"

	"github.com/matthewmazzanti/wordgame/srv/game"
	"github.com/matthewmazzanti/wordgame/srv/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	DB *sql.DB
	Games map[string]*game.Game
	Users map[string]*model.User
}

func NewResolver(db *sql.DB) *Resolver {
	return &Resolver{
		DB: db,
		Games: make(map[string]*game.Game),
		Users: make(map[string]*model.User),
	}
}

func (resolver *Resolver) SetCookies(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("user-id")

		var user *model.User = nil;

		// Try to get the user
		if err == nil && c != nil {
			user = resolver.Users[c.Value]
		}

		// User may be nil
		ctx := context.WithValue(r.Context(), "user", user)
		ctx = context.WithValue(ctx, "writer", w)

		// and call the next with our new context
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
