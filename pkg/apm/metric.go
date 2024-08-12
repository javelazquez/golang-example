package apm

import (
	"context"
	"fmt"
	"github.com/newrelic/go-agent/v3/newrelic"
	"os"
	"time"
)

type NewRelic interface {
	GetApplication() *newrelic.Application
}

type newRelic struct {
	application *newrelic.Application
}

func NewNewRelic(ctx context.Context) (NewRelic, error) {
	var nrApp *newrelic.Application
	var err error
	isLocal := os.Getenv("ENV") == "local"

	if !isLocal {
		nrApp, err = newrelic.NewApplication(
			newrelic.ConfigAppName("My Go App"),
			newrelic.ConfigLicense("your_new_relic_license_key"),
			newrelic.ConfigDistributedTracerEnabled(true),
		)
		if err != nil {
			fmt.Println("New Relic initialization failed:", err)
			return newRelic{}, err
		}
		err = nrApp.WaitForConnection(10 * time.Second)
		if err != nil {
			fmt.Println("New Relic connection failed:", err)
			return nil, err
		}
	}

	return &newRelic{application: nrApp}, nil
}

func (n newRelic) GetApplication() *newrelic.Application {
	return n.application
}
