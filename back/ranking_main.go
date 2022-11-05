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
    usersScores := repo.GetScores(svc, leagueId, nil, userId)
    leagueScores := repo.GetScores(svc, leagueId, &leagueId, userId)
    scores := usersScores
    for _, score := range(leagueScores) {
        scores = append(scores, score)
    }
    fmt.Println(len(usersScores))
    fmt.Println(usersScores)
    fmt.Println(len(leagueScores))
    fmt.Println(leagueScores)
    scores = repo.Initialize(svc, leagueId, userId, scores)
    fmt.Println(len(scores))
    fmt.Println(scores)
    grouped := GroupByUser(scores, leagueId, userId)
    fmt.Println(grouped)
}
