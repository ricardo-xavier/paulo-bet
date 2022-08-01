all: login

clean:
	rm -f login login.zip

login: login.go model/user.go repo/user.go utils/response.go
	go build -o login login.go
	zip login.zip login
