package storage

type SetRatingDbModel struct {
	ServiceProviderId     int `json:"ServiceProviderId" validate:"required,gte=0"`
	ServiceProviderRating int `json:"ServiceProviderRating" validate:"required,gte=0,lte=5"`
}

type GetRatingDbModel struct {
	ServiceProviderId int `json:"ServiceProviderId" validate:"required,gte=0"`
}
