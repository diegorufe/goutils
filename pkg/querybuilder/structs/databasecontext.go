package structs

import "github.com/diegorufe/goutils/pkg/querybuilder/types"

type DataBaseContext struct {
	Type types.BuilderType
	DB   any
}

func InstanceDBContext(builderType types.BuilderType, db any) *DataBaseContext {
	return &DataBaseContext{Type: builderType, DB: db}
}
