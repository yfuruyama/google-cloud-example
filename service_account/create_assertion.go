package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type ServiceAccount struct {
	ClientEmail string `json:"client_email"`
	PrivateKey  string `json:"private_key"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <SERVICE_ACCOUNT_KEY.json>\n", os.Args[0])
		os.Exit(-1)
	}

	serviceAccountPath := os.Args[1]
	serviceAccountJson, err := ioutil.ReadFile(serviceAccountPath)
	if err != nil {
		panic(err)
	}

	var serviceAccount ServiceAccount
	err = json.Unmarshal(serviceAccountJson, &serviceAccount)
	if err != nil {
		panic(err)
	}

	header := &struct {
		Alg string `json:"alg"`
		Typ string `json:"typ"`
	}{
		Alg: "RS256",
		Typ: "JWT",
	}

	headerJson, _ := json.Marshal(header)
	headerJwt := base64.RawURLEncoding.EncodeToString(headerJson)

	body := &struct {
		Iss   string `json:"iss"`
		Aud   string `json:"aud"`
		Scope string `json:"scope"`
		Iat   string `json:"iat"`
		Exp   string `json:"exp"`
	}{
		Iss:   serviceAccount.ClientEmail,
		Aud:   "https://oauth2.googleapis.com/token",
		Scope: "https://www.googleapis.com/auth/cloud-platform",
		Iat:   fmt.Sprintf("%d", time.Now().Unix()),
		Exp:   fmt.Sprintf("%d", time.Now().Unix()+900),
	}
	bodyJson, _ := json.Marshal(body)
	bodyJwt := base64.RawURLEncoding.EncodeToString(bodyJson)

	headerAndBody := fmt.Sprintf("%s.%s", headerJwt, bodyJwt)

	block, _ := pem.Decode([]byte(serviceAccount.PrivateKey))
	if block == nil {
		panic(errors.New("invalid private key"))
	}
	keyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	privateKey, ok := keyInterface.(*rsa.PrivateKey)
	if !ok {
		panic(errors.New("not RSA private key"))
	}
	privateKey.Precompute()

	headerAndBodyHash := sha256.Sum256([]byte(headerAndBody))
	headerAndBodyHashSlice := headerAndBodyHash[:]
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, headerAndBodyHashSlice)
	if err != nil {
		panic(err)
	}
	signatureJwt := base64.RawURLEncoding.EncodeToString(signature)

	assertion := fmt.Sprintf("%s.%s", headerAndBody, signatureJwt)
	fmt.Println(assertion)
}
