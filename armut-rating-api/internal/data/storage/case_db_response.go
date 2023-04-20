package storage

type SetRatingDbResponse struct {
	Error                 error `json:"-"`
	ServiceProviderId     int   `json:"ServiceProviderId" validate:"required,gte=0"`
	ServiceProviderRating int   `json:"ServiceProviderRating" validate:"required,gte=0,lte=5"`
}

type GetRatingDbResponse struct {
	Error            error   `json:"-"`
	AverageRating    float32 `json:"average_rating"`
	TotalRatingCount int     `json:"total_rating_count"`
}
