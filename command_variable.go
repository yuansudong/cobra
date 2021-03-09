package cobra

// FlagType 定义一个Flag类型
type FlagType string

func (ft FlagType) String() string {
	return string(ft)
}

const (
	// FlagInt int类型
	FlagInt     FlagType = "int"
	FlagInt64   FlagType = "int64"
	FlagInt32   FlagType = "int32"
	FlagInt8    FlagType = "int8"
	FlagString  FlagType = "string"
	FlagFloat32 FlagType = "float32"
	FlagFloat64 FlagType = "float64"
)

// GeneratorFlag 生成的Flag对象
type GeneratorFlag struct {
	FlagName   string   `yaml:"FlagName"`
	FlagType   FlagType `yaml:"FlagType"`
	FlagValue  string   `yaml:"FlagValue"`
	FlagUseage string   `yaml:"FlagUseage"`
}
