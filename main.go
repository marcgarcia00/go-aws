package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.Println("Start Lambda GO")
	lambda.Start(handler)
}

func handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Begin User Retrieval")
	log.Println("Incoming req header: ")
	fmt.Printf("%v", event.Headers)
	fmt.Printf("%v", event.Body)
	response := events.APIGatewayProxyResponse{
		Body:       event.Body,
		StatusCode: 200,
	}

	return response, nil
}
