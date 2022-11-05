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
    fmt.Printf("login %s\n", loginRequest.Login);
    userEntity := repo.GetUser(svc, loginRequest.Login)
    if userEntity == nil {
        resp := utils.ErrorResponse(fmt.Errorf("%s", loginRequest.Login), http.StatusNotFound)
        resp.Headers = make(map[string]string)
        resp.Headers["Access-Control-Allow-Origin"] = "*"
        return resp, nil
    }
    if userEntity.Password != utils.Crypt(loginRequest.Password) && !(userEntity.Password == "" && loginRequest.Password == "") {
        resp := utils.ErrorResponse(fmt.Errorf("Invalid password"), http.StatusBadRequest)
        resp.Headers = make(map[string]string)
        resp.Headers["Access-Control-Allow-Origin"] = "*"
        return resp, nil
    }
    loginResponse := model.LoginResponse {
        UserName: userEntity.Name,
        Token: utils.BuildToken(loginRequest.Login),
    }
    body, err := json.Marshal(loginResponse)
    if err != nil {
        panic(err)
    }
    resp := events.APIGatewayProxyResponse {
        Body: string(body),
        StatusCode: http.StatusOK,
    }
    resp.Headers = make(map[string]string)
    resp.Headers["Access-Control-Allow-Origin"] = "*"
    return resp, nil
}

func main() {
    lambda.Start(HandleRequest)
}
