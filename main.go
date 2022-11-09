package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.Println("Start Lambda GO")
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Begin User Retrieval")

	response := events.APIGatewayProxyResponse{
		Body:       "Woo hoo!",
		StatusCode: 200,
	}

	return response, nil
}
