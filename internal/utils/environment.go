package utils

import (
	"github.com/hugokishi/go-unleash/internal/domain/memory"
)

func GetPrefix() string {
	switch memory.APP_ENV_AppEnvironment {
	case "development":
		return "development"
	case "staging":
		return "staging"
	}
	return ""
}