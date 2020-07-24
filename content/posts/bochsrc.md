+++
title = "bochsrc"
author = ["rx-78-kum0"]
description = "nasm编译配置"
date = 2020-05-26
lastmod = 2020-07-02T14:01:42+08:00
tags = ["nasm"]
categories = ["configuration"]
draft = false
+++

```yaml
###############################################################
# Configuration file for Bochs
###############################################################

# how much memory the emulated machine will have
megs: 32

# filename of ROM images
romimage: file=/usr/local/share/bochs/BIOS-bochs-latest
vgaromimage: file=/usr/local/share/bochs/VGABIOS-lgpl-latest

# what disk images will be used
floppya: 1_44=a.img, status=inserted

# choose the boot disk.
boot: floppy

# where do we send log messages?
# log: bochsout.txt

# disable the mouse
mouse: enabled=0

# enable key mapping, using US layout as default.
# keyboard_mapping: enabled=1, map=/usr/local/share/bochs/keymaps/x11-pc-us.map
keyboard:  keymap=/usr/local/share/bochs/keymaps/x11-pc-us.map
```
