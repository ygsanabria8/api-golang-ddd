package finder

import (
	"api.ddd/pkgs/mediator"
	"encoding/json"
)

func (h *GetUserByIdQueryHandler) Handler(message *mediator.Message) (interface{}, error) {

	query := message.GetMessage().(*GetUserByIdQuery)
	commandJson, _ := json.Marshal(query)
	h.logger.Info("Invoked GetUserByIdQueryHandler: " + string(commandJson))

	exampleDb, err := h.service.GetUserById(query.Id)
	if err != nil {
		return nil, err
	}

	return exampleDb, nil
}
