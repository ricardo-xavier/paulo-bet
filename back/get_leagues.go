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
    userId := req.PathParameters["userId"]
    token := req.QueryStringParameters["token"]
    ok := utils.CheckToken(userId, token)
    if !ok {
        resp := utils.ErrorResponse(fmt.Errorf("Invalid token"), http.StatusUnauthorized)
        resp.Headers = make(map[string]string)
        resp.Headers["Access-Control-Allow-Origin"] = "*"
        return resp, nil
    }
    svc := repo.Connect()
    leagues := repo.GetUserLeagues(svc, userId)
    getLeaguesResponse := model.GetLeaguesResponse {
        Leagues: leagues,
    }
    body, err := json.Marshal(getLeaguesResponse)
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
