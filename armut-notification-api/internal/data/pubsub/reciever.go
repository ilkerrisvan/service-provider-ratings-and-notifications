package pubsub

import (
	"armut-notification-api/internal/data/pubsub/receiver"
	"armut-notification-api/internal/util/env"
	"armut-notification-api/internal/util/logger"
	"armut-notification-api/internal/util/validator"
)

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

func AddReceivers(environment env.IEnvironment, loggr logger.ILogger, validatr validator.IValidator) {
	go receiver.NewNotificationReceiver(environment, loggr, validatr).InitReceivers(1)
}
