package persistence

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log/slog"
	"pointsale.example/GoPointSail/pkg/model"
)

var dynamoDbClient *dynamodb.Client
var userTable = aws.String("user")
var groupTable = aws.String("group")

func ProvideSdkConfig(config aws.Config) {
	dynamoDbClient = dynamodb.NewFromConfig(config)
}

func AddUser(user model.User) error {
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		slog.Error("Error mapping user model to dynamodb av", err)
		return err
	}
	_, err = dynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{TableName: userTable, Item: item})
	if err != nil {
		slog.Error("Error saving user to dynamodb", err)
		return err
	}
	return nil
}

func AllUsers() ([]model.User, error) {
	var users []model.User
	response, err := dynamoDbClient.Scan(context.TODO(), &dynamodb.ScanInput{TableName: userTable})
	if err != nil {
		slog.Error("Error retrieving users from dynamodb", err)
		return users, err
	}
	err = attributevalue.UnmarshalListOfMaps(response.Items, &users)
	if err != nil {
		slog.Error("Error unmarshalling results from dynamodb for AllUsers", err)
	}
	return users, err
}

func GetUserByEmail(email string) (model.User, error) {
	var user model.User
	response, err := dynamoDbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{Key: map[string]types.AttributeValue{
		"email": &types.AttributeValueMemberS{Value: email},
	}, TableName: userTable})
	if err != nil {
		slog.Error("Error finding user by email", err)
		return user, err
	}
	err = attributevalue.UnmarshalMap(response.Item, &user)
	if err != nil {
		slog.Error("Error unmarshalling results from dynamodb for GetUserByEmail", err)
	}
	return user, err
}
