+++
title = "一个包含 nil 指针的接口不是 nil 接口"
author = ["zakudriver"]
description = "golang 接口值: 一个包含 nil 指针的接口不是 nil 接口."
date = 2019-03-25
lastmod = 2021-07-29T18:05:58+08:00
tags = ["golang"]
categories = ["code"]
draft = false
+++

> golang 接口值: 一个包含 nil 指针的接口不是 nil 接口.
>
> 一个不包含任何值的 nil 接口值和一个刚好包含 nil 指针的接口值是不同的.

接口值由两个部分组成:

-   一个具体的类型
-   那个类型的值

它们被称为接口的动态类型和动态值.

| type | value |
|------|-------|
| x    | x     |

接口的零值是指动态类型为nil, 动态值也为nil.

| type | value |
|------|-------|
| nil  | nil   |

这样的接口才能满足 **接口值 == nil**.

```go
var a interface{}
fmt.Println(a == nil) // true
```


## 一个不包含任何值的 nil 接口值: {#一个不包含任何值的-nil-接口值}

interface 类型变量的动态类型和动态值都为 nil.
比如 nil, 或者:

```go
var a interface{}
fmt.Println(a) // a为nil
```


## 一个刚好包含 nil 指针的接口值: {#一个刚好包含-nil-指针的接口值}

赋值给后，interface 类型变量的动态类型不为 nil，动态值为 nil.

```go
var a interface{}
var b *string
a = b
fmt.Println(a) // 此时a的动态类型为*string，动态值为nil
```


## 这两种对象比较: {#这两种对象比较}

```go
func main(){
  var a interface{} // nil
  var b *string     // nil
  a = b

  fmt.Println(a == nil) // false
  fmt.Println(b == nil) // true
  fmt.Println(b == a)   // true
}
```

1.  a `= nil 为 false:
       b 赋值给 a, a 的动态类型为 *string, 动态值为 nil, 所以 a =` nil 为 false.

    ```go
    // 空接口
    type eface struct {
      _type *_type         // 类型信息
      data  unsafe.Pointer // 指向数据的指针
    }

    // 带有方法的接口
    type iface struct {
      tab  *itab           // 存储type信息还有结构实现方法的集合
      data unsafe.Pointer  // 指向数据的指针
    }
    ```

根据 interface 的底层实现,  a = b 实则是 a.data = unsafe.Pointer(&b). a = nil 才是 a.data = nil.

1.  b == nil 为 true: b 是一个空的指针(非接口)类型.

2.  b == a 为 true: 值都为 nil, b 是一个空的指针(非接口)类型.
