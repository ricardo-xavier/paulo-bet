package main

import (
    "fmt"
    "os"
    "encoding/json"
    "back/model"
    "back/repo"
    "back/utils"
)

func main() {
    userId := os.Args[1]
    password := os.Args[2]
    svc := repo.Connect()
    userEntity := repo.GetUser(svc, userId)
    if userEntity == nil {
        panic(fmt.Errorf("%s", userId))
    }
    if userEntity.Password != utils.Crypt(password) && !(userEntity.Password == "" && password == "") {
        panic(fmt.Errorf("[%s] [%s]", userEntity.Password, utils.Crypt(password)))
    }
    loginResponse := model.LoginResponse {
        UserName: userEntity.Name,
        Token: utils.BuildToken(userEntity.Name),
    }
    body, err := json.Marshal(loginResponse)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(body))
}
