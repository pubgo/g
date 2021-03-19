package netutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalIp(t *testing.T) {
	assert.True(t, len(InternalIp()) > 0)
}
