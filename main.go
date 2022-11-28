package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	username  string
	firstName string
	lastName  string
	isAdmin   bool
	password  string
}

func GetUser(ctx context.Context, credentials UserCredentials, dynaClient dynamodbiface.DynamoDBAPI) (
	*dynamodb.GetItemOutput, error) {

	log.Println("[START] GET USER")
	const tableName = "users"

	username := aws.String(credentials.Username)
	fmt.Printf("Username extracted from req body: %v \n", credentials.Username)

	result, err := dynaClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: username,
			},
		},
	})

	fmt.Printf("GET USER RETURNED: %v", result)
	return result, err
}

func Handler(ctx context.Context, credentials UserCredentials) error {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)})

	if err != nil {
		log.Println("Error connected to DynamoDB")
	}

	dynaClient := dynamodb.New(awsSession)
	result, err := GetUser(ctx, credentials, dynaClient)

	// if err != nil {
	// 	log.Fatalf("Error fetching User: %s", err)
	// 	return result
	// }

	user := User{}

	fmt.Printf("\nBEGIN UNMARSHAL with : %v and User Interface: %v", result.Item, user)
	jsonResult := dynamodbattribute.UnmarshalMap(result.Item, &user)

	fmt.Printf("\noutgoing payload: %v", jsonResult)

	return jsonResult

}

func main() {
	log.Println("Start Lambda GO")
	lambda.Start(Handler)
}
