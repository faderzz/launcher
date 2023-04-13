package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthRequest struct {
	Properties struct {
		AuthMethod string `json:"AuthMethod"`
		SiteName   string `json:"SiteName"`
		RpsTicket  string `json:"RpsTicket"`
	} `json:"Properties"`
	RelyingParty string `json:"RelyingParty"`
	TokenType    string `json:"TokenType"`
}

func authenticateWithXboxLive(accessToken string) error {
	// Create the request payload
	authRequest := AuthRequest{
		Properties: struct {
			AuthMethod string `json:"AuthMethod"`
			SiteName   string `json:"SiteName"`
			RpsTicket  string `json:"RpsTicket"`
		}{
			AuthMethod: "RPS",
			SiteName:   "user.auth.xboxlive.com",
			RpsTicket:  accessToken,
		},
		RelyingParty: "http://auth.xboxlive.com",
		TokenType:    "JWT",
	}

	// Convert the request payload to JSON
	payload, err := json.Marshal(authRequest)
	if err != nil {
		return err
	}

	// Create a POST request
	req, err := http.NewRequest("POST", "https://user.auth.xboxlive.com/user/authenticate", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle the response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}

	// TODO: Parse and use the response body as needed

	return nil
}
