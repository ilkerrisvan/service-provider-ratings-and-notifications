package model

type NotificationReceiverEventModel struct {
	ServiceProviderId     int     `json:"ServiceProviderId" validate:"required,gte=0"`
	ServiceProviderRating float32 `json:"ServiceProviderRating" validate:"required,gte=0,lte=5"`
}
