package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("zhang cheng")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成Token
func GenToken(userID int64, username string) (string, error) {
	c := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "community",
		},
	}
	//使用指定的签名方式创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获取完整的编码后的的字符串
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析Token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
