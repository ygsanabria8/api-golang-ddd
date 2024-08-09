package mediator

// IMessage interface that allows to access
// To Message struct and struct message string
type IMessage interface {
	GetMessage() interface{}
	GetMessageString() string
}

// Message that has a interface with data
// And name of that interface
type Message struct {
	message       interface{}
	messageString string
}

// GetMessage implementation that return
// Interface with the message
func (r *Message) GetMessage() interface{} {
	return r.message
}

// GetMessageString implementation that return
// string with the message struct
func (r *Message) GetMessageString() string {
	return r.messageString
}

// CreateMessage function that create a Message struct
// With a specific interface provided and get its name string
func CreateMessage(message interface{}) *Message {
	return &Message{
		message:       message,
		messageString: TypeOf(message),
	}
}
