package service

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"pointsale.example/GoPointSail/pkg/model"
)

var dynamoDbClient *dynamodb.Client

func ProvideSdkConfig(config aws.Config) {
	dynamoDbClient = dynamodb.NewFromConfig(config)
}

func AddUser(user model.User) {
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		panic(err)
	}
	_, err = dynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{TableName: aws.String("user"), Item: item})
	if err != nil {
		log.Printf("Error saving user to dynamodb, %v\n", err)
	}
}

func AllUsers() []model.User {
	var users []model.User
	response, err := dynamoDbClient.Scan(context.TODO(), &dynamodb.ScanInput{TableName: aws.String("user")})
	if err != nil {
		log.Printf("Error retrieving users from dynamodb, %v\n", err)
		return users
	}
	err = attributevalue.UnmarshalListOfMaps(response.Items, &users)
	if err != nil {
		log.Printf("Error unmarshalling results from dynamodb, %v\n", err)
	}
	return users
}

//func RemoveUser(id int) {
//	foundIndex := -1
//	for i, user := range users {
//		if user.Id == id {
//			foundIndex = i
//			break
//		}
//	}
//	if foundIndex != -1 {
//		users = append(users[:foundIndex], users[foundIndex+1:]...)
//	}
//}
