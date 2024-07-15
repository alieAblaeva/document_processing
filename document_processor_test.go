package main

import (
	"log"
	"testing"

	document "github.com/alieAblaeva/document_processing/proto"
)

func TestMainFunction(t *testing.T) {
	var config MongoConfig
	config.getConf()

	processor, err := NewDocumentProcessor(&config)
	if err != nil {
		log.Printf("Failed to connect to mongo:\n%v", err)
	}

	if err != nil {
		log.Printf("Failed to connect to mongo:\n%v", err)
	}

	doc1 := &document.TDocument{
		Url:       "http://example.com",
		PubDate:   0,
		FetchTime: 2,
		Text:      "Initial text",
	}

	doc2 := &document.TDocument{
		Url:       "http://example.com",
		PubDate:   1,
		FetchTime: 3,
		Text:      "Updated text",
	}

	doc3 := &document.TDocument{
		Url:       "http://example.com",
		PubDate:   2,
		FetchTime: 1,
		Text:      "Earlier text",
	}

	_, err = processor.Process(doc1)
	if err != nil {
		t.Fatalf("Process failed: %v", err)
	}

	res, err := processor.Process(doc2)
	if err != nil {
		t.Fatalf("Process failed: %v", err)
	}

	if res.Text != "Updated text" {
		t.Errorf("Expected text to be 'Updated text', got '%s'", res.Text)
	}

	if res.PubDate != 0 {
		t.Errorf("Expected PubDate to be 0, got %d", res.PubDate)
	}

	res, err = processor.Process(doc3)
	if err != nil {
		t.Fatalf("Process failed: %v", err)
	}

	if res.FirstFetchTime != 1 {
		t.Errorf("Expected FirstFetchTime to be 1, got %d", res.FirstFetchTime)
	}

}
