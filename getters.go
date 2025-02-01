package main

import (
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/option"
)

func readConfig() (config, error) {
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return config{}, err
	}

	var Config config
	err = toml.Unmarshal(data, &Config)
	if err != nil {
		return config{}, err
	}

	return Config, nil
}

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

func getClient(Config config) (*cloudflare.Client, error) {
	// Create and return cloudflare client
	client := cloudflare.NewClient(
		option.WithAPIKey(Config.Token),
		option.WithAPIEmail(Config.Email),
	)

	return client, nil
}
