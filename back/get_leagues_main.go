package main

import (
    "fmt"
    "os"
    "encoding/json"
    "back/model"
    "back/repo"
)

func main() {
    userId := os.Args[1]
    fmt.Println(userId)
    svc := repo.Connect()
    leagues := repo.GetUserLeagues(svc, userId)
    getLeaguesResponse := model.GetLeaguesResponse {
        Leagues: leagues,
    }
    body, err := json.Marshal(getLeaguesResponse)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(body))
}
