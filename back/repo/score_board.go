package repo

import (
    "fmt"
    "strings"
    "strconv"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
    "back/model"
)

func GetScores(svc *dynamodb.DynamoDB, leagueId string, userId *string, login string) []model.ScoreBoard {
    var scores []model.ScoreBoard

	sort := expression.Key("sort").Equal(expression.Value(leagueId))
	proj := expression.NamesList(expression.Name("hash"),
        expression.Name("date"),
        expression.Name("home"),
        expression.Name("visitors"))

    if userId != nil {
	    hash := expression.Key("hash").BeginsWith(*userId + "_")
	    sort = expression.KeyAnd(sort, hash)
    }

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
        userMatch := strings.Split(*i["hash"].S, "_")
        home, _ := strconv.Atoi(*i["home"].N)
        visitors, _ := strconv.Atoi(*i["visitors"].N)
		scoreBoard := model.ScoreBoard {
            UserId: userMatch[0],
            MatchId: userMatch[1],
            Home: home,
            Visitors: visitors,
        }
        if i["date"] != nil {
            scoreBoard.Date = *i["date"].S
            if strings.HasSuffix(scoreBoard.Date, ":00") {
                scoreBoard.Date = scoreBoard.Date[0:len(scoreBoard.Date)-3]
            }
        }
		scores = append(scores, scoreBoard)
	}

    return scores
}

func Initialize(svc *dynamodb.DynamoDB, leagueId string, userId string, scores []model.ScoreBoard) []model.ScoreBoard {
    var adminBets []model.ScoreBoard
    var userBets []model.ScoreBoard
    for _, score := range(scores) {
        if score.UserId == userId {
            userBets = append(userBets, score)
        }
        if score.UserId == leagueId {
            adminBets = append(adminBets, score)
        }
    }
    if len(userBets) == len(adminBets) {
        return scores
    }
    for _, admin := range(adminBets) {
        found := false
        for _, user := range(userBets) {
            if user.MatchId == admin.MatchId {
                found = true
                break
            }
        }
        if !found {
            addBet(svc, leagueId, userId, admin.MatchId)
            bet := model.ScoreBoard {
                UserId: userId,
                MatchId: admin.MatchId,
                Home: 0,
                Visitors: 0,
            }
            scores = append(scores, bet)
        }
    }
    return scores
}

func addBet(svc *dynamodb.DynamoDB, leagueId string, userId string, matchId string) {
    fmt.Printf("PAULOBET:addBet [%s] [%s] [%s]\n", leagueId, userId, matchId)
    item := map[string]*dynamodb.AttributeValue {
        "hash": { S: aws.String(userId + "_" + matchId) },
        "sort": { S: aws.String(leagueId) },
        "home": { N: aws.String("0") },
        "visitors": { N: aws.String("0") },
    }
	params := &dynamodb.PutItemInput{
		Item:      item,
        TableName: aws.String("PAULOBET"),
	}
	_, err := svc.PutItem(params)
	if err != nil {
		panic(err)
	}
}

func UpdateBet(svc *dynamodb.DynamoDB, leagueId string, userId string, matchId string, home int, visitors int) {
    fmt.Printf("PAULOBET:updBet [%s] [%s] [%s] %d %d\n", leagueId, userId, matchId, home, visitors)
    key := map[string]*dynamodb.AttributeValue {
        "hash": { S: aws.String(userId + "_" + matchId) },
        "sort": { S: aws.String(leagueId) },
    }
    values := map[string]*dynamodb.AttributeValue {
        ":home": { N: aws.String(strconv.Itoa(home)) },
        ":visitors": { N: aws.String(strconv.Itoa(visitors)) },
    }

	params := &dynamodb.UpdateItemInput{
		Key:                       key,
        TableName:                 aws.String("PAULOBET"),
		UpdateExpression:          aws.String("set home=:home, visitors=:visitors"),
        ExpressionAttributeValues: values,
		ReturnValues:              aws.String("UPDATED_NEW"),
	}

	_, err := svc.UpdateItem(params)
	if err != nil {
		panic(err)
	}
}
