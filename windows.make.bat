# make file for window
@echo off
set APP_NAME=myapp
set BUILD_DIR=build

echo Creating build directory...
if not exist %BUILD_DIR% mkdir %BUILD_DIR%

echo Building the Go application...
set GOOS=windows
set GOARCH=amd64
go build -o %BUILD_DIR%\%APP_NAME%.exe main.go

echo Build complete! Executable is located in %BUILD_DIR%.
