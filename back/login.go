package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "back/model"
    "back/repo"
    "back/utils"
)

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    fmt.Printf("login %s %s\n", req.Resource, req.HTTPMethod)
    var user model.User
    err := json.Unmarshal([]byte(req.Body), &user)
    if err != nil {
        panic(err)
    }
    fmt.Printf("user %s\n", user.Login)
    userEntity := repo.GetUser(user.Login)
    if userEntity == nil {
        return utils.ErrorResponse(err, http.StatusNotFound), nil
    }
    fmt.Printf("userEntity %s\n", userEntity.Login)
    body, err := json.Marshal(*userEntity)
    if err != nil {
        panic(err)
    }
    fmt.Printf("body %s\n", body)
    return events.APIGatewayProxyResponse {
        Body: string(body),
        StatusCode: http.StatusOK,
    }, nil
}

func main() {
    lambda.Start(HandleRequest)
}
