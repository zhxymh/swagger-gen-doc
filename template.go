package gendoc

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	swagger "github.com/zhxymh/swagger-gen-doc/swagger"
)

type TemplateData struct {
	Title    string          `json:"title"`
	Version  string          `json:"version"`
	OpenApi  string          `json:"openapi"`
	Api      orderedApiInfo  `json:"api"`
	ApiModel orderedApiModel `json:"apiModel"`
}

type ApiInfo struct {
	Url     string `json:"url"`
	Mehtod  string `json:"method"`
	Summary string `json:"summary"`

	HasParameters  bool `json:"hasParameters"`
	HasRequestBody bool `json:"hasRequestBody"`
	HasResponses   bool `json:"hasResponses"`
	HasTags        bool `json:"hasTags"`

	Parameters  []*ParameterInfo `json:"parameters"`
	RequestBody *RequestBodyInfo `json:"requestBody"`
	Responses   *ResponseInfo    `json:"responses"`
	Tags        []string         `json:"tags"`
}

type ParameterInfo struct {
	Description string `json:"description"`
	In          string `json:"in"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Nullable    bool   `json:"nullable"`
	Default     string `json:"default"`
}

type RequestBodyInfo struct {
	Description string   `json:"description"`
	Type        string   `json:"type"`
	Content     []string `json:"content"`
}

type ResponseInfo struct {
	StateCode   int      `json:"stateCode"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	RefType     string   `json:"refType"`
	ItemType    string   `json:"itemType"`
	ItemRefType string   `json:"itemRefType"`
	Content     []string `json:"content"`
}

type ApiModel struct {
	Name       string          `json:"name"`
	Properties []*PropertyInfo `json:"properties"`
}

type PropertyInfo struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	RefType     string `json:"refType"`
	ItemType    string `json:"itemType"`
	ItemRefType string `json:"itemRefType"`
	Nullable    bool   `json:"nullable"`
	Description string `json:"description"`
}

func GenerateTemplateData(swaggerJson []byte) *TemplateData {
	var swaggerData swagger.SwaggerData
	json.Unmarshal(swaggerJson, &swaggerData)

	//log.Println(&swaggerData)

	templateData := &TemplateData{
		Title:   swaggerData.Info.Title,
		Version: swaggerData.Info.Version,
		OpenApi: swaggerData.OpenApi,
	}

	for keyPath, valPath := range swaggerData.Paths {
		for k, v := range valPath {
			apiInfo := &ApiInfo{
				Url:         keyPath,
				Mehtod:      k,
				Summary:     v.Summary,
				Tags:        v.Tags,
				Parameters:  parseParameters(v.Parameters),
				RequestBody: parseRequestBody(v.RequestBody),
				Responses:   parseRespone(v.Responses),
			}

			apiInfo.HasParameters = len(apiInfo.Parameters) > 0
			apiInfo.HasRequestBody = apiInfo.RequestBody.Description != "" || len(apiInfo.RequestBody.Content) > 0
			apiInfo.HasResponses = apiInfo.Responses.Description != "" || len(apiInfo.Responses.Content) > 0
			apiInfo.HasTags = len(apiInfo.Tags) > 0

			templateData.Api = append(templateData.Api, apiInfo)

		}
	}

	templateData.ApiModel = parseModel(swaggerData.Components)

	sort.Sort(templateData.Api)
	sort.Sort(templateData.ApiModel)

	return templateData
}

func parseParameters(parameters []swagger.ParameterInfo) []*ParameterInfo {
	result := make([]*ParameterInfo, 0, len(parameters))
	for _, val := range parameters {
		parameter := &ParameterInfo{
			Description: val.Description,
			In:          val.In,
			Name:        val.Name,
			Type:        val.Schema.Type,
			Nullable:    val.Schema.Nullable,
		}
		if val.Schema.Default != nil {
			parameter.Default = fmt.Sprintf("%v", val.Schema.Default)
		}
		result = append(result, parameter)
	}

	return result
}

func parseRequestBody(requestBody swagger.RequestBodyInfo) *RequestBodyInfo {
	requestBodyInfo := &RequestBodyInfo{
		Description: requestBody.Description,
	}

	requestType := ""
	content := make([]string, 0, len(requestBody.Content))

	for k, v := range requestBody.Content {
		if requestType == "" {
			ref := strings.Split(v["schema"]["$ref"], "/")
			requestType = ref[len(ref)-1]
		}

		content = append(content, k)
	}

	requestBodyInfo.Type = requestType
	requestBodyInfo.Content = content

	return requestBodyInfo
}

func parseRespone(responses map[int]swagger.ResponseInfo) *ResponseInfo {
	swaggerResponse := responses[200]

	result := &ResponseInfo{
		StateCode:   200,
		Description: swaggerResponse.Description,
	}

	responseItmeType := ""
	responseItmeRefType := ""
	responseType := ""
	responseRefType := ""
	content := make([]string, 0, len(swaggerResponse.Content))

	for k, v := range swaggerResponse.Content {
		flag := true
		if flag {
			responseType = v.Schema.Type
			ref := strings.Split(v.Schema.Ref, "/")
			responseRefType = ref[len(ref)-1]

			responseItmeType = v.Schema.Items.Type
			itemRef := strings.Split(v.Schema.Items.Ref, "/")
			responseItmeRefType = itemRef[len(itemRef)-1]

			flag = false
		}

		content = append(content, k)
	}

	result.Type = responseType
	result.RefType = responseRefType
	result.ItemType = responseItmeType
	result.ItemRefType = responseItmeRefType
	result.Content = content

	return result
}

func parseModel(components swagger.Components) []*ApiModel {
	models := make([]*ApiModel, 0, len(components.Schemas))

	for k, v := range components.Schemas {
		model := &ApiModel{
			Name: k,
		}

		for k, v := range v.Properties {
			ref := strings.Split(v.RefType, "/")
			itemRef := strings.Split(v.Items.RefType, "/")
			property := &PropertyInfo{
				Name:        k,
				Type:        v.Type,
				RefType:     ref[len(ref)-1],
				ItemType:    v.Items.Type,
				ItemRefType: itemRef[len(itemRef)-1],
				Nullable:    v.Nullable,
				Description: v.Description,
			}
			model.Properties = append(model.Properties, property)
		}

		models = append(models, model)
	}

	return models
}

type orderedApiInfo []*ApiInfo

func (o orderedApiInfo) Len() int           { return len(o) }
func (o orderedApiInfo) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
func (o orderedApiInfo) Less(i, j int) bool { return o[i].Url < o[j].Url }

type orderedApiModel []*ApiModel

func (o orderedApiModel) Len() int           { return len(o) }
func (o orderedApiModel) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
func (o orderedApiModel) Less(i, j int) bool { return o[i].Name < o[j].Name }
