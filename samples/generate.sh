#! /bin/bash

set -eu

WORKDIR=$(dirname $0)
cd $WORKDIR
WORKDIR=$(pwd) # Make $WORKDIR abolute
SWAG=${PWD}/../bin/swag

export PATH=$WORKDIR/../bin:$PATH

clients=(ts dart)
# clients=(ts go dart)
targets=(empty_root standard)

for t in ${targets[@]}; do
    pushd ${t}

    mkdir -p server clients
    # cd server && api_gen s ../api
    # cd ../clients
    cd clients
    for c in ${clients[@]}; do
        mkdir -p ${c}
        pushd ${c}

        api_gen c $c ../../api

        popd
    done
    cd ../
    # ${SWAG} init --parseDependency

    popd
done
