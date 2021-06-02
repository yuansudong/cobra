package cobra

var _FuncsMaps map[string]interface{} = map[string]interface{}{
	"GetVariable": GetVariable,
}

// _DockerfileTemplate
const _DockerfileTemplate string = `
# 编译镜像
FROM golang:1.15.3 AS builderImage
WORKDIR /go/src/project
COPY . .
# 配置免密
RUN echo 'https://{{.GitAccount}}/{{.GitPassword}}//{{.GitHost}}/{{.GitPro}}' > ~/.git-credentials
RUN git config --global credential.helper store
RUN go env -w GO111MODULE=on
RUN go env -w GONOPROXY={{.GitHost}}
RUN go env -w GOPRIVATE={{.GitHost}}
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go mod init
RUN go mod tidy
RUN chmod +x ./build_linux.sh
RUN ./build_linux.sh
# 运行镜像
FROM alpine:3.12.0
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezon && rm -rf /tmp/* \
    && rm -rf /var/cache/apk/* && mkdir -p /root/project
WORKDIR /root/project 
COPY --from=builderImage /go/src/project/application .
COPY --from=builderImage /go/src/project/etc .
EXPOSE 8080 
ENTRYPOINT [ "./{{.Use}}" ]
`

// _JenkinsTemaplate jenkinsfile的生成文档
const _JenkinsTemaplate string = `
// 以下代码,由代码生成工具生成,请不要手写
pipeline {
  agent any
  environment {
    // VERSION 当前版本
    PROJECT = "{{.Use}}"
    VERSION = "{{.AppVersion}}"
    // APP_NAME app的名称
    APP_NAME = "{{.use}}"
    // BASE_URL docker基础的url
    BASE_URL = "{{.DockerHub}}"
    // QY_WEIXIN_WEBHOOK 企业微信发起的webhook
    QY_WEIXIN_WEBHOOK = ""  
    // GO_PRO GO的工程 
    GO_PRO = "" 
  }
  stages {
    // build 编译阶段
    stage("build") {
      parallel {
        // 开发环境构建
         stage ('build_develop') {
           when {
             environment name:"GIT_BRANCH",value:"origin/develop"
           }
           steps {
             sh (label:'build_develop', script:
                ''' 
                  echo "build_develop"
                '''
              ) 
           }
         }
         // 测试环境构建
         stage ('build_release') {
           when {
             environment name:"GIT_BRANCH",value:"origin/release"
           }
           steps {
             sh "echo 'build,release'"
           }
         }
         // 产品级别的发布 
         stage ('build_prod') {
           when {
             environment name:"GIT_BRANCH",value:"origin/master"
           }
           steps {
             sh "echo 'helllo,prod'"
           }
         }
      }  
    }
    // test 测试阶段.
    stage("test") {
      parallel {
        // 开发环境中的测试
        stage ("test_develop") {
          when {
            environment name:"GIT_BRANCH",value:"origin/develop" 
          }
          steps{
            sh "echo 'test,develop,开发包测试'"
          }
        }
        stage ("test_release") {
          when {
            environment name:"GIT_BRANCH",value:"origin/release"
          }
          steps{
            sh "echo 'test,develop,预发布测试'"
          }
        }
        stage ("test_prod") {
          when {
            environment name:"GIT_BRANCH",value:"origin/master"
          }
          steps {
            sh "echo 'test,prod,你是没得测试的'"
          }
        }
      }
    }
    stage("deploy"){
      parallel {
        stage ("deploy_develop") {
          when {
            environment name:"GIT_BRANCH",value:"origin/develop" 
          }
          steps {
            sh (label:"deploy_develop_script",script:
              '''
                echo "启动脚本"
                sudo docker-compose up -d
              '''
            )
          }
        }
        stage ("deploy_release") {
          when {
            environment name:"GIT_BRANCH",value:"origin/release" 
          }
          steps {
            sh "echo 'deploy,release 测试环境部署'"
          }
        }
        stage ("deploy_prod") {
           when {
             environment name:"GIT_BRANCH",value:"origin/master" 
           } 
           steps {
             sh "echo 'deploy,prod 生产环境部署'"
           }
        }
      }
    }
  }
}
`

const _MainTemplate string = `
// // 以下代码,由代码生成工具生成,请不要手写
package main

import (
	"{{.Package}}/cmdline"
)

func main() {
	cmdline.Execute()
}
`

const _BuildLinuxTamplate string = `
#!/bin/bash
os=$(go env GOOS)
arch=$(go env GOARCH)
goversion=$(go version | awk '{print $3}')
commitid=$(git rev-parse --short HEAD)
account=$(git log --pretty=format:"%%an" -1)
branch=$(git branch --show-current)
nowtime=$(date +%Y-%m-%d.%H:%M:%S)
go build -ldflags "-X {{.Package}}/cmdline._GitBranch=${branch} -X {{.Package}}/cmdline._OS=${os} -X {{.Package}}/cmdline._Arch=${arch} -X {{.Package}}/cmdline._GoVersion=${goversion} -X {{.Package}}/cmdline._GitCommit=${commitid} -X {{.Package}}/cmdline._GitAccount=${account} -X {{.Package}}/cmdline._DateTime=${nowtime}" -o {{.Use}}
`

