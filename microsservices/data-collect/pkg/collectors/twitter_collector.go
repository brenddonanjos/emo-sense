package collectors

type TwitterCollector struct {
	APIKey    string
	APISecret string
}

func NewTwitterCollector(apiKey, apiSecret string) *TwitterCollector {
	return &TwitterCollector{APIKey: apiKey, APISecret: apiSecret}
}

func (t *TwitterCollector) Collect() error {
	return nil
}
