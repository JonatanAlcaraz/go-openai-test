package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	authToken := os.Getenv("OPENAI_AUTH_TOKEN")
	if authToken == "" {
		log.Fatal("Error: missing OPENAI_AUTH_TOKEN")
	}
	client := openai.NewClient(authToken)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo0125,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "you are a helpful chatbot",
			},
		},
	}
	fmt.Println("Conversation")
	fmt.Println("---------------------")
	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: s.Text(),
		})
		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}
		fmt.Printf("%s\n\n", resp.Choices[0].Message.Content)
		req.Messages = append(req.Messages, resp.Choices[0].Message)
		fmt.Print("> ")
	}
}
