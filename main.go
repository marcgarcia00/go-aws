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
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id        int
	Username  string
	FirstName string
	LastName  string
	IsAdmin   bool
}

// func HandleRequest(ctx context.Context, credentials UserCredentials) (events.APIGatewayProxyResponse, error) {
// 	log.Println("Begin User Retrieval")
// 	log.Println("Incoming req header: ")
// 	fmt.Printf("%v", credentials.Username)
// 	fmt.Printf("%v", credentials.Password)

// 	response := events.APIGatewayProxyResponse{
// 		Headers:    map[string]string{"Content-Type": "application/json"},
// 		StatusCode: 200,
// 	}
// 	responseBody, _ := json.Marshal(credentials)
// 	response.Body = string(responseBody)
// 	log.Println("End of Function")
// return &response, nil
// }

func GetUser(ctx context.Context, credentials UserCredentials, dynaClient dynamodbiface.DynamoDBAPI) (
	*dynamodb.GetItemOutput, error) {
	log.Println("[START] GET USER")
	fmt.Printf("CONTEXT: %v", ctx)
	const tableName = "Users"

	// username := aws.String(credentials.Username)
	result, err := dynaClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String("1"),
			},
		},
	})
	// "username": {
	// 	S: aws.String("marcgarcia"),
	// },
	// "firstName": {
	// 	S: aws.String("MARC"),
	// },
	// "lastName": {
	// 	S: aws.String("GARCIA"),
	// },
	// "isAdmin": {
	// 	BOOL: aws.Bool(true),
	// },
	if err != nil {
		log.Fatalf("Error fetching User: %s", err)
		return result, err
	}
	return result, err
}

func Handler(ctx context.Context, credentials UserCredentials) (*dynamodb.GetItemOutput, error) {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)})

	if err != nil {
		log.Println("error returned")
	}

	dynaClient := dynamodb.New(awsSession)
	response, err := GetUser(ctx, credentials, dynaClient)

	return response, err
}

func main() {
	log.Println("Start Lambda GO")
	lambda.Start(Handler)
	// lambda.Start(HandleRequest)
}
