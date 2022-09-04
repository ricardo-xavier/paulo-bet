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
    scores := repo.GetScores(svc, leagueId, nil, userId)
    fmt.Println(scores)
    scores = repo.Initialize(svc, leagueId, userId, scores)
    fmt.Println(scores)
    grouped := GroupByUser(scores, leagueId, userId)
    fmt.Println(grouped)
}
