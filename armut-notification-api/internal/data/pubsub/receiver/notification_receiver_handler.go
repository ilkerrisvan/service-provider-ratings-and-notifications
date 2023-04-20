package receiver

import (
	receiverEventModel "armut-notification-api/internal/data/pubsub/model"
	"armut-notification-api/internal/data/storage"
	"armut-notification-api/internal/util/env"
	"armut-notification-api/internal/util/logger"
	"armut-notification-api/internal/util/validator"
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"time"
)

type INotificationReceiver interface {
	InitReceivers(count int)
}
type NotificationReceiver struct {
	projectId      string
	subscriptionId string
	timeout        time.Duration
	db             storage.INotificationDb
	environment    env.IEnvironment
}

func NewNotificationReceiver(environment env.IEnvironment, loggr logger.ILogger, validatr validator.IValidator) INotificationReceiver {
	receiver := NotificationReceiver{
		projectId:      environment.Get(env.NotificationReceiverProjectId),
		subscriptionId: environment.Get(env.SubscriptionId),
		timeout:        10,
	}
	receiver.db = storage.NewNotificationDb(environment, loggr, validatr)
	return &receiver
}
func (r *NotificationReceiver) InitReceivers(count int) {
	for i := 0; i < count; i++ {
		go r.ReceiveMessages()
	}
}
func (r *NotificationReceiver) ReceiveMessages() {
	msg := make(chan *pubsub.Message, 1)
	ctx := context.Background()
	client, _ := pubsub.NewClient(ctx, r.projectId)
	sub := client.Subscription(r.subscriptionId)
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go sub.Receive(cctx, func(ctx context.Context, m *pubsub.Message) {
		msg <- m
	})

	for {
		select {
		case res := <-msg:
			res.Ack()
			var model receiverEventModel.NotificationReceiverEventModel
			_ = json.Unmarshal(res.Data, &model)
			setRatingCh := make(chan *storage.SetRatingDbResponse)
			go r.db.SetRatings(
				setRatingCh, &storage.SetRatingDbModel{
					ServiceProviderId:     model.ServiceProviderId,
					ServiceProviderRating: model.ServiceProviderRating,
				})
		}
	}
}
