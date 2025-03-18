# Makefile

APP_NAME := myapp
BUILD_DIR := build

.PHONY: all clean

all: $(BUILD_DIR)/$(APP_NAME).exe

$(BUILD_DIR)/$(APP_NAME).exe: main.go
	mkdir -p $(BUILD_DIR)
	set GOOS=windows
	set GOARCH=amd64
	go build -o $(BUILD_DIR)/$(APP_NAME).exe main.go

clean:
	rm -rf $(BUILD_DIR)
