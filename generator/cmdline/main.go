package main

import (
	"encoding/json"
	"io"

	"gopkg.in/yaml.v2"
)


type _Student struct {
}

func (s *_Student) Read(p []byte) (n int, err error) {
	return 0, nil
}
func (s *_Student) Write(p []byte) (n int, err error) {
	return 0, nil
}

func main() {
	// doc.GenYaml(_Root, os.Stdout)
	// doc.GenMarkdown(_Root, os.Stdout)
	// for _,item := _Root.Commands() {

	// }
	json.Marshal(v interface{})
	yaml.Marshal(in interface{})
	p := new(_Student)
	V(p)
	Execute()
}

func V(rc io.ReadWriter) {

}
