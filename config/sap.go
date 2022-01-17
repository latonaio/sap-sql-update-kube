package config

import "os"

type SAP struct {
	baseURL string
}

func newSAP() *SAP {
	return &SAP{
		baseURL: os.Getenv("SAP_API_BASE_URL"),
	}
}
func (c *SAP) BaseURL() string {
	return c.baseURL
}
