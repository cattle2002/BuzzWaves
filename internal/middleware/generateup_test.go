package middleware

import (
	"BuzzWaves/pkkg"
	"fmt"
	"testing"
)

func TestGenernateUp(t *testing.T) {
	jwtup, _ := pkkg.GenerateJWTUP("gone", "123")
	fmt.Println(jwtup)
}
