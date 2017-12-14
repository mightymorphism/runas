# Copyright (c) 2017 Trough Creek Holdings, LLC.  All Rights Reserved

ROOT=$(abspath ${CURDIR}/..)

export GOPATH:=${ROOT}

GO_PROGS=runas

include ${ROOT}/mk/Makefile.vars
include ${ROOT}/mk/Makefile.golang

all: go_prep go_build

update: go_update

depend: go_depend

fmt:
	find . -type d \( -name vendor -o -name golang.org \) -prune -o -name '*.go' -print | xargs -n 1 gofmt -w


go_prep:
	go get golang.org/x/sys/unix

go_update:
	go get -u golang.org/x/sys/unix
