package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/gofrs/uuid"
)

func main() {
	// Generate a UUID version 7
	uuidV7, err := uuid.NewV7()
	if err != nil {
		fmt.Println("Error generating UUID:", err)
		return
	}

	fmt.Printf("UUIDv7: %s\n", uuidV7.String())

	// Remove hyphens and convert the UUID to a big.Int
	uuidWithoutHyphens := strings.Replace(uuidV7.String(), "-", "", -1)
	tokenID, success := new(big.Int).SetString(uuidWithoutHyphens, 16)

	if !success {
		fmt.Println("Error converting UUID to number string.")
		return
	}

	fmt.Printf("Token ID: %s\n", tokenID.String())
}
