#!/bin/bash
mkdir -p output
go mod tidy
SUPPORT_GOOS=("freebsd" "darwin" "linux" "windows")
BSD_SUPPORT_GOARCH=("386" "amd64" "arm64" "arm")
LINUX_SUPPORT_GOARCH=("386" "amd64" "arm64" "arm")
MAC_SUPPORT_GOARCH=("amd64" "arm64")
WIN_SUPPORT_GOARCH=("386" "amd64")

NAME="terraform-provider-volcengine"
set -ea

for goos in "${SUPPORT_GOOS[@]}"
do
  if [ "$goos" == "freebsd" ]
  then
     for goarch in "${BSD_SUPPORT_GOARCH[@]}"
     do
        echo "build $goos/$goarch/$NAME"
        CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build  -o $NAME
        mkdir -p  output/"$goos"/"$goarch"
        mv "$NAME" output/"$goos"/"$goarch"/"$NAME"
     done
  fi
  if [ "$goos" == "darwin" ]
    then
       for goarch in "${MAC_SUPPORT_GOARCH[@]}"
       do
          echo "build $goos/$goarch/$NAME"
          CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build  -o $NAME
          mkdir -p  output/"$goos"/"$goarch"
          mv "$NAME" output/"$goos"/"$goarch"/"$NAME"
       done
  fi
  if [ "$goos" == "linux" ]
    then
       for goarch in "${LINUX_SUPPORT_GOARCH[@]}"
       do
          echo "build $goos/$goarch/$NAME"
          CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build  -o $NAME
          mkdir -p  output/"$goos"/"$goarch"
          mv "$NAME" output/"$goos"/"$goarch"/"$NAME"
       done
  fi
  if [ "$goos" == "windows" ]
    then
       for goarch in "${WIN_SUPPORT_GOARCH[@]}"
       do
          echo "build $goos/$goarch/$NAME"
          CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build  -o $NAME
          mkdir -p  output/"$goos"/"$goarch"
          mv "$NAME" output/"$goos"/"$goarch"/"$NAME"
       done
    fi
done