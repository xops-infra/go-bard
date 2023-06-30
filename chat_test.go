package go_bard

import (
	"context"
	"testing"
)

func TestCreateChat(t *testing.T) {
	var client = NewClient("", "http://127.0.0.1:5000")

	var req = ChatRequest{
		Message: "Why did not I show up at my parents wedding?",
	}
	chat, err := client.CreateChat(context.Background(), &req)
	if err != nil {
		t.Error(err)
	}
	t.Log(chat)
}
