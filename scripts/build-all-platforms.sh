#!/bin/bash

# Build binary for all platforms for cli $1 & version $2.

set -e

DIR=`dirname $0`
ROOT=$DIR/..

CLI=$1
VSN=$2

if [ -z "$CLI" ]; then
    echo "require cli as first parameter" 2>&1
    exit 1
fi

if [ -z "$VSN" ]; then
    echo "require version as second parameter" 2>&1
    exit 1
fi

if [ ! -d "$CLI" ]; then
    echo "\"$cli\" is not a cli" 2>&1
    exit 1
fi

for t in                                                                      \
    darwin_amd64                                                              \
    linux_amd64 ;
do
    os="${t%_*}"
    arch="${t#*_}"
    output="${CLI}_${VSN}_${os}_${arch}"

    if [ "$os" == "windows" ] ; then
        output+=".exe"
    fi

    echo "building ${output}"
    GOOS=$os GOARCH=$arch go build                                            \
        -o $ROOT/build/${output}                                              \
        $ROOT/${CLI}
done

wait
