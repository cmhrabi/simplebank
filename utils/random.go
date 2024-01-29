package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min int, max int) int {
	return (rand.Intn((max - min)) + min)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		num := rand.Intn(k)
		sb.WriteByte(alphabet[num])

	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(RandomInt(0, 10))
}

func RandomMoney() int64 {
	return int64(RandomInt(20, 1000))
}

func RandomCurrency() string {
	avail := []string{"USD", "CAD", "EUR"}
	return avail[RandomInt(0, len(avail)-1)]
}
