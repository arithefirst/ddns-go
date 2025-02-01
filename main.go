package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go/v3/zones"
	"github.com/joho/godotenv"
)

func main() {
	// ip, err := getIpAddress()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// Load the env vars into the program
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	client, err := getClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	page, err := client.Zones.List(context.TODO(), zones.ZoneListParams{})
	if err != nil {
		log.Fatalf("Error in request: %s", err.Error())
	}
	fmt.Printf("%v\n", page)
}
