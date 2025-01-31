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

	token, err := getAccessToken()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("IP: %sToken: %s\n", ip, token)
}
