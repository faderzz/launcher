package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/browser"
)

func microsoftAuth() {
	// Define client credentials
	const ClientID = "" // Replace with your own client ID
	// Microsoft authentication endpoint with Xbox Live scope
	const AuthURL = "https://login.live.com/oauth20_authorize.srf?response_type=code&client_id=" + ClientID + "&redirect_uri=https://login.live.com/oauth20_desktop.srf&scope=XboxLive.signin%20offline_access"

	// Create a server
	server := http.NewServeMux()

	// Create a handler for the server
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Close the connection
		defer r.Body.Close()

		// Get the code from the URL
		code := r.URL.Query().Get("code")
		// Create a new POST request
		req, err := http.NewRequest("POST", "https://login.live.com/oauth20_token.srf", strings.NewReader("client_id="+ClientID+"&code="+code+"&grant_type=authorization_code&redirect_uri=https://login.live.com/oauth20_desktop.srf&scope=XboxLive.signin%20offline_access"))
		if err != nil {
			fmt.Println(err)
			return
		}
		// Set the headers
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		// Send the request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Close the connection
		defer resp.Body.Close()

		// Create a new decoder
		dec := json.NewDecoder(resp.Body)
		// Create a new map
		m := make(map[string]interface{})
		// Decode the JSON
		err = dec.Decode(&m)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Print the access token
		fmt.Println("Access token:", m["access_token"])
		// Print the refresh token
		fmt.Println("Refresh token:", m["refresh_token"])
		// Print the user hash
		fmt.Println("User hash:", m["user_hash"])

		// Store access token in variable
		accessToken := m["access_token"].(string)
		// Run function authenticateWithXboxLive in func/getXLSXToken.go
		authenticateWithXboxLive(accessToken)
	})

	// Open the browser
	err := browser.OpenURL(AuthURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen on port 8080
	fmt.Println("Listening on port 8080")
	err = http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Println(err)
		return
	}
}
