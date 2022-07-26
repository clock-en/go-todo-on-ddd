package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/clock-en/go-todo-on-ddd-on-ddd/adapter/controller"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/dao"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/infrastructure/graph/model"
)

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUser) (*model.User, error) {
	userDao := &dao.UserDao{}
	userController := controller.NewUserController(userDao)
	viewModel, err := userController.RegisterUser(input.Name, input.Email, input.Password)
	if err != nil {
		return nil, err
	}
	dataModel := &model.User{
		ID:    viewModel.ID(),
		Name:  viewModel.Name(),
		Email: viewModel.Email(),
	}
	return dataModel, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
