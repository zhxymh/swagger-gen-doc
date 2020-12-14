package swagger

type SwaggerData struct {
	Info       Info                                `json:"info"`
	OpenApi    string                              `json:"openapi"`
	Paths      map[string]map[string]InterfaceInfo `json:"paths"`
	Components Components                          `json:"components"`
}

type Info struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

type InterfaceInfo struct {
	Parameters  []ParameterInfo      `json:"parameters"`
	RequestBody RequestBodyInfo      `json:"requestBody"`
	Responses   map[int]ResponseInfo `json:"responses"`
	Summary     string               `json:"summary"`
	Tags        []string             `json:"tags"`
}

type ParameterInfo struct {
	Description string              `json:"description"`
	In          string              `json:"in"`
	Name        string              `json:"name"`
	Schema      ParameterSchemaInfo `json:"schema"`
}

type ParameterSchemaInfo struct {
	Description string      `json:"description"`
	Nullable    bool        `json:"nullable"`
	Type        string      `json:"type"`
	Default     interface{} `json:"default"`
}

type RequestBodyInfo struct {
	Content     map[string]map[string]map[string]string `json:"content"`
	Description string                                  `json:"description"`
}

type ResponseInfo struct {
	Content     map[string]SchemaInfo `json:"content"`
	Description string                `json:"description"`
}

type SchemaInfo struct {
	Schema SchemaInfoDtail `json:"schema"`
}

type SchemaInfoDtail struct {
	Ref   string              `json:"$ref"`
	Type  string              `json:"type"`
	Items SchemaItemInfoDtail `json:"items"`
}

type SchemaItemInfoDtail struct {
	Ref  string `json:"$ref"`
	Type string `json:"type"`
}

type Components struct {
	Schemas map[string]ApiModel `json:"schemas"`
}

type ApiModel struct {
	Type                 string                  `json:"type"`
	Properties           map[string]PropertyInfo `json:"properties"`
	AdditionalProperties bool                    `json:"additionalProperties"`
}

type PropertyInfo struct {
	Type                 string           `json:"type"`
	AdditionalProperties string           `json:"additionalProperties"`
	Nullable             bool             `json:"nullable"`
	Description          string           `json:"description"`
	Format               string           `json:"format"`
	RefType              string           `json:"$ref"`
	Items                PropertyItemInfo `json:"items"`
}

type PropertyItemInfo struct {
	Type    string `json:"type"`
	RefType string `json:"$ref"`
}
