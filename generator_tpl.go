package cobra

const _CodesTemplate string = `
package {{.Package}}
import(
	"github.com/yuansudong/cobra"
	{{range .Commands}}
	{{.Use}} "{{.PkgPath}}"
	{{end}}
)
// GlobalFlag 全局Flag
type GlobalFlag struct {
	{{range .GlobalFlags}}
	{{.FlagName}} *{{.FlagType}}
	{{end}}
}
{{range .Commands }}
// Local{{.Use}}Flag {{.Use}}的Flag
type Local{{.Use}}Flag struct {
	{{range .Flags}}
	{{.FlagName}} *{{.FlagType}}	
	{{end}}
}
{{end}}
var _Root *cobra.Command = _InitRoot()
func _InitRoot() *cobra.Command {
	mGlobalsFlags :=  new(GlobalFlag)
	mRootCommand := new(cobra.Command)
	mRootCommand.Use = "{{.Use}}" 
	mRootCommand.Long = "{{.Long}}"
	{{range .GlobalFlags}}
		{{if eq .FlagType "float32"}}
		mGlobalsFlags.{{.FlagName}} = mRootCommand.PersistentFlags().Float32(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "float64"}}
		mGlobalsFlags.{{.FlagName}} = mRootCommand.PersistentFlags().Float64(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int8"}}
		mGlobalsFlags.{{.FlagName}} = mRootCommand.PersistentFlags().Int8(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int16"}}
		mGlobalsFlags.{{.FlagName}} = mRootCommand.PersistentFlags().Int16(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int32"}}
		mGlobalsFlags.{{.FlagName}} = mRootCommand.PersistentFlags().Int32(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int64"}}
		mGlobalsFlags.{{.FlagName}} = mRootCommand.PersistentFlags().Int64(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int"}}
		mGlobalsFlags.{{.FlagName}} = mRootCommand.PersistentFlags().Int(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "string"}}
		mGlobalsFlags.{{.FlagName}} = mRootCommand.PersistentFlags().String(
			"{{.FlagName}}",
			"{{.FlagValue}}", 
			"{{.FlagUsage}}",
		)
		{{end}}
	{{end}}
	{{range .Commands}}
    mRootCommand.AddCommand(_InitSub{{.Use}}(mGlobalsFlags))
	{{end}}
	return mRootCommand
}

// Execute 执行入口
func Execute() error {
	return _Root.Execute()
}
{{range .Commands}}
func _InitSub{{.Use}}(mGlobal *GlobalFlag) *cobra.Command {
	mLocal := new(Local{{.Use}}Flag)
	mCommand :=  new(cobra.Command)
	mCommand.Use = "{{.Use}}"
	mCommand.Long = "{{.Long}}"
	mCommand.Short = "{{.Short}}"
	mCommand.Run = func(cmd *cobra.Command, args []string) {
		{{.Use}}.{{.Function}}(mGlobal,mLocal)
	} 
	{{range .Flags}}
		{{if eq .FlagType "float32"}}
		mLocal.{{.FlagName}} = mCommand.PersistentFlags().Float32(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "float64"}}
		mLocal.{{.FlagName}} = mCommand.PersistentFlags().Float64(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int8"}}
		mLocal.{{.FlagName}} = mCommand.PersistentFlags().Int8(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int16"}}
		mLocal.{{.FlagName}} = mCommand.PersistentFlags().Int16(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int32"}}
		mLocal.{{.FlagName}} = mCommand.PersistentFlags().Int32(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int64"}}
		mLocal.{{.FlagName}} = mCommand.PersistentFlags().Int64(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int"}}
		mLocal.{{.FlagName}} = mCommand.PersistentFlags().Int(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "string"}}
		mLocal.{{.FlagName}} = mCommand.PersistentFlags().String(
			"{{.FlagName}}",
			"{{.FlagValue}}", 
			"{{.FlagUsage}}",
		)
		{{end}}
	{{end}}
	return mCommand
}
{{end}}

`
