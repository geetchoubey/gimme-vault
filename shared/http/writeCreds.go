package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type AWSCredentials struct {
	AccessKey     string `json:"access_key"`
	SecretKey     string `json:"secret_key"`
	SecurityToken string `json:"security_token"`
}

type Response struct {
	Data AWSCredentials `json:"data"`
}

func WriteCredentials(url string, token string) (AWSCredentials, error) {
	ttl, _ := json.Marshal(map[string]string{
		"ttl": "240m",
	})
	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(ttl))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Vault-Token", token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return AWSCredentials{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		if err != nil {
			return AWSCredentials{}, err
		}
	}
	var response = Response{}
	if err := json.Unmarshal(body, &response); err != nil {
		if err != nil {
			return AWSCredentials{}, err
		}
	}
	return response.Data, nil
}
