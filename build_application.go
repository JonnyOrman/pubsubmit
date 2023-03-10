package pubsubmit

import "github.com/jonnyorman/fireworks"

func BuildApplication[T any]() *fireworks.Application {
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

	routerBuilder := fireworks.NewGinRouterBuilder()

	routerBuilder.AddPost("/", requestHandler.Handle)

	router := routerBuilder.Build()

	application := fireworks.NewApplication(router)

	return application
}
