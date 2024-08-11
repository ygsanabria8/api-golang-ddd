package finder

import (
	"api.ddd/src/domain/aggregates"
	"errors"
)

func (f *MySqlUserFinder) GetUserById(userId string) (*aggregates.User, error) {
	f.logger.Infof("Getting User: " + userId)
	user := &aggregates.User{
		Id: userId,
	}
	result := f.database.First(&user)
	if result.Error != nil {
		f.logger.Errorw("User Not Found", result.Error)
		return nil, errors.New("some error happens getting user")
	}

	return user, nil
}

func (f *MySqlUserFinder) GetAllUsers() ([]*aggregates.User, error) {
	f.logger.Infof("Getting Users")
	var users []*aggregates.User

	result := f.database.First(&users)
	if result.Error != nil {
		f.logger.Errorw("Users Not Found", result.Error)
		return nil, errors.New("some error happens getting user")
	}

	if result.RowsAffected == 0 {
		f.logger.Errorw("Empty Users In Data Base")
		return nil, errors.New("users not found")
	}

	return users, nil
}
