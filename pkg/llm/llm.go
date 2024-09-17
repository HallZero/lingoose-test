package llm

import (
	"context"
	"log"

	"github.com/henomis/lingoose/llm/ollama"
	"github.com/henomis/lingoose/thread"
)


func GenerateLLMResponse(userInput string) string {
	
	myThread := thread.New()
	myThread.AddMessage(thread.NewUserMessage().AddContent(
		thread.NewTextContent(userInput),
	))

	var llmResponse string
	
	err := ollama.New().WithEndpoint("http://localhost:11434/api").WithModel("dolphin2.2-mistral").
		WithStream(func(s string) {
			llmResponse += s
		}).Generate(context.Background(), myThread)

	if err != nil {
		log.Fatalf("Failed to generate LLM response: %s", err)
	}

	return llmResponse
}