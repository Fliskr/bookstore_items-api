package queries

import (
	"fmt"

	"github.com/olivere/elastic/v7"
)

func (q EsQuery) Build() elastic.Query {
	query := elastic.NewBoolQuery()

	var esQueries = make([]elastic.Query, 0)
	for _, f := range q.Equals {
		esQueries = append(esQueries, elastic.NewMatchQuery(f.Field, f.Value))
	}
	query.Must(esQueries...)
	fmt.Println("QUERY", esQueries)
	return query
}
