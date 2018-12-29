package services

import (
	"github.com/emwalker/digraph/models"
	"github.com/volatiletech/sqlboiler/boil"
)

// Connection holds fields used by service calls.
type Connection struct {
	Exec  boil.ContextExecutor
	Actor *models.User
}

// CleanupFunc is a function that can be called to roll back the effects of a service call.
type CleanupFunc func() error

// New returns a new service connection
func New(exec boil.ContextExecutor, actor *models.User) *Connection {
	return &Connection{exec, actor}
}
