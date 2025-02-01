package main

import (
	"encoding/json"
	"errors"
	"fmt"
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

func getRecord(Config config, recordID string, zoneID string) (CloudflareDNSResponse, error) {

	// Yes this uses the reqular HTTP instead of the SDK
	// I started getting a bunch of errors and stuff out of the SDK
	// even after taking verbatium from the docs, so here we are.
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", zoneID, recordID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CloudflareDNSResponse{}, err
	}

	req.Header.Set("X-Auth-Email", Config.Email)
	req.Header.Set("X-Auth-Key", Config.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return CloudflareDNSResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CloudflareDNSResponse{}, err
	}

	var response CloudflareDNSResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return CloudflareDNSResponse{}, err
	}

	return response, nil
}
