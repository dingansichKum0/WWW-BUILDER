+++
title = "golang gc优化"
author = ["rx-78-kum0"]
description = "golang gc优化"
date = 2020-06-23
lastmod = 2020-07-02T14:00:29+08:00
tags = ["golang"]
categories = ["code"]
draft = false
+++

## 小对象要合并 {#小对象要合并}


## 函数频繁创建的简单的对象，直接返回对象，效果比返回指针效果要好 {#函数频繁创建的简单的对象-直接返回对象-效果比返回指针效果要好}


## 类型转换要注意，官方用法消耗特别大。 {#类型转换要注意-官方用法消耗特别大}

```go
package string_util

import (
  "unsafe"
)

func str2bytes(s string) []byte {
  x := (*[2]uintptr)(unsafe.Pointer(&s))
  h := [3]uintptr{x[0], x[1], x[1]}
  return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
  return *(*string)(unsafe.Pointer(&b))
}
```


## 避免反复创建slice，map {#避免反复创建slice-map}

```go
func(r*Reader)Read()([]byte,error)
// 此函数没有形参，每次调用的时候返回一个[]byte。
func(r*Reader)Read(buf[]byte)(int,error)
// 此函数个函数在每次迪调用的时候，会重用形参声明。

```


## 避免使用"+"拼接字符串 {#避免使用-plus-拼接字符串}

```go
package string_utils

import (
  "strings"
)

func strAppend(s string, ss ...string) string {
  var r strings.Builder
  r.WriteString(s)
  for _, v := range ss {
    r.WriteString(v)
  }

  return r.String()
}
```
