package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	name := request.PathParameters["name"]
	return events.APIGatewayProxyResponse{Body: "{'Message': 'Hello " + name + "' }", StatusCode: 200}, nil

}

func main() {
	lambda.Start(Handler)
}
