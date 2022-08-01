package repo

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "back/model"
)

func GetUser(login string) *model.User {
    user := model.User{}
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))
    svc := dynamodb.New(sess)
    key := map[string]*dynamodb.AttributeValue { "login": { S: aws.String(login) } }

    params := &dynamodb.GetItemInput {
        Key:             key,
        TableName:       aws.String("USERS"),
    }

    result, err := svc.GetItem(params)
    if err != nil {
        panic(err)
    }

    if result.Item == nil {
        return nil
    }

    err = dynamodbattribute.UnmarshalMap(result.Item, &user)
    if err != nil {
        panic(err)
    }
    return &user
}
