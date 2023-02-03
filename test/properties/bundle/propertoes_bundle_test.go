package bundle

import (
	"testing"

	"github.com/diegorufe/goutils/pkg/assertions"
	"github.com/diegorufe/goutils/pkg/properties/bundle"
)

func TestLoadFile(t *testing.T) {
	bd := bundle.InstancePropertiesBundle()
	assertions.AssertError(t, bd.LoadFile("./test_properties.properties"))
	assertions.AssertEqual(t, "ABC", bd.Get("hello"))
}

func TestLoadMultipeFiles(t *testing.T) {
	bd := bundle.InstancePropertiesBundle()
	assertions.AssertError(t, bd.LoadFiles("./testdir"))
	assertions.AssertEqual(t, "ABC", bd.Get("hello"))
}
