#!/usr/bin/env bash

VERSION=$(git describe --tags --abbrev=0)
DEB_VERSION=$(echo $VERSION | sed 's/v//' | sed 's/\./-/2')

sed -i "s/VERSION_TAG/$DEB_VERSION/" deb/DEBIAN/control

GOARCH=amd64
GOOS=linux

go build

cp custom-prompt deb/usr/local/bin

DEB_FOLDER="custom-prompt_${DEB_VERSION}_${GOARCH}"

echo $DEB_FOLDER
