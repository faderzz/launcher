package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthResponse struct {
	IssueInstant  string `json:"IssueInstant"`
	NotAfter      string `json:"NotAfter"`
	Token         string `json:"Token"`
	DisplayClaims struct {
		XUI []struct {
			UHS string `json:"uhs"`
		} `json:"xui"`
	} `json:"DisplayClaims"`
}

type AuthRequest struct {
	Properties struct {
		AuthMethod string `json:"AuthMethod"`
		SiteName   string `json:"SiteName"`
		RpsTicket  string `json:"RpsTicket"`
	} `json:"Properties"`
	RelyingParty string `json:"RelyingParty"`
	TokenType    string `json:"TokenType"`
}

type XSTSResponse struct {
	IssueInstant  string `json:"IssueInstant"`
	NotAfter      string `json:"NotAfter"`
	Token         string `json:"Token"`
	DisplayClaims struct {
		XUI []struct {
			UHS string `json:"uhs"`
		} `json:"xui"`
	} `json:"DisplayClaims"`
}

type XSTSRequest struct {
	Properties struct {
		SandboxID  string   `json:"SandboxId"`
		UserTokens []string `json:"UserTokens"`
	} `json:"Properties"`
	RelyingParty string `json:"RelyingParty"`
	TokenType    string `json:"TokenType"`
}

type AuthenticationError struct {
	Message string
	Err     error
}

func (e *AuthenticationError) Error() string {
	return e.Message
}

func authenticateWithXboxLive(accessToken string) (string, string, error) {
	var xstsToken string
	var userHashMC string

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
		return xstsToken, userHashMC, err
	}

	// Create a POST request
	req, err := http.NewRequest("POST", "https://user.auth.xboxlive.com/user/authenticate", bytes.NewBuffer(payload))
	if err != nil {
		return xstsToken, userHashMC, err
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return xstsToken, userHashMC, err
	}
	defer resp.Body.Close()

	// Handle the response
	if resp.StatusCode != http.StatusOK {
		return xstsToken, userHashMC, fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}

	// Parse the response body
	var authResponse AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		return xstsToken, userHashMC, err
	}

	// Store and print the Xbox Live token and user hash
	xblToken := authResponse.Token
	userHash := authResponse.DisplayClaims.XUI[0].UHS
	fmt.Println("Xbox Live Token:", xblToken)
	fmt.Println("User Hash:", userHash)

	// Call getXSTSToken function
	xstsToken, userHashMC, err = getXSTSToken(xblToken)
	if err != nil {
		return xstsToken, userHashMC, err
	}
	print(xstsToken)

	// Return the XSTS token and user hash and nil error
	return xstsToken, userHashMC, err
}

func getXSTSToken(xblToken string) (string, string, error) {
	// Create the request payload
	xstsRequest := XSTSRequest{
		Properties: struct {
			SandboxID  string   `json:"SandboxId"`
			UserTokens []string `json:"UserTokens"`
		}{
			SandboxID:  "RETAIL",
			UserTokens: []string{xblToken},
		},
		RelyingParty: "rp://api.minecraftservices.com/",
		TokenType:    "JWT",
	}

	// Convert the request payload to JSON
	payload, err := json.Marshal(xstsRequest)
	if err != nil {
		return "", "", err
	}

	// Create a POST request
	req, err := http.NewRequest("POST", "https://xsts.auth.xboxlive.com/xsts/authorize", bytes.NewBuffer(payload))
	if err != nil {
		return "", "", err
	}

	// Set the Content-Type and Accept headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// Handle the response
	if resp.StatusCode != http.StatusOK {
		// Check for error response
		if resp.StatusCode == http.StatusUnauthorized {
			var errorResponse struct {
				Identity string `json:"Identity"`
				XErr     int    `json:"XErr"`
				Message  string `json:"Message"`
				Redirect string `json:"Redirect"`
			}
			err = json.NewDecoder(resp.Body).Decode(&errorResponse)
			if err != nil {
				return "", "", err
			}
			return "", "", fmt.Errorf("Xbox Live authorization error: Identity: %s, XErr: %d, Message: %s, Redirect: %s",
				errorResponse.Identity, errorResponse.XErr, errorResponse.Message, errorResponse.Redirect)
		} else {
			return "", "", fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
		}
	}

	// Parse the response body
	var xstsResponse XSTSResponse
	err = json.NewDecoder(resp.Body).Decode(&xstsResponse)
	if err != nil {
		return "", "", err
	}

	// Return the XSTS token and user hash
	return xstsResponse.Token, xstsResponse.DisplayClaims.XUI[0].UHS, nil
}
