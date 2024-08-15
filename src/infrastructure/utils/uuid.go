package utils

import "github.com/google/uuid"

func GetUUID() string {
	return uuid.New().String()
}

func IsValidUUID(id string) (string, error) {
	uuidObject, err := uuid.Parse(id)
	return uuidObject.String(), err
}
