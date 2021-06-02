package cobra

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

// YamlFlag Yaml的flags配置
type YamlFlag struct {
	FlagName  string   `yaml:"Name"`
	FlagType  FlagType `yaml:"Type"`
	FlagValue string   `yaml:"Default"`
	FlagUsage string   `yaml:"Usage"`
}

// YamlCommand 用于描述一个命令
type YamlCommand struct {
	Use      string     `yaml:"Use"`
	Long     string     `yaml:"Long"`
	Short    string     `yaml:"Short"`
	PkgPath  string     `yaml:"PkgPath"`
	Function string     `yaml:"Function"`
	Flags    []YamlFlag `yaml:"Flags"`
}

// YamlModels 描述一个配置文件
type YamlModels struct {
	Use                string        `yaml:"Application"`
	Long               string        `yaml:"Description"`
	GeneratorDirectory string        `yaml:"GeneratorDirectory"`
	GlobalFlags        []YamlFlag    `yaml:"GlobalFlags"`
	Commands           []YamlCommand `yaml:"SubCommand"`
	Package            string        `yaml:"PackageName"`
	AppVersion         string        `yaml:"Version"`
	// DockerHub hub镜像仓库
	DockerHub   string `yaml:"DockerHub"`
	GitAccount  string `yaml:"GitAccount"`
	GitPassword string `yaml:"GitPassword"`
	GitPro      string `yaml:"GitPro"`
	GitHost     string `yaml:"GitHost"`
}

// LoadDataFromFile 用于从文件中加载数据
func (ym *YamlModels) LoadDataFromFile(file string) *YamlModels {
	bDataBody, mErr := ioutil.ReadFile(file)
	if mErr != nil {
		log.Fatalln(mErr.Error())
	}
	mErr = yaml.Unmarshal(bDataBody, ym)
	if mErr != nil {
		log.Fatalln(mErr.Error())
	}
	return ym
}

// GeneratorBuildWin 生成windows编译脚本
func (ym *YamlModels) GeneratorBuildWin() {
	if ym.GeneratorDirectory == "" {
		ym.GeneratorDirectory = "."
	}
	mFileHandle, mFileErr := os.Create(filepath.Join(
		ym.GeneratorDirectory,
		"build_windows.bat"))
	if mFileErr != nil {
		log.Fatalln(mFileErr.Error())
	}
	defer mFileHandle.Close()
	mTp, mErr := template.New("windows").Funcs(_FuncsMaps).Parse(_BuildWinTemplate)
	if mErr != nil {
		log.Fatalln(mErr)
	}
	mExecuteErr := mTp.Execute(mFileHandle, ym)
	if mExecuteErr != nil {
		log.Fatalln()
	}
}

// GeneratorBuildLinux 生成编译linux的脚本
func (ym *YamlModels) GeneratorBuildLinux() {
	if ym.GeneratorDirectory == "" {
		ym.GeneratorDirectory = "."
	}
	mFileHandle, mFileErr := os.Create(filepath.Join(
		ym.GeneratorDirectory,
		"build_linux.sh"))
	if mFileErr != nil {
		log.Fatalln(mFileErr.Error())
	}
	defer mFileHandle.Close()

	mTp, mErr := template.New("linux").Funcs(_FuncsMaps).Parse(_BuildLinuxTamplate)
	if mErr != nil {
		log.Fatalln(mErr)
	}
	mExecuteErr := mTp.Execute(mFileHandle, ym)
	if mExecuteErr != nil {
		log.Fatalln()
	}
}

func (ym *YamlModels) GeneratorMain() {
	if ym.GeneratorDirectory == "" {
		ym.GeneratorDirectory = "."
	}
	mFileHandle, mFileErr := os.Create(filepath.Join(
		ym.GeneratorDirectory,
		"main.go"))
	if mFileErr != nil {
		log.Fatalln(mFileErr.Error())
	}
	defer mFileHandle.Close()
	mTp, mErr := template.New("main").Funcs(_FuncsMaps).Parse(_MainTemplate)
	if mErr != nil {
		log.Fatalln(mErr)
	}
	mExecuteErr := mTp.Execute(mFileHandle, ym)
	if mExecuteErr != nil {
		log.Fatalln()
	}
}

// Generator 用于生成代码
func (ym *YamlModels) Generator() {
	ym.GeneratorCodes()
	ym.GeneratorBuildLinux()
	ym.GeneratorBuildWin()
	ym.GeneratorMain()
	ym.GeneratorFlags()
}

// GetVariable
func GetVariable(str string) string {
	arr := strings.Split(str, "_")

	for index, val := range arr {
		arr[index] = strings.ToUpper(val[0:1]) + val[1:]
	}
	return strings.Join(arr, "")
}

// GeneratorFlags 用于获得代码
func (ym *YamlModels) GeneratorFlags() {
	if ym.GeneratorDirectory == "" {
		ym.GeneratorDirectory = "."
	}
	dir := filepath.Join(ym.GeneratorDirectory, "flags")
	os.MkdirAll(dir, 0755)
	mFileHandle, mFileErr := os.Create(filepath.Join(dir, "flags.go"))
	if mFileErr != nil {
		log.Fatalln(mFileErr.Error())
	}
	defer mFileHandle.Close()
	mTp, mErr := template.New("flags").Funcs(_FuncsMaps).Parse(_FlagsTemplate)
	if mErr != nil {
		log.Fatalln(mErr)
	}
	mExecuteErr := mTp.Execute(mFileHandle, ym)
	if mExecuteErr != nil {
		log.Fatalln()
	}
}

// GeneratorCodes 用于获得代码
func (ym *YamlModels) GeneratorCodes() {
	if ym.GeneratorDirectory == "" {
		ym.GeneratorDirectory = "."
	}
	dir := filepath.Join(ym.GeneratorDirectory, "cmdline")
	os.MkdirAll(dir, 0755)
	mFileHandle, mFileErr := os.Create(filepath.Join(dir, "cmdline.go"))
	if mFileErr != nil {
		log.Fatalln(mFileErr.Error())
	}
	defer mFileHandle.Close()
	mTp, mErr := template.New("codes").Funcs(_FuncsMaps).Parse(_CodesTemplate)
	if mErr != nil {
		log.Fatalln(mErr)
	}
	mExecuteErr := mTp.Execute(mFileHandle, ym)
	if mExecuteErr != nil {
		log.Fatalln()
	}
}
