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
    var loginRequest model.LoginRequest
    err := json.Unmarshal([]byte(req.Body), &loginRequest)
    if err != nil {
        panic(err)
    }
    svc := repo.Connect()
    fmt.Printf("PAULOBET:login [%s]\n", loginRequest.Login)
    userEntity := repo.GetUser(svc, loginRequest.Login)
    if userEntity == nil {
        return utils.ErrorResponse(fmt.Errorf("%s", loginRequest.Login), http.StatusNotFound), nil
    }
    if userEntity.Password != utils.Crypt(loginRequest.Password) && !(userEntity.Password == "" && loginRequest.Password == "") {
        fmt.Printf("PAULOBET:login [%s] Invalid password\n", loginRequest.Login)
        return utils.ErrorResponse(fmt.Errorf("Invalid password"), http.StatusBadRequest), nil
    }
    loginResponse := model.LoginResponse {
        UserName: userEntity.Name,
        Token: utils.BuildToken(loginRequest.Login),
    }
    body, err := json.Marshal(loginResponse)
    if err != nil {
        panic(err)
    }
    fmt.Printf("PAULOBET:login [%s] [%s]\n", loginRequest.Login, loginResponse.Token)
    return events.APIGatewayProxyResponse {
        Body: string(body),
        StatusCode: http.StatusOK,
    }, nil
}

func main() {
    lambda.Start(HandleRequest)
}
