package main

// YamlFlag Yaml的flags配置
type YamlFlag struct {
	Name string `yaml:"Name"`
	Type string `yaml:"Type"`
}

// YamlExecute 用于描述一个执行
type YamlExecute struct {
	PkgPath  string `yaml:"PkgPath"`
	Function string `yaml:"Function"`
}

// YamlCommand 用于描述一个命令
type YamlCommand struct {
	Use   string      `yaml:"Use"`
	Long  string      `yaml:"Long"`
	Short string      `yaml:"Short"`
	Exec  YamlExecute `yaml:"Execute"`
	Flags []YamlFlag  `yaml:"Flags"`
}

// YamlModels 描述一个配置文件
type YamlModels struct {
	Use                string        `yaml:"Application"`
	Long               string        `yaml:"Description"`
	GeneratorDirectory string        `yaml:"GeneratorDirectory"`
	GlobalFlags        []YamlFlag    `yaml:"GlobalFlags"`
	Commands           []YamlCommand `yaml:"SubCommand"`
}
