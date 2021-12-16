package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"golangOne/golang-one/homework8/config"
	"log"
)

func main() {
	s := config.GetConfig()
	err := envconfig.Process("myapp", &s)
	format := "Port:%v\n DB_URL:%s\n Jaeger_URL:%s\n Sentry_URL:%s\n Kafka_Broker:%s\n Some_App_Id:%s\n Some_App_Key:%s\n"
	_, err = fmt.Printf(format, s.Port, s.DbUrl, s.JaegerUrl, s.SentryUrl, s.KafkaBroker, s.SomeAppId, s.SomeAppKey)
	if err != nil {
		log.Fatal(err.Error())
	}
}
