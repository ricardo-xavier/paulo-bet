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
        resp := utils.ErrorResponse(fmt.Errorf("Invalid token"), http.StatusUnauthorized)
        resp.Headers = make(map[string]string)
        resp.Headers["Access-Control-Allow-Origin"] = "*"
        return resp, nil
    }
    svc := repo.Connect()
    usersScores := repo.GetScores(svc, leagueId, nil, login)
    leagueScores := repo.GetScores(svc, leagueId, &leagueId, login)
    scores := usersScores
    for _, score := range(leagueScores) {
        scores = append(scores, score)
    }
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
