package util

import (
	"math/rand"
	"strings"
	"time"
)

const str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTU"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(str)

	for i := 0; i < n; i++ {
		c := str[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomStringHP(n int) string {
	var sb strings.Builder

	k := len(str)

	for i := 0; i < n; i++ {
		c := str[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomHashspassword() string {
	return RandomStringHP(20)
}

func RandomFullname() string {
	return RandomString(6)
}

func RandomEmail() string {
	return RandomString(6)
}

func RandomCurrency() string {
	cureencies := []string{"EUR", "USDT", "BTC"}
	n := len(cureencies)
	return cureencies[rand.Intn(n)]

}
