package main

import (
	"github.com/yuansudong/cobra"
)

func main() {
	Read()
}

type RootCfg struct {
	Cfg *string
}

// Read
func Read() {

	RootCfg := new(RootCfg)

	rootCmd := cobra.Command{
		Use:   "cmdline",
		Long:  "命令行代码生成工具",
		Short: "命令行代码生成工具",
		Run: func(cmd *cobra.Command, args []string) {
			mM := new(cobra.YamlModels)
			mM.LoadDataFromFile(*RootCfg.Cfg).Generator()
		},
	}
	RootCfg.Cfg = rootCmd.PersistentFlags().String(
		"cfg",
		"./cmdline.yaml",
		"--cfg=./cmdline.yaml",
	)
	rootCmd.Execute()
}

// Test 用于读取
func Test() {
	mM := new(cobra.YamlModels)
	mM.LoadDataFromFile("./generator_example.yaml").GeneratorCodes()
}
