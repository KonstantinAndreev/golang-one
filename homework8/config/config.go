package config

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Port        int    `valid:"-"`
	DbUrl       string `valid:"url"`
	JaegerUrl   string `valid:"url"`
	SentryUrl   string `valid:"url"`
	KafkaBroker string `valid:"-"`
	SomeAppId   string `valid:"-"`
	SomeAppKey  string `valid:"-"`
}

func CheckConfig(cfg Config) error {
	if _, err := govalidator.ValidateStruct(&cfg); err != nil {
		return errors.New("Not valid config!")
	}
	return nil
}

func GetConfig() Config {
	var cfg Config
	err := envconfig.Process("myapp", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return cfg
}
