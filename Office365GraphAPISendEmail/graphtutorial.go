package main

import (
	"fmt"
	"graphtutorial/graphhelper"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Go Graph Tutorial")
	fmt.Println()

	// Load .env files
	// .env.local takes precedence (if present)
	godotenv.Load(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	graphHelper := graphhelper.NewGraphHelper()

	initializeGraph(graphHelper)
	sendMail(graphHelper)
}

func initializeGraph(graphHelper *graphhelper.GraphHelper) {
	err := graphHelper.InitializeGraphForUserAuth()
	if err != nil {
		log.Panicf("Error initializing Graph for user auth: %v\n", err)
	}
}

func sendMail(graphHelper *graphhelper.GraphHelper) {
	// Send mail to the signed-in user
	// Get the user for their email address
	subject := "Testing Microsoft Graph"
	body := "Hello world!"
	email := "samuel.venturi@giannotti.com.br"

	err := graphHelper.SendMail(&subject, &body, &email)
	if err != nil {
		log.Panicf("Error sending mail: %v", err)
	}

	fmt.Println("Mail sent.")
	fmt.Println()
}
