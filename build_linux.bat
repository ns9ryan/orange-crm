@echo off
echo Start building for Go Linux.
if exist orange-crm-v2 (
    echo Old files exist.
    del /f /s /q orange-crm-v2
    echo The old file has been deleted.
) else (
    echo The old file does not exist.
)
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -o orange-crm-v2 main.go
echo Compilation completed.