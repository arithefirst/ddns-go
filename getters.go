package main

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/option"
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

func getZones() ([]string, error) {
	zones := os.Getenv("ZONES")
	if zones == "" {
		return nil, errors.New("Error getting zones: Not set")
	}

	return strings.Split(zones, ","), nil
}
