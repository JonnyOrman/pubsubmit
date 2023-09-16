package pubsubmit

func RunTyped[T any](allowedOrigins []string) {
	application := BuildApplication[T](allowedOrigins)

	application.Run()
}
