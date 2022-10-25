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
    var changePasswordRequest model.ChangePasswordRequest
    err := json.Unmarshal([]byte(req.Body), &changePasswordRequest)
    if err != nil {
        panic(err)
    }
    ok := utils.CheckToken(changePasswordRequest.Login, changePasswordRequest.Token)
    if !ok {
        resp := utils.ErrorResponse(fmt.Errorf("Invalid token"), http.StatusUnauthorized)
        resp.Headers = make(map[string]string)
        resp.Headers["Access-Control-Allow-Origin"] = "*"
        return resp, nil
    }
    svc := repo.Connect()
    repo.ChangePassword(svc, changePasswordRequest.Login, utils.Crypt(changePasswordRequest.Password))
    resp := events.APIGatewayProxyResponse {
        StatusCode: http.StatusNoContent,
    }
    resp.Headers = make(map[string]string)
    resp.Headers["Access-Control-Allow-Origin"] = "*"
    return resp, nil
}

func main() {
    lambda.Start(HandleRequest)
}
