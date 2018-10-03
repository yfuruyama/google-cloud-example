package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"time"

	kms "cloud.google.com/go/kms/apiv1"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s KEY_ID\n", os.Args[0])
		os.Exit(-1)
	}
	keyId := os.Args[1]

	header := struct {
		Typ string `json:"typ"`
		Alg string `json:"alg"`
	}{
		Typ: "JWT",
		Alg: "RS256",
	}

	headerJson, err := json.Marshal(header)
	if err != nil {
		panic(err)
	}

	body := struct {
		Iat int64 `json:"iat"`
	}{
		Iat: time.Now().Unix(),
	}

	bodyJson, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	headerAndBody := fmt.Sprintf("%s.%s", base64.RawURLEncoding.EncodeToString(headerJson), base64.RawURLEncoding.EncodeToString(bodyJson))

	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		panic(err)
	}

	digest := sha256.Sum256([]byte(headerAndBody))
	digestSlice := digest[:]
	req := &kmspb.AsymmetricSignRequest{
		Name: keyId,
		Digest: &kmspb.Digest{
			Digest: &kmspb.Digest_Sha256{digestSlice},
		},
	}

	resp, err := client.AsymmetricSign(ctx, req)
	if err != nil {
		panic(err)
	}
	signature := resp.Signature

	fmt.Printf("%s.%s", string(headerAndBody), base64.RawURLEncoding.EncodeToString(signature))
}
