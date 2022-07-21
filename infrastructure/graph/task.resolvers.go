package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/dao"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/graph/generated"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/graph/model"
	"github.com/google/uuid"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTask) (*model.Task, error) {
	generatedID, _ := uuid.NewRandom()
	id, err := vo.NewID(generatedID.String())
	title, err := vo.NewTaskTitle(input.Title)
	content, err := vo.NewTaskContent(input.Content)
	userID, err := vo.NewID(input.UserID)

	taskEntity := entity.NewTask(*id, *title, *content, *userID)

	taskDao := dao.NewTaskDao()
	defer taskDao.Close()
	task, err := taskDao.CreateTask(*taskEntity)
	if err != nil {
		return nil, err
	}
	taskDataModel := &model.Task{
		ID:      task.ID().Value(),
		Title:   task.Title().Value(),
		Content: task.Content().Value(),
		UserID:  task.UserID().Value(),
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
