package main

import (
	"armut-notification-api/docs"
	"armut-notification-api/internal/api"
	controllerV1 "armut-notification-api/internal/api/v1/controller/notification"
	ps "armut-notification-api/internal/data/pubsub"
	"armut-notification-api/internal/util/env"
	"armut-notification-api/internal/util/logger"
	"armut-notification-api/internal/util/validator"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	environment := env.New()
	environment.Init()
	loggr := logger.New(environment)
	validatr := validator.New()
	router := gin.Default()
	addSwagger(router, environment)
	router.Use(api.RateLimiter(loggr))
	addRoutes(router, environment, loggr, validatr)
	ps.AddReceivers(environment, loggr, validatr)
	err := router.Run(environment.Get(env.AppHost))
	if err != nil {
		loggr.Error("Application fail to start.")
	}
}

func addRoutes(router *gin.Engine, environment env.IEnvironment, loggr logger.ILogger, validatr validator.IValidator) {
	api := router.Group("api")
	v1 := api.Group("v1")
	controllerV1.NewNotificationController(environment, loggr, nil, validatr).RegisterRoutes(v1)
}

func addSwagger(router *gin.Engine, environment env.IEnvironment) {
	docs.SwaggerInfo.Title = fmt.Sprintf("Armut Case Study Notification API (%v)", environment.Get(env.AppEnvironment))
	docs.SwaggerInfo.Host = environment.Get(env.AppHost)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
