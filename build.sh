#!/bin/bash
OS=$1
ARCH=$2

NAME="terraform-provider-volcengine"
set -ea

if [ "$OS" == "" ]
then
  OS="darwin"
fi

if [ "$ARCH" == "" ]
then
  ARCH="arm64"
fi

CGO_ENABLED=0 GOOS=$OS GOARCH=$ARCH go build  -o $NAME
# rm -f $GOPATH/bin/$NAME
# cp $NAME $GOPATH/bin/
# 如果terraform版本高于或者等于0.13
# 需要执行如下三条指令 来映射CLI到本地路径
# 如果小于此版本可以不做这三个操作
# shellcheck disable=SC2154
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/hashicorp/volcengine/0.0.1/"$OS"_"$ARCH"/
rm -f ~/.terraform.d/plugins/registry.terraform.io/hashicorp/volcengine/0.0.1/"$OS"_"$ARCH"/"$NAME"_v0.0.1
cp $NAME ~/.terraform.d/plugins/registry.terraform.io/hashicorp/volcengine/0.0.1/"$OS"_"$ARCH"/"$NAME"_v0.0.1

rm -f $NAME
