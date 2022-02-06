package canbus

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
			Standard int16    `yaml:"standard,omitempty"`
			Extended int16    `yaml:"extended,omitempty"`
			J1939 struct {
				Priority      uint8 `yaml:"priority,omitempty"`
				DataPage      uint8 `yaml:"data_page,omitempty"`
				PDUFormat     uint8 `yaml:"pdu_format,omitempty"`
				PDUSpecific   uint8 `yaml:"pdu_specific,omitempty"`
				SourceAddress uint8 `yaml:"source_address,omitempty"`
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

// Scaling, Limit, Transfer, Offset
type SLOT struct {
	Object `yaml:",inline"`
	Spec   struct {
		Min    float32 `yaml:"min"`
		Max    float32 `yaml:"max"`
		Offset float32 `yaml:"offset"`
		Size   string  `yaml:"size"`
	} `yaml:"metadata"`
}
