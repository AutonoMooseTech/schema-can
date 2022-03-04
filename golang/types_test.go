package schemacan

import (
	"testing"
)

func TestMessageInvalid(t *testing.T) {
	msg := Message{}

	if err := msg.Validate(); err == nil {
		t.Errorf("Defualt initialised Message should not be valid: %s", err)
	}
}

func TestSlotInvalid(t *testing.T) {
	slot := SLOT{}

	if err := slot.Validate(); err == nil {
		t.Errorf("Defualt initialised SLOT should not be valid: %s", err)
	}
}

