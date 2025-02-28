PROJECT=lazy
VERSION=V1.0.0

INSTALLPATH = /usr/local/bin/
TIMESTAMP   = $(shell date '+%Y-%m-%d %H:%M:%S')
UNIX        = $(shell date -d '${TIMESTAMP}' +%s)
GITBRANCH   = $(shell git rev-parse --abbrev-ref HEAD)
GITCOMMIT   = $(shell git rev-parse HEAD)
GITGOFIELS ?= $(shell git ls-files '*.go')

TARGETDIRBASE = build
TARGETDIR     = ${TARGETDIRBASE}/${PROJECT}-${UNIX}
