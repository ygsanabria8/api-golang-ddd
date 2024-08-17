package services

import (
	"api.ddd/src/domain/aggregates"
	"api.ddd/src/domain/events"
)

func (s *UserService) GetUserById(userId string) (*aggregates.User, error) {
	return s.NoSqlFinder.GetUserById(userId)
}

func (s *UserService) GetAllUsers() ([]*aggregates.User, error) {
	return s.NoSqlFinder.GetAllUsers()
}

func (s *UserService) CreateUser(user *aggregates.User) (*aggregates.User, error) {
	userSaved, err := s.SqlRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	event := events.NewCreatedUserEvent(userSaved)
	s.Kafka.SendMessage(event)
	go func() {
		_, _ = s.NoSqlRepository.CreateUser(userSaved)
	}()
	return userSaved, nil
}

func (s *UserService) UpdateUser(user *aggregates.User) (*aggregates.User, error) {
	newUser, err := s.SqlRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	event := events.NewUpdatedUserEvent(newUser)
	s.Kafka.SendMessage(event)
	go func() {
		_, _ = s.NoSqlRepository.UpdateUser(newUser)
	}()
	return newUser, nil

}

func (s *UserService) DeleteUser(userId string) error {
	err := s.SqlRepository.DeleteUser(userId)
	if err != nil {
		return err
	}

	event := events.NewDeletedUserEvent(userId)
	s.Kafka.SendMessage(event)
	go func() {
		_ = s.NoSqlRepository.DeleteUser(userId)
	}()
	return nil
}
