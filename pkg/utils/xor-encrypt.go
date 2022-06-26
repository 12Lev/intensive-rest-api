package utils

import "encoding/base64"

//для рандомизации пароля
const (
	lower  = "abcdefghijklmnopqrstuvwxyz"
	upper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits = "0123456789"
)

func NewHash(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}
func FromHash(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}
