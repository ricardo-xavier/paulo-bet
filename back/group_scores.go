package main

import (
    "sort"
    "time"
    "back/model"
)

func GroupByUser(scores []model.ScoreBoard, leagueId string, login string) []model.ScoreBoard {
    m := make(map[string]model.ScoreBoard)
    for _, scoreBoard := range(scores) {
        userId := scoreBoard.UserId
        if userId == leagueId {
            matchId := scoreBoard.MatchId
            m[matchId] = scoreBoard
        }
    }

    u := make(map[string]int)
    for i, _ := range(scores) {
        userId := scores[i].UserId
        loc, _ := time.LoadLocation("America/Sao_Paulo")
        currentTime := time.Now().In(loc)
        date := currentTime.Format("2006-01-02 15:04:05")
        matchId := scores[i].MatchId
        admin := m[matchId]
        add := date >= admin.Date
        if userId != login {
            scores[i].Editable = false
            scores[i].Visible = date >= admin.Date
        } else {
            scores[i].Visible = true
            if login != leagueId {
                scores[i].Editable = date < admin.Date
            } else {
                scores[i].Editable = true
            }
        }
        if userId != leagueId || login == leagueId {
            n := u[userId]
            matchId := scores[i].MatchId
            admin := m[matchId]

            wAdmin := 0
            if admin.Home > admin.Visitors {
                wAdmin = 1
            } else if admin.Home < admin.Visitors {
                wAdmin = -1
            }
            wUser := 0
            if scores[i].Home > scores[i].Visitors {
                wUser = 1
            } else if scores[i].Home < scores[i].Visitors {
                wUser = -1
            }

            matches := 0
            if scores[i].Home == admin.Home {
                matches++
            }
            if scores[i].Visitors == admin.Visitors {
                matches++
            }

            if add {
                if matches == 2 {
                    n = n + 10
                } else if wUser == wAdmin {
                    n = n + 3 + (matches * 2)
                } else {
                    n = n + matches
                }
            }

            scores[i].Date = admin.Date[5:16]
            scores[i].Score = n - u[userId]
            u[userId] = n
        }
    }

    var list []model.ScoreBoard
    for k, v := range(u) {
        score := model.ScoreBoard {
            UserId: k,
            Score: v,
        }
        list = append(list, score)
    }

    sort.SliceStable(list, func(i, j int) bool {
        return list[i].Score > list[j].Score
    })

    return list
}
