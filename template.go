package gendoc

import (
	"encoding/json"
)

type TemplateData struct {
	Info    Info   `json:"info"`
	OpenApi string `json:"openapi"`
}

type Info struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

func GenerateTemplateData(swaggerJson []byte) *TemplateData {
	var data TemplateData
	json.Unmarshal(swaggerJson, &data)

	return &data
}
