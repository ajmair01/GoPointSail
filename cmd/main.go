package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"log/slog"
	"pointsale.example/GoPointSail/pkg/service/persistence"
	"pointsale.example/GoPointSail/pkg/service/router"
)

func main() {
	awsSdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		slog.Error("Unable to load AWS SDK config", err)
	}
	persistence.ProvideSdkConfig(awsSdkConfig)

	router.StartRouter()
}
