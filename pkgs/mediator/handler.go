package mediator

import (
	"api.ddd/src/api/server"
	"errors"
)

// IHandler a base interface to create any handler
type IHandler interface {
	Handler(*Message) (interface{}, error)
}

// IDispatcher an interface that allow to access to
// Register and Send methods that allows to register command and command handler
// And send it without dependency just using de command struct
type IDispatcher interface {
	RegisterHandler(IHandler, interface{}) error
	Send(*Message) (interface{}, error)
}

// DispatcherMemory allow to save into a map
// pair command and command handler registered into the app
type DispatcherMemory struct {
	handlers map[string]IHandler
	logger   *server.Logger
}

// NewDispatcherMemory Provide a new DispatcherMemory to save handlers
// And return IDispatcher interface to allow access to RegisterHandler
// And Sen methods
func NewDispatcherMemory(
	logger *server.Logger,
) IDispatcher {
	return &DispatcherMemory{
		handlers: make(map[string]IHandler),
		logger:   logger,
	}
}

// RegisterHandler implementation where validate if a command was registered
// And register command that not was registered before
func (m *DispatcherMemory) RegisterHandler(handler IHandler, message interface{}) error {

	handlerName := TypeOf(message)
	if _, err := m.handlers[handlerName]; err {
		m.logger.Warnf("Error registering Handler: %s", handlerName)
		return errors.New("handler was already registered")
	}

	m.handlers[handlerName] = handler
	return nil
}

// Send allow to execute handler methods to specific
// handler that was registered previously and validate if it exists
func (m *DispatcherMemory) Send(message *Message) (interface{}, error) {

	handlerName := message.GetMessageString()
	handler, exist := m.handlers[message.GetMessageString()]
	if !exist {
		m.logger.Warnf("Error Executing Handler: %s", handlerName)
		return nil, errors.New("handler was not registered")
	}

	responseChan := make(chan interface{})
	errorChan := make(chan error)

	defer close(responseChan)
	defer close(errorChan)

	go ExecuteHandler(handler, message, responseChan, errorChan)

	response := <-responseChan
	err := <-errorChan

	return response, err
}

func ExecuteHandler(handler IHandler, message *Message, responseChan chan interface{}, errorChan chan error) {
	response, err := handler.Handler(message)
	responseChan <- response
	errorChan <- err
}
