package query_utils

import "strings"

//TODO: @make Query Builder using Builder Pattern

type query struct {
	//Query
	FinalQuery *strings.Builder
}

type Value struct {
}

type QueryBuilder struct {
	Query *query
}