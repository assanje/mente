#!/bin/zsh

versionName="1.0"
#gitRevision=$(git rev-list -1 HEAD)
buildDate=$(date +"%Y-%m-%d")

# make goos goarch extension
function make {
        GOOS=$1 GOARCH=$2 go build \
                -ldflags "-X 'main.buildDate=$buildDate' -X 'main.versionName=$versionName'" \
                -o mente-$versionName-$1-$2$3 \
                ./main.go
}

#make  darwin amd64 ""
make  darwin arm64 ""
#make   linux amd64 ""
#make windows amd64 ".exe"