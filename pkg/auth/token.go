package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	mongoDb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mongo "intensive-rest-api/pkg/database"
	"intensive-rest-api/pkg/utils"
	"time"
)

var secretKey = []byte("secretGoJwt")

type AuthRequest struct {
	Email    string `json:"email" bson:"email" example:"user@test.ru"`
	Password string `json:"password" bson:"password" example:"TestPwd123"`
}

type AuthResponse map[string]interface{}

func (request *AuthRequest) Validated() error {
	pwd := fmt.Sprintf("%v", request.Password)
	if len(pwd) <= 7 {
		return errors.New("password - not strong! Min 8 length")
	}
	return nil
}

func (request *AuthRequest) Auth() (*AuthResponse, error) {
	err := request.Validated()
	if err != nil {
		return nil, err
	}
	login := request.Email
	password := request.Password
	pwd := utils.NewHash(password)
	client, err := mongo.Connect()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())
	findOptions := options.FindOne()
	one := client.Database(mongo.DefaultDB()).
		Collection(mongo.USERS_TABLE).
		FindOne(context.Background(), bson.M{"_id": login, "password": pwd}, findOptions)
	err = one.Err()
	if err != nil {
		return nil, err
	}
	var decodeType map[string]interface{}
	err = one.Decode(&decodeType)
	if err != nil {
		return nil, err
	}
	/*get token*/
	company := fmt.Sprintf("%v", decodeType["company"])
	createTime := time.Now().UTC()
	expiration := createTime.Add(time.Hour * 24)
	type myCustomClaims struct {
		Expiration int64  `json:"token_expiration_date"`
		Create     int64  `json:"token_create_date"`
		Login      string `json:"login"`
		Company    string `json:"company"`
		jwt.StandardClaims
	}
	claims := myCustomClaims{
		Expiration:     expiration.Unix(),
		Create:         createTime.Unix(),
		Login:          login,
		Company:        company,
		StandardClaims: jwt.StandardClaims{},
	}
	tokenData := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenData.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
	result := AuthResponse{}
	result["id"] = uuid.NewString()
	result["login"] = login
	result["email"] = login
	result["company"] = company
	result["token_create_date"] = createTime.Unix()
	result["token_expiration_date"] = expiration.Format(time.RFC3339)
	result["token"] = token
	/*add to db*/
	var models []mongoDb.WriteModel
	key := bson.M{"_id": login}
	value := bson.M{"$set": result}
	models = append(models, mongoDb.NewUpdateOneModel().
		SetFilter(key).
		SetUpdate(value).
		SetUpsert(true))
	bulkOption := options.BulkWriteOptions{}
	bulkOption.SetOrdered(true)
	_, err = client.Database(mongo.DefaultDB()).
		Collection(mongo.USER_TOKENS_TABLE).
		BulkWrite(context.TODO(), models, &bulkOption)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func Validate(token string) bool {
	exist, err := mongo.Exist(mongo.DefaultDB(), mongo.USER_TOKENS_TABLE, bson.M{"token": token})
	if err != nil {
		return false
	}
	return exist
}
