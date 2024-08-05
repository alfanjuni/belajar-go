package responses

// Response is the standard response structure
type Response struct {
	Data   interface{} `json:"data"`
	Meta   Meta        `json:"meta"`
	Status Status      `json:"status"`
}

// Meta contains additional metadata
type Meta struct {
	Total int `json:"total"`
}

// Status contains the status code and message
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
