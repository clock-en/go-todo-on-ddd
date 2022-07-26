package repository

import (
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/entity"
	"github.com/clock-en/go-todo-on-ddd-on-ddd/domain/vo"
)

type userRepository struct {
	dao IUserDao
}

func NewUserRepository(dao IUserDao) userRepository {
	return userRepository{dao: dao}
}

func (r userRepository) CreateUser(user entity.User) (*entity.AuthUser, error) {
	err := r.dao.CreateUser(
		user.ID().Value(),
		user.Name().Value(),
		user.Email().Value(),
		user.Password().Hash(),
	)
	if err != nil {
		return nil, err
	}
	authUser := entity.NewAuthUser(user.ID(), user.Name(), user.Email())
	return authUser, nil
}
func (r userRepository) FindRegisteredUser(name vo.UserName, email vo.Email) (*entity.AuthUser, error) {
	row, err := r.dao.FindRegisteredUser(name.Value(), email.Value())
	if err != nil {
		return nil, err
	}
	authUser := createAuthUserEntity(row)
	return &authUser, nil
}

type IUserRow interface {
	ID() string
	Name() string
	Email() string
	Password() string
}

type IUserDao interface {
	CreateUser(id string, name string, email string, hash string) error
	FindRegisteredUser(name string, email string) (IUserRow, error)
}

func createAuthUserEntity(userRow IUserRow) entity.AuthUser {
	id, _ := vo.NewID(userRow.ID())
	name, _ := vo.NewUserName(userRow.Name())
	email, _ := vo.NewEmail(userRow.Email())
	// TODO: voのエラー処理を実装
	return *entity.NewAuthUser(*id, *name, *email)
}
