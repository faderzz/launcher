package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/browser"
)

type LoginWithXboxRequest struct {
	IdentityToken string `json:"identityToken"`
}

type LoginWithXboxResponse struct {
	Username    string   `json:"username"`
	Roles       []string `json:"roles"`
	AccessToken string   `json:"access_token"`
	TokenType   string   `json:"token_type"`
	ExpiresIn   int      `json:"expires_in"`
}

func microsoftAuth() {
	// Define client credentials
	const ClientID = "3aa1025b-23c2-4a38-9faf-3d5f795fa58b"         // Replace with your own client ID
	const ClientSecret = "VTq8Q~9a6q.bZsL1AqQIPnhQspKo25Qa2n2l6b2X" // Replace with your own client secret
	var code string

	// Redirect for server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the code from the request
		code = r.URL.Query().Get("code")
		fmt.Println("code: " + code)
	})
	// Start the server
	go http.ListenAndServe(":8080", nil)

	// Get code flow code
	// Microsoft authentication endpoint with Xbox Live scope
	const AuthURL = "https://login.live.com/oauth20_authorize.srf?client_id=" + ClientID + "&scope=XboxLive.signin%20offline_access&redirect_uri=http://localhost:8080&response_type=code"
	// Open the browser
	err := browser.OpenURL(AuthURL)
	if err != nil {
		// fmt.Println(err)
		return
	}

	// Get access token
	const url = "https://login.live.com/oauth20_token.srf"
	// Data
	data := "client_id=" + ClientID + "&redirect_uri=http://localhost:8080&client_secret=" + ClientSecret + "&grant_type=authorization_code&code=" + code

	// Create a POST request with url, headers and data
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return
	}
	// Set the headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// Print the response
	fmt.Println(resp)
	// Check if the response is OK
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: " + resp.Status)
		return
	}
	// Decode the response
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}
	// Print the access token
	fmt.Println("Access token: " + result["access_token"].(string))
	// Print the refresh token
	fmt.Println("Refresh token: " + result["refresh_token"].(string))
}
