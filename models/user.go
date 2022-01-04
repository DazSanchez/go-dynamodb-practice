package models

import (
	"context"

	"github.com/DazSanchez/go-dynamodb-practice/db"
	"github.com/DazSanchez/go-dynamodb-practice/forms"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type User struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	FirstSurname string `json:"firstSurname"`
	Email        string `json:"email"`
}

func (h User) GetAll() ([]User, error) {
	db := db.GetDB()

	res, err := db.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("Users"),
	})
	if err != nil {
		return nil, err
	}

	users := []User{}

	err = attributevalue.UnmarshalListOfMaps(res.Items, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (h User) GetByID(id string) (*User, error) {
	db := db.GetDB()

	res, err := db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("Users"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		return nil, err
	}

	var user *User
	if err = attributevalue.UnmarshalMap(res.Item, &user); err != nil {
		return nil, err
	}

	return user, nil
}

func (h User) CreateUser(payload forms.CreateUserDTO) (*User, error) {
	db := db.GetDB()
	user := User{
		ID:           uuid.NewString(),
		Name:         payload.Name,
		FirstSurname: payload.FirstSurname,
		Email:        payload.Email,
	}

	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		return nil, err
	}

	params := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("Users"),
	}

	if _, err := db.PutItem(context.TODO(), params); err != nil {
		return nil, err
	}

	return &user, nil
}
