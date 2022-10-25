package model

type ScoreBoard struct {
    UserId string `json:"userId"`
    MatchId string `json:"matchId"`
    Home int `json:"home"`
    Visitors int `json:"visitors"`
    Score int `json:"score"`
    Date string `json:"date"`
    Editable bool `json:"editable"`
    Visible bool `json:"visible"`
}
