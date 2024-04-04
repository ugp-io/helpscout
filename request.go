package helpscout

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	APIKey        string
	APISecret     string
	Conversations ConversationsService
	Threads       ThreadsService
	Tags          TagsService
}

type TokenRefreshResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func NewClient(apiKey string, apiSecret string) *Client {

	c := &Client{
		APIKey:    apiKey,
		APISecret: apiSecret,
	}

	c.Conversations = &ConversationsServiceOp{client: c}
	c.Threads = &ThreadsServiceOp{client: c}
	c.Tags = &TagsServiceOp{client: c}

	return c

}

func (c *Client) TokenAccess(ctx context.Context) error {

	data := url.Values{}
	data.Set("client_id", c.APIKey)
	data.Set("client_secret", c.APISecret)
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(
		"POST",
		"https://api.helpscout.net/v2/oauth2/token",
		strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var token TokenRefreshResponse
	err = json.Unmarshal(body, &token)
	if err != nil {
		return err
	}

	accessCode = token.AccessToken

	return nil
}
