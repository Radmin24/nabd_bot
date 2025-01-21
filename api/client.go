package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"not_a_boring_date_bot/internal/models"
	"strconv"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string, timeout int) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

func (c *Client) SendCommand(ctx context.Context, data interface{}, updateType string) (*models.ControllerResponce, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+updateType, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api returned non-200 status code")
	}

	var apiResp models.ControllerResponce

	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		log.Printf("error decoding response body into models.ControllerResponce struct: %v\n", err)
		return nil, err
	}

	return &apiResp, nil
}

func (c *Client) SendID(ctx context.Context, id int) (*models.ControllerResponce, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL+"messages/", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("id", strconv.Itoa(id))
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// log.Println(resp)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("api returned non-200 status code")
	}

	var apiResp models.ControllerResponce

	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		log.Printf("error decoding response body into models.ControllerResponce struct: %v\n", err)
		return nil, err
	}

	return &apiResp, nil
}
