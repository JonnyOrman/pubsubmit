package pubsubmit

import (
	"github.com/gin-gonic/gin"
	"github.com/jonnyorman/fireworks"
)

type GinRequestBodyReader[T any] struct {
	reader           fireworks.Reader
	bodyDeserialiser fireworks.DataDeserialiser[T]
}

func NewGinRequestBodyReader[T any](
	reader fireworks.Reader,
	bodyDeserialiser fireworks.DataDeserialiser[T],
) *GinRequestBodyReader[T] {
	this := new(GinRequestBodyReader[T])

	this.reader = reader
	this.bodyDeserialiser = bodyDeserialiser

	return this
}

func (this GinRequestBodyReader[T]) Read(ginContext *gin.Context) T {
	bodyByteArray := this.reader.Read(ginContext.Request.Body)

	body := this.bodyDeserialiser.Deserialise(bodyByteArray)

	return body
}
