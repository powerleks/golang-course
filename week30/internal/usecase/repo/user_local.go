package repo

import (
	"context"
	"fmt"

	"restapi/internal/entity"
)

type UserRepo struct {
	Store        map[int]*entity.User
	FriendsStore map[int]map[int]*entity.User
	Counter      int
}

func New() *UserRepo {
	return &UserRepo{
		Store: make(map[int]*entity.User),
		FriendsStore: make(map[int]map[int]*entity.User),
		Counter: 0,
	}
}

func (r *UserRepo) GetAllUsers(ctx context.Context) (*[]entity.User, error) {
	users := make([]entity.User, len(r.Store))
	i := 0
	for _, user := range r.Store {
		users[i] = *user
		i++
	}
	return &users, nil
}

func (r *UserRepo) GetUserById(ctx context.Context, userId int) (*entity.User, error) {
	user, ok := r.Store[userId]
	if !ok {
		return nil, fmt.Errorf("пользователь с ID=%d не найден", userId)
	}
	return user, nil
}

func (r *UserRepo) CreateUser(ctx context.Context, user *entity.User) (int, error) {
	user.Id = r.Counter
	r.Counter++
	r.Store[user.Id] = user
	r.FriendsStore[user.Id] = map[int]*entity.User{}
	return user.Id, nil
}

func (r *UserRepo) MakeFriends(ctx context.Context, userId1, UserId2 int) error {
	user1, err := r.GetUserById(ctx, userId1)
	if err != nil {
		return err
	}
	user2, err := r.GetUserById(ctx, UserId2)
	if err != nil {
		return err
	}
	r.FriendsStore[userId1][UserId2] = user2
	r.FriendsStore[UserId2][userId1] = user1
	return nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, userId int, user *entity.User) error {
	r.Store[userId] = user
	return nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, userId int) error {
	_, ok := r.Store[userId]
	if !ok {
		return fmt.Errorf("пользователь с ID=%d не найден", userId)
	}
	delete(r.Store, userId)
	return nil
}

func (r *UserRepo) DeleteFriend(ctx context.Context, userId1, userId2 int) error {
	userFriends, err := r.GetUserFriends(ctx, userId1)
	if err != nil {
		return err
	}
	delete(userFriends, userId1)

	userFriends, err = r.GetUserFriends(ctx, userId2)
	if err != nil {
		return err
	}
	delete(userFriends, userId2)
	return  nil
}

func (r *UserRepo) GetUserFriends(ctx context.Context, userId int) (map[int]*entity.User, error) {
	userFriends, ok := r.FriendsStore[userId]
	if !ok {
		return nil, fmt.Errorf("пользователь с ID=%d не найден", userId)
	}
	return userFriends, nil
}