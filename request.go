package helpscout

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	APIKey      string
	APISecret   string
	AccessToken *string
	ExpiresIn   *int

	Conversations ConversationsService
	Threads       ThreadsService
	Tags          TagsService
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

func (c *Client) Request(method string, url string, body interface{}, v interface{}) error {
	// c.AccessToken = new(string)
	// *c.AccessToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ik1qaEJORVF3TkRaQk56ZENORFV3TmpjeVJUSkZOMFU1TTBZMk5EYzRPVFUyUVRjMlFUSkVOZyJ9.eyJodHRwOi8vc3BzL2VudiI6InByb2QiLCJodHRwOi8vc3BzL2lzLXRyaWFsIjpmYWxzZSwiaHR0cDovL3Nwcy9vcmctaWQiOiIzMDU1MzgzNjkyMjg1MDE4MjIyNTMxMTIxODg2MzM4Mzk5NjU0OTUiLCJodHRwOi8vc3BzL29yZy1uYW1lIjoiVW5kZXJncm91bmQgUHJpbnRpbmciLCJodHRwOi8vc3BzL293bmVyLW9yZy1pZCI6IjMwNTUzODM2OTIyODUwMTgyMjI1MzExMjE4ODYzMzgzOTk2NTQ5NSIsImh0dHA6Ly9zcHMvdGVzdCI6Im5vIiwiaHR0cDovL3Nwcy90b2tpZCI6Ikl5T2tHdTNzS3dsNDRPd0wiLCJpc3MiOiJodHRwczovL2F1dGguc3BzY29tbWVyY2UuY29tLyIsInN1YiI6InhhQ1l6TGtRb0s4QU5GQkN4dVVHWHB1ajF0Q3dQYUkxQGNsaWVudHMiLCJhdWQiOiJodHRwczovL3Nwc2NvbW1lcmNlLmNvbSIsImlhdCI6MTc4MjMzMTg4MCwiZXhwIjoxNzgyMzM1NDgwLCJndHkiOiJjbGllbnQtY3JlZGVudGlhbHMiLCJhenAiOiJ4YUNZekxrUW9LOEFORkJDeHVVR1hwdWoxdEN3UGFJMSJ9.A1O65wDE-ujjVhLXxMS4y23U44iIavu75PlUEUmLTf7bT2_u8KzseARme7lS9JQLlQ9-ncTENBfb1v6vKQHg7AReZZbT-42lg16UdARmxfI1qTka9S1stR_x6ebVdyV9-BTH60ytpS4Aiesp6KbdsQlS9MNYDdM6kV2f2W-5sW9sxeIDxcOBK0KBSjObKPH1U_AshE0ytxEBTDx6zrU8TnPveLSQGEBs_1Sk8g_8K4VciKwXj0qxx6MLZK8pCKdXOo5izfFeWZscqNdeDYURkjgQSCicB5L8a-G7IfLL2RzOp1HQJIZNUFrI4aSlYAlY5WkEI4SjR1dSiOAvWA71iA"
	if c.AccessToken == nil || c.ExpiresIn == nil || *c.ExpiresIn <= 30 {
		if err := c.GetAccessToken(); err != nil {
			return fmt.Errorf("failed to get access token: %w", err)
		}
	}

	var bodyReader io.Reader
	if body != nil {
		requestJson, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewBuffer(requestJson)
	}

	httpReq, errNewRequest := http.NewRequest(method, url, bodyReader)
	if errNewRequest != nil {
		return errNewRequest
	}
	httpReq.Header.Set("Authorization", "Bearer "+*c.AccessToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	// fmt.Println("Resp:", string(respBody))
	// fmt.Println(url)
	decoder := json.NewDecoder(bytes.NewReader(respBody))
	errDecode := decoder.Decode(&v)
	if errDecode != nil {
		return errDecode
	}

	return nil
}

// func (c *Client) TokenAccess(ctx context.Context) error {
func (c *Client) GetAccessToken() error {

	body := map[string]string{
		"client_id":     c.APIKey,
		"client_secret": c.APISecret,
		"grant_type":    "client_credentials",
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest(
		"POST",
		"https://api.helpscout.net/v2/oauth2/token",
		bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("token request returned %d: %s", resp.StatusCode, string(respBody))
	}

	tokenResp := make(map[string]interface{})
	if err := json.Unmarshal(respBody, &tokenResp); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if accessToken, ok := tokenResp["access_token"].(string); ok {
		fmt.Println(accessToken)
		expiration := tokenResp["expires_in"].(int)
		c.ExpiresIn = &expiration
		c.AccessToken = &accessToken
	}

	return nil
}
