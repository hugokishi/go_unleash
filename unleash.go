package go_unleash

import (
	"fmt"

	"github.com/Unleash/unleash-client-go/v3"
	"github.com/hugokishi/go_unleash/pkg/domain/memory"
	"github.com/hugokishi/go_unleash/pkg/domain/structs"

	"github.com/Unleash/unleash-client-go/v3/context"
	"github.com/hugokishi/go_unleash/pkg/driver/logs"
	unleashDriver "github.com/hugokishi/go_unleash/pkg/driver/unleash"
	"github.com/sirupsen/logrus"

	"github.com/go-playground/validator/v10"
)

type Unleash struct {
}

var (
	validate *validator.Validate
)

func InitUnleash(config structs.UnleashConfig) *Unleash {
	validate = validator.New()

	if err := validate.Struct(&config); err != nil {
		logrus.Error(err)
		return nil
	}

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
		return nil
	}

	logrus.Infof("The unleash configuration was successful, using the %v environment", memory.APP_ENV_AppEnvironment)
	return &Unleash{}
}

func (u *Unleash) IsEnabled(flagName string) bool {
	var composedFlag string

	if memory.APP_ENV_AppEnvironment != "production" {
		composedFlag = fmt.Sprintf("%v-%v", memory.APP_ENV_AppEnvironment, flagName)
	}

	isEnabled := unleash.IsEnabled(composedFlag)
	logrus.Infof("Checked %v, Enabled: %v", composedFlag, isEnabled)

	return isEnabled
}

func (u *Unleash) IsEnabledContext(flagName string, ctx context.Context) bool {
	var composedFlag string

	if memory.APP_ENV_AppEnvironment != "production" {
		composedFlag = fmt.Sprintf("%v-%v", memory.APP_ENV_AppEnvironment, flagName)
	}

	isEnabled := unleash.IsEnabled(composedFlag, unleash.WithContext(ctx))
	logrus.Infof("Checked %v, Enabled: %v", composedFlag, isEnabled)

	return isEnabled
}
