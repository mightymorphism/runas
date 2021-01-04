# Copyright (c) 2017-2021 Trough Creek Holdings, LLC.  All Rights Reserved.

ROOT=$(abspath ${CURDIR})

SUBDIR=src\
	docker

RECURSIVE_TARGETS := init depend clean nuke

include ${ROOT}/mk/Makefile.vars
include ${ROOT}/mk/Makefile.golang
include ${ROOT}/mk/Makefile.docker
include ${ROOT}/mk/Makefile.package

check:

all:
	@echo "Please specify target: build, docker, release"; exit 1

build:
	${MAKE} -C src all

release_depend:
	$(MAKE) -C docker all

release_version:
	bash mk/packaging/mkversion.sh src/runas
	bash mk/packaging/mkversion.sh src/waitfor

release: git_no_untracked release_depend
	docker-compose -f docker/build/compose-build.yml up
