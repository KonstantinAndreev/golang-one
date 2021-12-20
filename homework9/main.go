package main

import (
	"flag"
	"golangOne/golang-one/homework9/config"
	"log"
)

var (
	jsonConfigPath = flag.String("json-file-path", "", "Путь до файла, где хранится json конфиг")
)

func main() {

	flag.Parse()
	cfg, err := config.GetJsonOrEnvConfig(*jsonConfigPath)
	if err != nil {
		log.Fatalf("Cannot get config: %v", err)
	}

	if err = config.PrintConfig(cfg); err != nil {
		log.Fatalf("Cannot print config: %v", err)
	}
}
