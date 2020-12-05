package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	gendoc "github.com/zhxymh/swagger-gen-doc"
)

func main() {
	inPtr := flag.String("doc_in", "resources/swagger.json", "the swagger json full path.")
	outPtr := flag.String("doc_out", "./swagger.md", "the generated file path.")
	optPtr := flag.String("doc_opt", "templates/markdown.tmpl", "the custome template file path.")

	flag.Parse()

	content, err := readJsonContent(*inPtr)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = generateDoc(&content, *outPtr, *optPtr); err != nil {
		log.Fatal(err)
		return
	}

	filePath, _ := filepath.Abs(*outPtr)
	log.Println("Generated successfully! file path: " + filePath)
}

func readJsonContent(file string) ([]byte, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	content, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func generateDoc(json *[]byte, outFile string, tmplFile string) error {
	templateData := gendoc.GenerateTemplateData(*json)
	tmpl, err := template.ParseFiles(tmplFile)

	if err != nil {
		return err
	}

	file, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	if err = tmpl.Execute(file, templateData); err != nil {
		return err
	}

	return nil
}
