package model

type ChangePasswordRequest struct {
    Id string `json:"id"`
    Password string `json:"password"`
    Token string `json:"token"`
}
