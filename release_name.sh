#!/usr/bin/env bash

VERSION=$(git describe --tags --abbrev=0)
DEB_VERSION=$(echo $VERSION | sed 's/v//' | sed 's/\./-/2')

GOARCH=amd64
GOOS=linux

DEB_FOLDER="custom-prompt_${DEB_VERSION}_${GOARCH}"

echo $DEB_FOLDER
