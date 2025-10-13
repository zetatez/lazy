APP_NAME := lazy
INSTALL_DIR := /usr/local/bin
CONFIG_DIR := $(HOME)/.config/lazy
CONFIG_FILE := $(CONFIG_DIR)/config.yaml

.PHONY: all build install clean

all: build

build:
	@echo "==> Building $(APP_NAME)..."
	go build -o ./$(APP_NAME) .

install: build
	@echo "==> Installing $(APP_NAME) to $(INSTALL_DIR)..."
	@sudo cp -f ./$(APP_NAME) $(INSTALL_DIR)/
	@echo "==> Initializing config..."
	@mkdir -p $(CONFIG_DIR)
	cp -f config.yaml $(CONFIG_DIR)/config.yaml

uninstall:
	@echo "==> Uninstalling $(APP_NAME)..."
	@sudo rm -f $(INSTALL_DIR)/$(APP_NAME)
	@echo "==> Removing config..."
	@rm -rf $(CONFIG_DIR)

clean:
	@echo "==> Cleaning..."
	@rm -rf ./$(APP_NAME)

