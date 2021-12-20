package config

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"os"
)

const (
	envPrefix = "myapp"
)

type Config struct {
	Port        int    `json:"port" envconfig:"port"`                //MYAPP_PORT
	DbUrl       string `json:"dbUrl" envconfig:"db_url"`             //MYAPP_DB_URL
	JaegerUrl   string `json:"jaegerUrl" envconfig:"jaeger_url"`     //MYAPP_JAEGER_URL
	SentryUrl   string `json:"sentryUrl" envconfig:"sentry_url"`     //MYAPP_SENTRY_URL
	KafkaBroker string `json:"kafkaBroker" envconfig:"kafka_broker"` //MYAPP_KAFKA_BROKER
	SomeAppId   string `json:"someAppId" envconfig:"some_app_id"`    //MYAPP_SOME_APP_ID
	SomeAppKey  string `json:"someAppKey" envconfig:"some_app_key"`  //MYAPP_SOME_APP_KEY
}

func GetJsonOrEnvConfig(jsonFilePath string) (Config, error) {
	if jsonFilePath != "" {
		return GetJsonConfig(jsonFilePath)
	}
	return GetEnvConfig()
}

func GetJsonConfig(jsonFilePath string) (cfg Config, err error) {
	fileReader, err := os.Open(jsonFilePath)
	if err != nil {
		return cfg, errors.Wrap(err, "Cannot open file")
	}
	defer func() {
		if err2 := fileReader.Close(); err == nil && err2 != nil {
			err = errors.Wrap(err, "Cannot close file")
		}
	}()

	if err = json.NewDecoder(fileReader).Decode(&cfg); err != nil {
		return cfg, errors.Wrap(err, "Cannot decode json config")
	}
	return cfg, nil
}

func GetEnvConfig() (cfg Config, err error) {
	if err = envconfig.Process(envPrefix, &cfg); err != nil {
		return cfg, errors.Wrap(err, "Cannot get env config")
	}
	return cfg, nil
}

func PrintConfig(cfg Config) error {
	_, err := fmt.Printf(
		" Port:%d\n DbUrl:%s\n JaegerUrl:%s\n SentryUrl:%s\n KafkaBroker:%s\n SomeAppId:%s\n SomeAppKey:%s\n",
		cfg.Port, cfg.DbUrl, cfg.JaegerUrl, cfg.SentryUrl, cfg.KafkaBroker, cfg.SomeAppId, cfg.SomeAppKey,
	)
	return err
}
