package database

import (
	"fmt"
	"todo-list/configs"

	"github.com/elastic/go-elasticsearch/v8"
)

var Elastic *elasticsearch.Client

type Elasticsearch interface {
	// Config elasticsearch
	Config() *elasticsearch.Config

	// Connect to elasticsearch client
	Connect() error
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

func ElasticConnect() error {
	elastic := NewElasticsearch()

	err := elastic.Connect()

	if err != nil {
		return err
	}

	fmt.Println("Connect Elasticsearch Client Successfully")

	return nil
}
