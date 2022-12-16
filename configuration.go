package pubsubmit

import "github.com/jonnyorman/fireworks"

type Configuration struct {
	fireworks.Configuration
	Operation string
}

func NewConfiguration(
	fireworksConfiguration fireworks.Configuration,
	operation string) *Configuration {
	this := new(Configuration)

	this.Configuration = fireworksConfiguration
	this.Operation = operation

	return this
}
