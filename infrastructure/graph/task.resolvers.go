package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/graph/generated"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/graph/model"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTask) (*model.Task, error) {
	id, _ := rand.Int(rand.Reader, big.NewInt(100))
	task := &model.Task{
		ID:      fmt.Sprintf("tasks%d", id),
		Title:   input.Title,
		Content: input.Content,
	}
	r.tasks = append(r.tasks, task)
	return task, nil
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id string) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	return r.tasks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
