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
	q.table = table
	q.Query.FinalQuery.WriteString(fmt.Sprintf("INSERT INTO %s", q.table))
	return q
}
func (q *QueryBuilder) Select(table string) *QueryBuilder {
	q.table = table
	q.Query.FinalQuery.WriteString("SELECT ")
	return q
}
func (q *QueryBuilder) Update(table string) *QueryBuilder {
	q.table = table
	q.Query.FinalQuery.WriteString(fmt.Sprintf("UPDATE %s ", q.table))
	return q
}

func (q *QueryBuilder) Set() *QueryBuilder {
	return q
}
func (q *QueryBuilder) Columns(cols ...string) *QueryBuilder {

	// SELECT *
	if strings.Contains(q.Query.FinalQuery.String(), "SELECT") && len(cols) == 0 {
		q.Query.FinalQuery.WriteString(fmt.Sprintf("* FROM %s ", q.table))
	} else if strings.Contains(q.Query.FinalQuery.String(), "SELECT") && len(cols) > 0 {
		// SELECT (id,name, email, dll)
		q.Query.FinalQuery.WriteString("(")
		for i, col := range cols {
			if i == len(cols)-1 {
				q.Query.FinalQuery.WriteString(fmt.Sprintf("%s", col))
				break
			}
			q.Query.FinalQuery.WriteString(fmt.Sprintf("%s,", col))
		}

		q.Query.FinalQuery.WriteString(")")
		q.totalValues = len(cols)
	}

	// Update
	if strings.Contains(q.Query.FinalQuery.String(), "UPDATE") {
		q.Query.FinalQuery.WriteString("SET ")
		for i, col := range cols {
			if i == len(cols)-1 {
				q.Query.FinalQuery.WriteString(fmt.Sprintf("%s=?", col))
				break
			}
			q.Query.FinalQuery.WriteString(fmt.Sprintf("%s=?,", col))
		}
	}

	return q

}
func (q *QueryBuilder) Where(cols ...string) *QueryBuilder {
	q.Query.FinalQuery.WriteString(" WHERE ")
	for _, col := range cols {
		q.Query.FinalQuery.WriteString(fmt.Sprintf("%s=?", col))
	}

	return q
}
func (q *QueryBuilder) Values() *QueryBuilder {
	q.Query.FinalQuery.WriteString(" VALUES(")

	for i := 0; i < q.totalValues; i++ {
		if i == q.totalValues-1 {
			q.Query.FinalQuery.WriteString("?")
		} else {
			q.Query.FinalQuery.WriteString("?,")
		}
	}
	q.Query.FinalQuery.WriteString(")")

	return q
}

func (q *QueryBuilder) And() *QueryBuilder {
	q.Query.FinalQuery.WriteString(" AND ")
	return q
}

func (q *QueryBuilder) Build() string {
	return q.Query.FinalQuery.String()
}
