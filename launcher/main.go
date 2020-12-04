package main

import (
	"io/ioutil"
	"os"

	gendoc "github.com/zhxymh/swagger-gen-doc/gendoc"
)

func main() {
	println("Hello gendoc!")

	file, err := os.Open("/Users/zhangxin/doc/test/swagger.json")
	if err != nil {
		println(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	//println(string(content))

	gendoc.NewTemplate(content)
}
