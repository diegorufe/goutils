package structs

import "github.com/diegorufe/goutils/pkg/querybuilder/types"

type Join struct {
	Type      types.JoinType
	Table     string
	Alias     string
	Condition string
}
