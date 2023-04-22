package config

type ApplicationEnv string

const (
	AppEnv = "APP_ENV"
)

const (
	ProductionEnv ApplicationEnv = "production"
	DevelopEnv    ApplicationEnv = "develop"
	TestEnv       ApplicationEnv = "test"
)
