package main

import (
    "sort"
    "back/model"
)

func GroupByUser(scores []model.ScoreBoard, leagueId string) []model.ScoreBoard {
    m := make(map[string]model.ScoreBoard)
    for _, scoreBoard := range(scores) {
        userId := scoreBoard.UserId
        if userId == leagueId {
            gameId := scoreBoard.GameId
            m[gameId] = scoreBoard
        }
    }

    u := make(map[string]int)
    for i, _ := range(scores) {
        userId := scores[i].UserId
        if userId != leagueId {
            n := u[userId]
            gameId := scores[i].GameId
            admin := m[gameId]

            wAdmin := 0
            if admin.Home > admin.Visitor {
                wAdmin = 1
            } else if admin.Home < admin.Visitor {
                wAdmin = -1
            }
            wUser := 0
            if scores[i].Home > scores[i].Visitor {
                wUser = 1
            } else if scores[i].Home < scores[i].Visitor {
                wUser = -1
            }

            matches := 0
            if scores[i].Home == admin.Home {
                matches++
            }
            if scores[i].Visitor == admin.Visitor {
                matches++
            }

            if matches == 2 {
                n = n + 10
            } else if wUser == wAdmin {
                n = n + 3 + (matches * 2)
            } else {
                n = n + matches
            }

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
