package main

import (
    "os"
    "strconv"
    "back/repo"
)

func main() {
    leagueId := os.Args[1]
    userId := os.Args[2]
    gameId := os.Args[3]
    home, _ := strconv.Atoi(os.Args[4])
    visitors, _ := strconv.Atoi(os.Args[5])
    svc := repo.Connect()
    repo.UpdateBet(svc, leagueId, userId, gameId, home, visitors)
}
