        package main

import (
    "fmt"
    "strings"
    "time"
    "net/http"
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "back/model"
    "back/repo"
    "back/utils"
    "github.com/aws/aws-sdk-go/service/dynamodb"
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
    if userId != leaguePrefix {
        checkDate(svc, leaguePrefix, betRequest.MatchId)
    }
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

func checkDate(svc *dynamodb.DynamoDB, leaguePrefix string, matchId string) {
    adminDate := repo.GetDate(svc, leaguePrefix, matchId)
    if adminDate == nil {
        panic("admin date: " + leaguePrefix + "_" + matchId)
    }
    loc, _ := time.LoadLocation("America/Sao_Paulo")
    currentTime := time.Now().In(loc)
    date := currentTime.Format("2006-01-02 15:04:05")
    if date >= *adminDate {
        panic("admin date denied: " + leaguePrefix + "_" + matchId)
    }
}
