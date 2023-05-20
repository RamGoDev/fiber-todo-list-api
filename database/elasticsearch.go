package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"todo-list/app/indices"
	"todo-list/configs"
	"todo-list/helpers"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var Elastic *elasticsearch.Client

type Elasticsearch interface {
	// Config elasticsearch
	Config() *elasticsearch.Config

	// Connect to elasticsearch client
	Connect() error

	// Init all index
	InitAllIndex() error

	// Create Index
	CreateIndex(indexName string) error

	// Delete Index
	DeleteIndex(indexName []string) error

	// Search all
	SearchAll(indexName string, search string, pagination helpers.Pagination) (map[string]interface{}, error)

	// Add document
	AddDocument(indexName string, data []byte) error

	// Update document
	UpdateDocument(indexName string, id string, data []byte) error

	// Delete document
	DeleteDocument(indexName string, id string) error
}

type elasticsearchImpl struct {
	//
}

func NewElasticsearch() Elasticsearch {
	return &elasticsearchImpl{}
}

func (impl elasticsearchImpl) Config() *elasticsearch.Config {
	return &elasticsearch.Config{
		Addresses: []string{
			configs.GetEnv("ELASTICSEARCH_HOST") + ":" + configs.GetEnv("ELASTICSEARCH_PORT"),
		},
		Username: configs.GetEnv("ELASTICSEARCH_USERNAME"),
		Password: configs.GetEnv("ELASTICSEARCH_PASSWORD"),
	}
}

func (impl elasticsearchImpl) Connect() error {
	var err error
	Elastic, err = elasticsearch.NewClient(*impl.Config())

	if err != nil {
		return err
	}

	// Test get response info
	_, err = Elastic.Info()
	if err != nil {
		return err
	}

	return nil
}

func (impl elasticsearchImpl) InitAllIndex() error {
	var err error
	var user indices.User
	var todo indices.Todo

	err = impl.CreateIndex(user.IndexName())
	if err != nil {
		return err
	}
	impl.CreateIndex(todo.IndexName())
	if err != nil {
		return err
	}

	return nil
}

func (impl elasticsearchImpl) CreateIndex(indexName string) error {
	_, err := Elastic.Indices.Create(indexName)

	if err != nil {
		return err
	}

	return nil
}

func (impl elasticsearchImpl) DeleteIndex(indexName []string) error {
	// Set up the request object.
	req := esapi.IndicesDeleteRequest{
		Index: indexName,
	}

	// Perform the request with the client.
	_, err := req.Do(ctx, Elastic)

	if err != nil {
		return err
	}

	return nil
}

func (impl elasticsearchImpl) SearchAll(indexName string, search string, pagination helpers.Pagination) (map[string]interface{}, error) {
	var err error
	var res *esapi.Response
	var buf bytes.Buffer
	var resp map[string]interface{}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query": search,
			},
		},
		"size": pagination.Limit,
		"from": pagination.Limit * pagination.Page,
	}

	if err = json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err = Elastic.Search(
		Elastic.Search.WithContext(ctx),
		Elastic.Search.WithIndex(indexName),
		Elastic.Search.WithBody(&buf),
		Elastic.Search.WithTrackTotalHits(true),
		Elastic.Search.WithPretty(),
	)

	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(bytes))

	resp, err = helpers.ByteToMapStringInterface(bytes)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (impl elasticsearchImpl) AddDocument(indexName string, data []byte) error {
	dataJson, _ := helpers.ByteToMapStringInterface(data)
	documentId := fmt.Sprintf("%.0f", dataJson["id"])
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: documentId,
		Body:       bytes.NewBuffer(data),
		Refresh:    "true",
	}

	// Perform the request with the client.
	_, err := req.Do(ctx, Elastic)
	if err != nil {
		return err
	}

	return nil
}

func (impl elasticsearchImpl) UpdateDocument(indexName string, id string, data []byte) error {
	// Delete old document
	errDeleted := impl.DeleteDocument(indexName, id)
	if errDeleted != nil {
		return errDeleted
	}

	// Create or Replace with new document
	errCreated := impl.AddDocument(indexName, data)
	if errCreated != nil {
		return errCreated
	}

	return nil
}

func (impl elasticsearchImpl) DeleteDocument(indexName string, id string) error {
	// Set up the request object.
	req := esapi.DeleteRequest{
		Index:      indexName,
		DocumentID: id,
	}

	// Perform the request with the client.
	_, err := req.Do(ctx, Elastic)
	if err != nil {
		return err
	}

	return nil
}

func ElasticConnect() error {
	elastic := NewElasticsearch()

	err := elastic.Connect()

	if err != nil {
		return err
	}

	fmt.Println("Connect Elasticsearch Client Successfully")

	err = elastic.InitAllIndex()
	if err != nil {
		return err
	}

	return nil
}
