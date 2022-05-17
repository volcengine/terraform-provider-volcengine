#!/bin/bash
OS=$1
NAME="terraform-provider-vestack"
set -ea

if [ "$OS" == "" ]
then
  OS="darwin"
fi

CGO_ENABLED=0 GOOS=$OS GOARCH=amd64 go build  -o $NAME
rm -f $GOPATH/bin/$NAME
cp $GOPATH/src/code.byted.org/iaasng/terraform-provider-vestack/$NAME $GOPATH/bin/
# 如果terraform版本高于或者等于0.13
# 需要执行如下三条指令 来映射CLI到本地路径
# 如果小于此版本可以不做这三个操作
# shellcheck disable=SC2154
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/hashicorp/vestack/0.0.1/"$OS"_amd64/
rm -f ~/.terraform.d/plugins/registry.terraform.io/hashicorp/vestack/0.0.1/"$OS"_amd64/"$NAME"_v0.0.1
cp $GOPATH/src/code.byted.org/iaasng/terraform-provider-vestack/$NAME ~/.terraform.d/plugins/registry.terraform.io/hashicorp/vestack/0.0.1/"$OS"_amd64/"$NAME"_v0.0.1

rm -f $NAME