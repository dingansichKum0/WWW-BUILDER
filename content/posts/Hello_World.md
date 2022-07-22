+++
title = "Hello_World"
author = ["zakudriver"]
description = "hello world. ;)"
date = 2020-01-11
lastmod = 2022-07-22T10:12:40+08:00
tags = ["helloworld"]
categories = ["balabala"]
draft = false
+++

## Power by Hugo {#power-by-hugo}


### **现在，该博客改由 [Hugo](https://gohugo.io/) 强力驱动啦！** {#现在-该博客改由-hugo-强力驱动啦}

> Hugo 是一个 golang 编写的静态网页生成器。同类型工具还有使用 node 编写的 [Hexo](https://hexo.io/)。


### Hugo的优点 {#hugo的优点}


#### - 快 {#快}

基于 golang 编写的


#### - 简单 {#简单}

不仅使用简单，甚至自己写主题也很简单。


#### - 支持org格式! (优秀) {#支持org格式--优秀}

hexo, jekyll 等同类工具大多只支持 markdown. 而 hugo 支持 org 这一仅在 emacs 上有的文档格式，感动~

<div class="EXPLAIN">

但不会直接用org来发布~ 目前支持程度远没有 markdown 格式好。因为 org 在 emacs 上原生支持太好，又不像 md 这样通用导致其他平台支持程度不是很高。好在可以使用 emacs 插件把 org 转成 md，虽然还是会丢失一些 org 的特性。

</div>


### Org -&gt; Markdown {#org-markdown}

Emacs 的 [ox-hugo](https://ox-hugo.scripter.co/)


#### 安装 {#安装}

```lisp
(use-package ox-hugo
  :after ox)
```


#### 使用 {#使用}

<!--list-separator-->

-  需要在org文档元数据上标明hugo博客的根目录 HUGO_BASE_DIR 和生成文档的目标目录 HUGO_SECTION

    ```org
    #+HUGO_BASE_DIR: ~/Hugo-DirName
    #+HUGO_SECTION: posts
    ```

<!--list-separator-->

-  **C-c C-e H h**

    导出当前org文档

<!--list-separator-->

-  **org-hugo-auto-export-mode**

    保存org文档自动导出为md

    在hugo博客根目录添加.dir-locals.el文件:

    ```lisp
    (("content-org/"
      . ((org-mode . ((eval . (org-hugo-auto-export-mode)))))))
    ```

    最终Hugo目录树:

    ```nil
    <HUGO_BASE_DIR>
      ├── config.toml
      ├── content
         ├── <HUGO_SECTION>
      ├── content-org      <-- Org files in there
      ├── static
      ├── themes
      └── .dir-locals.el
    ```
