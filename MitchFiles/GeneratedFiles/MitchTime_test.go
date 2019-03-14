package GeneratedFiles

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadMessage_MitchMessage_Time(t *testing.T) {
	msg := NewTimeMessage()
	assert.NotNil(t, msg)
}
