package utils

import "github.com/diegorufe/goutils/pkg/querybuilder/structs"

func Select(name string, alias string) structs.Field {
	return structs.Field{Name: name, Alias: alias}
}
