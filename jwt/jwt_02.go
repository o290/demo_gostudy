package main

import (
	"demo_gostudy/jwt/claims"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func main() {
	secretKey := []byte("my secret key")
	//jwt_01生成的token
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsaWNlIiwiaXNzIjoibXlhcHAiLCJleHAiOjE3MzExNDAzNTF9.iVE_j_JSW5mZyb4IyLyAfV5aeXMVd1xkuoHlApHWrsE"
	//解析jwt
	//func(token *jwt.Token) (interface{}, error)用来验证签名算法是否匹配，匹配则返回秘钥进行签名验证
	token, err := jwt.ParseWithClaims(tokenString, &claims.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		log.Println("error parsing token:", err)
		//return
	}
	//从token中获取到claims
	if claims, ok := token.Claims.(*claims.MyClaims); ok && token.Valid {
		fmt.Println("username:", claims.Username)
		fmt.Println("issuer:", claims.Issuer)
	} else {
		fmt.Println("invalid token")
	}
	fmt.Println("Raw:", token.Raw)
	fmt.Println("Header:", token.Header)
	fmt.Println("Valid:", token.Valid)
	fmt.Println("Method:", token.Method)
	fmt.Println("Claims:", token.Claims)
	fmt.Println("Signature:", token.Signature)

	//错误处理
	if token.Valid {
		fmt.Println("token合法")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		fmt.Println("不是一个token")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		fmt.Println("token已经过期或者还没有生效")
	} else {
		fmt.Println("token处理异常...")
	}
}
