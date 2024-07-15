package main

import (
	"gateway/api"
	cf "gateway/config"
	"gateway/grpc/clients"
	"log"
)

func main() {
	config := cf.Load()

	services, err := clients.NewGrpcClients(&config)
	if err != nil {
		log.Fatalf("error while connecting clients. err: %s", err.Error())
	}

	engine := api.NewRouter(services)

	err = engine.Run(config.HTTP_PORT)
	if err != nil {
		log.Fatalf("error while running server. err: %s", err.Error())
	}
}
