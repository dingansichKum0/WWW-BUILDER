+++
title = "常用的typescript类型推导公式"
author = ["dingansichKum0"]
description = "类型推导公式汇总"
date = 2021-05-21
lastmod = 2021-09-30T15:42:14+08:00
tags = ["typescript"]
categories = ["code"]
draft = false
+++

## 函数参数类型 {#函数参数类型}

```typescript
type TFuncParameterType<T> = T extends (arg: infer P) => void ? P : string;

// e.g
function func(arg: string) {}
type TParamsType = TFuncParameterType<typeof func>; // string
```


## 函数返回值类型 {#函数返回值类型}

```typescript
type TFuncReturnType<T> = T extends (arg: any) => infer P ? P : string;

// e.g
function func(arg: string): number {}
type TReturnType = TFuncReturnType<typeof func>; // number
```


## 数组成员作为键约束 {#数组成员作为键约束}

```typescript
const keys = ["a", "b", "c"] as const;
type TKeysMap = Record<typeof keys[number], string> // KeysMap: { a: string; b: string; c: string; }
```


## 数组元素类型 {#数组元素类型}

```typescript
type ArrayElement<T extends readonly unknown[]> = T extends readonly (infer P)[]
  ? P
  : never;

// e.g
const arr = [0, 1];
type TArr = ArrayElement<typeof arr> // number
```


## 属性覆写 {#属性覆写}

```typescript
type Overwrite<T, R> = Omit<T, keyof R> & R;

// e.g
interface IA {
  a: number;
  b: string;
}
type TA = Overwrite<IA, {a: string}> // {a: string; b:string}
```


## 元祖作为键约束 {#元祖作为键约束}

```typescript
type TuplesAsKey<T extends readonly any[], P> = {
  [K in T[number]]: P;
};

// e.g
const keys = ["x", "y"] as const;
type TA = TuplesAsKey<typeof keys, string>; // {x: string, y: string}
```


## 函数类型省略this {#函数类型省略this}

```typescript
type OmitThisParameter<T> = unknown extends ThisParameterType<T>
  ? T
  : T extends (...args: infer A) => infer R
  ? (...args: A) => R
  : number;


// e.g
const func = (this: void, arg: string) => {};
type TFunc = OmitThisParameter<typeof func>; // type TFunc = (arg: string) => void
```


## 继承父类方法调用返回子类 {#继承父类方法调用返回子类}

```typescript
class Foo {
  static instance: any;

  static getInstance<T extends typeof Foo>(this: T): InstanceType<T> {
    if (this.instance) {
      return this.instance as InstanceType<T>;
    }
    return new this() as InstanceType<T>;
  }
}

class Bar extends Foo {}

Bar.getInstance(); // Bar
```
