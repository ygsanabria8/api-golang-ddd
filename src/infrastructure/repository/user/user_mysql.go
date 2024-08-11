package repository

import (
	"api.ddd/src/domain/aggregates"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

func (r *MySqlUserRepository) CreateUser(user *aggregates.User) (*aggregates.User, error) {
	user.Id = uuid.New().String()
	userJson, _ := json.Marshal(user)
	r.logger.Infof("Saving User: " + string(userJson))

	result := r.database.Create(&user)
	if result.Error != nil {
		r.logger.Errorw("Error Creating User", result.Error)
		return nil, errors.New("some error happens saving into database")
	}

	return user, nil
}

func (r *MySqlUserRepository) DeleteUser(userId string) error {
	r.logger.Infof("Deleting User: " + userId)
	user := &aggregates.User{
		Id: userId,
	}

	result := r.database.Delete(user)
	if result.Error != nil {
		r.logger.Errorw("Error Deleting User", result.Error)
		return errors.New("some error happens deleting into database")
	}

	if result.RowsAffected == 0 {
		r.logger.Errorw("Not Deleted User")
		return errors.New("some error happens deleting into database")
	}

	return nil
}

func (r *MySqlUserRepository) UpdateUser(user *aggregates.User) (*aggregates.User, error) {
	userJson, _ := json.Marshal(user)
	r.logger.Info("Updating User: " + string(userJson))

	result := r.database.Save(&user)
	if result.Error != nil {
		r.logger.Errorw("Error in Updating User", result.Error)
		return nil, errors.New("some error happens updating into database")
	}

	if result.RowsAffected == 0 {
		r.logger.Errorw("User Not Updated")
		return nil, errors.New("some error happens updating into database")
	}

	return user, nil
}
