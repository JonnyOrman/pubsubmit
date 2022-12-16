package pubsubmit

import "github.com/jonnyorman/fireworks"

func BuildApplication[T any]() *fireworks.Application {
	configuration := GenerateConfiguration("firesert-config")

	pubSubBodyDeserialiser := fireworks.JsonDataDeserialiser[fireworks.PubSubBody]{}

	ioutilReader := fireworks.IoutilReader{}

	bodyReader := fireworks.NewGinPubSubBodyReader(
		ioutilReader,
		pubSubBodyDeserialiser)

	dataDeserialiser := fireworks.JsonDataDeserialiser[T]{}

	dataReader := fireworks.NewHttpRequestBodyDataReader[T](
		bodyReader,
		dataDeserialiser)

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
