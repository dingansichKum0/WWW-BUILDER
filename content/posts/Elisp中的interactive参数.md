+++
title = "Elisp中的interactive参数"
author = ["dingansichKum0"]
description = "interactive 参数的含义"
date = 2021-07-20
lastmod = 2021-07-29T18:05:57+08:00
tags = ["emacs"]
categories = ["code"]
draft = false
+++

## interactive 参数的含义 {#interactive-参数的含义}

若一个函数带有交互模式声明，则它是一个命令函数。即可以通过 M-x(execute-command) 来调用。


### 声明格式 {#声明格式}

```lisp
(defun kumo-interactive-func ()
  (interactive "code-string")
  ;; do something...)
```


### 种类 {#种类}

-   Completion: 提供补全。TAB，SPC 和 RET 完成补全。
-   Existing: 必须是存在的对象名。不接受无效名称。如果输入无效则不会退出 minibuffer。
-   Default: 如果未输入则使用默认值。
-   No I/O: 不读取任何输入。因此不会使用提示符。
-   Prompt: 放在提示符或\n之前。
-   Special: 只能放在交互式字符参数前。


### 含义 {#含义}

-   接收多个输入以 \n 来分隔。

-   \*: 如果当前buffer处于read-only-mode时提示。 [Special]
-   @: 在第一个鼠标事件触发的window调用。[Special]
-   ^: 通过 shift 调用前需要标记区域，没有公共 shift 调用则停止标记。[Special]

-   a: 一个函数定义的符号名。[Existing, Completion, Prompt]

    ```lisp
    (defun with-func-arg (arg)
      (interactive "aEnter a function:")
      (funcall arg))
    ```
-   b: buffer名（已存在）。[Existing, Completion, Default, Prompt]
-   B: buffer名（可以不存在）。[Completion, Default, Prompt]
-   c: 字符。（接收任意输入，不用回车，不能使用输入法）。[Prompt]
-   C: 一个interactive函数的符号名。即满足(commandp xx)为t。[Existing, Completion, Prompt]
-   d: 光标位置，提供 int 类型的参数。[No I/O]
-   D: 目录。[Existing, Completion, Default, Prompt]
-   e: 必须绑定一个非键盘事件。提供event的list形式参数。[Input Events](https://www.gnu.org/software/emacs/manual/html%5Fnode/elisp/Input-Events.html) [No I/O]

    ```lisp
    (defun with-not-keyboard-event-arg (arg)
      (interactive "e")
      (print arg)) ;; ((down-mouse-1 (#<window 130 on a.el> 2845 (664 . 365) 460355375 nil 2845 (94 . 15) nil (664 . 19) (7 . 23))))

    (global-set-key (kbd "<down-mouse-1>") 'with-not-keyboard-event-arg)
    ```
-   f: 文件名（已存在）。[Existing, Completion, Default, Prompt]
-   F: 文件名（可以不存在）。[Completion, Default, Prompt]
-   G: 文件名（可以不存在），如果只输入目录名不含有文件名则使用目录名。[Completion, Default, Prompt]
-   i: 总是提供nil作为参数。[No I/O]
-   k: 按键序列。一直读取知道触发按键映射的指令，或直到没定义的按键序列。提供 string 或 vector 的参数。k 只会读取down-event事件，而忽律之后的up-event(主要是指鼠标点击后松开的event)。可以使用 U code 读取up-event事件。[Prompt]

    ```lisp
    (kbd "C-x") ⇒ "\C-x"
    (kbd "C-x C-f") ⇒ "\C-x\C-f"
    (kbd "C-x 4 C-f") ⇒ "\C-x4\C-f"
    (kbd "X") ⇒ "X"
    (kbd "RET") ⇒ "\^M"
    (kbd "C-c SPC") ⇒ "\C-c "
    (kbd "<f1> SPC") ⇒ [f1 32]
    (kbd "C-M-<down>") ⇒ [C-M-down]
    ```
-   K: 和 k 类似。改变已定义的按键序列。
-   m: mark位置。提供 int 类型的参数。[No I/O]
-   M: 任意文本。使用当前 buffer 的输入法在 minibuffer 中读取，并作为 string 返回。[Prompt]
-   n: int 类型参数。如果输入的不是 int 会提示再次输入。n 之前几乎不使用前缀参数。[Prompt]
-   N: 读取数字前缀参数。如果没有前缀参数则读取一个 int 作为参数。[Prefix Command Arguments](https://www.gnu.org/software/emacs/manual/html%5Fnode/elisp/Prefix-Command-Arguments.html) [Prompt]
-   p: 数字前缀参数。也可以不用p参数，直接在代码中判断 current-prefix-arg 的值。[No I/O]

    ```lisp
    (defun with-num-arg (arg)
      (interactive "p")
      (print arg)) ;; 4 (numberp arg) t


    (defun with-expression-arg (arg)
      (interactive
       (list (prefix-numeric-value current-prefix-arg)))
      (print arg)) ;; 4 (numberp arg) t
    ;; Same as 'with-num-arg
    ```
-   P: 原始前缀参数。[No I/O]

    ```lisp
    (defun with-num-arg (arg)
      (interactive "P")
      (print arg)) ;; (4) (lisp arg) t
    ```
-   r: region 的开始/结束位置。提供两个参数(beg end)，唯一提供两个参数的code。如果调用该命令时没有触发 region 则会报错[No I/O]
-   s: 任意文本。读入 minibuffer 并作为 string 返回。使用 C-j 或 RET 终止输入。[Prompt]
-   S: 在 minibuffer 中读取输入的 interned symbol 名。使用 C-j 或 RET 终止输入。[Prompt]
-   U: 一个按键序列或者 nil。可以在 k 或 K 之后使用，以获取在 k 或 K 读取down-event后被忽略的up-event(如果有。主要是指鼠标点击后松开的event)。如果没有up-event被忽略则提供 nil 参数。[No I/O]
-   v: 一个用户声明的变量。即满足(custom-variable-p)为t。[Existing, Completion, Prompt]
-   x: 一个 list。不会被 evaluated。使用C-j 或 RET 终止输入。[Prompt]

    ```lisp
    (defun with-list-arg (arg)
      (interactive "x")
      (print arg))
    ;; M-x with-list-arg is invoked; input (+ 1 2), print "(+ 1 2)"
    ```
-   x: 一个 list。会被 evaluated。使用C-j 或 RET 终止输入。[Prompt]

    ```lisp
    (defun with-list-arg (arg)
      (interactive "X")
      (print arg))
    ;; M-x with-list-arg is invoked; input (+ 1 2), print "3"
    ```
-   z: 一个 code system 名(symbol)。如果输入为无效 code system，则参数为 nil。 [Completion, Existing, Prompt]

    ```lisp
    (defun save-buffer-as-utf8 (coding-system)
      "Revert a buffer with `CODING-SYSTEM' and save as UTF-8."
      (interactive "zCoding system for visited file (default nil):")
      (revert-buffer-with-coding-system coding-system)
      (set-buffer-file-coding-system 'utf-8))
    ```
-   Z: 一个 code system 名(symbol)。仅当该命令有前缀参数时，否则参数为 nil。 [Completion, Existing, Prompt]
