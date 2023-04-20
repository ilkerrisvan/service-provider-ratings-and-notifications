package notification

type IReceiver interface {
	InitReceivers(count int)
}

type IReceiverHandler interface {
	Handle(ch chan error, model *ReceiverHandlerModel)
}

type ReceiverHandlerModel struct {
	MessageId  string `validate:"required"`
	Data       []byte `validate:"required"`
	Attributes map[string]string
}
