package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func ESQuery(q map[string]any) bytes.Buffer {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(q); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	return buf
}

func JSONDecoder[T any](r io.Reader, result *T) {
	if err := json.NewDecoder(r).Decode(result); err != nil {
		log.Fatalf("Error parsing response: %s", err)
	}
}

func main() {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "foo",
		Password: "bar",
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	query := map[string]any{
		"size": 1,
		"query": map[string]any{
			"match": map[string]any{
				"genre": "anime",
			},
		},
	}

	buf := ESQuery(query)

	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex("anime"),
		client.Search.WithBody(&buf),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error response: %s", res.String())
	}

	var result map[string]any
	JSONDecoder(res.Body, &result)

	fmt.Printf("Found %v hits\n", result["hits"].(map[string]any)["hits"])
}
