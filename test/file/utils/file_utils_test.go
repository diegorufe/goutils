package utils

import (
	"testing"

	"github.com/diegorufe/goutils/pkg/assertions"
	"github.com/diegorufe/goutils/pkg/file/utils"
)

func TestFileExistFalse(t *testing.T) {
	assertions.AssertFalse(t, utils.FileExist("TEST"))
}

func TestFileExist(t *testing.T) {
	assertions.AssertTrue(t, utils.FileExist("./test.txt"))
}

func TestDirExistsFalse(t *testing.T) {
	assertions.AssertFalse(t, utils.DirExists("dir"))
}

func TestDirExist(t *testing.T) {
	assertions.AssertTrue(t, utils.DirExists("./testdir"))
}
