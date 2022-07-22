+++
title = "nyan"
author = ["zakudriver"]
description = "(=｀ω´=)~"
date = 2020-03-14
lastmod = 2022-07-22T10:16:21+08:00
tags = ["shell"]
categories = ["code"]
draft = false
+++

```bash
#!/usr/bin/env bash

# Print nyan cat
# https://github.com/steckel/Git-Nyan-Graph/blob/master/nyan.sh
# If you want big animated version: `telnet miku.acm.uiuc.edu`

e='\033'
RESET="$e[0m"
BOLD="$e[1m"
CYAN="$e[0;96m"
RED="$e[0;91m"
YELLOW="$e[0;93m"
GREEN="$e[0;92m"

echo
if [ $[$RANDOM%2] -eq "0" ]; then
    echo -en $RED'`·.,¸,.·*·.'
    echo -e $RESET$BOLD'╭━━━━╮'$RESET
    echo -en $YELLOW'`·.,¸,.·*·.'
    echo -e $RESET$BOLD'|::: /\_/\\'$RESET
    echo -en $GREEN'`·.,¸,.·*·.'
    echo -e $RESET$BOLD'|:::( ◕ᴥ◕)'$RESET
    echo -en $CYAN'`·.,¸,.·*·.'
    echo -e $RESET$BOLD'u-u━━-u--u'$RESET
else
    echo -en $RED'-_-_-_-_-_-_-_'
    echo -e $RESET$BOLD',------,'$RESET
    echo -en $YELLOW'_-_-_-_-_-_-_-'
    echo -e $RESET$BOLD'|   /\_/\\'$RESET
    echo -en $GREEN'-_-_-_-_-_-_-'
    echo -e $RESET$BOLD'~|__( ^ .^)'$RESET
    echo -en $CYAN'-_-_-_-_-_-_-'
    echo -e $RESET$BOLD'""  ""'$RESET
fi
echo
```

{{< figure src="/ox-hugo/2020-06-29_23-28-37_nyan-img.png" >}}
