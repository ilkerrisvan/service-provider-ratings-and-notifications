package rating

import (
	"armut-rating-api/internal/api"
	pbRating "armut-rating-api/internal/data/pubsub/publisher/rating"
	"armut-rating-api/internal/service/rating"
	"armut-rating-api/internal/util/env"
	"armut-rating-api/internal/util/logger"
	"armut-rating-api/internal/util/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IRatingController interface {
	RegisterRoutes(routerGroup *gin.RouterGroup)
	GetAverage(context *gin.Context)
	AddRating(context *gin.Context)
}

type RatingController struct {
	path          string
	environment   env.IEnvironment
	loggr         logger.ILogger
	ratingService rating.IRatingService
	validatr      validator.IValidator
}

func NewRatingController(
	environment env.IEnvironment,
	loggr logger.ILogger,
	ratingService rating.IRatingService,
	publisher pbRating.IRatingPublisher,
	validatr validator.IValidator,
) IRatingController {
	controller := RatingController{
		path:        "rating",
		environment: environment,
		loggr:       loggr,
		validatr:    validatr,
	}
	if ratingService != nil {
		controller.ratingService = ratingService
	} else {
		controller.ratingService = rating.NewRatingService(environment, loggr, nil, validatr, publisher)
	}
	return &controller
}

func (r *RatingController) RegisterRoutes(routerGroup *gin.RouterGroup) {
	routes := routerGroup.Group(r.path)
	routes.GET("average", r.GetAverage)
	routes.POST("", r.AddRating)
}

// AddRating Give Rate
//
//	@basePath		/api
//	@router			/v1/rating [post]
//	@tags			Ratings
//	@summary		Rate can be given with this endpoint.
//	@description	Rate can be given with this endpoint.
//	@accept			json
//	@produce		json
//	@success		201			{object}	api.ApiResponse
//	@failure		400			{object}	api.ApiResponse
//	@Param			Model		body		AddRatingModel	true	"Request model"
func (r *RatingController) AddRating(context *gin.Context) {
	var model AddRatingModel

	err := context.ShouldBindJSON(&model)
	if err != nil {
		context.JSON(http.StatusBadRequest, api.Error("Not valid request."))
		r.loggr.Error("Not valid request")
		return
	}

	ch := make(chan *rating.AddRatingServiceResponse)
	defer close(ch)
	go r.ratingService.AddRating(
		ch, &rating.AddRatingServiceModel{
			ServiceProviderId:     model.ServiceProviderId,
			ServiceProviderRating: model.ServiceProviderRating,
		},
	)
	serviceResponse := <-ch
	if serviceResponse.Error != nil {
		context.JSON(http.StatusBadRequest, api.Error("Request is not valid. Check the body."))
		r.loggr.Error("Not valid request")
		return
	}
	context.JSON(http.StatusCreated, api.Ok(serviceResponse.RateData))
	r.loggr.Info("New rate is created.")

}

// GetAverage Get Average Information
//
//	@basePath		/api
//	@router			/v1/rating/average [get]
//	@tags			Ratings
//	@summary		You can get the rate average information with the Id information.
//	@description	You can get the rate average information with the Id information.
//	@accept			json
//	@produce		json
//	@success		200			{object}	api.ApiResponse
//	@failure		400			{object}	api.ApiResponse
//	@Param			id	query		int	false	"id of the service provider"
func (r *RatingController) GetAverage(context *gin.Context) {
	Id, err := strconv.Atoi(context.Query("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, api.Error("Query is not valid."))
		r.loggr.Error("Query is not valid")
		return
	}

	ch := make(chan *rating.GetAverageServiceResponse)
	defer close(ch)
	go r.ratingService.GetAverageService(
		ch, &rating.GetAverageServiceRequestModel{
			ServiceProviderId: Id,
		},
	)
	serviceResponse := <-ch
	if serviceResponse.Error != nil {
		context.JSON(http.StatusBadRequest, api.Error("This id has no rate information."))
		r.loggr.Error("Request about id that has no rate information")
		return
	}
	context.JSON(http.StatusOK, api.Ok(serviceResponse.AverageData))
	r.loggr.Info("Average data request is ok.")
}
