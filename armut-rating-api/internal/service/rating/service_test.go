package rating

import (
	"armut-rating-api/internal/data/pubsub/publisher"
	mockPubSubPublisher "armut-rating-api/internal/data/pubsub/publisher/rating"
	mockDb "armut-rating-api/internal/data/storage"
	"armut-rating-api/internal/util/env"
	"armut-rating-api/internal/util/logger"
	"armut-rating-api/internal/util/validator"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RatingServiceTestSuite struct {
	suite.Suite
	mockRatingService         IRatingService
	mockEnvironment           *env.MockIEnvironment
	mockLogger                *logger.MockILogger
	mockValidator             *validator.MockIValidator
	mockRatingDb              *mockDb.MockIRatingDb
	mockRatingPubSubPublisher *mockPubSubPublisher.MockIRatingPublisher
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
	s.mockRatingPubSubPublisher = mockPubSubPublisher.NewMockIRatingPublisher(ctrl)
	s.mockRatingDb = mockDb.NewMockIRatingDb(ctrl)

	s.mockRatingService = NewRatingService(s.mockEnvironment, s.mockLogger, s.mockRatingDb, s.mockValidator, s.mockRatingPubSubPublisher)
}

// Runs after each test in the suite.
func (s *RatingServiceTestSuite) TearDownTest() {
	s.T().Log("Teardown")
}
func (s *RatingServiceTestSuite) TestAddRating() {
	model := &AddRatingServiceModel{
		ServiceProviderId:     1,
		ServiceProviderRating: 5,
	}
	expected := &AddRatingServiceResponse{
		RateData: model,
	}

	s.mockValidator.EXPECT().ValidateStruct(gomock.Any()).Return(nil)
	s.mockRatingPubSubPublisher.EXPECT().Publish(gomock.Any(), gomock.Any()).DoAndReturn(func(mockCh chan publisher.PublisherResponse, model *mockPubSubPublisher.RatingPublisherModel) {})
	s.mockRatingDb.EXPECT().SetRatings(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ch chan *mockDb.SetRatingDbResponse, model *mockDb.SetRatingDbModel) {
			ch <- &mockDb.SetRatingDbResponse{
				ServiceProviderId:     expected.RateData.ServiceProviderId,
				ServiceProviderRating: expected.RateData.ServiceProviderRating,
			}
		})

	ch := make(chan *AddRatingServiceResponse)
	defer close(ch)
	go s.mockRatingService.AddRating(ch, model)

	response := <-ch
	s.Nil(response.Error)
	s.EqualValues(response.RateData.ServiceProviderId, expected.RateData.ServiceProviderId)
	s.EqualValues(response.RateData.ServiceProviderRating, expected.RateData.ServiceProviderRating)
}

func (s *RatingServiceTestSuite) TestGetAverageService() {
	model := &GetAverageServiceResponseModel{
		AverageRating:    1,
		TotalRatingCount: 1,
	}

	testReq := &GetAverageServiceRequestModel{
		ServiceProviderId: 1,
	}

	expected := &GetAverageServiceResponse{
		AverageData: model,
	}

	s.mockValidator.EXPECT().ValidateStruct(gomock.Any()).Return(nil)
	s.mockRatingDb.EXPECT().GetRatingAverage(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ch chan *mockDb.GetRatingDbResponse, model *mockDb.GetRatingDbModel) {
			ch <- &mockDb.GetRatingDbResponse{
				AverageRating:    expected.AverageData.AverageRating,
				TotalRatingCount: expected.AverageData.TotalRatingCount,
			}
		})

	ch := make(chan *GetAverageServiceResponse)
	defer close(ch)
	go s.mockRatingService.GetAverageService(ch, testReq)

	response := <-ch
	s.Nil(response.Error)
	s.EqualValues(response.AverageData.AverageRating, expected.AverageData.AverageRating)
	s.EqualValues(response.AverageData.TotalRatingCount, expected.AverageData.TotalRatingCount)
}
