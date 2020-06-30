package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	res "github.com/mahbubzulkarnain/golang-serverless-response"
)

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	productID := req.PathParameters["productID"]

	result, err := deleteByID(productID)
	if err != nil {
		return res.Errors(err)
	}

	return res.Success(map[string]interface{}{
		"message": "success",
		"data":    result,
	})
}

func main() {
	lambda.Start(handler)
}
