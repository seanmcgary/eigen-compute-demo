package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	mnemonic := os.Getenv("MNEMONIC")
	if mnemonic == "" {
		log.Print("MNEMONIC environment variable is not set")
	}

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Printf("Error creating wallet from mnemonic: %v", err)
	} else {
		path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
		account, err := wallet.Derive(path, false)
		if err != nil {
			log.Printf("Error deriving account: %v", err)
		} else {
			fmt.Printf("First wallet address: %s\n", account.Address.Hex())
		}
	}

	for {
		fmt.Println(time.Now().Format(time.RFC3339))
		time.Sleep(time.Minute)
	}
}
