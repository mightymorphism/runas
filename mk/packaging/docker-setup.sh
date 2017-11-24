#! /bin/bash
# Copyright (c) 2017 Trough Creek Holdings, LLC.  All Rights Reserved

set -e

src="/src"
over="/over"
home="/home/build"

mkdir ${over}
mount -t tmpfs ${over} ${over}
mkdir -p ${over}/upper ${over}/work

useradd -m -s /bin/bash -d $home build
mkdir -p ${home}/src
chown -R build:build ${home}

mount -t overlay overlay -o lowerdir=${src},upperdir=${over}/upper,workdir=${over}/work ${home}/src
