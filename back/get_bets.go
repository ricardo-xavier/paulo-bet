package main

import (
    "fmt"
    "sort"
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
    userId := req.PathParameters["userId"]
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
    userScores := repo.GetScores(svc, leagueId, &userId, login)
    leagueScores := repo.GetScores(svc, leagueId, &leagueId, login)
    scores := userScores
    if login != leagueId {
        for _, score := range(leagueScores) {
            scores = append(scores, score)
        }
    }
    GroupByUser(scores, leagueId, login)
    userScores = nil
    for _, score := range(scores) {
        if score.UserId == userId {
            if !score.Visible {
                score.Home = 9;
                score.Visitors = 9;
            }
            userScores = append(userScores, score)
        }
    }
    sort.SliceStable(userScores, func(i, j int) bool {
        return userScores[i].Date < userScores[j].Date
    })
    getBetsResponse := model.GetBetsResponse {
        Bets: userScores,
    }
    body, err := json.Marshal(getBetsResponse)
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
