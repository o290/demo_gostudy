package main

import (
	"crypto/rand"
	"crypto/rsa"
	claims2 "demo_gostudy/jwt/claims"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

// RSA签名与解析
func main() {
	//1.创建密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey
	if err != nil {
		log.Printf("generate key error:%s", err)
	}

	//2.创建claims
	claims := claims2.MyClaims{
		Username: "alice",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 3)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Server",
		},
	}

	//3.创建token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	//4.私钥加密
	signedString, err := token.SignedString(privateKey)
	if err != nil {
		log.Printf("sign error:%s", err)
	}
	fmt.Println("Token:", signedString)
	//5.公钥验证
	token, err = jwt.ParseWithClaims(signedString, &claims2.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		fmt.Println(err)
	} else if claims, ok := token.Claims.(*claims2.MyClaims); ok && token.Valid {
		fmt.Println(claims)
	}
}
