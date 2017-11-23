# Copyright (c) 2017 Trough Creek Holdings, LLC.  All Rights Reserved

ROOT=$(abspath ${CURDIR}/..)

SUBDIR=

DOCKER_CONTAINERS=$(shell ${ROOT}/bin/docker-deps -p ${DOCKER_NS} -f make list .)

include ${ROOT}/mk/Makefile.vars
include ${ROOT}/mk/Makefile.docker

-include Makefile.deps

all: docker_build

depend: docker_depend

list:
	@echo ${DOCKER_CONTAINERS}
