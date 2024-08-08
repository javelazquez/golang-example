package main

import (
	"context"
	"fmt"
	"golang-example/internal/core/repositories"
	"golang-example/internal/core/services"
	"golang-example/internal/entrypoint/http"
	"golang-example/internal/entrypoint/http/handler"
	"golang-example/pkg/apm"
	"golang-example/pkg/credential"
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

	configCredential := credential.ConfigCredential{
		AWSRegion:    "ARG",
		AWSAccessKey: "TEST",
		AWSSecretKey: "TEST-KEY",
	}

	//pkg
	nr, err := apm.NewNewRelic(ctx)
	if err != nil {
		fmt.Println("se rompio newrelic", err)
	}

	credentialAWS, err := credential.NewCredential(ctx, configCredential)
	if err != nil {
		fmt.Println("se rompio credential", err)
	}

	configKVS := kvs.Config{
		TableName:  "payout_table",
		Credential: credentialAWS,
	}

	storageKVS := kvs.NewDynamoKVS(configKVS)
	webServer := server.NewWebServer(ctx, configServer, nr)

	//Repositories
	payoutRepository := repositories.NewPayoutRepository(storageKVS)

	//Services
	payoutService := services.NewPayoutService(payoutRepository)

	//Router
	payoutHandler := handler.NewPayoutHTTPHandlers(payoutService)
	http.NewRouter(payoutHandler).RouterURLs(webServer)

	//RUN
	webServer.Run(ctx)
}
