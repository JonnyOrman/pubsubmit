package pubsubmit

func Run[T any]() {
	application := BuildApplication[T]()

	application.Run()
}
