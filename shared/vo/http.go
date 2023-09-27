package vo

type Reject struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	Metadata any    `json:"metadata,omitempty"`
}

type Resolve struct {
	Data     interface{} `json:"data"`
	Message  string      `json:"message,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
}
