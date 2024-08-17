package query

import (
	"api.ddd/pkgs/mediator"
	"encoding/json"
)

func (h *GetUserByIdQueryHandler) Handler(message *mediator.Message) (interface{}, error) {

	query := message.GetMessage().(*GetUserByIdQuery)
	commandJson, _ := json.Marshal(query)
	h.logger.Info("Invoked GetUserByIdQueryHandler: " + string(commandJson))

	userDb, err := h.service.GetUserById(query.Id)
	if err != nil {
		return nil, err
	}

	return userDb, nil
}

func (h *GetAllUsersQueryHandler) Handler(message *mediator.Message) (interface{}, error) {

	h.logger.Info("Invoked GetAllUsersQueryHandler")

	usersDb, err := h.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return usersDb, nil
}
