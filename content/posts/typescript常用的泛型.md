+++
title = "typescript常用的泛型"
author = ["zakudriver"]
description = "常用的泛型汇总"
date = 2021-05-21
lastmod = 2023-02-27T17:01:20+08:00
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


## 元祖成员作为键约束 {#元祖成员作为键约束}

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


## 值类型覆写 {#值类型覆写}

```typescript
type Overwrite<T, R> = Omit<T, keyof R> & R;

// e.g
interface IA {
  a: number;
  b: string;
}
type TA = Overwrite<IA, {a: string}> // {a: string; b:string}
```


## 键覆写 {#键覆写}

```typescript
type OverwriteKey<T, K extends keyof T, P extends keyof any> = Omit<T, K> & { [S in P]: T[K] };

// e.g
interface IA {
  a: number;
  b: string;
}
type TA = OverwriteKey<IA, "a", "c"> // {c: number; b:string}
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


## 元祖类型 {#元祖类型}

```typescript
type _TupleOf<T, N extends number, R extends unknown[]> = R["length"] extends N ? R : _TupleOf<T, N, [T, ...R]>;
type Tuple<T, N extends number> = N extends N ? (number extends N ? T[] : _TupleOf<T, N, []>) : never;

// e.g
type TTuple = Tuple<string, 4>; // [string, string, string. string]
```


## 函数第二个参数类型有第一个参数决定(第一个参数是Key，第二个参数是Value) {#函数第二个参数类型有第一个参数决定--第一个参数是key-第二个参数是value}

```typescript
export class Foo {
  private _data = {
    a: 1,
    b: "2",
  };

  public set<K extends keyof Foo["_data"]>(key: K, value: Foo["_data"][K]): void {
    this._data[key] = value;
  }
}
```


## 元祖转联合类型 {#元祖转联合类型}

```typescript
const arr = <const>["foo", "bar", "baz"];

type Ts = typeof arr[number]; // "foo" | "bar" | "baz"
```


## 联合类型转元祖 {#联合类型转元祖}

```typescript
type UnionToIntersection<U> = (U extends any ? (k: U) => void : never) extends (k: infer I) => void ? I : never;
type LastOf<T> = UnionToIntersection<T extends any ? () => T : never> extends () => infer R ? R : never;

type Push<T extends any[], V> = [...T, V];

type TuplifyUnion<T, L = LastOf<T>, N = [T] extends [never] ? true : false> = true extends N
  ? []
  : Push<TuplifyUnion<Exclude<T, L>>, L>;

type Tabc = "a" | "b" | "c";
type TTuple = TuplifyUnion<abc>; // ["a", "b", "c"]
```


## 键值对的键构建联合类型 {#键值对的键构建联合类型}

```typescript
const kv = {
  foo: 1,
  bar: 2,
  baz: 3,
} as const;

type Ts = keyof typeof kv; // "foo" | "bar" | "baz;"
```


## 键值对的值构建联合类型 {#键值对的值构建联合类型}

```typescript
const kv = {
  foo: 1,
  bar: 2,
  baz: 3,
} as const;

type Ts = typeof kv[keyof typeof kv]; // 1 | 2 | 3
```


## 数组项的某一字段值构建联合类型 {#数组项的某一字段值构建联合类型}

```typescript
const kvs = [
  { name: "foo", other: "something" },
  { name: "bar", other: "something" },
  { name: "baz", other: "something" },
] as const;

type Ts = typeof kvs[number]["name"]; // "foo" | "bar" | "baz"
```


## 合并联合键值对类型 {#合并联合键值对类型}

```typescript
type FnReturnType<T> = T extends (...args: any) => infer R ? (R extends void ? {} : R) : {};

type AllKeys<T> = T extends any ? keyof T : never;

type CommonKeys<T extends object> = keyof T;

type Subtract<A, C> = A extends C ? never : A;

type NonCommonKeys<T extends object> = Subtract<AllKeys<T>, CommonKeys<T>>;

type PickType<T, K extends AllKeys<T>> = T extends { [k in K]?: any } ? T[K] : undefined;

type PickTypeOf<T, K extends string | number | symbol> = K extends AllKeys<T> ? PickType<T, K> : never;

type Merge<T extends object> = {
  [k in CommonKeys<T>]: PickTypeOf<T, k>;
} & {
  [k in NonCommonKeys<T>]?: PickTypeOf<T, k>;
};

type PartialMerge<T extends object> = {
  [k in CommonKeys<T>]?: PickTypeOf<T, k>;
} & {
  [k in NonCommonKeys<T>]?: PickTypeOf<T, k>;
};

type A = { a: string } | { a: string; b: string } | { a: string; c: boolean };

type B = PartialMerge<A>; // { a?: string; b?:string; c?: boolean; }
type C = Merge<A>; // { a: string; b?:string; c?: boolean}
```
