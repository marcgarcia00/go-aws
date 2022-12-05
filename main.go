package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
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
	Username  string `json:username`
	FirstName string `json:password`
	LastName  string `json:lastname`
	IsAdmin   bool   `json:isAdmin`
	Password  string `json:password`
}

type Response struct {
	Message *string
	User    *User
}

func GetUserbyKey(credentials UserCredentials, body string, dynaClient dynamodbiface.DynamoDBAPI) (
	*dynamodb.GetItemOutput, error) {
	log.Println("[START] GET USER")
	const tableName = "users"
	username := aws.String(credentials.Username)

	result, err := dynaClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: username,
			},
		},
	})

	if err != nil {
		log.Println("Error fetching user")
		return result, err
	}
	fmt.Printf("GET USER RETURNED: %v", result)
	return result, err
}

func Handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userMap := User{}

	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)})

	if err != nil {
		log.Println("Error connected to DynamoDB")
		return events.APIGatewayProxyResponse{}, err
	}

	dynaClient := dynamodb.New(awsSession)

	var credentialsJson UserCredentials
	credentialsBytes := []byte(event.Body)

	json.Unmarshal(credentialsBytes, &credentialsJson)
	fmt.Printf("\n unmarshalled event body: %v", credentialsJson)

	user, err := GetUserbyKey(credentialsJson, event.Body, dynaClient)

	if err != nil {
		log.Fatalf("Error fetching User: %s", err)
		return events.APIGatewayProxyResponse{}, err
	}

	jsonFromAWSErr := dynamodbattribute.UnmarshalMap(user.Item, &userMap)

	if jsonFromAWSErr != nil {
		fmt.Print("error with json unmarshaling")
		return events.APIGatewayProxyResponse{}, jsonFromAWSErr
	}
	responseBody := Response{
		Message: aws.String("Success!"),
		User:    &userMap,
	}

	log.Println("Successfully completed get method")
	responseBytes, err := json.Marshal(responseBody)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		Body:       string(responseBytes),
	}
	return response, nil
}

func main() {
	log.Println("Start Lambda GO")
	lambda.Start(Handler)
}
