package utils

import (
	"github.com/diegorufe/goutils/pkg/querybuilder/structs"
	"github.com/diegorufe/goutils/pkg/querybuilder/types"
)

func defaultFilter(name string, value any, filterType types.FilterType) structs.Filter {
	return structs.Filter{Operator: types.And, OpenBrackets: 1, CloseBrackets: 1, Name: name, Value: value, Type: filterType}
}

func Eq(columnName string, value any) structs.Filter {
	return defaultFilter(columnName, value, types.Eq)
}

func Like(columnName string, value any) structs.Filter {
	return defaultFilter(columnName, value, types.Like)
}

func LikeEnd(columnName string, value any) structs.Filter {
	return defaultFilter(columnName, value, types.LikeEnd)
}

func LikeStart(columnName string, value any) structs.Filter {
	return defaultFilter(columnName, value, types.LikeStart)
}

func Ge(columnName string, value any) structs.Filter {
	return defaultFilter(columnName, value, types.Ge)
}

func Gt(columnName string, value any) structs.Filter {
	return defaultFilter(columnName, value, types.Gt)
}

func Lt(columnName string, value any) structs.Filter {
	return defaultFilter(columnName, value, types.Lt)
}

func Le(columnName string, value any) structs.Filter {
	return defaultFilter(columnName, value, types.Le)
}

func In(columnName string, value any) structs.Filter {
	return defaultFilter(columnName, value, types.In)
}

func Ne(columnName string, value any) structs.Filter {
	return defaultFilter(columnName, value, types.Ne)
}

func NotIn(columnName string, value any) structs.Filter {
	return defaultFilter(columnName, value, types.NotIn)
}

func IsNull(columnName string) structs.Filter {
	return defaultFilter(columnName, nil, types.IsNull)
}

func IsNotNull(columnName string) structs.Filter {
	return defaultFilter(columnName, nil, types.IsNotNull)
}
