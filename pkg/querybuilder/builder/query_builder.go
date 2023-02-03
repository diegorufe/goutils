package builder

import (
	"github.com/diegorufe/goutils/pkg/querybuilder/structs"
	"github.com/diegorufe/goutils/pkg/querybuilder/types"
)

type QueryBuilder interface {
	SetType(builderType types.BuilderType) QueryBuilder
	GetType() types.BuilderType
	And(filter structs.Filter) QueryBuilder
	Or(filter structs.Filter) QueryBuilder
	RawQuery(query string, arguments ...any) (*structs.QueryResult, error)
	Exec(query string, arguments ...any) (*structs.QueryResult, error)
	AddParam(name string, value any) QueryBuilder
	Clear() QueryBuilder
	GetFilters() []structs.Filter
	GetParams() map[string]any
}

type DefaultQueryBuilder struct {
	Type       types.BuilderType
	Fields     []structs.Field
	Filters    []structs.Filter
	Joins      []structs.Join
	Orders     []structs.Order
	Groups     []structs.Group
	Trasaction *structs.Transaction
	Params     map[string]any
}

func InstanceBuilder(builderType types.BuilderType, transaction *structs.Transaction) QueryBuilder {
	queryBuilder := DefaultQueryBuilder{Type: builderType, Trasaction: transaction}
	queryBuilder.Params = make(map[string]any)
	return queryBuilder
}

func (builder DefaultQueryBuilder) SetType(builderType types.BuilderType) QueryBuilder {
	builder.Type = builderType
	return builder
}

func (builder DefaultQueryBuilder) GetType() types.BuilderType {
	return builder.Type
}

func (builder DefaultQueryBuilder) appendFilter(filter structs.Filter, filterOperator types.FilterOperator) QueryBuilder {
	filter.Operator = filterOperator
	builder.Filters = append(builder.Filters, filter)
	return builder
}

func (builder DefaultQueryBuilder) And(filter structs.Filter) QueryBuilder {
	return builder.appendFilter(filter, types.And)
}

func (builder DefaultQueryBuilder) Or(filter structs.Filter) QueryBuilder {
	return builder.appendFilter(filter, types.Or)
}

func (builder DefaultQueryBuilder) RawQuery(query string, arguments ...any) (*structs.QueryResult, error) {
	return nil, nil
}

func (builder DefaultQueryBuilder) Exec(query string, arguments ...any) (*structs.QueryResult, error) {
	return nil, nil
}

func (builder DefaultQueryBuilder) AddParam(name string, value any) QueryBuilder {
	builder.Params[name] = value
	return builder
}

func (builder DefaultQueryBuilder) Clear() QueryBuilder {
	builder.Fields = nil
	builder.Joins = nil
	builder.Fields = nil
	builder.Orders = nil
	builder.Groups = nil
	builder.Params = make(map[string]any)
	return builder
}

func (builder DefaultQueryBuilder) GetFilters() []structs.Filter {
	return builder.Filters
}

func (builder DefaultQueryBuilder) GetParams() map[string]any {
	return builder.Params
}
