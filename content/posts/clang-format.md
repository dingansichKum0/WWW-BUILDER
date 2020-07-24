+++
title = "clang-format"
author = ["rx-78-kum0"]
description = "clang-format格式化工具配置"
date = 2020-02-25
lastmod = 2020-07-02T14:01:53+08:00
tags = ["tools"]
categories = ["configuration"]
draft = false
+++

```yaml
# .clang-format

BasedOnStyle: LLVM

Language:	Cpp

IndentWidth : 2

# 开括号(开圆括号、开尖括号、开方括号)后的对齐: Align, DontAlign, AlwaysBreak(总是在开括号后换行)
AlignAfterOpenBracket:	Align

# 大括号换行，只有当BreakBeforeBraces设置为Custom时才有效
BraceWrapping:
  # class定义后面
  AfterClass:	false
  # 控制语句后面
  AfterControlStatement:	false
  # enum定义后面
  AfterEnum:	false
  # 函数定义后面
  AfterFunction:	true
  # 命名空间定义后面
  AfterNamespace:	false
  # ObjC定义后面
  AfterObjCDeclaration:	false
  # struct定义后面
  AfterStruct:	false
  # union定义后面
  AfterUnion:	false
  # catch之前
  BeforeCatch:	true
  # else之前
  BeforeElse:	true
  # 缩进大括号
  IndentBraces:	false

BreakBeforeBraces:	Custom
```
