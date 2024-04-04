package helpscout

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type ConversationsServiceOp struct {
	client *Client
}

type ConversationsService interface {
	BrowseConversations(context.Context, string) (*HelpScoutConversationsResponse, error)
	UpdateTag(context.Context, HelpScoutTagUpdate) error
}

func (c *ConversationsServiceOp) BrowseConversations(ctx context.Context, fullURL string) (*HelpScoutConversationsResponse, error) {

	// Build initial URL or use given URL
	if fullURL == "" {
		escapedEmail := url.QueryEscape(`(email:(notifications@brandmanager360.com OR system@brandcomply.com OR designs@affinity-gateway.com))`)
		// escapedStatus := "status=all"
		// fullURL = fmt.Sprintf("%s?query=%s&%s", conversationsURL, escapedEmail, escapedStatus)
		fullURL = fmt.Sprintf("%s?query=%s", conversationsURL, escapedEmail)
	} else {

		parsedURL, err := url.Parse(fullURL)
		if err != nil {
			return nil, err
		}

		queryParams, err := url.ParseQuery(parsedURL.RawQuery)
		if err != nil {
			return nil, err
		}

		encodedQuery := url.Values{}
		encodedQuery.Set("query", queryParams.Get("query"))
		// encodedQuery.Set("status", queryParams.Get("status"))
		encodedQuery.Set("page", queryParams.Get("page"))

		parsedURL.RawQuery = encodedQuery.Encode()
		fullURL = parsedURL.String()
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessCode)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response HelpScoutConversationsResponse
	// var response interface{}
	decoder := json.NewDecoder(resp.Body)
	errDecode := decoder.Decode(&response)
	if errDecode != nil {
		return nil, errDecode
	}

	if response.Links.Next.Href != nil && response.Page.Number != nil && *response.Page.Number < 2 {
		newLink := *response.Links.Next.Href
		newResponse, err := c.BrowseConversations(ctx, newLink)
		if err != nil {
			return nil, err
		}
		response.Embedded.Conversations = append(response.Embedded.Conversations, newResponse.Embedded.Conversations...)
	}

	return &response, nil
}

func (c *ConversationsServiceOp) UpdateTag(ctx context.Context, update HelpScoutTagUpdate) error {

	fullURL := fmt.Sprintf("%v/%v/tags", conversationsURL, update.ConversationID)

	payload := update.Tags
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling payload:", err)
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Authorization", "Bearer "+accessCode)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	// var response HelpScoutThreadsResponse
	var response interface{}
	decoder := json.NewDecoder(resp.Body)
	errDecode := decoder.Decode(&response)
	if errDecode != nil {
		return errDecode
	}

	return nil
}
