package utils

import (
	"github.com/diegorufe/goutils/pkg/querybuilder/structs"
	"github.com/diegorufe/goutils/pkg/querybuilder/types"
)

func defaultJoin(table string, alias string, condition string, joinType types.JoinType) structs.Join {
	return structs.Join{Table: table, Alias: alias, Condition: condition, Type: joinType}
}

func Inner(table string, alias string, condition string) structs.Join {
	return defaultJoin(table, alias, condition, types.Inner)
}

func Left(table string, alias string, condition string) structs.Join {
	return defaultJoin(table, alias, condition, types.Left)
}

func Right(table string, alias string, condition string) structs.Join {
	return defaultJoin(table, alias, condition, types.Right)
}
