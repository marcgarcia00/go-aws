package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// struct User {
// 	Id int64
// }

func main() {
	log.Println("Start Lambda GO")
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Println("Begin User Retrieval")
	log.Println("Incoming Payload: ")
	fmt.Printf("%v", request)

	response := events.APIGatewayProxyResponse{
		Body:       "Update: 12:43p",
		StatusCode: 200,
	}
	log.Println("End of Function")
	return response, nil
}
