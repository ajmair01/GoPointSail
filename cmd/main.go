package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"log"
	"pointsale.example/GoPointSail/pkg/model"
	"pointsale.example/GoPointSail/pkg/service"
	"time"
)

func main() {

	awsSdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Unable to load AWS SDK config, %v", err)
	}
	service.ProvideSdkConfig(awsSdkConfig)

	service.AddUser(model.User{
		FirstName:      "Some",
		LastName:       "Body",
		MiddleInitial:  "J",
		UserName:       "someJbody",
		Email:          "somebody@gmail.com",
		Phone:          "5558675309",
		Password:       "817y6sdfkljbasvbg78435",
		CreatedAt:      time.Now(),
		LastUpdated:    time.Now(),
		SecurityGroups: nil,
		OtpCode:        "",
		OtpExpiration:  time.Now(),
		OptIn:          model.OptIn{AllowEmail: false, AllowSMS: true},
	})

	fmt.Println(service.AllUsers())
}
