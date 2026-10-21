@echo off
echo Start building for Go Windows.
if exist orange-crm.exe (
    echo Old files exist.
    del /f /s /q orange-crm.exe
    echo The old file has been deleted.
) else (
    echo The old file does not exist.
)
set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go build -o orange-crm.exe main.go
echo Compilation completed.