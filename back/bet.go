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
    fmt.Println("=====PauloBet:bet")
    leagueId := req.PathParameters["leagueId"]
    userId := req.PathParameters["userId"]
    fmt.Println(leagueId)
    fmt.Println(userId)
    var betRequest model.BetRequest
    err := json.Unmarshal([]byte(req.Body), &betRequest)
    if err != nil {
        panic(err)
    }
    fmt.Println(betRequest)
    ok := utils.CheckToken(userId, betRequest.Token)
    if !ok {
        return utils.ErrorResponse(fmt.Errorf("Invalid token"), http.StatusUnauthorized), nil
    }
    svc := repo.Connect()
    repo.UpdateBet(svc, leagueId, userId, betRequest.MatchId, betRequest.Home, betRequest.Visitors)
    return events.APIGatewayProxyResponse {
        StatusCode: http.StatusNoContent,
    }, nil
}

func main() {
    lambda.Start(HandleRequest)
}
