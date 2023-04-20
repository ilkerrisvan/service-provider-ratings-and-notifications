package notification

type GetNotificationServiceResponse struct {
	NotificationData *GetNotificationServiceResponseModel
	Error            error `json:"-"`
}
type PublishPubSubMessageServiceResponse struct {
	Error      error `json:"-"`
	MessageIds []string
}
