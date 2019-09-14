#!/bin/bash

cd "$(dirname $0)"
DIRS=". configuration"
set -e
for subdir in $DIRS; do
  pushd $subdir
  go vet
  popd
done