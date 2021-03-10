#!/bin/bash
os=$(go env GOOS)
arch=$(go env GOARCH)
goversion=$(go version | awk '{print $3}')
commitid=$(git rev-parse --short HEAD)
account=$(git log --pretty=format:"%%an" -1)
branch=$(git branch --show-current)
nowtime=$(date +%Y-%m-%d.%H:%M:%S)
appversion=1.6
appname=example

go build -ldflags "-X main._GitBranch=${branch} -X main._AppName=${appname} -X main._AppVersion=${appversion} -X main._OS=${os} -X main._Arch=${arch} -X main._GoVersion=${goversion} -X main._GitCommit=${commitid} -X main._GitAccount=${account} -X main._DateTime=${nowtime}" -o ${appname}
