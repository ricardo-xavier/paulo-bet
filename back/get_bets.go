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
    leagueId := req.PathParameters["leagueId"]
    userId := req.PathParameters["userId"]
    svc := repo.Connect()
    userScores := repo.GetScores(svc, leagueId, &userId)
    leagueScores := repo.GetScores(svc, leagueId, &leagueId)
    scores := userScores
    for _, score := range(leagueScores) {
        scores = append(scores, score)
    }
    GroupByUser(scores, leagueId)
    userScores = nil
    for _, score := range(scores) {
        if score.UserId == userId {
            userScores = append(userScores, score)
        }
    }
    getBetsResponse := model.GetBetsResponse {
        Bets: userScores,
    }
    body, err := json.Marshal(getBetsResponse)
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
