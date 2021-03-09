package main

import (
	"github.com/yuansudong/cobra"
)

func main() {
	Test()
}

// Read
func Read() {

}

// Test 用于读取
func Test() {
	mM := new(cobra.YamlModels)
	mM.LoadDataFromFile("./generator_example.yaml").GeneratorCodes()
}
