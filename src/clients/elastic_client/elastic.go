package elastic_client

import (
	"context"
	"fmt"
	"time"

	"github.com/gervi/bookstore_utils-go/logger"
	"github.com/olivere/elastic/v7"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
	Delete(string, string) (*elastic.DeleteResponse, error)
	Update(string, string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
	fmt.Println("Elastic connected OK")
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, i interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().Index(index).BodyJson(i).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error trying index %s in es", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Get(index, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().Index(index).Id(id).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error trying to get index %s in es", id), err)
		return nil, err
	}

	return result, nil
}

func (c *esClient) Search(index string, q elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := c.client.Search(index).Query(q).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error trying to find in %s", index), err)
		return nil, err
	}

	return result, nil
}

func (c *esClient) Delete(index string, id string) (*elastic.DeleteResponse, error) {
	ctx := context.Background()
	result, err := c.client.Delete().Index(index).Id(id).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error trying to find in %s", index), err)
		return nil, err
	}

	return result, nil
}

func (c *esClient) Update(index string, id string, body interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().Index(index).Id(id).BodyJson(body).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error trying update %s in es", index), err)
		return nil, err
	}
	return result, nil
}
