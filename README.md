Package go_unleash implements the possibility to connect your app to unleash in a better way, making it possible to use mocks and interfaces.

## Installation
Use go get.
```
go get github.com/hugokishi/go_unleash
```
Then import the validator package into your own code.
```
import "github.com/hugokishi/go_unleash"
```

## Usage
Sample code:
```go
package main

import (
	go_unleash "github.com/hugokishi/go_unleash"
	"github.com/hugokishi/go_unleash/internal/domain/structs"
)

func  main() {
	unClient := go_unleash.InitUnleash(
		structs.UnleashConfig{
			AppEnvironment: "YOUR_APP_ENVIRONMENT",
			AppName: "YOUR_APP_NAME",
			UnleashURL: "YOUR_UNLEASH_URL",
			UnleashAuthorizationToken: "YOUR_AUTHORIZATION_TOKEN",
			LoggingLevel: "YOUR_LOGGING_LEVEL",
		},
	)
	
	isEnabled := unClient.IsEnabled("MY_FLAG")
}
```

## Parameters:
| Parameter | Description | Value |
| :-------------: |:--------:| :-------------: |
| AppEnvironment | Refers to the environment that unleash is being configured for  | local, development, staging, production
| AppName | Refers to the name that will be registered in unleash | Defined by User |
| UnleashURL | Refers to the url of connection with unleash | Defined by User |
| UnleashAuthorizationToken | Refers to the authorization token for connecting to unleash | Defined by User |
| LoggingLevel | Refers to the logging level used in library | warn, info, debug, error (default info) |
 
