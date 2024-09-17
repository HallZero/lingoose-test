package llm

import "github.com/henomis/lingoose/llm/ollama"

func InstantiateOllama() (*ollama.Ollama){
	
	ollama := ollama.New().WithEndpoint("http://localhost:11434/api").WithModel("dolphin2.2-mistral")

	return ollama
}