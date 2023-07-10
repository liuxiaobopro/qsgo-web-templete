BIN_FILE=qsgo-web-templete
BIN_DIR=.
BIN_PATH=$(BIN_DIR)/$(BIN_FILE)

tidy:
	@export GO111MODULE=on
	@export GOPROXY=https://goproxy.cn
	@go mod tidy

build: windows

check:
	@echo "Checking..."
	@go fmt
	@golangci-lint run
	@echo "Done! $(shell date "+%Y-%m-%d %H:%M:%S")"

windows:
	@echo "Building for Windows..."
	@go fmt
	@go build -o $(BIN_PATH).exe
	@echo "Done! $(shell date "+%Y-%m-%d %H:%M:%S")"

linux:
	@echo "Building for Linux..."
	@go fmt
	@go build -o $(BIN_PATH)
	@echo "Done! $(shell date "+%Y-%m-%d %H:%M:%S")"

clean:
	@echo "Cleaning..."
	@rm -rf $(BIN_PATH).exe && rm -rf $(BIN_PATH)

run: tidy clean windows
	@echo "Running..."
	@$(BIN_PATH).exe $(ARGS)

prod:
	@echo "正在重置分支并拉取代码..."
	@git reset --hard && git pull
	@echo "正在编译..."
	@cd /www/wwwroot/qsgo-web-templete
	@go fmt && go mod tidy && go build -o $(BIN_PATH)
	@echo "Done! $(shell date "+%Y-%m-%d %H:%M:%S")"
