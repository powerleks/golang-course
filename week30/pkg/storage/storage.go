package storage

import (
	"fmt"
	"restapi/pkg/user"
)

type Service struct {
	Store        map[int]*user.User
	FriendsStore map[int]map[int]*user.User
	Counter      int
}

func (s *Service) GetUser(userID int) (*user.User, error) {
	user, ok := s.Store[userID]
	if !ok {
		return nil, fmt.Errorf("пользователь с ID=%d не найден", userID)
	}
	return user, nil
}

func (s *Service) AddUser(newUser *user.User) {
	newUser.Id = s.Counter
	s.Counter++
	s.Store[newUser.Id] = newUser
	s.FriendsStore[newUser.Id] = map[int]*user.User{}
}

func (s *Service) DeleteUser(userID int) error {
	user, ok := s.Store[userID]
	if !ok {
		return fmt.Errorf("пользователь с ID=%d не найден", userID)
	}
	delete(s.Store, user.Id)
	return nil
}

func (s *Service) GetFriends(userID int) (map[int]*user.User, error) {
	userFriends, ok := s.FriendsStore[userID]
	if !ok {
		return nil, fmt.Errorf("пользователь с ID=%d не найден", userID)
	}
	return userFriends, nil
}

func (s *Service) MakeFriends(userID1, UserID2 int) error {
	user1, err := s.GetUser(userID1)
	if err != nil {
		return err
	}
	user2, err := s.GetUser(UserID2)
	if err != nil {
		return err
	}
	s.FriendsStore[userID1][UserID2] = user2
	s.FriendsStore[UserID2][userID1] = user1
	return nil
}

func (s *Service) DeleteFriend(userID1, UserID2 int) error {
	userFriends, ok := s.FriendsStore[userID1]
	if !ok {
		return fmt.Errorf("пользователь с ID=%d не найден", userID1)
	}
	delete(userFriends, UserID2)
	return  nil
}
