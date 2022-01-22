package usecase

import (
	"context"

	"restapi/internal/entity"
)

type (
	User interface {
		GetAllUsers(context.Context) (*[]entity.User, error)
		GetUserById(context.Context, int) (*entity.User, error)
		AddUser(context.Context, *entity.User, *entity.UserFriends) (int, error)
		UpdateUser(context.Context, int, *entity.User) error
		DeleteUser(context.Context, int) error
		GetUserFriends(context.Context, int) (map[int]*entity.User, error)
		MakeFriends(context.Context, int, int) error
	}

	UserRepo interface {
		GetAllUsers(context.Context) (*[]entity.User, error)
		GetUserById(context.Context, int) (*entity.User, error)
		CreateUser(context.Context, *entity.User) (int, error)
		MakeFriends(context.Context, int, int) error
		UpdateUser(context.Context, int, *entity.User) error
		DeleteUser(context.Context, int) error
		DeleteFriend(context.Context, int, int) error
		GetUserFriends(context.Context, int) (map[int]*entity.User, error)
	}
)
