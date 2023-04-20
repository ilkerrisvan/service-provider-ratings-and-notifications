package notification

import (
	"armut-notification-api/internal/util/logger"
	"armut-notification-api/internal/util/validator"
	"cloud.google.com/go/pubsub"
	"context"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type IRatingReceiver interface {
	InitReceivers(count int)
}
type RatingReceiver struct {
	loggr                 logger.ILogger
	validatr              validator.IValidator
	projectId             string
	subscriptionId        string
	timeout               time.Duration
	defaultAttributes     map[string]string
	messageCacheKeyPrefix string
	handler               IReceiverHandler
	receiverName          string
}

// InitReceivers
// Initializes multiple receivers that listen given topic for new messages.
func (r *RatingReceiver) InitReceivers(count int) {
	r.loggr.Info(r.receiverName + " Initializing receivers.")
	for i := 0; i < count; i++ {
		go r.receive()
	}
	r.loggr.Info(r.receiverName + " Initialized " + strconv.Itoa(count) + " receivers.")
}

func (r *RatingReceiver) receive() {
	// Register another receiver if one of them fails.
	defer func() {
		if rec := recover(); rec != nil {
			r.loggr.Error(r.receiverName+" Recovered the panic. Trying to receive again.", zap.Any("panic", rec))
			r.receive()
		}
	}()

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, r.projectId)
	if err != nil {
		r.loggr.Panic(r.receiverName + " Panicked while creating pub sub client.")
	}
	defer client.Close()

	subscription := client.Subscription(r.subscriptionId)
	err = subscription.Receive(ctx, r.eventHandler)
	if err != nil {
		r.loggr.Panic(r.receiverName + " Panicked while receiving messages from subscription.")
	}
}

func (r *RatingReceiver) eventHandler(ctx context.Context, msg *pubsub.Message) {
	defer func() {
		if rec := recover(); rec != nil {
			r.loggr.Error(r.receiverName+" "+msg.ID+" ID message is panicked during execution in event handler.",
				zap.String("messageId", msg.ID),
				zap.String("data", string(msg.Data)),
				zap.Any("attributes", msg.Attributes),
				zap.Any("panic", rec),
			)
		}
	}()

	ch := make(chan error)
	defer close(ch)
	go r.handler.Handle(ch, &ReceiverHandlerModel{
		MessageId:  msg.ID,
		Data:       msg.Data,
		Attributes: msg.Attributes,
	})

	err := <-ch
	if err != nil {
		r.loggr.Error(r.receiverName+" "+msg.ID+" ID message is failed to process.",
			zap.String("messageId", msg.ID),
			zap.String("data", string(msg.Data)),
			zap.Any("attributes", msg.Attributes),
			zap.Error(err),
		)
		return
	}

	// Success
	msg.Ack()
	r.loggr.Info(r.receiverName+" "+msg.ID+" ID message is processed successfully.",
		zap.String("messageId", msg.ID),
		zap.String("data", string(msg.Data)),
		zap.Any("attributes", msg.Attributes),
	)
}
