package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/dao"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/graph/generated"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/graph/model"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/usecase"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTask) (*model.Task, error) {
	inputData := usecase.NewCreateTaskInput(input.Title, input.Content, input.UserID)
	taskDao := dao.NewTaskDao()
	interactor := usecase.NewCreateTaskUsecase(inputData, *taskDao)
	output, err := interactor.Handler()
	if err != nil {
		return nil, err
	}
	taskDataModel := &model.Task{
		ID:      output.Task().ID().Value(),
		Title:   output.Task().Title().Value(),
		Content: output.Task().Content().Value(),
		UserID:  output.Task().UserID().Value(),
	}

	return taskDataModel, nil
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
