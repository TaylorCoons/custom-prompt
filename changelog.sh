#/usr/bin/env bash

VERSION=$(git describe --tags --abbrev=0)
git log --format="%s" HEAD...$VERSION