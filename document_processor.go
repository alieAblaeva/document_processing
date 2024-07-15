package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	document "github.com/alieAblaeva/document_processing/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Processor interface {
	Process(d *document.TDocument) (*document.TDocument, error)
}

type DocumentProcessor struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewDocumentProcessor(config *MongoConfig) (*DocumentProcessor, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", config.User, config.Pass, config.Host, config.Port)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	collection := client.Database(config.DB).Collection(config.Collect)
	return &DocumentProcessor{client: client, collection: collection}, nil
}

func (db *DocumentProcessor) Process(d *document.TDocument) (*document.TDocument, error) {
	if d == nil {
		return nil, errors.New("document not found")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var existingDoc document.TDocument
	filter := bson.M{"url": d.Url}
	err := db.collection.FindOne(ctx, filter).Decode(&existingDoc)

	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	if err == mongo.ErrNoDocuments {
		fmt.Println(d.Url, ": First time fetched")
		d.FirstFetchTime = d.FetchTime
		_, err = db.collection.InsertOne(ctx, d)
		if err != nil {
			return nil, err
		}
		return d, nil
	}

	fmt.Println(d.Url, ": duplicate")

	if d.FetchTime > existingDoc.FetchTime {
		existingDoc.Text = d.Text
		existingDoc.FetchTime = d.FetchTime
	}

	if d.FetchTime < existingDoc.FetchTime {
		existingDoc.PubDate = d.PubDate
	}

	if d.FetchTime < existingDoc.FirstFetchTime {
		existingDoc.FirstFetchTime = d.FetchTime
	}

	_, err = db.collection.ReplaceOne(ctx, filter, existingDoc)
	if err != nil {
		return nil, err
	}

	return &existingDoc, nil
}
