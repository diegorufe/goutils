package utils

import (
	"testing"

	"github.com/diegorufe/goutils/pkg/assertions"
	"github.com/diegorufe/goutils/pkg/strings/utils"
)

func TestStarWidth(t *testing.T) {
	assertions.AssertTrue(t, utils.StartWith("TEST", "T"))
	assertions.AssertTrue(t, utils.StartWith("TEST", "A", "B", "T"))
	assertions.AssertFalse(t, utils.StartWith("TEST", "A", "B", "E"))
}

func TestEndWith(t *testing.T) {
	assertions.AssertTrue(t, utils.EndtWith("TEST", "T"))
	assertions.AssertTrue(t, utils.EndtWith("TEST", "A", "B", "T"))
	assertions.AssertFalse(t, utils.EndtWith("TEST", "A", "B", "E"))
}

func TestContains(t *testing.T) {
	assertions.AssertTrue(t, utils.Contains("TEST", "T"))
	assertions.AssertTrue(t, utils.Contains("TEST", "A", "B", "T"))
	assertions.AssertFalse(t, utils.Contains("TEST", "A", "B", "C"))
}
