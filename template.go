package gendoc

import (
	"encoding/json"
)

type Template struct {
	Info Info `json:"info"`
}

type Info struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

func NewTemplate(swaggerJson []byte) {
	var data Template
	json.Unmarshal(swaggerJson, &data)

	println(data.Info.Title)
	println(data.Info.Version)
}
