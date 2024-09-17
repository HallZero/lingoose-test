package llm

import (
	"context"
	"log"

	"github.com/henomis/lingoose/thread"
)

func GenerateLLMResponse(userInput string) string {
	
	myThread := thread.New()
	myThread.AddMessage(thread.NewUserMessage().AddContent(
		thread.NewTextContent(userInput),
	))

	ollama := InstantiateOllama()
	rag := InstantiateRag()

	a := InstantiateAssistant(ollama, rag)
		
	a.WithThread(
		thread.New().AddMessages(
			thread.NewUserMessage().AddContent(
				thread.NewTextContent(userInput),
			),
		),
	)

	err := a.Run(context.Background())

	if err != nil {
		log.Fatalf("Failed to generate LLM response: %s", err)
	}

	var llmResponse string

	lastMessage := a.Thread().LastMessage()

	if len(lastMessage.Contents) > 0 {
		for _, content := range lastMessage.Contents {
			if content.Type == thread.ContentTypeText {
				llmResponse = content.AsString() 
			}
		}
	}

	return llmResponse
}