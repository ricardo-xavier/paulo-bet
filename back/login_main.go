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
    svc := repo.Connect()
    userEntity := repo.GetUser(svc, userId)
    if userEntity == nil {
        panic(fmt.Errorf("%s", userId))
    }
    loginResponse := model.LoginResponse {
        UserName: userEntity.Name,
    }
    body, err := json.Marshal(loginResponse)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(body))
}
