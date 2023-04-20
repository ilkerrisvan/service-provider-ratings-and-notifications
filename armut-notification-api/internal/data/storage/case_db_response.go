package storage

type SetRatingDbResponse struct {
	Error                 error `json:"-"`
	ServiceProviderId     int   `json:"ServiceProviderId" validate:"required,gte=0"`
	ServiceProviderRating int   `json:"ServiceProviderRating" validate:"required,gte=0,lte=5"`
}

type GetRatingDbResponse struct {
	Error         error `json:"-"`
	Notifications []int `json:"average_rating"`
}

type DeleteRatingDbResponse struct {
	Error error `json:"-"`
}
