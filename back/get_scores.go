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
    svc := repo.Connect()
    scores := repo.GetScores(svc, leagueId, nil)
    grouped := GroupByUser(scores, leagueId)
    getScoresResponse := model.GetScoresResponse {
        Scores: grouped,
    }
    body, err := json.Marshal(getScoresResponse)
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
