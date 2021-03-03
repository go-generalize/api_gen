#! /bin/bash

set -eu

WORKDIR=$(dirname $0)

export PATH=$WORKDIR/../bin:$PATH
cd $WORKDIR

clients=(ts go)
targets=(empty_root)

for t in ${targets[@]}; do
    pushd ${t}

    cd server && server_generator .
    cd ../clients
    for c in ${clients[@]}; do
        mkdir -p ${c}
        pushd ${c}

        api_gen c $c ../../server

        popd
    done

    popd
done
