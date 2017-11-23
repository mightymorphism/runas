# Copyright (c) 2017 Trough Creek Holdings, LLC.  All Rights Reserved

ROOT=$(abspath ${CURDIR})

SUBDIR=src

include ${ROOT}/mk/Makefile.vars
include ${ROOT}/mk/Makefile.golang

check:
