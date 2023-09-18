package pubsubmit

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jonnyorman/fireworks"
)

func BuildApplication[T any](allowedOrigins []string) *fireworks.Application {
	configuration := GenerateConfiguration("pubsubmit-config")

	bodyDeserialiser := fireworks.JsonDataDeserialiser[T]{}

	ioutilReader := fireworks.IoutilReader{}

	dataReader := NewGinRequestBodyReader[T](
		ioutilReader,
		bodyDeserialiser,
	)

	dataPublisher := NewPubSubDataPublisher[T](configuration)

	requestHandler := NewHttpRequestHandler[T](
		dataReader,
		dataPublisher,
	)

	// routerBuilder := fireworks.NewGinRouterBuilder()

	// routerBuilder.AddPost("/", requestHandler.Handle)

	// router := routerBuilder.Build()

	router := gin.Default()

	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = allowedOrigins
	// router.Use(cors.New(corsConfig))

	router.Use(cors.Default())

	router.POST("/", requestHandler.Handle)

	application := fireworks.NewApplication(router)

	return application
}
