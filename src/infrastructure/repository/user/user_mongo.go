package repository

import (
	"api.ddd/src/domain/aggregates"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (r *MongoUserRepository) CreateUser(user *aggregates.User) (*aggregates.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	user.Id = uuid.New().String()
	userJson, _ := json.Marshal(user)
	r.logger.Infof("Saving User: " + string(userJson))

	collection := r.database.Database("api-golang-ddd").Collection("user")
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		r.logger.Errorw("Error in CreateMongoUser", err)
		return nil, errors.New("some error happens saving into database")
	}
	return user, nil
}

func (r *MongoUserRepository) DeleteUser(userId string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	r.logger.Infof("Deleting User: " + userId)
	filter := bson.M{
		"_id": userId,
	}

	collection := r.database.Database("api-golang-ddd").Collection("user")
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		r.logger.Errorw("Error Deleting User", err)
		return errors.New("some error happens deleting into database")
	}

	if result.DeletedCount == 0 {
		r.logger.Errorw("Not Deleted User")
		return errors.New("some error happens deleting into database")
	}
	return nil
}

func (r *MongoUserRepository) UpdateUser(user *aggregates.User) (*aggregates.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	productJson, _ := json.Marshal(user)
	r.logger.Info("Updating User Mongo: " + string(productJson))
	filter := bson.M{
		"_id": user.Id,
	}

	collection := r.database.Database("api-golang-ddd").Collection("user")
	result, err := collection.ReplaceOne(ctx, filter, user)
	if err != nil {
		r.logger.Errorw("Error in Updating User", err)
		return nil, errors.New("some error happens updating into database")
	}

	if result.ModifiedCount == 0 {
		r.logger.Errorw("User Not Updated", err)
		return nil, errors.New("some error happens updating into database")
	}

	return user, nil
}
