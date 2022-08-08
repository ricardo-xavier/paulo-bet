all: login get_leagues

clean:
	rm -f login get_leagues *.zip

login: login.go model/user.go repo/session.go repo/user.go utils/response.go
	go build -o login login.go
	zip login.zip login

get_leagues: get_leagues.go model/league.go repo/session.go repo/user.go utils/response.go
	go build -o get_leagues get_leagues.go
	zip get_leagues.zip get_leagues
