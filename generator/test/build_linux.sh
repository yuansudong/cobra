
#!/bin/bash
os=$(go env GOOS)
arch=$(go env GOARCH)
goversion=$(go version | awk '{print $3}')
commitid=$(git rev-parse --short HEAD)
account=$(git log --pretty=format:"%%an" -1)
branch=$(git branch --show-current)
nowtime=$(date +%Y-%m-%d.%H:%M:%S)
go build -ldflags "-X cmdline._GitBranch=${branch} -X cmdline._OS=${os} -X cmdline._Arch=${arch} -X cmdline._GoVersion=${goversion} -X cmdline._GitCommit=${commitid} -X cmdline._GitAccount=${account} -X cmdline._DateTime=${nowtime}" -o ${appname}
