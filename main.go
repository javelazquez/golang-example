package main

import (
	"context"
	"fmt"
	"golang-example/internal/core/adapters"
	"golang-example/internal/core/services"
	"golang-example/internal/entrypoint/http"
	"golang-example/internal/entrypoint/http/handler"
	"golang-example/pkg/apm"
	"golang-example/pkg/credentials"
	"golang-example/pkg/kvs"
	"golang-example/pkg/server"
	"time"
)

func main() {
	ctx := context.Background()

	//configs
	configServer := server.ConfigServer{
		Port:            "8080",
		IdleTimeout:     5 * time.Minute,
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    10 * time.Second,
		ShutdownTimeout: 10 * time.Second,
	}

	configCredential := credentials.ConfigCredential{
		AWSRegion:    "eu-west-1",
		AWSAccessKey: "user-test",
		AWSSecretKey: "pass-test",
		RoleARN:      "",
	}

	//pkg
	nr, err := apm.NewNewRelic(ctx)
	if err != nil {
		fmt.Println("se rompio newrelic", err)
	}

	credentialAWS, err := credentials.NewCredential(ctx, configCredential)
	if err != nil {
		fmt.Println("se rompio credentials", err)
	}

	configKVS := kvs.Config{
		TableName:   "payout_table",
		Credential:  credentialAWS,
		AWSEndpoint: "http://localhost:4566",
	}

	storageKVS := kvs.NewDynamoKVS(configKVS)
	webServer := server.NewWebServer(ctx, configServer, nr)

	//Repositories
	payoutRepository := adapters.NewPayoutRepository(storageKVS)

	//Services
	payoutService := services.NewPayoutService(payoutRepository)

	//Router
	payoutHandler := handler.NewPayoutHTTPHandlers(payoutService)
	http.NewRouter(payoutHandler).RouterURLs(webServer)

	//RUN
	webServer.Run(ctx)
}
