package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	POSTGRES_USER   string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASS   string `mapstructure:"POSTGRES_PASS"`
	POSTGRES_DBNAME string `mapstructure:"POSTGRES_DBNAME"`
	POSTGRES_HOST   string `mapstructure:"POSTGRES_HOST"`
	POSTGRES_PORT   int    `mapstructure:"POSTGRES_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	appEnvironmentVersion := ApplicationEnv(os.Getenv(AppEnv))

	filename, err := getEnvironmentFilename(appEnvironmentVersion)
	if err != nil {
		return
	}
	err = loadFromEnvFile(path, filename, &config)
	if err != nil {
		return
	}
	err = setEnvironmentValuesOnFly(&config)
	if err != nil {
		return
	}
	setEmptyEnvironmentValuesByDefault(&config)
	return
}

func setEnvironmentValuesOnFly(config *Config) error {
	configType := reflect.TypeOf(*config)
	currentConfig := reflect.ValueOf(config).Elem()
	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		currentField := currentConfig.Field(i)
		foundEnvironment, err := getEnvironmentValueByField(&field, "mapstructure")
		if err != nil {
			return err
		}
		if foundEnvironment == "" {
			continue
		}
		setEnvironmentToField(&currentField, foundEnvironment)
	}
	return nil
}

func setEmptyEnvironmentValuesByDefault(config *Config) {
	if config.POSTGRES_DBNAME == "" {
		config.POSTGRES_DBNAME = "postgres"
	}
	if config.POSTGRES_PASS == "" {
		config.POSTGRES_PASS = "postgres"
	}
	if config.POSTGRES_USER == "" {
		config.POSTGRES_USER = "postgres"
	}
	if config.POSTGRES_PORT == 0 {
		config.POSTGRES_PORT = 5432
	}
	if config.POSTGRES_HOST == "" {
		config.POSTGRES_HOST = "0.0.0.0"
	}
}

func getEnvironmentFilename(appEnvironmentVersion ApplicationEnv) (configFilename string, resultError error) {
	switch appEnvironmentVersion {
	case ProductionEnv:
		configFilename = "production"
	case DevelopEnv:
		configFilename = "develop"
	case TestEnv:
		configFilename = "test"
	default:
		errorText := fmt.Sprintf("bad environment var: %s", AppEnv)
		resultError = errors.New(errorText)
	}
	return
}

func getEnvironmentValueByField(field *reflect.StructField, tag string) (string, error) {
	environmentName, ok := field.Tag.Lookup(tag)
	if !ok {
		return "", errors.New("tag is not found")
	}
	return os.Getenv(environmentName), nil
}

func setEnvironmentToField(field *reflect.Value, newValue string) (retErr error) {
	if !field.IsValid() || !field.CanSet() {
		return errors.New("config field is not reachable")
	}
	switch kind := field.Type().Kind(); kind {
	case reflect.String:
		field.SetString(newValue)
		retErr = nil
	case reflect.Int:
		stringedInteger, err := strconv.Atoi(newValue)
		retErr = err
		if err != nil {
			return
		}
		field.SetInt(int64(stringedInteger))
	default:
		retErr = errors.New("unexpected type")
	}
	return
}

func loadFromEnvFile(path string, filename string, config *Config) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(filename)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return viper.Unmarshal(&config)
}
