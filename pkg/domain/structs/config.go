package structs

type UnleashConfig struct {
	// Refers to the environment that unleash is being configured for (local, development, staging, production)
	AppEnvironment string `validate:"required,oneof='local' 'development' 'staging' 'production'"`
	// Refers to the name that will be registered in unleash
	AppName string `validate:"required"`
	// Refers to the url of connection with unleash
	UnleashURL string `validate:"required"`
	// Refers to the authorization token for connecting to unleash
	UnleashAuthorizationToken string `validate:"required"`
	// Refers to the logging level used in library
	LoggingLevel string ``
}
