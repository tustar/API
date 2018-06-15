package helpers

import (
	"time"
	"math/rand"
	"fmt"
)

func GenerateCaptcha() string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	captcha := fmt.Sprintf("%06v", rand.Int31n(1000000))
	return captcha
}
