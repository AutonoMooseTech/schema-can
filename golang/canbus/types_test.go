package schemacan

import (
	"testing"
)

func TestMessageInvalid(t *testing.T) {
	msg := Message{}

	if err := msg.Validate(); err == nil {
		t.Errorf("Message should not be valid: %s", err)
	}
}
