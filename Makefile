# Copyright (c) 2017 Trough Creek Holdings, LLC.  All Rights Reserved

ROOT=$(abspath ${CURDIR})

SUBDIR=src\
	docker

RECURSIVE_TARGETS := init depend clean nuke

include ${ROOT}/mk/Makefile.vars
include ${ROOT}/mk/Makefile.golang
include ${ROOT}/mk/Makefile.package

check:

all:
	@echo "Please specify target: build, docker, release"; exit 1

build:
	${MAKE} -C src all

release_depend:
	$(MAKE) -C docker all

release: git_no_untracked release_depend
release:
	bash mk/packaging/mkversion.sh src/runas
