package helpers

import (
	"time"
	"math/rand"
	"fmt"
)

func GenerateCode() string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rand.Int31n(1000000))
	return code
}
