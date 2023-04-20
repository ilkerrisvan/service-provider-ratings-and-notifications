package notification

type AddRatingServiceModel struct {
	ServiceProviderId     int `json:"ServiceProviderId" validate:"required,gte=0"`
	ServiceProviderRating int `json:"ServiceProviderRating" validate:"required,gte=0,lte=5"`
}

type GetNotificationServiceModel struct {
	ServiceProviderId int `json:"ServiceProviderId" validate:"required,gte=0"`
}

type GetNotificationServiceResponseModel struct {
	NotificationData []int `json:"NotificationData" validate:"required,gte=0"`
}
