package notification

import (
	"armut-notification-api/internal/data/storage"
	"armut-notification-api/internal/util/env"
	"armut-notification-api/internal/util/logger"
	"armut-notification-api/internal/util/validator"
)

type INotificationService interface {
	GetNotificationService(ch chan *GetNotificationServiceResponse, model *GetNotificationServiceModel)
}

type NotificationService struct {
	validatr    validator.IValidator
	db          storage.INotificationDb
	environment env.IEnvironment
	loggr       logger.ILogger
}

func NewNotificationService(environment env.IEnvironment, loggr logger.ILogger, db storage.INotificationDb, validatr validator.IValidator) INotificationService {
	service := NotificationService{
		validatr:    validatr,
		environment: environment,
		loggr:       loggr,
	}
	if db != nil {
		service.db = db
	} else {
		service.db = storage.NewNotificationDb(environment, loggr, validatr)
	}
	return &service
}

func (s *NotificationService) GetNotificationService(ch chan *GetNotificationServiceResponse, model *GetNotificationServiceModel) {
	getNotificationsCh := make(chan *storage.GetRatingDbResponse)
	go s.db.GetRatings(
		getNotificationsCh,
		&storage.GetRatingDbModel{
			ServiceProviderId: model.ServiceProviderId,
		})
	response := <-getNotificationsCh

	if response.Error != nil {
		ch <- &GetNotificationServiceResponse{Error: response.Error}
		s.loggr.Error("GetNotificationService has an error caused by GetRatings")
		return
	}
	ch <- &GetNotificationServiceResponse{
		NotificationData: &GetNotificationServiceResponseModel{
			NotificationData: response.Notifications,
		},
	}

	deleteNotificationsCh := make(chan *storage.DeleteRatingDbResponse)
	go s.db.DeleteRatings(
		deleteNotificationsCh,
		&storage.DeleteRatingDbModel{
			ServiceProviderId: model.ServiceProviderId,
		})
}
