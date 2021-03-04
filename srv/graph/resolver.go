package graph

import (
	"database/sql"
	"github.com/matthewmazzanti/wordgame/srv/game"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	DB *sql.DB
	Games map[string]*game.Game
}
