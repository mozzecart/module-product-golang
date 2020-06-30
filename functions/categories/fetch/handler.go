package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	res "github.com/mahbubzulkarnain/golang-serverless-response"
)

func handler(_ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	results, err := fetch()
	if err != nil {
		return res.Errors(err)
	}

	return res.Success(map[string]interface{}{
		"message": "success",
		"data":    results,
	})
}

func main() {
	lambda.Start(handler)
}
