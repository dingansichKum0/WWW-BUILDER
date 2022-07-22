+++
title = "获取用户home目录路径"
author = ["zakudriver"]
description = "golang 获取用户home目录路径"
date = 2020-06-10
lastmod = 2022-07-22T10:18:15+08:00
tags = ["golang"]
categories = ["code"]
draft = false
+++

```go
func Home() (string, error) {
  user, err := user.Current()
  if nil == err {
    return user.HomeDir, nil
  }

  // cross compile support

  if runtime.GOOS == "windows"  {
    return homeWindows()
  }

  // Unix-like system, so just assume Unix
  return homeUnix()
}

func homeUnix() (string, error) {
  // First prefer the HOME environmental variable
  if home := os.Getenv("HOME"); home != "" {
    return home, nil
  }

  // If that fails, try the shell
  var stdout bytes.Buffer
  cmd := exec.Command("sh", "-c", "eval echo ~$USER")
  cmd.Stdout = &stdout
  if err := cmd.Run(); err != nil {
    return "", err
  }

  result := strings.TrimSpace(stdout.String())
  if result == "" {
    return "", errors.New("blank output when reading home directory")
  }

  return result, nil
}

func homeWindows() (string, error) {
  drive := os.Getenv("HOMEDRIVE")
  path := os.Getenv("HOMEPATH")
  home := drive + path
  if drive == "" || path == "" {
    home = os.Getenv("USERPROFILE")
  }
  if home == "" {
    return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
  }

  return home, nil
}
```
