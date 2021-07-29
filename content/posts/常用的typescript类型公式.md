+++
title = "常用的typescript类型推导公式"
author = ["dingansichKum0"]
description = "类型推导公式汇总"
date = 2021-05-21
lastmod = 2021-07-29T18:06:00+08:00
tags = ["typescript"]
categories = ["code"]
draft = false
+++

## 函数参数类型 {#函数参数类型}

```typescript
type FuncParamsType<T> = T extends (arg: infer P) => void ? P : string;

function func(arg: string) {}

type ParamsType = FuncParamsType<typeof func>; // ParamsType: string
```


### 函数返回值类型 {#函数返回值类型}

```typescript
type FuncReturnType<T> = T extends (arg: any) => infer P ? P : string;

function func(arg: string): number {}

type ReturnType = FuncReturnType<typeof func>; // ReturnType: number
```


## 数组成员作为键约束 {#数组成员作为键约束}

```typescript
const keys = ["a", "b", "c"] as const;

type Keys = typeof keys[number]

type KeysMap = Record<Keys, string> // KeysMap: { a: string; b: string; c: string; }

```


## 可选键约束 {#可选键约束}

```typescript
type Key = "a" | "b";

type OptionalKeyMap={ // optionalKeyMap: { "a" | "b": number }
  [K in Key]?: number
};
```
