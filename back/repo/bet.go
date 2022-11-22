package repo

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetDate(svc *dynamodb.DynamoDB, userId string, matchId string) *string {
    key := map[string]*dynamodb.AttributeValue {
        "hash": { S: aws.String(userId + "_" + matchId) },
        "sort": { S: aws.String(userId) },
    }

    params := &dynamodb.GetItemInput {
        Key:       key,
        TableName: aws.String("PAULOBET"),
    }

    result, err := svc.GetItem(params)
    if err != nil {
        panic(err)
    }

    if result.Item == nil {
        panic("admin match not found")
    }

    return result.Item["date"].S
}
