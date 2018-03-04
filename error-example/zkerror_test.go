package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorZNodeInvalid(t *testing.T) {
	assert.EqualError(t, errZNodeInvalid{"foo error"}, "invalid ZNode URI: foo error")
	assert.EqualError(t, errZNodeInvalid{}, "invalid ZNode URI: ")
}
