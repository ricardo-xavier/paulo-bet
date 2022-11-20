package main

import (
    "fmt"
    "strings"
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
    userId := strings.ToLower(req.PathParameters["userId"])
    leaguePrefix := strings.Split(leagueId, "-")[0]
    if userId == leaguePrefix {
        leagueId = leaguePrefix
    }
    var betRequest model.BetRequest
    err := json.Unmarshal([]byte(req.Body), &betRequest)
    if err != nil {
        panic(err)
    }
    ok := utils.CheckToken(userId, betRequest.Token)
    if !ok {
        resp := utils.ErrorResponse(fmt.Errorf("Invalid token"), http.StatusUnauthorized)
        resp.Headers = make(map[string]string)
        resp.Headers["Access-Control-Allow-Origin"] = "*"
        return resp, nil
    }
    fmt.Printf("DEBUG bet %v %v %v %v %v\n", leagueId, userId, betRequest.MatchId, betRequest.Home, betRequest.Visitors);
    svc := repo.Connect()
    repo.UpdateBet(svc, leagueId, userId, betRequest.MatchId, betRequest.Home, betRequest.Visitors)
    resp := events.APIGatewayProxyResponse {
        StatusCode: http.StatusNoContent,
    }
    resp.Headers = make(map[string]string)
    resp.Headers["Access-Control-Allow-Origin"] = "*"
    return resp, nil
}

func main() {
    lambda.Start(HandleRequest)
}
