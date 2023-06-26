include config.mk

default: clean build

list:
	@echo "## list"
	@echo "--------"
	@echo "### local packages"
	@echo "-------------------"
	go list ./...
	@echo
	@echo "### all dep packages"
	@echo "-------------------"
	go list -m all
	@echo
	@echo "### upgradeable packages"
	@echo "-------------------"
	go list -m -u all

upgrade: list
	@echo "## upgrade"
	@echo "--------"
	go get -u

fmt-show-diff:
	@echo "## fmt-show-diff"
	@echo "--------"
	gofmt -s -l -d ${GITGOFIELS}

fmt:
	@echo "## fmt"
	@echo "--------"
	gofmt -s -w ${GITGOFIELS}

tidy:
	@echo "## tidy"
	@echo "--------"
	go mod tidy

test:
	@echo "## test"
	@echo "--------"
	go test -cpu=1,2,4 -v -tags integration ./...

vet:
	@echo "## vet"
	@echo "--------"
	go vet $(echo ${GITGOFIELS}|grep -v /vendor/)

build: fmt-show-diff fmt tidy test vet
	@echo "## build"
	@echo "--------"
	go mod download
	go build .
	@chmod +x ./lazy
	@mkdir -p ${TARGETDIR}/
	@cp lazy ${TARGETDIR}/
	@echo "-----------------"
	@echo "# summary:"
	@echo -e "project: ${PROJECT}\nversion: ${VERSION}\ntimestamp: ${TIMESTAMP}\nbranch: ${GITBRANCH}\ncommit: ${GITCOMMIT}"  > ${TARGETDIR}/release-${PROJECT}-${VERSION}-${GITBRANCH}-${GITCOMMIT}-${UNIX}
	@cat  ${TARGETDIR}/release*

install: clean build uninstall
	@echo "## install"
	@echo "--------"
	rm -rf ~/.config/lazy
	mkdir -p ~/.config/lazy && cp -f lazy.yaml ~/.config/lazy/
	sudo cp -f ./lazy ${INSTALLPATH}
	@echo "install done"

uninstall:
	@echo "## uninstall"
	@echo "--------"
	rm -rf ~/.config/lazy
	sudo rm -f ${INSTALLPATH}/${PROJECT}
	@echo "uninstall done"

clean:
	@echo "## clean"
	@echo "--------"
	rm -rf ${PROJECT} ${TARGETDIRBASE}-*

.PHONY: default list upgrade fmt-show-diff fmt tidy test vet build clean
