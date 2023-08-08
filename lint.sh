#!/bin/bash

ROOT=$(pwd)

function localLint(){
  flag=0
  for file in `ls $ROOT"/"$1`; do
       if [ ! -d $ROOT"/"$1"/"$file ]; then
         flag=1
         break
       fi
  done
  if [ $flag == 1 ];then
      echo 'golangci-lint  '$ROOT'/'$1''
      golangci-lint run "$ROOT"/"$1"
  fi
  for file in `ls $ROOT"/"$1`; do
     if [  -d $ROOT"/"$1"/"$file ]; then
       localLint "$1"/"$file"
     fi
  done
}

localLint 'common'
localLint 'volcengine'