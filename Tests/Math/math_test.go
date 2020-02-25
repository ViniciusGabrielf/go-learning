package math

import (
	"testing"
)

const defaultError = "Expected value %v, but the result found was %v."

// TestAverage is a function to test result of average()
func TestAverage(t *testing.T) {
	expectedValue := 7.28
	value := Average(7.2, 9.9, 6.1, 5.9)

	if value != expectedValue {
		t.Errorf(defaultError, expectedValue, value)
	}
}
