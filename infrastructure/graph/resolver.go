package graph

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	tasks []*model.Task
}
