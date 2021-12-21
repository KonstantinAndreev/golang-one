package main

import (
	"fmt"
	"golangOne/golang-one/homework8/config"
	"log"
)

func main() {
	cfg := config.GetConfig()
	var err error

	if err = config.CheckConfig(cfg); err != nil { // проверяется валидность конфига
		log.Fatalf("Config is invalid: %v", err)
	}
	format := "Port:%v\n DB_URL:%s\n Jaeger_URL:%s\n Sentry_URL:%s\n Kafka_Broker:%s\n Some_App_Id:%s\n Some_App_Key:%s\n"
	_, err = fmt.Printf(format, cfg.Port, cfg.DbUrl, cfg.JaegerUrl, cfg.SentryUrl, cfg.KafkaBroker, cfg.SomeAppId, cfg.SomeAppKey)

	if err != nil {
		log.Fatal(err.Error())
	}
}
