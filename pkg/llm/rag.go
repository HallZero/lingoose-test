package llm

import (
	"context"
	"os"
	"log"

	ollamaembedder "github.com/henomis/lingoose/embedder/ollama"
	"github.com/henomis/lingoose/index"
	"github.com/henomis/lingoose/index/vectordb/jsondb"
	"github.com/henomis/lingoose/rag"
)

func InstantiateRag() (*rag.RAG) {
	
	r := rag.New(
		index.New(
			jsondb.New().WithPersist("db.json"),
			ollamaembedder.New().WithModel("dolphin2.2-mistral"),
		),
	).WithTopK(3)

	_, err := os.Stat("db.json")
	if os.IsNotExist(err) {
		err = r.AddSources(context.Background(), "document.txt")
		if err != nil {
			log.Fatalf("Error creating RAG: %v", err)
		}
	}

	return r
}