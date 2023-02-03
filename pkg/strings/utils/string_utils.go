package utils

import (
	"strings"
)

func StartWith(value string, values ...string) bool {
	var starWith bool = false

	for _, s := range values {
		if strings.HasPrefix(value, s) {
			starWith = true
			break
		}
	}

	return starWith
}

func EndtWith(value string, values ...string) bool {
	var endWith bool = false

	for _, s := range values {
		if strings.HasSuffix(value, s) {
			endWith = true
			break
		}
	}

	return endWith
}

func Contains(value string, values ...string) bool {
	var containsValue bool = false

	for _, s := range values {
		if strings.Contains(value, s) {
			containsValue = true
			break
		}
	}

	return containsValue
}
