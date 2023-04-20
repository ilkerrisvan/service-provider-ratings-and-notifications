package rating

import (
	"armut-rating-api/internal/data/pubsub/publisher"
	"armut-rating-api/internal/util/env"
	"armut-rating-api/internal/util/logger"
	"armut-rating-api/internal/util/validator"
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"time"
)

type IRatingPublisher interface {
	Publish(ch chan publisher.PublisherResponse, model *RatingPublisherModel)
}

type RatingPublisher struct {
	environment env.IEnvironment
	loggr       logger.ILogger
	validatr    validator.IValidator
	projectId   string
	topicId     string
	timeout     time.Duration
}

func NewRatingPublisher(environment env.IEnvironment, loggr logger.ILogger, validatr validator.IValidator) IRatingPublisher {
	projectId := environment.Get(env.RatingPublisherProjectId)
	topicId := environment.Get(env.RatingPublisherTopicId)

	publisher := RatingPublisher{
		loggr:     loggr,
		validatr:  validatr,
		projectId: projectId,
		topicId:   topicId,
		timeout:   time.Second * 60,
	}
	return &publisher
}

func (p *RatingPublisher) Publish(ch chan publisher.PublisherResponse, model *RatingPublisherModel) {
	ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
	defer cancel()

	client, err := pubsub.NewClient(ctx, p.projectId)
	if err != nil {
		ch <- publisher.PublisherResponse{Error: err}
		p.loggr.Error("publisher has error.")
		return
	}
	defer client.Close()

	bytes, err := json.Marshal(model)
	if err != nil {
		ch <- publisher.PublisherResponse{Error: err}
		p.loggr.Error("publish model could not marshal")
		return
	}

	topic := client.Topic(p.topicId)
	result := topic.Publish(ctx, &pubsub.Message{
		Data: bytes,
	})

	messageId, err := result.Get(ctx)
	if err != nil {
		ch <- publisher.PublisherResponse{Error: err}
		p.loggr.Error("publisher has error.")
		return
	}
	ch <- publisher.PublisherResponse{MessageId: &messageId}
	p.loggr.Info("publish is ok.")
	return
}
