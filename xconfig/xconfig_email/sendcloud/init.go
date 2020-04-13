package sendcloud

type emailResult struct {
	Result     bool `json:"result"`
	StatusCode int  `json:"statusCode"`
	Message    string
}
