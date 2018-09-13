package util

import (
	// standard library
	// "time"
	"fmt"
	"log"
	// "bytes"
	// "net/smtp"
	sp "github.com/SparkPost/gosparkpost"
)

func SendEmail(emailChannel chan string, userEmail string){
	fmt.Println("inside sendEmail - starting!")

	apiKey := "1cc105670c9d6c49ac3e08c0be236ae914a9f42f"
	cfg := &sp.Config{
	  BaseUrl:    "https://api.sparkpost.com",
	  ApiKey:     apiKey,
	  ApiVersion: 1,
	}
	var client sp.Client
	err := client.Init(cfg)
	if err != nil {
	  log.Fatalf("SparkPost client init failed: %s\n", err)
	}
  
	// Create a Transmission using an inline Recipient List
	// and inline email Content.
	tx := &sp.Transmission{
	  Recipients: []string{"pweyand@gmail.com"},
	  Content: sp.Content{
		HTML:    "<p>Hello world</p>",
		From:    "platypus@email.zennify.me",
		Subject: "Hello from gosparkpost",
	  },
	}
	id, _, err := client.Send(tx)
	if err != nil {
	  log.Fatal(err)
	}
  
	// The second value returned from Send
	// has more info about the HTTP response, in case
	// you'd like to see more than the Transmission id.
	log.Printf("Transmission sent with id [%s]\n", id)
	
	emailChannel<-"email notification sent!"
}
