package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/clock-en/go-todo-on-ddd-on-ddd/adapter/controller"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/dao"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/graph/custom_errors"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/graph/generated"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/graph/model"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTask) (*model.Task, error) {
	taskDao := &dao.TaskDao{}
	taskController := controller.NewTaskController(taskDao)
	outputData, err := taskController.CreateTask(input.Title, input.Content, input.UserID)
	if err != nil {
		return nil, err
	}

	if inputErrors := outputData.InputErrors(); inputErrors != nil {
		return nil, custom_errors.GenerateGraphqlInputError(ctx, inputErrors)
	}

	dataModel := &model.Task{
		ID:      outputData.Task().ID(),
		Title:   outputData.Task().Title(),
		Content: outputData.Task().Content(),
		UserID:  outputData.Task().UserID(),
	}
	return dataModel, nil
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id string) (*model.Task, error) {
	taskDao := &dao.TaskDao{}
	taskController := controller.NewTaskController(taskDao)
	outputData, err := taskController.FindTask(id)
	if err != nil {
		return nil, err
	}

	if inputErrors := outputData.InputErrors(); inputErrors != nil {
		return nil, custom_errors.GenerateGraphqlInputError(ctx, inputErrors)
	}

	dataModel := &model.Task{
		ID:      outputData.Task().ID(),
		Title:   outputData.Task().Title(),
		Content: outputData.Task().Content(),
		UserID:  outputData.Task().UserID(),
	}
	return dataModel, nil
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, userID string) ([]*model.Task, error) {
	dataModel := []*model.Task{}

	taskDao := &dao.TaskDao{}
	taskController := controller.NewTaskController(taskDao)
	outputData, err := taskController.FetchUserTasks(userID)
	if err != nil {
		return nil, err
	}

	if inputErrors := outputData.InputErrors(); inputErrors != nil {
		return nil, custom_errors.GenerateGraphqlInputError(ctx, inputErrors)
	}

	for _, task := range outputData.Tasks() {
		dataModel = append(dataModel, &model.Task{
			ID:      task.ID(),
			Title:   task.Title(),
			Content: task.Content(),
			UserID:  task.UserID(),
		})
	}

	return dataModel, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
