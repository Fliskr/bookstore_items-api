package services

import (
	"github.com/gervi/bookstore_items-api/src/domain/items"
	"github.com/gervi/bookstore_items-api/src/domain/queries"
	"github.com/gervi/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	Update(item items.Item) (*items.Item, rest_errors.RestErr)
	Search(q queries.EsQuery) (*items.Items, rest_errors.RestErr)
	Delete(string) rest_errors.RestErr
}

type itemsService struct {
}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}
func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{Id: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Search(q queries.EsQuery) (*items.Items, rest_errors.RestErr) {
	var items items.Items
	if err := items.Search(q); err != nil {
		return nil, err
	}
	return &items, nil
}

func (s *itemsService) Delete(id string) rest_errors.RestErr {
	item := items.Item{Id: id}
	if err := item.Delete(); err != nil {
		return err
	}
	return nil
}

func (s *itemsService) Update(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Update(); err != nil {
		return nil, err
	}
	return &item, nil
}
