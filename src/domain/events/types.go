package events

import "api.ddd/src/domain/aggregates"

type CreatedUserEvent struct {
	User *aggregates.User
}

type UpdatedUserEvent struct {
	User *aggregates.User
}

type DeletedUserEvent struct {
	UserId string
}
