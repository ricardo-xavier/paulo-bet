package model

type BetRequest struct {
    MatchId string `json:"matchId"`
    Home int `json:"home"`
    Visitors int `json:"visitors"`
    Token string `json:"token"`
}
