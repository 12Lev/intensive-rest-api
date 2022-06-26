package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"intensive-rest-api/pkg/utils"
	"time"
)

const (
	USERS_TABLE       = `users`
	PROFILE_TABLE     = `user_profiles`
	USER_TOKENS_TABLE = `user_tokens`
)

type MongoConfig struct {
	mongoHost string
	mongoPort string
	mongoDb   string
}

//создаем дефолтну конфитгурацию
var (
	mongoConfig = MongoConfig{
		mongoHost: utils.GetEnv("MONGO_HOST", "localhost"),
		mongoPort: utils.GetEnv("MONGO_PORT", "27017"),
		mongoDb:   utils.GetEnv("MONGO_DB", "test"),
	}
)

func InitConfig() {
	mongoConfig = MongoConfig{
		mongoHost: utils.GetEnv("MONGO_HOST", "localhost"),
		mongoPort: utils.GetEnv("MONGO_PORT", "27017"),
		mongoDb:   utils.GetEnv("MONGO_DB", "test"),
	}
}

func DefaultDB() string {
	return mongoConfig.mongoDb
}

func Connect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().
		SetMaxPoolSize(300). //максимальное количество соединений
		ApplyURI(fmt.Sprintf(`mongodb://%s:%s/%s`, mongoConfig.mongoHost, mongoConfig.mongoPort, mongoConfig.mongoDb))
	return mongo.Connect(ctx, clientOptions)
}

func Exist(database, table string, filter interface{}) (bool, error) {
	findOptions := options.FindOne()
	client, err := Connect()
	if err != nil {
		return false, err
	}
	one := client.Database(database).
		Collection(table).
		FindOne(context.Background(), filter, findOptions)
	defer client.Disconnect(context.Background())
	err = one.Err()
	if err != nil {
		return false, err
	}
	return true, nil
}
