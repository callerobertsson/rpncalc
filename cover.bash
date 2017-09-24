#!/bin/bash

echo Generating...

go test --coverprofile cover.out ./rpncalc

if [ $? != 0 ] ; then
    echo "Failed to generate coverage!"
    exit
fi

echo "Opening..."

go tool cover -html=cover.out

echo "Done"
