package main

import (
	"errors"
	"io"
	"net/http"
	"os"

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

func getAccessToken() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	token := os.Getenv("CLOUDFLARE_ACCESS_TOKEN")
	if token == "" {
		return "", errors.New("Error getting access token: Token not set")
	}

	return token, nil
}
