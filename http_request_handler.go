package pubsubmit

import (
	"github.com/gin-gonic/gin"
	"github.com/jonnyorman/fireworks"
)

type HttpRequestHandler[T any] struct {
	dataReader    fireworks.DataReader[T]
	dataPublisher DataPublisher[T]
}

func NewHttpRequestHandler[T any](
	dataReader fireworks.DataReader[T],
	dataPublisher DataPublisher[T]) *HttpRequestHandler[T] {
	this := new(HttpRequestHandler[T])

	this.dataReader = dataReader
	this.dataPublisher = dataPublisher

	return this
}

func (this HttpRequestHandler[T]) Handle(ginContext *gin.Context) {
	data := this.dataReader.Read(ginContext)

	this.dataPublisher.Publish(data)
}
