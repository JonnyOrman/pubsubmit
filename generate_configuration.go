package pubsubmit

import "github.com/jonnyorman/fireworks"

func GenerateConfiguration(fileName string) *Configuration {
	configurationFilePathProvider := fireworks.NewConfigurationFilePathProvider(fileName)

	configurationFileReader := fireworks.NewConfigurationFileReader(configurationFilePathProvider)

	configurationJsonFileReader := fireworks.NewConfigurationJsonFileReader(configurationFileReader)

	configurationJson := configurationJsonFileReader.Read()

	projectIDProvider := fireworks.CreateConfigurationValueProvider("projectID", "PROJECT_ID", configurationJson)

	collectionNameProvider := fireworks.CreateConfigurationValueProvider("collectionName", "COLLECTION_NAME", configurationJson)

	operationProvider := fireworks.CreateConfigurationValueProvider("operation", "OPERATION", configurationJson)

	configurationLoader := fireworks.NewApplicationConfigurationLoader(
		projectIDProvider,
		collectionNameProvider,
	)

	fireworksConfiguration := configurationLoader.Load()

	operation, _ := operationProvider.Get()

	configuration := NewConfiguration(fireworksConfiguration, operation)

	return configuration
}
