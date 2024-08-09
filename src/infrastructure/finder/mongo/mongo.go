package finder

import (
	"api.ddd/src/domain/aggregates"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (f Finder) GetUserById(userId string) (*aggregates.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	f.logger.Info("Getting User: " + userId)
	collection := f.database.Database("api-golang-ddd").Collection("user")
	filter := bson.M{
		"_id": userId,
	}

	var user *aggregates.User
	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		f.logger.Errorw("Error Getting User Id", err)
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (f Finder) GetAllUsers() ([]*aggregates.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	f.logger.Info("Getting All Users")
	collection := f.database.Database("api-golang-ddd").Collection("user")

	findOptions := options.Find()
	filter := bson.M{}

	var users []*aggregates.User

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		f.logger.Errorw("Error Getting Users", err)
		return nil, errors.New("user not found")
	}

	for cursor.Next(context.TODO()) {
		var record *aggregates.User
		err := cursor.Decode(&record)
		if err != nil {
			f.logger.Errorw("Error Adding User To List", err)
			return nil, errors.New("error decoding user")
		}
		users = append(users, record)
	}

	return users, nil
}
