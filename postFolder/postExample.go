package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
// Hit api with json/application type
// { "name" : "insert name here"}

type Response events.APIGatewayProxyResponse

type BodyRequest struct {
	RequestName string `json:"name"`
}

type BodyResponse struct {
	ResponseName string `json:"name"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	bodyRequest := BodyRequest{
		RequestName: "",
	}
	err := json.Unmarshal([]byte(request.Body), &bodyRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 401}, nil
	}

	bodyResponse := BodyResponse{
		ResponseName: bodyRequest.RequestName + " LastName",
	}

	response, err := json.Marshal(&bodyResponse)

	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil

}

func main() {

	lambda.Start(Handler)
}
