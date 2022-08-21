package main

import (
    "fmt"
    "os"
    "back/repo"
)

func main() {
    leagueId := os.Args[1]
    userId := os.Args[2]
    svc := repo.Connect()
    userScores := repo.GetScores(svc, leagueId, &userId)
    leagueScores := repo.GetScores(svc, leagueId, &leagueId)
    scores := userScores
    for _, score := range(leagueScores) {
        scores = append(scores, score)
    }
    GroupByUser(scores, leagueId, userId)
    userScores = nil
    for _, score := range(scores) {
        if score.UserId == userId {
            userScores = append(userScores, score)
        }
    }
    fmt.Println(userScores)
}
