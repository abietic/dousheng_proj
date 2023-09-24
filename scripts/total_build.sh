#!/usr/bin/env bash


for i in $(find . -maxdepth 1 -type d)
do
    if [ "$i" = "." ]
    then
        continue
    fi
    cd $i
    if test -e ./output
    then
        echo `pwd`'有历史输出,执行清除'
        rm -r ./output
    fi
    if test -e ./build.sh
    then
        echo `pwd`'构建中'
        sh ./build.sh
    fi
    cd ..
done