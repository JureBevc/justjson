@echo off
SETLOCAL ENABLEDELAYEDEXPANSION

REM Define output directory
SET OUTPUT_DIR=build

REM Create output directory if it doesn't exist
IF NOT EXIST "%OUTPUT_DIR%" (
    mkdir "%OUTPUT_DIR%"
)

REM Build for Linux
SET GOOS=linux
SET GOARCH=amd64
go build -o "%OUTPUT_DIR%\JustJSON-linux-amd64"
SET GOARCH=arm64
go build -o "%OUTPUT_DIR%\JustJSON-linux-arm64"

REM Build for macOS
SET GOOS=darwin
SET GOARCH=amd64
go build -o "%OUTPUT_DIR%\JustJSON-darwin-amd64"
SET GOARCH=arm64
go build -o "%OUTPUT_DIR%\JustJSON-darwin-arm64"

REM Build for Windows
SET GOOS=windows
SET GOARCH=amd64
go build -o "%OUTPUT_DIR%\JustJSON-windows-amd64.exe"
SET GOARCH=386
go build -o "%OUTPUT_DIR%\JustJSON-windows-386.exe"

REM Clean up environment variables
SET GOOS=
SET GOARCH=

echo Build completed! Binaries are located in the %OUTPUT_DIR% directory.

ENDLOCAL
