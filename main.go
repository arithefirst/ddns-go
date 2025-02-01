package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/dns"
	"github.com/cloudflare/cloudflare-go/v3/zones"
)

type config struct {
	Token   string
	Email   string
	Records []string
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

	fmt.Print(Config)

	print("Creating Client...\n")
	client, err := getClient(Config)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Fetch all of your account's Zone IDs
	zones, err := client.Zones.List(context.TODO(), zones.ZoneListParams{})
	// For each Zone ID, search for the requested domains
	for _, zone := range zones.Result {
		page, err := client.DNS.Records.List(context.TODO(), dns.RecordListParams{
			ZoneID: cloudflare.F(zone.ID),
		})
		if err != nil {
			log.Fatal(err.Error())
		}

		for _, record := range Config.Records {
			for _, v := range page.Result {
				if v.Name == record {
					fmt.Print("AHH")
				}
			}
		}
	}
}
