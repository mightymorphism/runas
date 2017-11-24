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

release:
	bash mk/packaging/mkversion.sh src/runas
	docker-compose -f docker/build/compose-build.yml run build
