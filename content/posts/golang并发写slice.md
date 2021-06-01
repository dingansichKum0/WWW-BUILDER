+++
title = "golang并发写slice"
author = ["rx-78-kum0"]
description = "线程安全的slice"
date = 2021-05-21
lastmod = 2021-05-24T10:29:07+08:00
tags = ["golang"]
categories = ["code"]
draft = false
+++

```go
package sliceSafe

type SliceSafe struct {
  channel chan int `desc:"即将加入到数据slice的数据"`
  data    []int    `desc:"数据slice"`
}

// 新建一个size大小缓存的active object对象
func New(size int, done func()) *SliceSafe {
  s := &SliceSafe{
    channel: make(chan int, size),
    data:    make([]int, 0),
  }

  go func() {
    s.schedule()
    done()
  }()
  return s
}

// 把管道中的数据append到slice中
func (s *SliceSafe) schedule() {
  for v := range s.channel {
    s.data = append(s.data, v)
  }
}

// 增加一个值
func (s *SliceSafe) Add(v int) {
  s.channel <- v
}

// 管道使用完关闭
func (s *SliceSafe) Close() {
  close(s.channel)
}

// 返回slice
func (s *Service) Slice() []int {
  return s.data
}
```
