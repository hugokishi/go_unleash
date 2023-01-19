package memory

var (
	APP_ENV_AppEnvironment string
	APP_ENV_LoggingLevel   string
)

func SetVariables(
	AppEnvironment, LoggingLevel string,
) {
	APP_ENV_AppEnvironment = AppEnvironment
	APP_ENV_LoggingLevel = LoggingLevel
}
