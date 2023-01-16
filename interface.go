package go_unleash

import "github.com/Unleash/unleash-client-go/v3/context"

type IUnleash interface {
	IsEnabled(flagName string) bool
	IsEnabledContext(flagName string, ctx context.Context) bool
}
