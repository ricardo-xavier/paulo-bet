package main

import (
    "net/http"
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "back/model"
    "back/repo"
)

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    userId := req.PathParameters["userId"]
    svc := repo.Connect()
    leagues := repo.GetUserLeagues(svc, userId)
    getLeaguesResponse := model.GetLeaguesResponse {
        Leagues: leagues,
    }
    body, err := json.Marshal(getLeaguesResponse)
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
