package notification

import (
	"armut-notification-api/internal/api"
	"armut-notification-api/internal/data/storage"
	"armut-notification-api/internal/service/notification"
	"armut-notification-api/internal/util/env"
	"armut-notification-api/internal/util/logger"
	"armut-notification-api/internal/util/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type INotificationController interface {
	RegisterRoutes(routerGroup *gin.RouterGroup)
	GetNotification(context *gin.Context)
}

type NotificationController struct {
	path          string
	environment   env.IEnvironment
	loggr         logger.ILogger
	ratingService notification.INotificationService
	validatr      validator.IValidator
	db            storage.INotificationDb
}

func NewNotificationController(
	environment env.IEnvironment,
	loggr logger.ILogger,
	notificationService notification.INotificationService,
	validatr validator.IValidator,

) INotificationController {
	controller := NotificationController{
		path:        "notification",
		environment: environment,
		loggr:       loggr,
		validatr:    validatr,
	}

	if notificationService != nil {
		controller.ratingService = notificationService
	} else {
		controller.ratingService = notification.NewNotificationService(environment, loggr, nil, validatr)
	}

	return &controller
}

func (c *NotificationController) RegisterRoutes(routerGroup *gin.RouterGroup) {
	routes := routerGroup.Group(c.path)
	routes.GET("", c.GetNotification)
}

// GetNotification GetNotification
//
//	@basePath		/api
//	@router			/v1/notification [get]
//	@tags			Notifications
//	@summary		Shows the rates since the previous request
//	@description	Shows the rates since the previous request
//	@accept			json
//	@produce		json
//	@success		200			{object}	api.ApiResponse
//	@failure		400			{object}	api.ApiResponse
//	@Param			id	query		int	false	"Size of the page."
func (c *NotificationController) GetNotification(context *gin.Context) {
	Id, err := strconv.Atoi(context.Query("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, api.Error("Query is not valid."))
		c.loggr.Error("Query is not valid")
		return
	}
	ch := make(chan *notification.GetNotificationServiceResponse)
	defer close(ch)
	go c.ratingService.GetNotificationService(
		ch, &notification.GetNotificationServiceModel{
			ServiceProviderId: Id,
		},
	)

	serviceResponse := <-ch
	if serviceResponse.Error != nil {
		context.JSON(http.StatusBadRequest, api.Error("This id has no rate information."))
		c.loggr.Error("Request about id that has no rate information")
		return
	}
	context.JSON(http.StatusOK, api.Ok(serviceResponse.NotificationData))
	c.loggr.Info("Average data request is ok.")
}
