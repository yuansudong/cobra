
package cmdline
import(
	"github.com/yuansudong/cobra"
	
	add "github.com/yuansudong/cobra"
	
	dec "github.com/yuansudong/cobra"
	
)
// GlobalFlag 全局Flag
type GlobalFlag struct {
	
	F1 *string
	
	F2 *int32
	
	F3 *int64
	
	F4 *float32
	
}

// LocaladdFlag add的Flag
type LocaladdFlag struct {
	
	A1 *string	
	
	A2 *int32	
	
	A3 *int64	
	
	A4 *float32	
	
}

// LocaldecFlag dec的Flag
type LocaldecFlag struct {
	
	F1 *string	
	
	F2 *int32	
	
	F3 *int64	
	
	F4 *float32	
	
}


var (
	// _Branch 分支名称
	_GitBranch string = "UNKNOWN"
	// _GitCommitID 最近一次的提交ID
	_GitCommit string = "UNKNOWN"
	// _GitAccount 提交人的名字
	_GitAccount string = "UNKNOWN"
	// _DateTime 编译的时间
	_DateTime string = "UNKNOWN"
	// _GoVersion GO的编译版本
	_GoVersion string = "UNKNOWN"
	// _OS 编译时的操作系统
	_OS string = "UNKNOWN"
	// _CPU类型
	_Arch string = "UNKNOWN"
)

// _InitVersion 用于初始化版本
func _InitVersion() *cobra.Command {
	mCommand := new(cobra.Command)
	mCommand.Use = "version"
	mCommand.Long = "查看编译以及版本信息"
	mCommand.Short = "查看编译以及版本信息"
	mCommand.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println("App Name     :   ", "example")
		fmt.Println("App Version  :   ", "")
		fmt.Println("Git Branch   :   ", _GitBranch)
		fmt.Println("Git Commit   :   ", _GitCommit)
		fmt.Println("Git Account  :   ", _GitAccount)
		fmt.Println("Go Version   :   ", _GoVersion)
		fmt.Println("Build System :   ", _OS)
		fmt.Println("Build Time   :   ", _DateTime)
		fmt.Println("Build Arch   :   ", _Arch)
	}
	return mCommand
}

var _Root *cobra.Command = _InitRoot()
func _InitRoot() *cobra.Command {
	mGlobalsFlags :=  new(GlobalFlag)
	mRootCommand := new(cobra.Command)
	mRootCommand.Use = "example" 
	mRootCommand.Long = "这是一个Example.exe的程序"
	
		
		mGlobalsFlags.F1 = mRootCommand.PersistentFlags().String(
			"F1",
			"F1_DEF", 
			"F1=V1",
		)
		
	
		
		mGlobalsFlags.F2 = mRootCommand.PersistentFlags().Int32(
			"F2",
			32, 
			"F2=V2",
		)
		
	
		
		mGlobalsFlags.F3 = mRootCommand.PersistentFlags().Int64(
			"F3",
			64, 
			"F3=V3",
		)
		
	
		
		mGlobalsFlags.F4 = mRootCommand.PersistentFlags().Float32(
			"F4",
			32.00, 
			"F4=V4",
		)
		
	
	
    mRootCommand.AddCommand(_InitSubadd(mGlobalsFlags))
	
    mRootCommand.AddCommand(_InitSubdec(mGlobalsFlags))
	
	mRootCommand.AddCommand(_InitVersion())
	return mRootCommand
}

// Execute 执行入口
func Execute() error {
	return _Root.Execute()
}

func _InitSubadd(mGlobal *GlobalFlag) *cobra.Command {
	mLocal := new(LocaladdFlag)
	mCommand :=  new(cobra.Command)
	mCommand.Use = "add"
	mCommand.Long = "add的详细描述"
	mCommand.Short = "add的简短描述"
	mCommand.Run = func(cmd *cobra.Command, args []string) {
		add.Haha(mGlobal,mLocal)
	} 
	
		
		mLocal.A1 = mCommand.PersistentFlags().String(
			"A1",
			"A1_DEF", 
			"A1=V1",
		)
		
	
		
		mLocal.A2 = mCommand.PersistentFlags().Int32(
			"A2",
			32, 
			"A2=V2",
		)
		
	
		
		mLocal.A3 = mCommand.PersistentFlags().Int64(
			"A3",
			64, 
			"A3=V3",
		)
		
	
		
		mLocal.A4 = mCommand.PersistentFlags().Float32(
			"A4",
			32.00, 
			"A4=V4",
		)
		
	
	return mCommand
}

func _InitSubdec(mGlobal *GlobalFlag) *cobra.Command {
	mLocal := new(LocaldecFlag)
	mCommand :=  new(cobra.Command)
	mCommand.Use = "dec"
	mCommand.Long = "dec的详细描述"
	mCommand.Short = "dec的简短描述"
	mCommand.Run = func(cmd *cobra.Command, args []string) {
		dec.Decription(mGlobal,mLocal)
	} 
	
		
		mLocal.F1 = mCommand.PersistentFlags().String(
			"F1",
			"F1_DEF", 
			"F1=V1",
		)
		
	
		
		mLocal.F2 = mCommand.PersistentFlags().Int32(
			"F2",
			32, 
			"F2=V2",
		)
		
	
		
		mLocal.F3 = mCommand.PersistentFlags().Int64(
			"F3",
			64, 
			"F3=V3",
		)
		
	
		
		mLocal.F4 = mCommand.PersistentFlags().Float32(
			"F4",
			32.00, 
			"F4=V4",
		)
		
	
	return mCommand
}


