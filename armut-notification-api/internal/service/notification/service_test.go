package notification

import (
	mockDb "armut-notification-api/internal/data/storage"
	"armut-notification-api/internal/util/env"
	"armut-notification-api/internal/util/logger"
	"armut-notification-api/internal/util/validator"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RatingServiceTestSuite struct {
	suite.Suite
	mockRatingService INotificationService
	mockEnvironment   *env.MockIEnvironment
	mockLogger        *logger.MockILogger
	mockValidator     *validator.MockIValidator
	mockRatingDb      *mockDb.MockINotificationDb
}

// Run suite.
func TestService(t *testing.T) {
	suite.Run(t, new(RatingServiceTestSuite))
}

// Runs before each test in the suite.
func (s *RatingServiceTestSuite) SetupTest() {
	s.T().Log("Setup")

	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	s.mockEnvironment = env.NewMockIEnvironment(ctrl)
	s.mockLogger = logger.NewMockILogger(ctrl)
	s.mockValidator = validator.NewMockIValidator(ctrl)
	s.mockRatingDb = mockDb.NewMockINotificationDb(ctrl)

	s.mockRatingService = NewNotificationService(s.mockEnvironment, s.mockLogger, s.mockRatingDb, s.mockValidator)
}

// Runs after each test in the suite.
func (s *RatingServiceTestSuite) TearDownTest() {
	s.T().Log("Teardown")
}
func (s *RatingServiceTestSuite) TestGetAverageService() {
	var tempRateValues = []int{1, 2, 3, 4}

	expected := &GetNotificationServiceResponseModel{
		NotificationData: tempRateValues,
	}

	model := &GetNotificationServiceModel{
		ServiceProviderId: 1,
	}

	s.mockValidator.EXPECT().ValidateStruct(gomock.Any()).Return(nil)
	s.mockRatingDb.EXPECT().GetRatings(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ch chan *mockDb.GetRatingDbResponse, model *mockDb.GetRatingDbModel) {
			ch <- &mockDb.GetRatingDbResponse{
				Notifications: expected.NotificationData,
			}
		})
	s.mockRatingDb.EXPECT().DeleteRatings(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ch chan *mockDb.DeleteRatingDbResponse, model *mockDb.DeleteRatingDbModel) {
			ch <- &mockDb.DeleteRatingDbResponse{
				Error: nil,
			}
		})

	ch := make(chan *GetNotificationServiceResponse)
	defer close(ch)
	go s.mockRatingService.GetNotificationService(ch, model)

	response := <-ch
	s.Nil(response.Error)
	s.EqualValues(response.NotificationData.NotificationData, expected.NotificationData)
}
