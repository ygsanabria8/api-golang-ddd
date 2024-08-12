package services

import "api.ddd/src/domain/aggregates"

func (s *UserService) GetUserById(userId string) (*aggregates.User, error) {
	return s.NoSqlFinder.GetUserById(userId)
}

func (s *UserService) GetAllUsers() ([]*aggregates.User, error) {
	return s.NoSqlFinder.GetAllUsers()
}

func (s *UserService) CreateUser(user *aggregates.User) (*aggregates.User, error) {
	userSaved, err := s.SqlRepository.CreateUser(user)
	return userSaved, err
}

func (s *UserService) UpdateUser(user *aggregates.User) (*aggregates.User, error) {
	newUser, err := s.SqlRepository.UpdateUser(user)
	return newUser, err
}

func (s *UserService) DeleteUser(userId string) error {
	err := s.SqlRepository.DeleteUser(userId)
	return err
}
