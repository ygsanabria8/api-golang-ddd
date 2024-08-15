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
		_, err := s.NoSqlRepository.CreateUser(userSaved)
		if err != nil {
			s.Logger.Errorw("Create", err)
		}
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
		_, err := s.NoSqlRepository.UpdateUser(newUser)
		if err != nil {
			s.Logger.Errorw("Update", err)
		}
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
		err := s.NoSqlRepository.DeleteUser(userId)
		if err != nil {
			s.Logger.Errorw("Delete", err)
		}
	}()
	return nil
}
