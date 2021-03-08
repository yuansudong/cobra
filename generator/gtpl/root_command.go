package gtpl

import "github.com/yuansudong/cobra"

var _RootCommand = cobra.Command{}

// ExampleGlobalFlags 全局配置
type ExampleGlobalFlags struct {
	F1 string
	F2 bool
}

func init() {
	_RootCommand.Use = "VarUse"
	_RootCommand.Long = "VarLong"

}

// Execute 执行函数
func Execute() error {
	return _RootCommand.Execute()
}
