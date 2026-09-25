package helpscout

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

type ConversationsServiceOp struct {
	client *Client
}

type ConversationsService interface {
	BrowseConversations(context.Context, HelpScoutConversationRequest) (*HelpScoutConversationsResponse, error)
	UpdateConversationTag(context.Context, HelpScoutTagUpdate) error
	UpdateConversation(context.Context, HelpScoutConversationUpdate) error
	GetThreadAttachment(ctx context.Context, req HelpScoutGetAttachmentRequest) (*HelpScoutAttachmentResponse, error)
}

func (c *ConversationsServiceOp) BrowseConversations(ctx context.Context, req HelpScoutConversationRequest) (*HelpScoutConversationsResponse, error) {

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
		if req.Tags != nil && len(*req.Tags) > 0 {
			tagJoin := strings.Join(*req.Tags, ",")
			buildURL += "&tag=" + url.QueryEscape(tagJoin)
		}
		fullURL = fmt.Sprintf("%s?%v", conversationsURL, buildURL)
	} else {

		parsedURL, err := url.Parse(*req.URL)
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

		if req.Mailboxes != nil {
			encodedQuery.Set("mailbox", queryParams.Get("mailbox"))
		}
		if req.Status != nil {
			encodedQuery.Set("status", queryParams.Get("status"))
		}

		if req.Tags != nil && len(*req.Tags) > 0 {
			encodedQuery.Set("tag", queryParams.Get("tag"))
		}

		parsedURL.RawQuery = encodedQuery.Encode()
		fullURL = parsedURL.String()
	}

	var response HelpScoutConversationsResponse
	if err := c.client.Request("GET", fullURL, nil, &response); err != nil {
		return nil, err
	}

	if response.Links.Next.Href != nil {
		newLink := *response.Links.Next.Href
		newResponse, err := c.BrowseConversations(ctx, HelpScoutConversationRequest{URL: &newLink})
		if err != nil {
			return nil, err
		}
		response.Embedded.Conversations = append(response.Embedded.Conversations, newResponse.Embedded.Conversations...)
	}

	return &response, nil
}

func (c *ConversationsServiceOp) UpdateConversation(ctx context.Context, update HelpScoutConversationUpdate) error {

	var payload map[string]interface{}
	switch {
	case update.Status != nil:
		payload = map[string]interface{}{
			"op":    "replace",
			"path":  "/status",
			"value": *update.Status,
		}
	case update.MailboxID != nil:
		payload = map[string]interface{}{
			"op":    "move",
			"path":  "/mailboxId",
			"value": *update.MailboxID,
		}
	}
	return c.client.Request(
		"PATCH",
		fmt.Sprintf("%v/%v", conversationsURL, update.ConversationID),
		payload,
		nil)
}

func (c *ConversationsServiceOp) GetThreadAttachment(ctx context.Context, req HelpScoutGetAttachmentRequest) (*HelpScoutAttachmentResponse, error) {

	var response HelpScoutAttachmentResponse
	fullURL := fmt.Sprintf("%v/%v/attachments/%v/file",
		conversationsURL,
		req.ConversationID,
		req.AttachmentID)

	err := c.client.Request("GET", fullURL, nil, &response.Attachment)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &response, nil
}

func (c *ConversationsServiceOp) UpdateConversationTag(ctx context.Context, update HelpScoutTagUpdate) error {

	return c.client.Request(
		"PUT",
		fmt.Sprintf("%v/%v/tags", conversationsURL, update.ConversationID),
		map[string]interface{}{"tags": update.Tags},
		nil)
}
