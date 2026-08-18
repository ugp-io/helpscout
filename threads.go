package helpscout

import (
	"context"
)

type ThreadsServiceOp struct {
	client *Client
}

type ThreadsService interface {
	Browse(context.Context, string) (*HelpScoutThreadsResponse, error)
}

func (c *ThreadsServiceOp) Browse(ctx context.Context, url string) (*HelpScoutThreadsResponse, error) {

	var response HelpScoutThreadsResponse
	if err := c.client.Request("GET", url, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
