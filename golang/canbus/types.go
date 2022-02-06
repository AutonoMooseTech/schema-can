package canbus

import (
	"errors"
)

// Common object attributes
type Object struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string            `yaml:"name"`
		Namespace string            `yaml:"namespace,omitempty"`
		Labels    map[string]string `yaml:"labels,omitempty"`
	} `yaml:"metadata"`
}

// Message
type Message struct {
	Object `yaml:",inline"`
	Spec   struct {
		Id struct {
			Standard *uint16 `yaml:"standard,omitempty"`
			Extended *uint32 `yaml:"extended,omitempty"`
			J1939    *struct {
				Priority         *uint8 `yaml:"priority,omitempty"`
				DataPage         *bool  `yaml:"data_page,omitempty"`
				ExtendedDataPage *bool  `yaml:"extended_data_page,omitempty"`
				PDUFormat        *uint8 `yaml:"pdu_format,omitempty"`
				PDUSpecific      *uint8 `yaml:"pdu_specific,omitempty"`
				SourceAddress    *uint8 `yaml:"source_address,omitempty"`
			} `yaml:"j1939,omitempty"`
		} `yaml:"id"`
		Length uint8 `yaml:"length"`
		Data   []struct {
			Name          string `yaml:"name,omitempty"`
			Size          string `yaml:"size,omitempty"`
			SLOTReference string `yaml:"slot,omitempty"`
			Padding       uint8  `yaml:"padding,omitempty"`
		} `yaml:"data" json:"data"`
	} `yaml:"spec"`
}

func (msg *Message) Validate() (err error) {
	// Check identifier is valid
	if msg.Spec.Id.Standard != nil {
		// id within bounds
		if *msg.Spec.Id.Standard > ((2 ^ 11) - 1) {
			err = errors.New("Standard identifier value is too large")
		}
	} else if msg.Spec.Id.Extended != nil {
		// id within bounds
		if *msg.Spec.Id.Extended > ((2 ^ 29) - 1) {
			err = errors.New("Extended identifier value is too large")
		}
	} else if msg.Spec.Id.J1939 != nil {
		// priority within bounds
		if *msg.Spec.Id.J1939.Priority <= 0b111 {
			err = errors.New("Priority is too high. Maximum is 3")
		}

		// other fields are constrained by their types
	} else {
		err = errors.New("At least one identifier type must be declared")
	}

	return err
}

// Scaling, Limit, Transfer, Offset
type SLOT struct {
	Object `yaml:",inline"`
	Spec   struct {
		Min    *float32 `yaml:"min"`
		Max    *float32 `yaml:"max"`
		Offset *float32 `yaml:"offset"`
		Size   *string  `yaml:"size"`
		Unit   *string  `yaml:"unit,omitempty"`
	} `yaml:"spec"`
}

func (msg *SLOT) Validate() (err error) {
	// Check identifier is valid
	if msg.Spec.Min == nil {
		return errors.New("Minimum value must be declared")
	} else if msg.Spec.Max == nil {
		return errors.New("Maximum value must be declared")
	} else if msg.Spec.Offset == nil {
		return errors.New("Maximum value must be declared")
	} else if msg.Spec.Size == nil {
		return errors.New("Size must be declared")
	}

	return err
}

