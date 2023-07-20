package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIfItGetsAnErrorIfIdIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")
}
