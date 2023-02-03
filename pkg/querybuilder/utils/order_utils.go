package utils

import (
	"github.com/diegorufe/goutils/pkg/querybuilder/structs"
	"github.com/diegorufe/goutils/pkg/querybuilder/types"
)

func defaultOrder(name string, orderType types.OrderType) structs.Order {
	return structs.Order{Name: name, Type: orderType}
}

func Asc(columnaName string) structs.Order {
	return defaultOrder(columnaName, types.Asc)
}

func Desc(columnaName string) structs.Order {
	return defaultOrder(columnaName, types.Desc)
}
