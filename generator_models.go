package cobra

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
	mTp, mErr := template.New("codes").Parse(_CodesTemplate)
	if mErr != nil {
		log.Fatalln(mErr)
	}
	mExecuteErr := mTp.Execute(mFileHandle, ym)
	if mExecuteErr != nil {
		log.Fatalln()
	}
}
