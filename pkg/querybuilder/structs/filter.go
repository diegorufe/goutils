package structs

import "github.com/diegorufe/goutils/pkg/querybuilder/types"

type Filter struct {
	Type          types.FilterType
	Operator      types.FilterOperator
	Filters       []Filter
	OpenBrackets  int8
	CloseBrackets int8
	Value         any
	Name          string
}
