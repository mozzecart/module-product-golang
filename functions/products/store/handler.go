package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	res "github.com/mahbubzulkarnain/golang-serverless-response"

	"github.com/mozze/module-product-golang/models"
)

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := &models.Product{}

	errParser := json.Unmarshal([]byte(req.Body), body)
	if errParser != nil {
		return res.Errors(errParser)
	}

	results, err := store(body)
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
