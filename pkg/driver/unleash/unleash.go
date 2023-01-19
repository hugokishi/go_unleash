package unleash

import (
	"net/http"

	"github.com/Unleash/unleash-client-go/v3"
	"github.com/hugokishi/go_unleash/pkg/domain/structs"
)

func InitializeUnleash(config structs.UnleashConfig) error {
	return unleash.Initialize(
		unleash.WithListener(
			&unleash.DebugListener{},
		),
		unleash.WithAppName(
			config.AppName,
		),
		unleash.WithUrl(
			config.UnleashURL,
		),
		unleash.WithCustomHeaders(
			http.Header{
				"Authorization": {config.UnleashAuthorizationToken},
			},
		),
	)
}
