package main

import (
	"errors"
	"io"
	"net/http"
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
