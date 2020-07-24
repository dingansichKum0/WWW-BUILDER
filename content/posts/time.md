+++
title = "golang time工具函数"
author = ["rx-78-kum0"]
description = "golang常用时间工具函数"
date = 2020-06-23
lastmod = 2020-07-02T13:59:13+08:00
tags = ["golang"]
categories = ["code"]
draft = false
+++

```go
package util-time

import (
  "strconv"
  "time"
)

// 获取当前的时间 - 字符串
func GetCurrentDate() string {
  return time.Now().Format("2006/01/02 15:04:05")
}

// 获取当前的时间 - Unix时间戳
func GetCurrentUnix() int64 {
  return time.Now().Unix()
}

// 获取当前的时间 - 毫秒级时间戳
func GetCurrentMilliUnix() int64 {
  return time.Now().UnixNano() / 1000000
}

// 获取当前的时间 - 纳秒级时间戳
func GetCurrentNanoUnix() int64 {
  return time.Now().UnixNano()
}

func GetCurrentTime() string {
  var cstSh, _ = time.LoadLocation("Asia/Shanghai")
  t := time.Now().In(cstSh).Format("2006/01/02/ 15:04:05")
  return t
}

func GetCurrentHour() int {
  var cstSh, _ = time.LoadLocation("Asia/Shanghai")
  t,_ := strconv.Atoi(time.Now().In(cstSh).Format("2006010215"))
  return t
}

func GetCurrentDay() int {
  var cstSh, _ = time.LoadLocation("Asia/Shanghai")
  t,_ := strconv.Atoi(time.Now().In(cstSh).Format("20060102"))
  return t
}


func GetLastDay() int {
  var cstSh, _ = time.LoadLocation("Asia/Shanghai")
  nTime := time.Now()
  t,_ := strconv.Atoi(nTime.AddDate(0,0,-1).In(cstSh).Format("20060102"))
  return t
}

func GetlastHour() int {
  nTime := time.Now()
  lastTime := nTime.Add(time.Hour * -1)
  t,_ := strconv.Atoi(lastTime.Format("2006010215"))
  return t
}
```
