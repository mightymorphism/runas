#!/bin/bash
# Copyright (c) 2017 Trough Creek Holdings, LLC.  All Rights Reserved

set -e
INVOKE_DIR=$(pwd)
PACKAGE_DIR=$(mktemp -d)
trap "{ cd ${INVOKE_DIR}; rm -Rf $PACKAGE_DIR; }" EXIT

cd ${PACKAGE_DIR}

echo "BUILD_ROOT: ${BUILD_ROOT}"
echo "RELEASE: ${RELEASE}"
echo "RELEASE_DATE: ${RELEASE_DATE}"

version=$(cat ${BUILD_ROOT}/REVISION)
if [ -n "${RELEASE_DATE}" ]; then
  version="${version}-${RELEASE_DATE}"
fi

package_name="runas-${version}"

echo "Version: $version"

mkdir -p ${package_name} && cd ${package_name}
mkdir -p DEBIAN \
  usr/bin

cp ${BUILD_ROOT}/bin/runas  usr/bin/

# Debian packaging files
#cp ${BUILD_ROOT}/dist/ubuntu/spec/conffiles DEBIAN/conffiles
cp ${BUILD_ROOT}/dist/ubuntu/spec/control   DEBIAN/control
#cp ${BUILD_ROOT}/dist/ubuntu/spec/postinst  DEBIAN/postinst
#cp ${BUILD_ROOT}/dist/ubuntu/spec/prerm     DEBIAN/prerm

# Replace package version placeholder with the version value
sed -i -r -e "s/Version: VERSION/Version: ${version}/g" DEBIAN/control

# deb building command =========================================================
cd ${PACKAGE_DIR}
dpkg-deb --build ${package_name}/

# store packages ===============================================================
cp ${package_name}.deb ${BUILD_ROOT}/dist/ubuntu/deb/runas-${version}.deb
