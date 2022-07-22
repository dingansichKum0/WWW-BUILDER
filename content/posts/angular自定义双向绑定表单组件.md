+++
title = "angular自定义双向绑定表单组件"
author = ["zakudriver"]
description = "angular使用ControlValueAccessor抽象类实现响应式表单组件"
date = 2019-07-08
lastmod = 2022-07-22T10:13:04+08:00
tags = ["angular"]
categories = ["code"]
draft = false
+++

angular的ControlValueAccessor是一个连接表单模型和视图DOM的抽象类接口

<style>.org-center { margin-left: auto; margin-right: auto; text-align: center; }</style>

<div class="org-center">

使自定义表单组件像原生input一样映射到form表单模型中, 拥有自定义表单组件的form也能使用响应式表单.
(也就是使自定义表单组件拥有formControlName属性和ngModel接口.)

</div>

毕竟响应式表单才是angular的利器.


## ControlValueAccessor {#controlvalueaccessor}

```typescript
export interface ControlValueAccessor {
  writeValue(obj: any): void;
  registerOnChange(fn: any): void;
  registerOnTouched(fn: any): void;
  setDisabledState?(isDisabled: boolean): void;
}
```


### writeValue(obj: any): {#writevalue--obj-any}

该方法是接收模版中的ngModel.

```typescript
writeValue(value: any): void {
 this._renderer.setProperty(this._elementRef.nativeElement, 'value', value);
}
```


### registerOnChange(fn: any): void: {#registeronchange--fn-any--void}

该方法是组件接收到 change 事件的回调, 可以用来通知外部达成双向绑定, 即ngModelChange.

```typescript
registerOnChange(fn: (_: any) => void): void {
 this._onChange = fn;
}
```


### registerOnTouched(fn: any): {#registerontouched--fn-any}

接收到 touched 事件的回调.

```typescript
registerOnTouched(fn: any): void {
  this._onTouched = fn;
}
```


### setDisabledState?(isDisabled: boolean): {#setdisabledstate--isdisabled-boolean}

该方法是组件输入状态 disable &lt;=&gt; enable 变化时的回调。该方法会根据参数值，启用或禁用指定的DOM元素.


## 以下组件类实现了ControlValueAccessor接口. {#以下组件类实现了controlvalueaccessor接口-dot}


### CheckboxControlValueAccessor {#checkboxcontrolvalueaccessor}

用于checkbox复选组件
选择器:

-   input[type=checkbox][formControlName]
-   input[type=checkbox][formControl]
-   input[type=checkbox][ngModel]


### NumberValueAccessor {#numbervalueaccessor}

用于number类型的输入组件
选择器:

-   input[type=number][formControlName]
-   input[type=number][formControl]
-   input[type=number][ngModel]


### DefaultValueAccessor {#defaultvalueaccessor}

用于 text 和 textarea 类型的输入组件
选择器:

-   input:not([type=checkbox])[formControlName]
-   textarea[formControlName]
-   input:not([type=checkbox])[formControl]
-   textarea[formControl]
-   input:not([type=checkbox])[ngModel]
-   textarea[ngModel]
-   [ngDefaultControl]


### RadioControlValueAccessor {#radiocontrolvalueaccessor}

用于radio单选组件
选择器:

-   input[type=radio][formControlName]
-   input[type=radio][formControl]
-   input[type=radio][ngModel]

扩展方法:

-   fireUncheck(value: any): void:取消选中的回调.


### RangeValueAccessor {#rangevalueaccessor}

用于范围输入组件
选择器:

-   input[type=range][formControlName]
-   input[type=range][formControl]
-   input[type=range][ngModel]


### SelectControlValueAccessor {#selectcontrolvalueaccessor}

用于select组件
选择器:

-   select:not([multiple])[formControlName]
-   select:not([multiple])[formControl]
-   select:not([multiple])[ngModel]

扩展方法:

-   compareWith: (o1: any, o2: any) =&gt; boolean:比较函数. 例如option的ngValue是一个对象, 当选中项填入表单时,需要编写一个比较函数来处理当前选中的对象是哪一个option.


### SelectMultipleControlValueAccessor {#selectmultiplecontrolvalueaccessor}

用于多选select组件
选择器:

-   select[multiple][formControlName]
-   select[multiple][formControl]
-   select[multiple][ngModel]

扩展方法:
compareWith: (o1: any, o2: any) =&gt; boolean:比较函数. 例如option的ngValue是一个对象, 当选中项填入表单时, 需要编写一个比较函数来处理当前选中的对象是哪一个option.


## EG {#eg}


### 自定义表单组件代码结构 {#自定义表单组件代码结构}

```typescript
import {
  Component,
  OnInit,
  HostListener,
  ViewEncapsulation,
  forwardRef,
  Input,
  OnDestroy,
  ChangeDetectorRef,
  ChangeDetectionStrategy
} from '@angular/core';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from '@angular/forms';
import { Subject } from 'rxjs';

@Component({
  selector       : '[app-radiobox]',
  templateUrl    : './radiobox.component.html',
  styleUrls      : ['./radiobox.component.styl'],
  encapsulation  : ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush,
  host           : {
    '[class.radio-wrapper]'        : 'true',
    '[class.radio-wrapper-checked]': 'checked'
  },
  providers: [
    {
      provide    : NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => RadioboxComponent),
      multi      : true
    }
  ]
})
export class RadioboxComponent implements OnInit, OnDestroy, ControlValueAccessor {
  @Input()
  value: boolean;
  checked: boolean;
  select$ = new Subject<RadioboxComponent>();

  onChange: (_: any) => void = () => null;
  onTouched: () => void = () => null;

  constructor(private _cdr: ChangeDetectorRef) {}

  @HostListener('click', ['$event'])
  onClick(e: MouseEvent): void {
    e.stopPropagation();
    e.preventDefault();
    this.checked = !this.checked;
    this.onChange(this.checked);
    this.select$.next(this);
  }

  writeValue(value: boolean) {
    this.checked = value;
  }
  registerOnChange(fn: (_: boolean) => {}): void {
    this.onChange = fn;
  }
  registerOnTouched(fn: () => {}): void {
    this.onTouched = fn;
  }

  ngOnInit() {}

  ngOnDestroy() {}
}
```


### 调用 {#调用}

```typescript
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';

@Component({
  selector: 'app-example',
  template: `
    <form [formGroup]="testForm">
      <label>试试</label>
      <label app-radiobox formControlName="check">check me</label>
    </form>
  `
})
export class ExampleComponent implements OnInit {
  testForm: FormGroup = this._fb.group({
    check: false
  });
  constructor(private _fb: FormBuilder) {}

  ngOnInit() {
    this.testForm.valueChanges.subscribe(d => {
      console.log(d);
    });
  }
}
```
