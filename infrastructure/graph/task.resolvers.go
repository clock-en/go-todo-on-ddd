package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/clock-en/go-todo-on-ddd-on-ddd/adapter/controller"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/graph/generated"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/graph/model"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTask) (*model.Task, error) {
	taskController := controller.TaskController{}
	viewModel, err := taskController.CreateTask(input.Title, input.Content, input.UserID)
	if err != nil {
		return nil, err
	}
	dataModel := &model.Task{
		ID:      viewModel.ID(),
		Title:   viewModel.Title(),
		Content: viewModel.Content(),
		UserID:  viewModel.UserID(),
	}
	return dataModel, nil
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id string) (*model.Task, error) {
	taskController := controller.TaskController{}
	viewModel, err := taskController.FindTask(id)
	if err != nil {
		return nil, err
	}

	dataModel := &model.Task{
		ID:      viewModel.ID(),
		Title:   viewModel.Title(),
		Content: viewModel.Content(),
		UserID:  viewModel.UserID(),
	}
	return dataModel, nil
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
