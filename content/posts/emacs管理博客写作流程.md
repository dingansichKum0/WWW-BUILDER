+++
title = "emacs管理博客写作流程"
author = ["dingansichKum0"]
description = "使用emacs作为hugo博客客户端"
date = 2020-07-08
lastmod = 2021-07-29T18:08:33+08:00
tags = ["emacs"]
categories = ["code"]
draft = false
+++

## 使用emacs作为hugo博客的客户端 {#使用emacs作为hugo博客的客户端}

<div class="EXPLAIN">
  <div></div>

Hugo 生成博客的静态页面虽然已经很方便了，ox-hugo 自动md -> org也很便捷了，但写作和发布博客需要经过若干命令行操作才能完成。没有客户端界面直观。

</div>


### 目前写作和发布的流程 {#目前写作和发布的流程}

```nil
新建org文件 -> balabala并自动转md -> 命令行: $ hugo // 生成静态页面 -> magit stage/commit/push -> 完成
```


### hugo 的 major-mode 插件: [easy-hugo](https://github.com/masasam/emacs-easy-hugo) {#hugo-的-major-mode-插件-easy-hugo}

> easy-hugo 是 emacs 上的 hugo 博客管理的插件。支持markdown or org-mode or AsciiDoc or reStructuredText or mmark or html 等格式文档，多博客多站点，以及多平台部署。

{{< figure src="/ox-hugo/2020-07-09_11-02-02_screencast.gif" width="600px" >}}


#### easy-hugo 配置 {#easy-hugo-配置}

```lisp
(use-package easy-hugo
  :commands easy-hugo
  :bind
  (:map easy-hugo-mode-map
        ("SPC" . general-simulate-C-c)
        ("G" . kumo-easy-hugo-github-deploy))
  :custom
  (easy-hugo-org-header t)
  (easy-hugo-basedir kumo/easy-hugo-basedir)
  (easy-hugo-postdir kumo/easy-hugo-postdir)
  (easy-hugo-url kumo/easy-hugo-url)
  (easy-hugo-preview-url kumo/easy-hugo-preview-url)
  (easy-hugo-github-deploy-script kumo/easy-hugo-github-deploy-script)
  (easy-hugo-default-ext ".org")
  :hook
  (easy-hugo-mode . (lambda ()
                      (evil-set-initial-state 'easy-hugo-mode 'emacs)))
  )
```

<div class="EXPLAIN">
  <div></div>

需要配置 easy-hugo 的 basedir/postdir/preview-url/deploy-script 等等。方便 easy-hugo 读取出博客列表、部署脚本，以及执行一键预览等。

</div>


#### 部署脚本 {#部署脚本}

```bash
#!/bin/sh

# If a command fails then the deploy stops
set -e

printf "\033[0;32mDeploying updates to GitHub...\033[0m\n"

# Build the project.
hugo # if using a theme, replace with `hugo -t <YOURTHEME>`

# Go To Public folder
cd public

# Add changes to git.
git add .

# Commit changes.
msg="rebuilding site $(date)"
if [ -n "$*" ]; then
  msg="$*"
fi
git commit -m "$msg"

# Push source and build repos.
git push origin master -f
```

<!--list-separator-->

-  配置的一些坑

    -   由于我使用了evil-mode，所以 easy-hugo-mode-map 的原生键位是不能用了。只有在 easy-hugo-mode 中屏蔽 evil-mode.

    <!--listend-->

    ```lisp
    :hook
    (easy-hugo-mode . (lambda ()
                        (evil-set-initial-state 'easy-hugo-mode 'emacs)))
    ```

    -   屏蔽 evil-mode 同时也屏蔽了 general 的触发键 SPC。 emacs 上没有了 SPC 就像走路蒙着双眼...

    <!--listend-->

    ```lisp
    :bind
    (:map easy-hugo-mode-map
          ("SPC" . general-simulate-C-c))
    ```

    -   easy-hugo 原生的 github page 部署函数不支持交互式shell脚本，导致 git push 时不能输入密码。需要自己写个函数来替换它。

    <div class="EXPLAIN">
      <div></div>

    这里使用 C-q 而不是 q 退出 async-shell-command 的 buffer。因为万一账号或密码中带有 "q" 就不好了...😅

    </div>

    ```lisp
    :bind
    (:map easy-hugo-mode-map
          ("G" . kumo-easy-hugo-github-deploy))


    (defun kumo-easy-hugo-github-deploy ()
      "Easy-Hugo deploy github page."
      (interactive)
      (let* ((output-buffer (get-buffer-create kumo/easy-hugo-github-deploy-buffer-name))
             (command-window (async-shell-command (expand-file-name (concat kumo/easy-hugo-basedir kumo/easy-hugo-github-deploy-script)) output-buffer nil)))
        (select-window command-window)
        (local-set-key (kbd "C-q") 'kill-buffer-and-window)))
    ```


### 现在的写作和发布的流程 {#现在的写作和发布的流程}

```nil
呼出 easy-hugo -> n键 新建org balabala并自动转md -> G键 发布部署
```

<p class="verse">
博客总算能有一个像样的管理界面了😅<br />
现在写作/发布可以直接依托 emacs，并且都不需要键入一个命令行。<br />
</p>
