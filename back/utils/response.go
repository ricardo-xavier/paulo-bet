package utils

import (
    "github.com/aws/aws-lambda-go/events"
)

func ErrorResponse(err error, status int) events.APIGatewayProxyResponse {
    return events.APIGatewayProxyResponse {
        Body: err.Error(),
        StatusCode: status,
    }
}
