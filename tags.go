package helpscout

import (
	"context"
	"encoding/json"
	"net/http"
)

type TagsServiceOp struct {
	client *Client
}

type TagsService interface {
	BrowseTags(context.Context) (*HelpScoutTagsResponse, error)
}

func (c *TagsServiceOp) BrowseTags(ctx context.Context) (*HelpScoutTagsResponse, error) {

	client := &http.Client{}
	reqhttp, err := http.NewRequest("GET", tagsURL, nil)
	if err != nil {
		return nil, err
	}
	reqhttp.Header.Add("Authorization", "Bearer "+accessCode)

	resp, err := client.Do(reqhttp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response HelpScoutTagsResponse
	decoder := json.NewDecoder(resp.Body)
	errDecode := decoder.Decode(&response)
	if errDecode != nil {
		return nil, errDecode
	}

	return &response, nil
}
