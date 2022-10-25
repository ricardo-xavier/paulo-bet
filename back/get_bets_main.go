package main

import (
    "fmt"
    "os"
    "back/repo"
)

func main() {
    leagueId := os.Args[1]
    userId := os.Args[2]
    login := os.Args[3]
    svc := repo.Connect()
    userScores := repo.GetScores(svc, leagueId, &userId, login)
    leagueScores := repo.GetScores(svc, leagueId, &leagueId, login)
    scores := userScores
    for _, score := range(leagueScores) {
        scores = append(scores, score)
    }
    fmt.Println(len(userScores))
    fmt.Println(userScores)
    GroupByUser(scores, leagueId, userId)
    userScores = nil
    for _, score := range(scores) {
        if score.UserId == userId {
            userScores = append(userScores, score)
        }
    }
    fmt.Println(len(userScores))
    fmt.Println(userScores)
}
