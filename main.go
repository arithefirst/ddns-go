package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go/v3/zones"
)

type config struct {
	Token string
	Email string
	Zones []string
}

func main() {
	// ip, err := getIpAddress()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	print("Reading Config...\n")
	Config, err := readConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	print("Creating Client...\n")
	client, err := getClient(Config)
	if err != nil {
		log.Fatal(err.Error())
	}

	print("Sending Request...\n")
	page, err := client.Zones.List(context.TODO(), zones.ZoneListParams{})
	if err != nil {
		log.Fatalf("Error in request: %s", err.Error())
	}
	fmt.Printf("%v\n", page)
}
