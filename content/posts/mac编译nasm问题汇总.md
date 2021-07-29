+++
title = "mac编译nasm问题汇总"
author = ["dingansichKum0"]
description = "mac编译nasm的坑"
date = 2020-05-26
lastmod = 2021-07-29T18:05:58+08:00
tags = ["汇编", "mac"]
categories = ["code"]
draft = false
+++

## ld: dynamic main executables must link with libSystem.dylib for architecture x86\_64 {#ld-dynamic-main-executables-must-link-with-libsystem-dot-dylib-for-architecture-x86-64}

```shell
# -lSystem
ld a.o -o a -lSystem
```
