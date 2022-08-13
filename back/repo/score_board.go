package repo

import (
    "strings"
    "strconv"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
    "back/model"
)

func GetScores(svc *dynamodb.DynamoDB, leagueId string) []model.ScoreBoard {
    var scores []model.ScoreBoard
	sort := expression.Key("sort").Equal(expression.Value(leagueId))

	proj := expression.NamesList(expression.Name("hash"),
        expression.Name("home"),
        expression.Name("visitor"))

	expr, err := expression.NewBuilder().WithKeyCondition(sort).WithProjection(proj).Build()
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
        userGame := strings.Split(*i["hash"].S, "_")
        home, _ := strconv.Atoi(*i["home"].N)
        visitor, _ := strconv.Atoi(*i["visitor"].N)
		scoreBoard := model.ScoreBoard {
            UserId: userGame[0],
            GameId: userGame[1],
            Home: home,
            Visitor: visitor,
        }
		scores = append(scores, scoreBoard)
	}

    return scores
}
