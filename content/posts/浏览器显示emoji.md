+++
title = "浏览器显示emoji"
author = ["rx-78-kum0"]
description = "浏览器显示emoji需要做的处理"
date = 2020-03-12
lastmod = 2020-07-02T13:58:54+08:00
tags = ["web"]
categories = ["code"]
draft = false
+++

```js
// js 表情emoji转码
// 发送请求时将uft16转为utf-8
function utf16toEntities(str) {
  var patt = /[\ud800-\udbff][\udc00-\udfff]/g; // 检测utf16字符正则
  return str.replace(patt, function(char) {
    var H, L, code;
    if (char.length === 2) {
      H = char.charCodeAt(0); // 取出高位
      L = char.charCodeAt(1); // 取出低位
      code = (H - 0xd800) * 0x400 + 0x10000 + L - 0xdc00; // 转换算法
      return '&#' + code + ';';
    } else {
      return char;
    }
  });
}

// 收到后端的数据时展示emoji
function entitiesToUtf16(str) {
  return str.replace(/&#(\d+);/g, function(match, dec) {
    let H = Math.floor((dec - 0x10000) / 0x400) + 0xd800;
    let L = (Math.floor(dec - 0x10000) % 0x400) + 0xdc00;
    return String.fromCharCode(H, L);
  });
}
```
