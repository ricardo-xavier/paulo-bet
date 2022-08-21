package model

type User struct {
    Login string `json:"login"`
    Name string `json:"name"`
    Password string `json:"password"`
}
