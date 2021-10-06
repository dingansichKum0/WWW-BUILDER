+++
title = "typescript的重载(Overload)"
author = ["dingansichKum0"]
description = "函数、箭头函数、方法（静态方法，实例方法）、react组件重载"
date = 2021-10-05
lastmod = 2021-10-06T10:35:39+08:00
tags = ["emacs"]
categories = ["code"]
draft = false
+++

<p class="verse">
重载允许函数（方法）接收不同类型或数量的参数时，做出不同的处理。<br />
js本身作为动态脚本语言是支持重载的，typescript的重载更多的是类型系统的补全。<br />
</p>


## 函数 {#函数}

```typescript
function func(arg: number): number;
function func(arg: string): string;
function func(arg: number | string): number | string {
  if (typeof arg === "number") {
    return 0;
  }
  return "0";
}
```


## lambda {#lambda}

lambda的重载需要通过定义 type 实现。

```typescript
type TFunc = {
  (x: number): number;
  (x: number, y: string): string;
};

export const func: TFunc = (x: number, y = ""): any => {
  if (y) {
    return `optional: ${y}`;
  }

  return x;
};
```


## 方法 {#方法}


### 静态方法 {#静态方法}

```typescript
class C {
  static method(arg: string): string;
  static method(): void;
  static method(arg?: string) {
    return arg;
  }
}
```


### 实例方法 {#实例方法}

```typescript
class C {
  public handle(arg: string): string;
  public handle(arg: boolean): void;
  public handle(arg: string | boolean) {
    if (typeof arg === "string") {
      return arg;
    }
    return;
  }
}
```


## React组件 {#react组件}


### 普通函数 {#普通函数}

```typescript
interface IInputSelectProps<T> {
  options: T[];
  getOptionsLabel?: (arg: T) => string;
  getOptionsValue?: (arg: T) => string;
}

export function InputSelect<T>(props: IInputSelectProps<T>): ReactElement;
export function InputSelect({
  options,
  getOptionsLabel = (arg) => arg,
  getOptionsValue = (arg) => arg,
}: IInputSelectProps<string>): ReactElement {
  return <div />;
}

<InputSelect
  options={[{ label: "1", value: "2" }]}
  getOptionsLabel={(arg) => arg.label}
  getOptionsValue={(arg) => arg.value}
/>
```


### lambda {#lambda}

```typescript
interface IInputSelectProps<T> {
  options: T[];
  getOptionsLabel?: (arg: T) => string;
  getOptionsValue?: (arg: T) => string;
}

type TInputSelect = {
  (props: IInputSelectProps<string>): ReactElement;
  <T>(props: IInputSelectProps<T>): ReactElement;
};

export const InputSelect: TInputSelect = ({
  options,
  getOptionsLabel = (arg) => arg,
  getOptionsValue = (arg) => arg,
}): ReactElement => {
  return <div />;
};

<InputSelect
  options={[{ label: "1", value: "2" }]}
  getOptionsLabel={(arg) => arg.label}
  getOptionsValue={(arg) => arg.value}
/>;

```
