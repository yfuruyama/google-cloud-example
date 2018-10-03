package main

import (
	"context"
	"fmt"
	"os"

	kms "cloud.google.com/go/kms/apiv1"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s KEY_ID\n", os.Args[0])
		os.Exit(-1)
	}
	keyId := os.Args[1]

	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		panic(err)
	}

	req := &kmspb.GetPublicKeyRequest{
		Name: keyId,
	}

	publicKey, err := client.GetPublicKey(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", publicKey.GetPem())
}
