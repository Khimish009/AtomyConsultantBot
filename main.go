package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	tgBotHost = "api.telegram.org"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	token := mustToken()
	fmt.Print("Work", token)
}

func mustToken() string {
	token := os.Getenv("TG_BOT_TOKEN")

	if token == "" {
		log.Fatal("token is not specified")
	}

	return token
}
