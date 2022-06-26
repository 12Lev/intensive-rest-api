package utils

import (
	"encoding/json"
	"os"
)

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == `` {
		return fallback
	}
	return value
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ToJsonBytes(input interface{}) []byte {
	marshal, err := json.Marshal(input)
	if err != nil {
		return []byte(``)
	}
	return marshal
}
