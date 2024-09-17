package llm

import (
	"github.com/henomis/lingoose/assistant"
	"github.com/henomis/lingoose/llm/ollama"
	"github.com/henomis/lingoose/rag"
)

func InstantiateAssistant(llm *ollama.Ollama, r *rag.RAG) (*assistant.Assistant) {
	a := assistant.New(
		llm,
	).WithParameters(
		assistant.Parameters{
			AssistantName:      "FAQ CS Assistant",
			AssistantIdentity:  "Assistant to help workers at Conta Simples with Frequent Asked Questions",
			AssistantScope:     "Responding in Portuguese as a halpful attendant",
			CompanyName:        "Filipi Kikuchi",
			CompanyDescription: "Ele é foda!",
		},
	).WithRAG(r)

	return a
}