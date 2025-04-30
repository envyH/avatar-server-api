package service

import (
	"avatar/config"
	"context"
	"fmt"
	"log"

	"google.golang.org/genai"
)

func GenerateText(userID int32, question string) (string, error) {
	cfg := config.LoadConfig()

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  cfg.APIKey_GOOGLE,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	stream := client.Models.GenerateContentStream(
		ctx,
		"gemini-2.0-flash",
		genai.Text(question),
		nil,
	)

	for chunk := range stream {
		part := chunk.Candidates[0].Content.Parts[0]
		fmt.Print(part.Text)
	}

	return "send success!!", nil
}
