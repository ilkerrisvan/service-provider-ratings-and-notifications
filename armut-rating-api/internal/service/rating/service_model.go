package rating

import (
	pb "armut-rating-api/internal/data/pubsub/publisher/rating"
)

type AddRatingServiceModel struct {
	ServiceProviderId     int `json:"ServiceProviderId" validate:"required,gte=0"`
	ServiceProviderRating int `json:"ServiceProviderRating" validate:"required,gte=0,lte=5"`
}

type GetAverageServiceRequestModel struct {
	ServiceProviderId int `json:"ServiceProviderId" validate:"required,gte=0"`
}

type GetAverageServiceResponseModel struct {
	AverageRating    float32 `json:"AverageRating" validate:"required,gte=0"`
	TotalRatingCount int     `json:"TotalRatingCount" validate:"required,gte=0"`
}

type PublishPubSubMessageServiceModel struct {
	Message pb.RatingPublisherModel `validate:"required"`
}
