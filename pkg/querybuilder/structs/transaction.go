package structs

import "github.com/diegorufe/goutils/pkg/querybuilder/types"

type Transaction struct {
	Type               types.BuilderType
	TransactionContext any
}
