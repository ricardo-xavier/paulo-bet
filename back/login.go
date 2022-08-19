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
fmt.Println("LOGIN")
fmt.Println(loginRequest.Id+".")
fmt.Println(loginRequest.Password+".")
    svc := repo.Connect()
    userEntity := repo.GetUser(svc, loginRequest.Id)
fmt.Println(userEntity.Password+".")
    if userEntity == nil {
        return utils.ErrorResponse(fmt.Errorf("%s", loginRequest.Id), http.StatusNotFound), nil
    }
    if userEntity.Password != utils.Crypt(loginRequest.Password) && !(userEntity.Password == "" && loginRequest.Password == "") {
fmt.Println("LOGIN INVALID PASSWORD")
        return utils.ErrorResponse(fmt.Errorf("Invalid password"), http.StatusUnauthorized), nil
    }
    loginResponse := model.LoginResponse {
        UserName: userEntity.Name,
        Token: utils.BuildToken(loginRequest.Id),
    }
    body, err := json.Marshal(loginResponse)
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
