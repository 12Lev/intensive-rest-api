package auth

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	mongoDb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mongo "intensive-rest-api/pkg/database"
	"intensive-rest-api/pkg/utils"
)

type CreateUserRequest struct {
	Email    string `json:"email" bson:"email" example:"user@test.ru"`
	Company  string `json:"company" bson:"company" example:"ООО IT"`
	Password string `json:"password" bson:"password" example:"TestPwd123"`
}

//больше 2  - создаем структуру

type CreateUserResponse map[string]interface{}

func (request *CreateUserRequest) Validated() error {
	pwd := fmt.Sprintf("%v", request.Password)
	if len(pwd) <= 7 {
		return errors.New("password - not strong! Min 8 length")
	}
	return nil
}

func (request *CreateUserRequest) AddOrUpdateUser() (*CreateUserResponse, error) {
	err := request.Validated()
	if err != nil {
		return nil, err
	}

	pwd := request.Password
	email := request.Email
	request.Password = utils.NewHash(pwd)
	client, err := mongo.Connect()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())
	var models []mongoDb.WriteModel
	key := bson.M{"_id": email}
	value := bson.M{"$set": request}
	models = append(models, mongoDb.NewUpdateOneModel().
		SetFilter(key).
		SetUpdate(value).
		SetUpsert(true))
	bulkOption := options.BulkWriteOptions{}
	bulkOption.SetOrdered(true)
	_, err = client.Database(mongo.DefaultDB()).
		Collection(mongo.USERS_TABLE).
		BulkWrite(context.TODO(), models, &bulkOption)
	if err != nil {
		return nil, err
	}
	return &CreateUserResponse{
		"user":   email,
		"status": "ok",
	}, nil
}
