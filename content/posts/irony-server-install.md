+++
title = "irony-server-install"
author = ["dingansichKum0"]
description = "安装c++补全后端irony"
date = 2020-05-29
lastmod = 2021-07-29T18:05:59+08:00
tags = ["emacs"]
categories = ["code"]
draft = false
+++

## Mac {#mac}

1.  首先需安装cmake和llvm

    ```shell
    brew install cmake llvm
    ```

2.  emacs里执行M-x irony-install-server得到安装执行命令

    ```shell
    cmake -DCMAKE_INSTALL_PREFIX\=/Users/kumotyou/.emacs.d/irony/ \
    /Users/kumotyou/.emacs.d/elpa/irony-20200130.849/server \
    && cmake --build . --use-stderr --config Release --target install
    ```

3.  添加环境变量

    ```shell
    -DCMAKE_PREFIX_PATH=/usr/local/opt/llvm
    -DCMAKE_INSTALL_RPATH_USE_LINK_PATH=ON
    -DLIBCLANG_INCLUDE_DIR=/usr/local/opt/llvm/include
    -DLIBCLANG_LIBRARY=/usr/local/opt/llvm/lib/libclang.dylib
    ```

4.  最终编译命令

    ```shell
    cmake -DCMAKE_INSTALL_PREFIX\=/Users/kumotyou/.emacs.d/irony/ \
    -DCMAKE_PREFIX_PATH=/usr/local/opt/llvm \
    -DCMAKE_INSTALL_RPATH_USE_LINK_PATH=ON \
    -DLIBCLANG_INCLUDE_DIR=/usr/local/opt/llvm/include \
    -DLIBCLANG_LIBRARY=/usr/local/opt/llvm/lib/libclang.dylib \
    /Users/kumotyou/.emacs.d/elpa/irony-20200130.849/server \
    && cmake --build . --use-stderr --config Release --target install
    ```


## Ubuntu {#ubuntu}

1.  拉取ccls源码, 并进入ccls根目录

    ```shell
    git clone https://github.com/MaskRay/ccls

    cd ccls
    ```

2.  拉取第三方依赖. (主要是rapidjson)

    ```shell
    git submodule update --init --recursive
    ```

3.  下载llvm的二进制包并解压

    ```shell
    wget -c http://releases.llvm.org/9.0.0/clang+llvm-9.0.0-x86_64-linux-gnu-ubuntu-18.04.tar.xz

    tar xf clang+llvm-9.0.0-x86_64-linux-gnu-ubuntu-18.04.tar.xz
    ```

4.  在根目录下执行cmake 执行结果保存到Release文件夹中

    ```shell
    cmake -H. -BRelease -DCMAKE_BUILD_TYPE=Release -DCMAKE_PREFIX_PATH=/path/to/clang+llvm-9.0.0-x86_64-linux-gnu-ubuntu-18.04

    cmake --build Release
    ```

5.  开始编译

    ```shell
    cd Release

    make -j4 #使用4线程编译
    ```

6.  编译完成, 安装

    ```shell
    sudo make install
    ```
