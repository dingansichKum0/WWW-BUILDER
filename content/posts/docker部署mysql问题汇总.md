+++
title = "docker部署mysql问题汇总"
author = ["zakudriver"]
description = "docker部署mysql问题汇总"
date = 2020-05-09
lastmod = 2022-07-22T10:13:29+08:00
tags = ["mysql", "docker"]
categories = ["code"]
draft = false
+++

## docker-compose 配置 {#docker-compose-配置}

```yaml
version: '3'

services:
  # mysql
  blog_mysql:
    image: mysql
    restart: always
    ports:
      - '3306:3306'
    environment:
      - MYSQL_ROOT_PASSWORD=xxxx
    volumes:
      - '/data/mysql:/var/lib/mysql'
```


## docker 安装完mysql 后客户端无法访问 {#docker-安装完mysql-后客户端无法访问}


### 进入镜像中的mysql {#进入镜像中的mysql}

```shell
docker exec -it container_id /bin/bash
```


### 登录mysql {#登录mysql}

```shell
mysql -u root -p
```


### 修改root账号, 可以通过任何客户端连接 {#修改root账号-可以通过任何客户端连接}

```shell
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'xxxxx';
```
