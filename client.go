package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	ApplicationJsonType = "application/json"
)

type client struct {
	BaseURL    string
	httpClient http.Client
}

type SetPasswordRequest struct {
	Password string `json:"password"`
}

type Response struct {
	Status int    `json:"-"`
	Link   string `json:"link"`
	Ttl    int    `json:"ttl"`
}

func (r *Response) IsSuccessful() bool {
	return r.Status == http.StatusOK
}

func NewClient(baseURL string) *client {
	return &client{BaseURL: baseURL, httpClient: http.Client{}}
}

func (c *client) SetPassword(password string) (*Response, error) {
	url := fmt.Sprintf("%s/set_password", c.BaseURL)
	request := SetPasswordRequest{Password: password}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error while marshaling set password body > %v", err)
	}

	resp, err := c.httpClient.Post(url, ApplicationJsonType, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("error making http request to %s > %v", url, err)
	}

	var response = &Response{Status: resp.StatusCode}
	responseBody, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if response.Status != http.StatusOK {
		return response, fmt.Errorf("error server returned: %s", resp.Status)
	}

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return response, fmt.Errorf("error unmarshaling json from body: %v", err)
	}

	return response, nil
}
