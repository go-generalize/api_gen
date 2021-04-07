#! /bin/bash

set -eu

WORKDIR=$(dirname $0)
cd $WORKDIR
WORKDIR=$(pwd) # Make $WORKDIR abolute

export PATH=$WORKDIR/../bin:$PATH

clients=(ts go)
targets=(empty_root standard)

for t in ${targets[@]}; do
    pushd ${t}

    mkdir -p server clients
    cd server && api_gen s ../api
    cd ../clients
    for c in ${clients[@]}; do
        mkdir -p ${c}
        pushd ${c}

        api_gen c $c ../../api

        popd
    done

    popd
done
