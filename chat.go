package go_bard

import (
	"context"
	"net/http"
)

type ChatRequest struct {
	Message string `json:"message"`
}

type UnitChoices struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Content           string        `json:"content"`
	ConversationID    string        `json:"conversation_id"`
	ResponseID        string        `json:"response_id"`
	FactualityQueries string        `json:"factuality_queries"`
	TextQuery         []any         `json:"text_query"`
	Choices           []UnitChoices `json:"choices"`
	Links             []any         `json:"links"`
	Images            []any         `json:"images"`
	Code              string        `json:"code"`
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
