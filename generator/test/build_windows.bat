
@echo off
for /F %%i in ('go env GOOS') do ( set os=%%i)
for /F %%i in ('go env GOARCH') do ( set arch=%%i)
for /F %%i in ('go env GOVERSION') do ( set goversion=%%i)
for /F %%i in ('git rev-parse --short HEAD') do ( set commitid=%%i)
for /F %%i in ('git log --pretty^=format:"%%an" -1') do ( set account=%%i)
for /F "tokens=* delims=" %%i in ('git branch --show-current') do ( set branch=%%i)
for /f "tokens=* delims=" %%i in ('echo %date:~0,4%-%date:~5,2%-%date:~8,2%.%time:~1,1%:%time:~3,2%:%time:~6,2%') do ( set nowtime="%%i")
go build -ldflags "-X cmdline._GitBranch=%branch% -X cmdline._OS=%os% -X cmdline._Arch=%arch% -X cmdline._GoVersion=%goversion% -X cmdline._GitCommit=%commitid% -X cmdline._GitAccount=%account% -X cmdline._DateTime=%nowtime%" -o %appname%.exe
