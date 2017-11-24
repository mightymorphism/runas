# Copyright (c) 2017 Trough Creek Holdings, LLC.  All Rights Reserved

ROOT=$(abspath ${CURDIR})

SUBDIR=src\
	docker

include ${ROOT}/mk/Makefile.vars
include ${ROOT}/mk/Makefile.golang
include ${ROOT}/mk/Makefile.package

check:

build:
	make -C src all

version:
	bash mk/packaging/mkversion.sh src/runas

release: git_no_untracked
	docker-compose -f docker/build/compose-build.yml run build
