#!/usr/bin/env bash

rm -rf custom-prompt_*_amd64

mkdir -p deb
mkdir -p deb/DEBIAN
mkdir -p deb/usr/local/bin
mkdir -p deb/etc/custom-prompt

cp control deb/DEBIAN

cp -r prompts deb/etc/custom-prompt

VERSION=$(git describe --tags --abbrev=0)
DEB_VERSION=$(echo $VERSION | sed 's/v//' | sed 's/\./-/2')

sed -i "s/VERSION_TAG/$DEB_VERSION/" deb/DEBIAN/control

GOARCH=amd64
GOOS=linux

go build

cp custom-prompt deb/usr/local/bin

DEB_FOLDER="custom-prompt_${DEB_VERSION}_${GOARCH}"

mv deb $DEB_FOLDER

dpkg-deb --build --root-owner-group $DEB_FOLDER