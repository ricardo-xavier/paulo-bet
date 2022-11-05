package repo

import (
    "fmt"
    "strings"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
    "back/model"
)

func GetUser(svc *dynamodb.DynamoDB, login string) *model.User {
    user := model.User{}
    key := map[string]*dynamodb.AttributeValue {
        "hash": { S: aws.String(login) },
        "sort": { S: aws.String("USER") },
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
	sort := expression.Key("sort").Equal(expression.Value("LEAGUE"))
	hash := expression.Key("hash").BeginsWith(userId + "_")

	proj := expression.NamesList(expression.Name("hash"))

	expr, err := expression.NewBuilder().WithKeyCondition(sort.And(hash)).WithProjection(proj).Build()
	if err != nil {
        panic(err)
	}

    params := &dynamodb.QueryInput {
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
        IndexName:                 aws.String("sort-hash-index"),
        TableName:                 aws.String("PAULOBET"),
    }

    result, err := svc.Query(params)
    if err != nil {
        panic(err)
    }

	for _, i := range result.Items {
        leagueId := strings.Split(*i["hash"].S, "_")[1]
		league := model.League {
            LeagueId: leagueId,
        }
		leagues = append(leagues, league)
	}

    return leagues
}

func ChangePassword(svc *dynamodb.DynamoDB, login string, password string) {
    fmt.Printf("PAULOBET:change password [%s]\n", login)
    key := map[string]*dynamodb.AttributeValue {
        "hash": { S: aws.String(login) },
        "sort": { S: aws.String("USER") },
    }
    values := map[string]*dynamodb.AttributeValue { ":password": { S: aws.String(password) } }

	params := &dynamodb.UpdateItemInput{
		Key:                       key,
        TableName:                 aws.String("PAULOBET"),
		UpdateExpression:          aws.String("set password=:password"),
        ExpressionAttributeValues: values,
		ReturnValues:              aws.String("UPDATED_NEW"),
	}

	_, err := svc.UpdateItem(params)
	if err != nil {
		panic(err)
	}
}
