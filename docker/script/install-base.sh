#! /bin/bash
# Copyright (c) 2017 Trough Creek Holdings, LLC.  All Rights Reserved

set -e

apt-get -y update

# Install basic shell utilities
apt-get install -y file wget --no-install-recommends
apt-get install -y vim-tiny --no-install-recommends
apt-get install -y libpq-dev postgresql-client --no-install-recommends
apt-get install -y ca-certificates --no-install-recommends
apt-get install -y locales

apt-get clean

localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8

export LOCALE="en_US.utf8"
echo 'export LOCALE=en_US.utf8' >> /etc/profile.d/locale.sh

# Update the set of installed CA certs
/usr/sbin/update-ca-certificates

# Install golang
golang_version=`/usr/local/bin/cfg-version golang`
wget --no-verbose https://storage.googleapis.com/golang/go${golang_version}.linux-amd64.tar.gz
tar -C /usr/local -xzf go${golang_version}.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile.d/golang.sh
