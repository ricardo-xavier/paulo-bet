package utils

import (
    "time"
    "crypto/md5"
    "encoding/hex"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ssm"
)

func Crypt(text string) string {
    hash := md5.Sum([]byte(text))
    return hex.EncodeToString(hash[:])
}

func getSecret(name string) string {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))
    ssmsvc := ssm.New(sess)
    parameter, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
        Name: &name,
    })
    if err != nil {
        panic(err)
    }
    return *parameter.Parameter.Value
}

func BuildToken(user string) string {
    currentTime := time.Now()
    date := currentTime.Format("02-01-2006")
    secret := getSecret("paulobet-secret")
    return Crypt(secret + ":" + user + ":" + date)
}

func CheckToken(user string, token string) bool {
    currentTime := time.Now()
    date := currentTime.Format("02-01-2006")
    secret := getSecret("paulobet-secret")
    return token == Crypt(secret + ":" + user + ":" + date)
}
