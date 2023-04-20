package rating

import "time"

type RatingPublisherModel struct {
	ServiceProviderId     int       `json:"ServiceProviderId" validate:"required,gte=0"`
	ServiceProviderRating int       `json:"ServiceProviderRating" validate:"required,gte=0,lte=5"`
	Date                  time.Time `json:"Date" validate:"required"`
}
