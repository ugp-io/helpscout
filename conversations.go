package helpscout

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type ConversationsServiceOp struct {
	client *Client
}

type ConversationsService interface {
	BrowseConversations(context.Context, HelpScoutConversationRequest) (*HelpScoutConversationsResponse, error)
	UpdateTag(context.Context, HelpScoutTagUpdate) error
}

func (c *ConversationsServiceOp) BrowseConversations(ctx context.Context, req HelpScoutConversationRequest) (*HelpScoutConversationsResponse, error) {

	// Build initial URL or use given URL
	var fullURL string
	if req.URL == nil {

		var buildURL string
		if req.Emails != nil {
			emailJoin := strings.Join(*req.Emails, " OR ")
			buildURL = "query=" + url.QueryEscape(`(email:(`+emailJoin+`))`)
		}

		if req.Status != nil {
			buildURL += "&status=" + *req.Status
		}
		if req.Mailboxes != nil {
			mailboxJoin := strings.Join(*req.Mailboxes, ",")
			buildURL += "&mailbox=" + mailboxJoin
		}
		fullURL = fmt.Sprintf("%s?%v", conversationsURL, buildURL)
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
		encodedQuery.Set("page", queryParams.Get("page"))
		encodedQuery.Set("mailbox", queryParams.Get("mailbox"))

		parsedURL.RawQuery = encodedQuery.Encode()
		fullURL = parsedURL.String()
	}

	fmt.Println(fullURL)
	client := &http.Client{}
	reqhttp, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	reqhttp.Header.Add("Authorization", "Bearer "+accessCode)

	resp, err := client.Do(reqhttp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response HelpScoutConversationsResponse
	decoder := json.NewDecoder(resp.Body)
	errDecode := decoder.Decode(&response)
	if errDecode != nil {
		return nil, errDecode
	}

	if response.Links.Next.Href != nil && response.Page.Number != nil && *response.Page.Number < 2 {
		newLink := *response.Links.Next.Href
		newResponse, err := c.BrowseConversations(ctx, HelpScoutConversationRequest{URL: &newLink})
		if err != nil {
			return nil, err
		}
		response.Embedded.Conversations = append(response.Embedded.Conversations, newResponse.Embedded.Conversations...)
	}

	return &response, nil
}

func (c *ConversationsServiceOp) UpdateTag(ctx context.Context, update HelpScoutTagUpdate) error {

	fullURL := fmt.Sprintf("%v/%v/tags", conversationsURL, update.ConversationID)

	payload := map[string]interface{}{"tags": update.Tags}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+accessCode)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var response interface{}
	decoder := json.NewDecoder(resp.Body)
	errDecode := decoder.Decode(&response)
	if errDecode != nil {
		return errDecode
	}

	return nil
}
