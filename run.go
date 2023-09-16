package pubsubmit

func Run(allowedOrigins []string) {
	RunTyped[map[string]interface{}](allowedOrigins)
}
