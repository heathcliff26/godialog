#!/bin/bash

set -e

echo "Check if source code is formatted"
make fmt
rc=0
git update-index --refresh && git diff-index --quiet HEAD -- || rc=1
if [ $rc -ne 0 ]; then
    echo "FATAL: Need to run \"make fmt\""
    exit 1
fi

echo "Check if dependencies are up to date"
make update-deps
rc=0
git update-index --refresh && git diff-index --quiet HEAD -- || rc=1
if [ $rc -ne 0 ]; then
    echo "FATAL: Need to run \"make update-deps\""
    exit 1
fi
