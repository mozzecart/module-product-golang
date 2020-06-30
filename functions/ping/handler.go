package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	res "github.com/mahbubzulkarnain/golang-serverless-response"
)

func handler(_ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return res.Success(map[string]interface{}{
		"message": "pong!!",
	})
}

func main() {
	lambda.Start(handler)
}
