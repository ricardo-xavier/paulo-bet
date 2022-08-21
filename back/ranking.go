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
    leagueId := req.PathParameters["leagueId"]
    login := req.QueryStringParameters["login"]
    token := req.QueryStringParameters["token"]
    ok := utils.CheckToken(login, token)
    if !ok {
        return utils.ErrorResponse(fmt.Errorf("Invalid token"), http.StatusUnauthorized), nil
    }
    svc := repo.Connect()
    scores := repo.GetScores(svc, leagueId, nil)
    scores = repo.Initialize(svc, leagueId, login, scores)
    grouped := GroupByUser(scores, leagueId, login)
    var ranking []model.Ranking
    for _, score := range grouped {
        r := model.Ranking {
            UserId: score.UserId,
            Score: score.Score,
        }
        ranking = append(ranking, r)
    }
    rankingResponse := model.RankingResponse {
        Ranking: ranking,
    }
    body, err := json.Marshal(rankingResponse)
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
