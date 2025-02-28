include config.mk

default: clean build

list:
	@echo "## list"
	@echo "### local packages"
	go list ./...
	@echo
	@echo "### all dep packages"
	go list -m all
	@echo
	@echo "### upgradeable packages"
	go list -m -u all

upgrade: list
	@echo "## upgrade"
	go get -u

fmt-show-diff:
	@echo "## fmt-show-diff"
	gofmt -s -l -d ${GITGOFIELS}

fmt:
	@echo "## fmt"
	gofmt -s -w ${GITGOFIELS}

tidy:
	@echo "## tidy"
	go mod tidy

test:
	@echo "## test"
	go test -cpu=1,2,4 -v -tags integration ./...

vet:
	@echo "## vet"
	go vet $(echo ${GITGOFIELS}|grep -v /vendor/)

build: fmt-show-diff fmt tidy test vet
	@echo "## build"
	go mod download
	go build .
	@chmod +x ./lazy
	@mkdir -p ${TARGETDIR}/
	@cp lazy ${TARGETDIR}/
	@echo "# summary:"
	@echo -e "project: ${PROJECT}\nversion: ${VERSION}\ntimestamp: ${TIMESTAMP}\nbranch: ${GITBRANCH}\ncommit: ${GITCOMMIT}"  > ${TARGETDIR}/release-${PROJECT}-${VERSION}-${GITBRANCH}-${GITCOMMIT}-${UNIX}
	@cat  ${TARGETDIR}/release*

install: clean build uninstall
	@echo "## install"
	rm -rf ~/.config/lazy
	mkdir -p ~/.config/lazy && cp -rf etc ~/.config/lazy/
	sudo cp -f ./lazy ${INSTALLPATH}
	@echo "install done"

uninstall:
	@echo "## uninstall"
	rm -rf ~/.config/lazy
	sudo rm -f ${INSTALLPATH}/${PROJECT}
	@echo "uninstall done"

clean:
	@echo "## clean"
	rm -rf ${PROJECT} ${TARGETDIRBASE}/*

.PHONY: default list upgrade fmt-show-diff fmt tidy test vet build clean
