all: login login_main get_leagues get_leagues_main scores scores_main

clean:
	rm -f login login_main get_leagues get_leagues_main scores scores_main *.zip

login: login.go model/user.go repo/session.go repo/user.go utils/response.go
	go build -o login login.go
	zip login.zip login

login_main: login_main.go model/user.go repo/session.go repo/user.go utils/response.go
	go build -o login_main login_main.go

get_leagues: get_leagues.go model/league.go repo/session.go repo/user.go utils/response.go
	go build -o get_leagues get_leagues.go
	zip get_leagues.zip get_leagues

get_leagues_main: get_leagues_main.go model/league.go repo/session.go repo/user.go utils/response.go
	go build -o get_leagues_main get_leagues_main.go

scores: scores.go model/score_board.go repo/score_board.go group_scores.go
	go build -o scores scores.go group_scores.go
	zip scores.zip scores

scores_main: scores_main.go model/score_board.go repo/score_board.go group_scores.go
	go build -o scores_main scores_main.go group_scores.go
