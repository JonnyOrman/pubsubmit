# pubsubmit

Create Go microservices that publish data to a Pub/Sub topic.

Very easy to use, create a working service by adding just one line of code to your `main` method:
```
pubsubmit.Run()
```

## Getting started

Create a new Go project
```
mkdir pubsubmit-example
cd pubsubmit-example    
go mod init pubsubmit/example
```

Get `pubsubmit`
```
go get github.com/jonnyorman/pubsubmit
```

Add a `main.go` file with the following
```
package main

import "github.com/jonnyorman/pubsubmit"

func main() {
	pubsubmit.Run()
}

```

Add a `pubsubmit-config.json` file with the following
```
{
    "projectID": "your-firebase-project",
    "collectionName": "entities",
    "operation": "create"
}
```

Tidy and run with access to a Firebase project or emulator
```
    go mod tidy
    go run .
```

Submit a `POST` to the service with a body. The body data will get published as a message to a topic called `entities-create-submit`.

You can also create a struct with the data you want to publish in the message. Create a struct and use it:
```
package main

import "github.com/jonnyorman/pubsubmit"

type EntityModel struct {
	Prop1 string
	Prop2 int
}

func main() {
	pubsubmit.RunTyped[EntityModel]()
}
```

## Environment configuration

The configuration can also be provided by the environment with the following keys:
- `projectID` - `PROJECT_ID`
- `collectionName` - `COLLECTION_NAME`
- `operation` - `OPERATION`

A combination of the `pubsubmit-config.json` file and environment variables can be used. For example, the project ID could be provided as the `PROJECT_ID` environment variable, while the collection name and operation are provided with the following configuration file:
```
{
    "collectionName": "entities",
    "operation": "create"
}
```

If a configuration value is provided in both `pubsubmit-config.json` and the environment, then the configuration file with take priority. For example, if the `PROJECT_ID` envronment varialbe has value "env-project-id" and the following `pubsubmit-config.json` file is provided:
```
{
    "projectID": "config-project-id"
}
```
then the project ID will be "config-project-id".