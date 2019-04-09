#!/usr/bin/env bash
BUILDVERSIONMAJOR=0
BUILDVERSIONMINOR=1
BUILDVERSIONPATCH=0
BUILDVERSIONFULL="$BUILDVERSIONMAJOR.$BUILDVERSIONMINOR.$BUILDVERSIONPATCH"
BUILDTIMESTAMP=$(date)

go build -o bin/ms -ldflags="-X 'github.com/tylerdmace/mumblestumble/version.BuildVersion=$BUILDVERSIONFULL' -X 'github.com/tylerdmace/mumblestumble/version.BuildTimestamp=$BUILDTIMESTAMP'" .