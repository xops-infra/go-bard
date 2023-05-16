package go_bard

import (
	"context"
	"net/http"
)

type ChatRequest struct {
	Input string `json:"input"`
}

type ChatResponse struct {
	Output string `json:"output"`
}

func (c *Client) CreateChat(
	ctx context.Context,
	request *ChatRequest,
) (response *ChatResponse, err error) {

	urlSuffix := "/chat"
	req, err := c.requestBuilder.build(ctx, http.MethodPost, c.fullURL(urlSuffix), request)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
