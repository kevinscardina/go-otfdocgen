#!/usr/bin/env bash

./clean.sh && ./build.sh
cp build/otfdocgen /usr/local/bin
