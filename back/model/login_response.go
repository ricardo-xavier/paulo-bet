package model

type LoginResponse struct {
    UserName string `json:"userName"`
    Token string `json:"token"`
}
