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

	fmt.Print(ip)
}
