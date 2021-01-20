package query_utils

import (
	"fmt"
	"strings"
)

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		Query: &query{
			FinalQuery: &strings.Builder{},
		},
	}
}

func (q *QueryBuilder) Insert(table string) *QueryBuilder {
	q.Query.FinalQuery.WriteString(fmt.Sprintf("INSERT INTO %s", table))
	return q
}

func (q *QueryBuilder) Columns(cols ...string) *QueryBuilder {
	q.Query.FinalQuery.WriteString("(")
	for i, col := range cols {
		if i == len(cols)-1 {
			q.Query.FinalQuery.WriteString(fmt.Sprintf("%s", col))
			break
		}
		q.Query.FinalQuery.WriteString(fmt.Sprintf("%s,", col))
	}

	q.Query.FinalQuery.WriteString(")")

	q.Query.FinalQuery.WriteString(" VALUES(")

	for index, _ := range cols {
		if index == len(cols)-1 {
			q.Query.FinalQuery.WriteString("?")
		} else {
			q.Query.FinalQuery.WriteString("?,")
		}
	}
	q.Query.FinalQuery.WriteString(")")

	// TODO: Implement with Channel

	return q
}

func (q *QueryBuilder) Get() *QueryBuilder {
	//TODO Write string builder SELECT
	return q
}

func (q *QueryBuilder) Build() string {
	return q.Query.FinalQuery.String()
}
