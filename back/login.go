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
    var user model.User
    err := json.Unmarshal([]byte(req.Body), &user)
    if err != nil {
        panic(err)
    }
    userEntity := repo.GetUser(user.Login)
    if userEntity == nil {
        return utils.ErrorResponse(fmt.Errorf("%s", user.Login), http.StatusNotFound), nil
    }
    body, err := json.Marshal(*userEntity)
    if err != nil {
        panic(err)
    }
    return events.APIGatewayProxyResponse {
        Body: string(body),
        StatusCode: http.StatusOK,
    }, nil
}

func main() {
    lambda.Start(HandleRequest)
}
