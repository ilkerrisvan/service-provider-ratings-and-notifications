package rating

import (
	pubsubPublisher "armut-rating-api/internal/data/pubsub/publisher/rating"
)

type AddRatingModel struct {
	ServiceProviderId     int `json:"ServiceProviderId" validate:"required,gte=0"`
	ServiceProviderRating int `json:"ServiceProviderRating" validate:"required,gte=0,lte=5"`
}
type GetAverageModel struct {
	ServiceProviderId    int     `json:"ServiceProviderId" validate:"required,gte=0"`
	ServiceAverageRating float32 `json:"ServiceAverageRating" validate:"required,gte=0,lte=5"`
}
type PublishPubSubMessageModel struct {
	Message pubsubPublisher.RatingPublisherModel `json:"Message" validate:"required"`
}
