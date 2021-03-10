@echo off
for /F %%i in ('go env GOOS') do ( set os=%%i)
for /F %%i in ('go env GOARCH') do ( set arch=%%i)
for /F %%i in ('go env GOVERSION') do ( set goversion=%%i)
for /F %%i in ('git rev-parse --short HEAD') do ( set commitid=%%i)
for /F %%i in ('git log --pretty^=format:"%%an" -1') do ( set account=%%i)
for /F "tokens=* delims=" %%i in ('git branch --show-current') do ( set branch=%%i)
for /f "tokens=* delims=" %%i in ('echo %date:~0,4%-%date:~5,2%-%date:~8,2%.%time:~1,1%:%time:~3,2%:%time:~6,2%') do ( set nowtime="%%i")
set appversion=1.6
set appname=example
go build -ldflags "-X main._GitBranch=%branch% -X main._AppName=%appname% -X main._AppVersion=%appversion% -X main._OS=%os% -X main._Arch=%arch% -X main._GoVersion=%goversion% -X main._GitCommit=%commitid% -X main._GitAccount=%account% -X main._DateTime=%nowtime%" -o %appname%.exe
