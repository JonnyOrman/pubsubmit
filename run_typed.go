package pubsubmit

func RunTyped[T any]() {
	application := BuildApplication[T]()

	application.Run()
}
