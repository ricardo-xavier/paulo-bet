package model

type ScoreBoard struct {
    UserId string `json:"userId"`
    GameId string `json:"gameId"`
    Home int `json:"home"`
    Visitor int `json:"visitor"`
    Score int `json:"score"`
}
