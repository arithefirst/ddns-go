package main

import (
	"fmt"
	"log"
)

func main() {
	ip, err := getIpAddress()
	if err != nil {
		log.Fatal(err)
	}

	client, err := getClient()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("IP: %sClient: %v\n", ip, client)
}
