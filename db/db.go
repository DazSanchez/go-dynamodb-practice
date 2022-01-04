package db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var db *dynamodb.Client

func Init() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		panic(fmt.Sprintf("unable to load SDK config, %v", err))
	}

	db = dynamodb.NewFromConfig(cfg)
}

func GetDB() *dynamodb.Client {
	return db
}
