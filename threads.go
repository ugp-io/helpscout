package helpscout

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ThreadsServiceOp struct {
	client *Client
}

type ThreadsService interface {
	Browse(context.Context, string) (*HelpScoutThreadsResponse, error)
}

func (c *ThreadsServiceOp) Browse(ctx context.Context, url string) (*HelpScoutThreadsResponse, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessCode)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	var response HelpScoutThreadsResponse
	// var response interface{}
	decoder := json.NewDecoder(resp.Body)
	errDecode := decoder.Decode(&response)
	if errDecode != nil {
		return nil, errDecode
	}

	return &response, nil
}
