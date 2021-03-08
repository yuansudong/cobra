package main

import (
	"io/ioutil"
	"log"

	"github.com/yuansudong/cobra"
	"gopkg.in/yaml.v2"
)

func main() {
	Read()
}

// Read
func Read() {

	data, err := ioutil.ReadFile("./generator_example.yaml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	mModels := new(YamlModels)
	if err := yaml.Unmarshal(data, mModels); err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("%+v\n", mModels)
}

// Test 用于读取
func Test() {
	rootCmd := cobra.Command{
		Use:   "yaml_cmd",
		Long:  "这是一个生成程序,用于从yaml文件",
		Short: "生成代码,从yaml文件中生成golang代码",
	}
	addCmd := cobra.Command{
		Use:   "add",
		Long:  "这是一个增加命令",
		Short: "增加命令,到命令行中",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("执行Add命令")
		},
	}
	addCmd.Flags().StringP("hello", "e", "world", "hello=world")
	rootCmd.AddCommand(&addCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Println(err.Error())
	}
}
