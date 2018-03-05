package metasyntactic

import (
	"strings"
	"testing"
)

func TestStringExample(t *testing.T) {
	t.Run("test to lower", func(t *testing.T) {
		testString := "FoO"
		wantString := "foo"

		newString := strings.ToLower(testString)
		if newString != wantString {
			t.Errorf("ToLower() = %v, want %v", newString, wantString)
		}
	})
}

func TestStringBadExample(t *testing.T) {
	t.Run("test to lower bad", func(t *testing.T) {
		testString := "jeFF Bean"
		wantString := "jeff bean"

		newString := strings.ToLower(testString)
		if newString != wantString {
			t.Errorf("ToLower() = %v, want %v", newString, wantString)
		}
	})
}
