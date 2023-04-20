package storage

type SetRatingDbModel struct {
	ServiceProviderId     int     `json:"ServiceProviderId" validate:"required,gte=0"`
	ServiceProviderRating float32 `json:"ServiceProviderRating" validate:"required,gte=0,lte=5"`
}
type DeleteRatingDbModel struct {
	ServiceProviderId int `json:"ServiceProviderId" validate:"required,gte=0"`
}

type GetRatingDbModel struct {
	ServiceProviderId int `json:"ServiceProviderId" validate:"required,gte=0"`
}
