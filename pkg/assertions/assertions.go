package assertions

import (
	"testing"
)

func AssertNotNill(t *testing.T, value any) {
	if value == nil {
		t.Fatal("Value is nill")
	}
}

func AssertNill(t *testing.T, value any) {
	if value != nil {
		t.Fatal("Value is not nill")
	}
}

func AssertEqual(t *testing.T, desire any, value any) {
	if desire != value {
		t.Fatalf("Value %s is not the same %s", desire, value)
	}
}

func AssertNotEqual(t *testing.T, desire any, value any) {
	if desire == value {
		t.Fatalf("Value %s is the same %s", desire, value)
	}
}

func AssertTrue(t *testing.T, condition bool) {
	if !condition {
		t.Fatalf("Expected result true and got false")
	}
}

func AssertFalse(t *testing.T, condition bool) {
	if condition {
		t.Fatalf("Expected result false and got true")
	}
}

func AssertError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Error has ocurred %s", err)
	}
}

func AssertThrows(t *testing.T, err error) {
	if err == nil {
		t.Fatalf("Error has not ocurred")
	}
}
