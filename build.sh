#!/bin/bash

OPT_DELETE=0

PROG_NAME=AOC22

while getopts c:d flag
do
    case "${flag}" in
        c) challenge=${OPTARG};;
        d) OPT_DELETE=1;;
    esac
done

build()
{
    echo "Copying source"
    cp $challengeDir/main.go .

    echo "Source copied"
    echo "Building"

    go mod init $PROG_NAME
    go mod tidy
    go build -v -o $PROG_NAME main.go
    chmod u+x $PROG_NAME

    echo "Build complete"
}

if [ $OPT_DELETE -eq 1 ]; then
    echo "Deleting source and build artifacts"
    rm $PROG_NAME
    rm go.mod
    rm main.go
    echo "Clean-up complete"
fi

if ! [ -z ${challenge+x} ]; then
    challengeDir=challenge-${challenge}
    echo "Building $challengeDir"
    build challengeDir
fi

