package structs

import "github.com/diegorufe/goutils/pkg/querybuilder/types"

type Order struct {
	Name string
	Type types.OrderType
}
