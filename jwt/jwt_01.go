package main

import (
	claims2 "demo_gostudy/jwt/claims"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

// HMAC的创建和签名

func main() {
	//1.定义秘钥
	secretKey := []byte("my secret key")
	//2.定义claims载荷
	claims := claims2.MyClaims{
		Username: "alice",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "myapp",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Minute)), // 设置有效期为 24 小时
		},
	}
	//2.使用HS256算法创建jwt——使用预定义的claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//3.使用密钥对token进行签名
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatalf("sign fail:%s", err)
	}
	fmt.Println("Generated JWT1:", tokenString)

	//创建jwt不使用预定义的claims
	token1 := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "alice",
		"id":       "20241109",
	})
	tokenString1, err := token1.SignedString(secretKey)
	if err != nil {
		log.Fatalf("sign fail:%s", err)
	}
	fmt.Println("Generated JWT2:", tokenString1)
}
