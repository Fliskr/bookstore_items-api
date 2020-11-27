package items

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gervi/bookstore_items-api/src/domain/queries"

	"github.com/gervi/bookstore_items-api/src/clients/elastic_client"
	"github.com/gervi/bookstore_utils-go/rest_errors"
)

const (
	itemIndex = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elastic_client.Client.Index(itemIndex, i)
	if err != nil {
		return rest_errors.NewDbError(err.Error())
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemId := i.Id
	result, err := elastic_client.Client.Get(itemIndex, itemId)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("Id %s not found", itemId))
		}
		return rest_errors.NewDbError(err.Error())
	}
	if err := json.Unmarshal(result.Source, &i); err != nil {
		return rest_errors.NewNotFoundError(fmt.Sprintf("Id %s not found", itemId))
	}
	i.Id = itemId
	return nil
}

func (i *Items) Search(q queries.EsQuery) rest_errors.RestErr {
	result, err := elastic_client.Client.Search(itemIndex, q.Build())
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError("Not founc")
		}
		return rest_errors.NewDbError(err.Error())
	}
	var newItems = make(Items, 0)
	for _, res := range result.Hits.Hits {
		var item Item
		if err := json.Unmarshal(res.Source, &item); err != nil {
			return rest_errors.NewNotFoundError("Invalid json body")
		}
		item.Id = res.Id
		newItems = append(newItems, item)
	}
	*i = newItems
	return nil
}

func (i *Item) Delete() rest_errors.RestErr {
	result, err := elastic_client.Client.Delete(itemIndex, i.Id)
	if err != nil {
		return rest_errors.NewDbError(err.Error())
	}
	fmt.Println(result.Result)
	return nil
}

func (i *Item) Update() rest_errors.RestErr {
	result, err := elastic_client.Client.Update(itemIndex, i.Id, i)
	if err != nil {
		return rest_errors.NewDbError(err.Error())
	}
	i.Id = result.Id
	return nil
}
