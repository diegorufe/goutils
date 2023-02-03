package types

type BuilderType int8

const (
	SqlLite  BuilderType = 0
	Postgres BuilderType = 1
	Mysql    BuilderType = 2
)

type OrderType int8

const (
	Asc  OrderType = 0
	Desc OrderType = 1
)

type FilterType int8
type FilterOperator int8

const (
	Eq        FilterType = 0
	Like      FilterType = 1
	Ne        FilterType = 2
	LikeStart FilterType = 3
	LikeEnd   FilterType = 4
	Le        FilterType = 5
	Lt        FilterType = 6
	Gt        FilterType = 7
	Ge        FilterType = 8
	In        FilterType = 9
	NotIn     FilterType = 10
	IsNull    FilterType = 11
	IsNotNull FilterType = 12
)

const (
	And FilterOperator = 0
	Or  FilterOperator = 1
)

type JoinType int8

const (
	Inner JoinType = 0
	Left  JoinType = 1
	Right JoinType = 2
)
