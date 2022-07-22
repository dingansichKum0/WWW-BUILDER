+++
title = "限制goroutine数量"
author = ["zakudriver"]
description = "golang限制goroutine数量"
date = 2020-07-02
lastmod = 2022-07-22T10:18:20+08:00
tags = ["golang"]
categories = ["code"]
draft = false
+++

```go
package main

import (
  "fmt"
  "sync"
  "time"
)

type Glimit struct {
  n int
  c chan struct{}
}

func New(n int) *Glimit {
  return &Glimit{
    n: n,
    c: make(chan struct{}, n),
  }
}

func (g *Glimit) Run(f func()) {
  g.c <- struct{}{}
  go func() {
    f()
    <-g.c
  }()
}

var wg = sync.WaitGroup{}

func main() {
  number := 10
  g := New(2)
  for i := 0; i < number; i++ {
    wg.Add(1)
    value := i

    g.Run(func() {
      // 做一些业务逻辑处理
      fmt.Printf("go func: %d\n", value)
      time.Sleep(time.Second)
      wg.Done()
    })
  }
  wg.Wait()
}
```
