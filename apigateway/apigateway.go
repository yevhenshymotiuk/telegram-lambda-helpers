// Package apigateway provides type aliases for AWS APIGateway types
package apigateway

import "github.com/aws/aws-lambda-go/events"

// Response is an alias for APIGatewayProxyResponse
type Response events.APIGatewayProxyResponse

var (
	// Response200 is Response object with status code 200
	Response200 = Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            "Ok",
	}
	// Response404 is Response object with status code 404
	Response404 = Response{StatusCode: 404}
)
