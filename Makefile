include config.mk

default: clean build

build:
	@echo "## build"
	@echo "--------"
	cargo build --release
	@echo "build done"

release: clean build
	@echo "## release"
	@echo "--------"
	cp -f lazy.yaml ./target/release/lazy.yaml
	mv ./target/release ./target/release-${PROJECT}-${VERSION}-${GITBRANCH}-${GITCOMMIT}-${UNIX}
	@echo "release done"

install: clean build uninstall
	@echo "## install"
	@echo "--------"
	rm -rf ~/.config/lazy
	mkdir -p ~/.config/lazy && cp -f lazy.yaml ~/.config/lazy/
	sudo cp -f ./target/release/lazy ${INSTALLPATH}
	@echo "install done"

uninstall:
	@echo "## uninstall"
	@echo "--------"
	rm -f ~/.config/lazy/lazy.yaml
	sudo rm -f ${INSTALLPATH}/${PROJECT}
	@echo "uninstall done"

clean:
	@echo "## clean"
	@echo "--------"
	rm -rf ./target
	@echo "clean done"

.PHONY: default build install clean
