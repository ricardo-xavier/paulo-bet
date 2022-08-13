package main

import (
    "fmt"
    "os"
    "back/repo"
)

func main() {
    leagueId := os.Args[1]
    svc := repo.Connect()
    scores := repo.GetScores(svc, leagueId)
    grouped := GroupByUser(scores, leagueId)
    fmt.Println(grouped)
}
