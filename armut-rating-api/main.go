package main

import (
	"armut-rating-api/docs"
	"armut-rating-api/internal/api"
	controllerV1 "armut-rating-api/internal/api/v1/controller/rating"
	pb "armut-rating-api/internal/data/pubsub/publisher/rating"
	"armut-rating-api/internal/util/env"
	"armut-rating-api/internal/util/logger"
	"armut-rating-api/internal/util/validator"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title					Armut Case Study Rating API
// @version					1.0
// @description				This is API where service providers can be given a rate and you can learn the rate averages.
// @termsOfService			http://swagger.io/terms/
// @contact.name			İlker Rişvan
// @contact.email			ilkerrisvan@outlook.com
// @license.name			Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host					0.0.0.0:8000
// @BasePath				/api
// @accept					json
// @produce					json
func main() {
	environment := env.New()
	environment.Init()
	loggr := logger.New(environment)
	validatr := validator.New()

	var publisher pb.IRatingPublisher

	router := gin.Default()
	addSwagger(router, environment)
	router.Use(api.RateLimiter(loggr))
	addRoutes(router, environment, loggr, publisher, validatr)
	err := router.Run(environment.Get(env.AppHost))
	if err != nil {
		loggr.Error("Application fail to start.")
	}
}

func addRoutes(router *gin.Engine, environment env.IEnvironment, loggr logger.ILogger, publisher pb.IRatingPublisher, validatr validator.IValidator) {
	api := router.Group("api")
	v1 := api.Group("v1")
	controllerV1.NewRatingController(environment, loggr, nil, publisher, validatr).RegisterRoutes(v1)
}

func addSwagger(router *gin.Engine, environment env.IEnvironment) {
	docs.SwaggerInfo.Title = fmt.Sprintf("Armut Case Study Rating API (%v)", environment.Get(env.AppEnvironment))
	docs.SwaggerInfo.Host = environment.Get(env.AppHost)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
