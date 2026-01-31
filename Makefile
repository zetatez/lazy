APP_NAME := lazy
INSTALL_DIR := /usr/local/bin
LDFLAGS := -s -w

.PHONY: all build install clean release

all: build

build:
	@echo "==> Building $(APP_NAME)..."
	go build -ldflags="$(LDFLAGS)" -o ./$(APP_NAME) .

install: build
	@echo "==> Installing $(APP_NAME) to $(INSTALL_DIR)..."
	@sudo cp -f ./$(APP_NAME) $(INSTALL_DIR)/

uninstall:
	@echo "==> Uninstalling $(APP_NAME)..."
	@sudo rm -f $(INSTALL_DIR)/$(APP_NAME)

clean:
	@echo "==> Cleaning..."
	@rm -f ./$(APP_NAME)

release: build
	@echo "==> Creating release for $(APP_NAME)..."
	@echo "Binary size: $$(du -h ./$(APP_NAME) | cut -f1)"
