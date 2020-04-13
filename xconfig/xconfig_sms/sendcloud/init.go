package sendcloud

type smsResult struct {
	Result     bool `json:"result"`
	StatusCode int  `json:"statusCode"`
	Message    string
}
