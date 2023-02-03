package builder

import (
	"testing"

	"github.com/diegorufe/goutils/pkg/assertions"
	bl "github.com/diegorufe/goutils/pkg/querybuilder/builder"
	"github.com/diegorufe/goutils/pkg/querybuilder/structs"
	"github.com/diegorufe/goutils/pkg/querybuilder/types"
)

func TestInstanceBuilder(t *testing.T) {
	assertions.AssertNotNill(t, bl.InstanceBuilder(types.SqlLite, &structs.Transaction{}))
}

func TestGetType(t *testing.T) {
	assertions.AssertNotNill(t, bl.InstanceBuilder(types.SqlLite, &structs.Transaction{}))
	assertions.AssertEqual(t, types.SqlLite, bl.InstanceBuilder(types.SqlLite, &structs.Transaction{}).GetType())
	assertions.AssertNotEqual(t, types.Mysql, bl.InstanceBuilder(types.SqlLite, &structs.Transaction{}).GetType())
}

func TestAnd(t *testing.T) {
	qb := bl.InstanceBuilder(types.SqlLite, &structs.Transaction{})
	assertions.AssertTrue(t, len(qb.And(structs.Filter{}).GetFilters()) == 1)
}

func TestOr(t *testing.T) {
	qb := bl.InstanceBuilder(types.SqlLite, &structs.Transaction{})
	assertions.AssertTrue(t, len(qb.Or(structs.Filter{}).GetFilters()) == 1)
}

func TestGetParams(t *testing.T) {
	qb := bl.InstanceBuilder(types.SqlLite, &structs.Transaction{})
	assertions.AssertTrue(t, len(qb.AddParam("TEST", "TEST").GetParams()) == 1)
}
