package go_bard

import (
	"context"
	"testing"
)

func TestCreateChat(t *testing.T) {
	var client = NewClient("xx")

	var req = ChatRequest{
		Input: "Why did not I show up at my parents wedding?",
	}
	chat, err := client.CreateChat(context.Background(), &req)
	if err != nil {
		t.Error(err)
	}
	t.Log(chat)
}
