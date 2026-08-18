package helpscout

import (
	"context"
)

type TagsServiceOp struct {
	client *Client
}

type TagsService interface {
	BrowseTags(context.Context) (*HelpScoutTagsResponse, error)
}

func (c *TagsServiceOp) BrowseTags(ctx context.Context) (*HelpScoutTagsResponse, error) {

	var response HelpScoutTagsResponse
	if err := c.client.Request("GET", tagsURL, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
