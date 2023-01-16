package go_unleash

import (
	"fmt"

	"github.com/Unleash/unleash-client-go/v3"
	"github.com/hugokishi/go-unleash/internal/domain/memory"
	"github.com/hugokishi/go-unleash/internal/domain/structs"

	"github.com/Unleash/unleash-client-go/v3/context"
	"github.com/hugokishi/go-unleash/internal/driver/logs"
	unleashDriver "github.com/hugokishi/go-unleash/internal/driver/unleash"
	"github.com/sirupsen/logrus"
)

type Unleash struct {
	unleash unleash.Client
}

func InitUnleash(config structs.UnleashConfig) {
	memory.SetVariables(
		config.AppEnvironment,
		config.LoggingLevel,
	)

	logs.Init()

	if err := unleashDriver.InitializeUnleash(structs.UnleashConfig{
		AppEnvironment:            config.AppEnvironment,
		AppName:                   config.AppName,
		UnleashURL:                config.UnleashURL,
		UnleashAuthorizationToken: config.UnleashAuthorizationToken,
	}); err != nil {
		logrus.Errorf("Unable to start connection to unleash, error: %v", err)
		return
	}

	logrus.Infof("The unleash configuration was successful, using the %v environment", "")
}

func (u *Unleash) IsEnabled(flagName string) bool {
	var composedFlag string

	if memory.APP_ENV_AppEnvironment != "production" {
		composedFlag = fmt.Sprintf("%v-%v", memory.APP_ENV_AppEnvironment, flagName)
	}

	return u.unleash.IsEnabled(composedFlag)
}

func (u *Unleash) IsEnabledContext(flagName string, ctx context.Context) bool {
	var composedFlag string

	if memory.APP_ENV_AppEnvironment != "production" {
		composedFlag = fmt.Sprintf("%v-%v", memory.APP_ENV_AppEnvironment, flagName)
	}

	return u.unleash.IsEnabled(composedFlag, unleash.WithContext(ctx))
}
