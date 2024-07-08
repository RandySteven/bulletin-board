package queries

import "fmt"

type (
	AggregateQuery struct {
		Field     string
		Operation AggregateOperation
	}

	AggregateOperation struct {
		Operational string
		Value       any
	}

	PaginationQuery struct {
		Limit  int
		Offset int
	}
)

func (query AggregateQuery) ToString() string {
	queryStr := fmt.Sprintf("%s %s %s ", query.Field, query.Operation.Operational, query.Operation.Value)
	return queryStr
}

func (query PaginationQuery) ToString() string {
	queryStr := fmt.Sprintf("limit %d offset %d", query.Limit, query.Offset)
	return queryStr
}
