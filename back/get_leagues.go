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
        return utils.ErrorResponse(fmt.Errorf("Invalid token"), http.StatusUnauthorized), nil
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
    return events.APIGatewayProxyResponse {
        Body: string(body),
        StatusCode: http.StatusOK,
    }, nil
}

func main() {
    lambda.Start(HandleRequest)
}
