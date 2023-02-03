package utils

import "github.com/diegorufe/goutils/pkg/querybuilder/structs"

func Group(name string) structs.Group {
	return structs.Group{Name: name}
}
