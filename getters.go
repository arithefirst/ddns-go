package main

import (
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/joho/godotenv"
)

func getIpAddress() (string, error) {
	resp, err := http.Get("https://icanhazip.com")
	if err != nil {
		return "", errors.New("Error getting IP Address: " + err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("Error getting IP Address: " + err.Error())
	}

	return string(body), nil
}

func getClient() (*cloudflare.Client, error) {
	// Load the env vars into the program
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	token := os.Getenv("CLOUDFLARE_ACCESS_TOKEN")
	if token == "" {
		return nil, errors.New("Error getting client: Token not set")
	}

	email := os.Getenv("CLOUDFLARE_USERNAME")
	if token == "" {
		return nil, errors.New("Error getting client: Email not set")
	}

	// Create and return cloudflare client
	client := cloudflare.NewClient(
		option.WithAPIKey(token),
		option.WithAPIEmail(email),
	)

	return client, nil
}
