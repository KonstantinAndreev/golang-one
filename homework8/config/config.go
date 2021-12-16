package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Configuration struct {
	Port        int    `valid:"Port"`
	DbUrl       string `valid:"DataBaseURL"`
	JaegerUrl   string `valid:"JaegerURL"`
	SentryUrl   string `valid:"SentryURL"`
	KafkaBroker string `valid:"KafkaBroker"`
	SomeAppId   string `valid:"SomeAppId"`
	SomeAppKey  string `valid:"SomeAppKey"`
}

func GetConfig() Configuration {
	var s Configuration
	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	return s
}
