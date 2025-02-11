package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

	fmt.Println("URL:", c.baseURL+updateType)

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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &models.ControllerResponce{}, err
	}
	fmt.Println("Ответ с клиента:", string(body))
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		log.Printf("error decoding response body into models.ControllerResponce struct: %v\n", err)
		return nil, err
	}

	fmt.Println("Answer :", apiResp)

	return &apiResp, nil
}

func (c *Client) SendID(ctx context.Context, id int) (models.ControllerResponce, error) {

	url := c.baseURL + "messages/" + strconv.Itoa(id)

	fmt.Println(url)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return models.ControllerResponce{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.ControllerResponce{}, errors.New("api returned non-200 status code")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.ControllerResponce{}, err
	}
	fmt.Println("Ответ с клиента:", string(body))
	var apiResp models.ControllerResponce

	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		log.Printf("error decoding response body into models.ControllerResponce struct: %v\n", err)
		return models.ControllerResponce{}, err
	}

	fmt.Println("Отправляю в условие:", apiResp)

	return apiResp, nil
}
