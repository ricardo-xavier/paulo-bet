package main

import (
    "os"
    "back/repo"
    "back/utils"
)

func main() {
    userId := os.Args[1]
    password := os.Args[2]
    svc := repo.Connect()
    repo.ChangePassword(svc, userId, utils.Crypt(password))
}