const _BuildWinTemplate string = `
@echo off
for /F %%i in ('go env GOOS') do ( set os=%%i)
for /F %%i in ('go env GOARCH') do ( set arch=%%i)
for /F %%i in ('go env GOVERSION') do ( set goversion=%%i)
for /F %%i in ('git rev-parse --short HEAD') do ( set commitid=%%i)
for /F %%i in ('git log --pretty^=format:"%%an" -1') do ( set account=%%i)
for /F "tokens=* delims=" %%i in ('git branch --show-current') do ( set branch=%%i)
for /f "tokens=* delims=" %%i in ('echo %date:~0,4%-%date:~5,2%-%date:~8,2%.%time:~1,1%:%time:~3,2%:%time:~6,2%') do ( set nowtime="%%i")
go build -ldflags "-X {{.Package}}/cmdline._GitBranch=%branch% -X {{.Package}}/cmdline._OS=%os% -X {{.Package}}/cmdline._Arch=%arch% -X {{.Package}}/cmdline._GoVersion=%goversion% -X {{.Package}}/cmdline._GitCommit=%commitid% -X {{.Package}}/cmdline._GitAccount=%account% -X {{.Package}}/cmdline._DateTime=%nowtime%" -o {{.Use}}.exe
`

// _FlagsTemplate 标签模板
const _FlagsTemplate string = `
package flags
// GlobalFlag 全局Flag
type GlobalFlag struct {
	{{range .GlobalFlags}}
	{{GetVariable .FlagName}} *{{.FlagType}}
	{{end}}
}
{{range .Commands }}
// Local{{GetVariable .Use}}Flag {{.Use}}的Flag
type Local{{GetVariable .Use}}Flag struct {
	{{range .Flags}}
	{{GetVariable .FlagName}} *{{.FlagType}}	
	{{end}}
}
{{end}}
`

const _CodesTemplate string = `
package cmdline
import(
	"fmt"
	"github.com/yuansudong/cobra"
	"{{.Package}}/flags"
	{{range .Commands}}
	{{.Use}} "{{.PkgPath}}"
	{{end}}
)

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
		fmt.Println("App Name     :   ", "{{.Use}}")
		fmt.Println("App Version  :   ", "{{.AppVersion}}")
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
	mGlobalsFlags :=  new(flags.GlobalFlag)
	mRootCommand := new(cobra.Command)
	mRootCommand.Use = "{{.Use}}" 
	mRootCommand.Long = "{{.Long}}"
	{{range .GlobalFlags}}
		{{if eq .FlagType "float32"}}
		mGlobalsFlags.{{GetVariable .FlagName}} = mRootCommand.PersistentFlags().Float32(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "float64"}}
		mGlobalsFlags.{{GetVariable .FlagName}} = mRootCommand.PersistentFlags().Float64(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int8"}}
		mGlobalsFlags.{{GetVariable .FlagName}} = mRootCommand.PersistentFlags().Int8(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int16"}}
		mGlobalsFlags.{{GetVariable .FlagName}} = mRootCommand.PersistentFlags().Int16(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int32"}}
		mGlobalsFlags.{{GetVariable .FlagName}} = mRootCommand.PersistentFlags().Int32(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int64"}}
		mGlobalsFlags.{{GetVariable .FlagName}} = mRootCommand.PersistentFlags().Int64(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int"}}
		mGlobalsFlags.{{GetVariable .FlagName}} = mRootCommand.PersistentFlags().Int(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "string"}}
		mGlobalsFlags.{{GetVariable .FlagName}} = mRootCommand.PersistentFlags().String(
			"{{.FlagName}}",
			"{{.FlagValue}}", 
			"{{.FlagUsage}}",
		)
		{{end}}
	{{end}}
	{{range .Commands}}
    mRootCommand.AddCommand(_InitSub{{GetVariable .Use}}(mGlobalsFlags))
	{{end}}
	mRootCommand.AddCommand(_InitVersion())
	return mRootCommand
}

// Execute 执行入口
func Execute() error {
	return _Root.Execute()
}
{{range .Commands}}
func _InitSub{{GetVariable .Use}}(mGlobal *flags.GlobalFlag) *cobra.Command {
	mLocal := new(flags.Local{{GetVariable .Use}}Flag)
	mCommand :=  new(cobra.Command)
	mCommand.Use = "{{.Use}}"
	mCommand.Long = "{{.Long}}"
	mCommand.Short = "{{.Short}}"
	mCommand.Run = func(cmd *cobra.Command, args []string) {
		{{.Use}}.{{.Function}}(mGlobal,mLocal)
	} 
	{{range .Flags}}
		{{if eq .FlagType "float32"}}
		mLocal.{{GetVariable .FlagName}} = mCommand.PersistentFlags().Float32(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "float64"}}
		mLocal.{{GetVariable .FlagName}} = mCommand.PersistentFlags().Float64(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int8"}}
		mLocal.{{GetVariable .FlagName}} = mCommand.PersistentFlags().Int8(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int16"}}
		mLocal.{{GetVariable .FlagName}} = mCommand.PersistentFlags().Int16(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int32"}}
		mLocal.{{GetVariable .FlagName}} = mCommand.PersistentFlags().Int32(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int64"}}
		mLocal.{{GetVariable .FlagName}} = mCommand.PersistentFlags().Int64(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "int"}}
		mLocal.{{GetVariable .FlagName}} = mCommand.PersistentFlags().Int(
			"{{.FlagName}}",
			{{.FlagValue}}, 
			"{{.FlagUsage}}",
		)
		{{else if eq .FlagType "string"}}
		mLocal.{{GetVariable .FlagName}} = mCommand.PersistentFlags().String(
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
