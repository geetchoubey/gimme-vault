package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	jww "github.com/spf13/jwalterweatherman"
)

type AuthResponse struct {
	ClientToken   string `json:"client_token"`
	LeaseDuration int    `json:"lease_duration"`
}

type LoginResponse struct {
	Auth AuthResponse `json:"auth"`
}

func Login(url string, password string) (AuthResponse, error) {
	reqBody, err := json.Marshal(map[string]string{
		"password": password,
	})
	if err != nil {
		return AuthResponse{}, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil || resp.StatusCode != http.StatusOK {
		if err == nil {
			return AuthResponse{}, fmt.Errorf("got error when logging in. Code: %d", resp.StatusCode)
		} else {
			jww.CRITICAL.Fatalf("error while logging in %v \n", err)
		}
		return AuthResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AuthResponse{}, err
	}
	var response = LoginResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		if err != nil {
			return AuthResponse{}, err
		}
	}
	return response.Auth, nil
}
