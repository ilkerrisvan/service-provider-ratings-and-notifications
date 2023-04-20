package rating

import (
	"armut-rating-api/internal/data/pubsub/publisher"
	pbRating "armut-rating-api/internal/data/pubsub/publisher/rating"
	"armut-rating-api/internal/data/storage"
	"armut-rating-api/internal/util/env"
	"armut-rating-api/internal/util/logger"
	"armut-rating-api/internal/util/validator"
)

type IRatingService interface {
	AddRating(ch chan *AddRatingServiceResponse, model *AddRatingServiceModel)
	GetAverageService(ch chan *GetAverageServiceResponse, model *GetAverageServiceRequestModel)
	PublishPubSubMessage(ch chan *PublishPubSubMessageServiceResponse, model *PublishPubSubMessageServiceModel)
}

type RatingService struct {
	environment     env.IEnvironment
	loggr           logger.ILogger
	validatr        validator.IValidator
	ratingDb        storage.IRatingDb
	ratingPublisher pbRating.IRatingPublisher
}

func NewRatingService(environment env.IEnvironment, loggr logger.ILogger, ratingDb storage.IRatingDb, validatr validator.IValidator, publisher pbRating.IRatingPublisher) IRatingService {
	service := RatingService{
		environment: environment,
		loggr:       loggr,
		validatr:    validatr,
	}
	if ratingDb != nil {
		service.ratingDb = ratingDb
	} else {
		service.ratingDb = storage.NewRatingDb(environment, loggr, validatr)
	}
	if publisher != nil {
		service.ratingPublisher = publisher
	} else {
		service.ratingPublisher = pbRating.NewRatingPublisher(environment, loggr, validatr)
	}

	return &service
}

func (s *RatingService) PublishPubSubMessage(ch chan *PublishPubSubMessageServiceResponse, model *PublishPubSubMessageServiceModel) {
	publisherCh := make(chan publisher.PublisherResponse)
	defer close(publisherCh)
	go s.ratingPublisher.Publish(publisherCh, &model.Message)
	var messageIds []string

	response := <-publisherCh
	if response.Error != nil {
		ch <- &PublishPubSubMessageServiceResponse{Error: response.Error}
		s.loggr.Error("Message could not publish")
		return
	}
	messageIds = append(messageIds, *response.MessageId)
	ch <- &PublishPubSubMessageServiceResponse{
		MessageIds: messageIds,
	}
	close(ch)
}

func (s *RatingService) AddRating(ch chan *AddRatingServiceResponse, model *AddRatingServiceModel) {
	modelErr := s.validatr.ValidateStruct(model)
	if modelErr != nil {
		ch <- &AddRatingServiceResponse{Error: modelErr}
		s.loggr.Error("addRating-> model is not valid.")
		return
	}

	var publisherModel PublishPubSubMessageServiceModel
	publisherModel.Message.ServiceProviderRating = model.ServiceProviderRating
	publisherModel.Message.ServiceProviderId = model.ServiceProviderId

	publisherCh := make(chan *PublishPubSubMessageServiceResponse)
	//closes when message arrived.
	go s.PublishPubSubMessage(
		publisherCh, &PublishPubSubMessageServiceModel{
			Message: publisherModel.Message,
		},
	)

	setRatingCh := make(chan *storage.SetRatingDbResponse)
	go s.ratingDb.SetRatings(
		setRatingCh, &storage.SetRatingDbModel{
			ServiceProviderId:     model.ServiceProviderId,
			ServiceProviderRating: model.ServiceProviderRating,
		})

	response := AddRatingServiceResponse{RateData: model}
	ch <- &response
	return
}

func (s *RatingService) GetAverageService(ch chan *GetAverageServiceResponse, model *GetAverageServiceRequestModel) {
	getAverageCh := make(chan *storage.GetRatingDbResponse)
	go s.ratingDb.GetRatingAverage(
		getAverageCh,
		&storage.GetRatingDbModel{
			ServiceProviderId: model.ServiceProviderId,
		})
	response := <-getAverageCh

	if response.Error != nil {
		ch <- &GetAverageServiceResponse{Error: response.Error}
		s.loggr.Error("GetAverageService has an error caused by getRatingAverage")
		return
	}

	ch <- &GetAverageServiceResponse{
		AverageData: &GetAverageServiceResponseModel{
			AverageRating:    response.AverageRating,
			TotalRatingCount: response.TotalRatingCount,
		},
	}
}
