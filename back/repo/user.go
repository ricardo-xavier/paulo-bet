package repo

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
    "back/model"
)

func GetUser(svc *dynamodb.DynamoDB, id string) *model.User {
    user := model.User{}
    key := map[string]*dynamodb.AttributeValue { "hash": { S: aws.String(id) }, "sort": { S: aws.String("USER") } }

    params := &dynamodb.GetItemInput {
        Key:       key,
        TableName: aws.String("PAULOBET"),
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

func GetUserLeagues(svc *dynamodb.DynamoDB, userId string) []model.League {
    var leagues []model.League
	hash := expression.Key("hash").Equal(expression.Value(userId))
	sort := expression.Key("sort").Equal(expression.Value("LEAGUE"))

	proj := expression.NamesList(expression.Name("leagueId"))

	expr, err := expression.NewBuilder().WithKeyCondition(hash.And(sort)).WithProjection(proj).Build()
	if err != nil {
        panic(err)
	}

    params := &dynamodb.QueryInput {
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
        TableName:                 aws.String("PAULOBET"),
    }

    result, err := svc.Query(params)
    if err != nil {
        panic(err)
    }

	for _, i := range result.Items {
		league := model.League{}
		err = dynamodbattribute.UnmarshalMap(i, &league)
		if err != nil {
            panic(err)
		}
		leagues = append(leagues, league)
	}

    return leagues
}
