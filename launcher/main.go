package main

import (
	"html/template"
	"io/ioutil"
	"os"

	gendoc "github.com/zhxymh/swagger-gen-doc"
)

func main() {
	println("Hello gendoc!")

	file, err := os.Open("resources/swagger.json")
	if err != nil {
		println(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)

	templateData := gendoc.GenerateTemplateData(content)
	println(templateData.Info.Version)

	t1, err := template.ParseFiles("templates/test.tmpl")

	if err != nil {
		println(err.Error())
	}

	_ = t1.Execute(os.Stdout, templateData)

	println()
}
