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

type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
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

	// Get access token and refresh token
	accessToken, refreshToken, err := getAccessTokenAndRefreshToken(code)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("accessToken: " + accessToken)
	fmt.Println("refreshToken: " + refreshToken)
}

func getAccessTokenAndRefreshToken(code string) (string, string, error) {
	const url = "https://login.live.com/oauth20_token.srf"

	// Define request payload
	data := map[string]string{
		"client_id": "3aa1025b-23c2-4a38-9faf-3d5f795fa58b",
		// "client_id":    "e06d23de-38f1-42a8-84bb-df59258ada2f",
		"redirect_uri": "http://localhost:8080",
		"secret_value": "VTq8Q~9a6q.bZsL1AqQIPnhQspKo25Qa2n2l6b2X",
		"grant_type":   "authorization_code",
		"code":         code,
	}

	// Marshal payload to JSON
	payload, err := json.Marshal(data)
	if err != nil {
		return "", "", err
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// Parse response JSON
	var accessTokenResponse AccessTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&accessTokenResponse)
	if err != nil {
		return "", "", err
	}

	// Extract access token and refresh token
	accessToken := accessTokenResponse.AccessToken
	refreshToken := accessTokenResponse.RefreshToken

	// Check status code
	if resp.StatusCode != 200 {
		fmt.Println(resp)
		return "", "", fmt.Errorf("Status code: %d", resp.StatusCode)
	}

	return accessToken, refreshToken, nil
}
