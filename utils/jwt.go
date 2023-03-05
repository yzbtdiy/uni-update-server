package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey []byte

// func init() {
// 	jwtKey = []byte(os.Getenv("JWT_SECRET"))
// }

func init() {
	jwtKey = []byte("0916cf39d65f66e3dbcba1d4ff3962d7")
}

type CustClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成Token
func GenerateToken(id int, username string) (string, error) {
	// 过期时间 默认7天
	// expireTime := time.Now().Add(7 * 24 * time.Hour)
	custClaims := CustClaims{
		Id:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, custClaims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	} else {
		return tokenStr, nil
	}
}

// 解析token
func ParseToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, err
}

// func ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
// 	claims := &Claims{}
// 	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		log.Println(err)
// 		return nil, nil, err
// 	}
// 	return token, claims, err
// }
