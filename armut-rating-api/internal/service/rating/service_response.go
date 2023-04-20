package rating

type AddRatingServiceResponse struct {
	RateData *AddRatingServiceModel
	Error    error `json:"-"`
}

type GetAverageServiceResponse struct {
	AverageData *GetAverageServiceResponseModel
	Error       error `json:"-"`
}
type PublishPubSubMessageServiceResponse struct {
	Error      error `json:"-"`
	MessageIds []string
}
