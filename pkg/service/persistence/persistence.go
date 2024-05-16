package persistence

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log/slog"
	"pointsale.example/GoPointSail/pkg/model"
)

var dynamoDbClient *dynamodb.Client

func ProvideSdkConfig(config aws.Config) {
	dynamoDbClient = dynamodb.NewFromConfig(config)
}

func AddUser(user model.User) error {
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		slog.Error("Error mapping user model to dynamodb av", err)
		return err
	}
	_, err = dynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{TableName: aws.String("user"), Item: item})
	if err != nil {
		slog.Error("Error saving user to dynamodb", err)
		return err
	}
	return nil
}

func AllUsers() ([]model.User, error) {
	var users []model.User
	response, err := dynamoDbClient.Scan(context.TODO(), &dynamodb.ScanInput{TableName: aws.String("user")})
	if err != nil {
		slog.Error("Error retrieving users from dynamodb", err)
		return users, err
	}
	err = attributevalue.UnmarshalListOfMaps(response.Items, &users)
	if err != nil {
		slog.Error("Error unmarshalling results from dynamodb", err)
	}
	return users, err
}
