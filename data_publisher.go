package pubsubmit

type DataPublisher[T any] interface {
	Publish(data T)
}
