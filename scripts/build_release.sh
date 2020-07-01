#!/bin/bash
set -e

DIST_DIR="../dist/"
OS=("darwin" "darwin" "linux" "linux" "linux" "windows" "windows")
ARCH=("amd64" "386" "amd64" "386" "arm" "amd64" "386")
TARGETS=("server_generator" "client_generator")

CURRENT_DIR=`pwd`


for target in "${TARGETS[@]}"
do
  cd $CURRENT_DIR
  cd $target

  for i in `seq 0 1 6`
  do
    GOOS=${OS[$i]}
    GOARCH=${ARCH[$i]}

    BUILD_PATH=${DIST_DIR}${target}_${GOOS}_${GOARCH}
    if [ "$GOOS" = "darwin" ]; then
      BUILD_PATH=${DIST_DIR}${target}r_macOS_${GOARCH}
    fi
    if [ "$GOOS" = "windows" ]; then
      BUILD_PATH=${DIST_DIR}${target}r_${GOOS}_${GOARCH}.exe
    fi

    env GOOS=${GOOS} GOARCH=${GOARCH}  go build -o ${BUILD_PATH}
  done
done