package usecase

import (
	"context"
	"fmt"

	"restapi/internal/entity"
)

type UserUseCase struct {
	repo   UserRepo
}

func New(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo:   r,
	}
}

func (uc *UserUseCase) GetAllUsers(ctx context.Context) (*[]entity.User, error) {
	users, err := uc.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetAllUsers - s.repo.GetAllUsers: %w", err)
	}

	return users, nil
}

func (uc *UserUseCase) GetUserById(ctx context.Context, userId int) (*entity.User, error) {
	user, err := uc.repo.GetUserById(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetUserById - s.repo.GetUserById: %w", err)
	}

	return user, nil
}

func (uc *UserUseCase) AddUser(ctx context.Context, user *entity.User, userFriends *entity.UserFriends) (int, error) {
	userId, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return -1, fmt.Errorf("UserUseCase - AddUser - s.repo.CreateUser: %w", err)
	}

	for _, friendId := range userFriends.Friends {
		err = uc.repo.MakeFriends(ctx, userId, friendId)
		if err != nil {
			return -1, fmt.Errorf("UserUseCase - AddUser - s.repo.MakeFriends: %w", err)
		}
	}

	return userId, nil
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, userId int, user *entity.User) error {
	err := uc.repo.UpdateUser(ctx, userId, user)
	if err != nil {
		return fmt.Errorf("UserUseCase - UpdateUser - s.repo.UpdateUser: %w", err)
	}

	return nil
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, userId int) error {
	friends, err := uc.repo.GetUserFriends(ctx, userId)
	if err != nil {
		return fmt.Errorf("UserUseCase - DeleteUser - s.repo.GetUserFriends: %w", err)
	}

	for _, u := range friends {
		err = uc.repo.DeleteFriend(ctx, u.Id, userId)
		if err != nil {
			return fmt.Errorf("UserUseCase - DeleteUser - s.repo.DeleteFriend: %w", err)
		}
	}

	err = uc.repo.DeleteUser(ctx, userId)
	if err != nil {
		return fmt.Errorf("UserUseCase - DeleteUser - s.repo.DeleteUser: %w", err)
	}

	return nil
}

func (uc *UserUseCase) GetUserFriends(ctx context.Context, userId int) (map[int]*entity.User, error) {
	user, err := uc.repo.GetUserFriends(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetUserFriends - s.repo.GetUserFriends: %w", err)
	}

	return user, nil
}

func (uc *UserUseCase) MakeFriends(ctx context.Context, userId1, UserId2 int) error {
	err := uc.repo.MakeFriends(ctx, userId1, UserId2)
	if err != nil {
		return fmt.Errorf("UserUseCase - MakeFriends - s.repo.MakeFriends: %w", err)
	}

	return nil
}