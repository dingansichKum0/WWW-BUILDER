+++
title = "vscode"
author = ["zakudriver"]
description = "vscode配置"
date = 2019-03-02
lastmod = 2021-07-29T18:05:59+08:00
tags = ["vscode"]
categories = ["configuration"]
draft = false
+++

```json
{
  // editor
  "editor.fontSize": 14,
  "editor.snippetSuggestions": "top",
  "editor.formatOnSave": true,
  "explorer.confirmDragAndDrop": true,
  "explorer.confirmDelete": false,
  "editor.detectIndentation": false,
  "files.autoSave": "off",
  // "editor.fontFamily": "'FuraCode', Menlo, Monaco, 'Courier New', monospace",
  "editor.fontLigatures": true,
  "workbench.iconTheme": "material-icon-theme",
  "workbench.colorTheme": "Nebula",

  // prettier
  "prettier.printWidth": 120,
  "editor.tabSize": 2,
  "prettier.singleQuote": false,
  "prettier.semi": true,

  // typescript
  "typescript.updateImportsOnFileMove.enabled": "always",

  // emmet
  "emmet.includeLanguages": {
    "javascript": "javascriptreact",
    "typescript": "typescriptreact"
  },
  "emmet.triggerExpansionOnTab": true,

  // 装饰器
  "javascript.implicitProjectConfig.experimentalDecorators": true,

  // stylus
  "stylusSupremacy.insertColons": false, // 是否插入冒号
  "stylusSupremacy.insertSemicolons": false, // 是否插入分好
  "stylusSupremacy.insertBraces": false, // 是否插入大括号
  "stylusSupremacy.insertNewLineAroundImports": true, // import之后是否换行
  "stylusSupremacy.insertNewLineAroundBlocks": true,
  "stylusSupremacy.insertSpaceAfterComment": true,
  "window.zoomLevel": 0,

  // Formatter
  "[javascript]": {
    "editor.defaultFormatter": "vscode.typescript-language-features"
  },
  "[typescript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[json]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[jsonc]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[typescriptreact]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[html]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },

  "search.followSymlinks": false,

  // vim
  "vim.useSystemClipboard": true,
  "vim.hlsearch": true,
  "vim.smartcase": true,
  "vim.leader": ",",
  "vim.highlightedyank.enable": true,
  "vim.highlightedyank.duration": 1000,
  "vim.highlightedyank.color": "rgba(250, 240, 170, 0.5)",
  "vim.history": 100,
  // "vim.cursorStylePerMode.insert": "line",
  // "vim.cursorStylePerMode.normal": "underline",
  // "vim.cursorStylePerMode.replace": "underline",
  // "vim.cursorStylePerMode.visual": "blink",
  // "vim.cursorStylePerMode.visualblock": "blink",
  // "vim.cursorStylePerMode.visualline": "underline",
  "vim.easymotion": true,
  "vim.easymotionMarkerFontSize": "16",
  "vim.easymotionMarkerHeight": 16,
  "vim.easymotionMarkerWidthPerChar": 9,
  "vim.normalModeKeyBindings": [],
  "vim.normalModeKeyBindingsNonRecursive": [
    {
      "before": ["d"],
      "after": ["\"", "_", "d"]
    },
    {
      "before": ["d", "d"],
      "after": ["\"", "_", "d", "d"]
    },
    {
      "before": ["D"],
      "after": ["\"", "_", "D"]
    },
    {
      "before": ["t"],
      "after": ["\"", "_", "x"]
    },
    {
      "before": ["X"],
      "after": ["\"", "_", "X"]
    },
    {
      "before": ["s"],
      "after": ["\"", "_", "s"]
    },
    {
      "before": ["S"],
      "after": ["\"", "_", "S"]
    },
    {
      "before": ["c"],
      "after": ["\"", "_", "c"]
    },
    {
      "before": ["C"],
      "after": ["\"", "_", "C"]
    }
  ],
  "vim.visualModeKeyBindings": [
    {
      "before": ["s"],
      "after": ["\"", "_", "s"]
    },
    {
      "before": ["S"],
      "after": ["\"", "_", "S"]
    },
    {
      "before": ["z", ")"],
      "after": ["c", "(", ")", "<Esc>", "h", "p"]
    },
    {
      "before": ["z", "}"],
      "after": ["c", "{", "}", "<Esc>", "h", "p"]
    },
    {
      "before": ["z", "]"],
      "after": ["c", "[", "]", "<Esc>", "h", "p"]
    },
    {
      "before": ["z", "'"],
      "after": ["c", "'", "'", "<Esc>", "h", "p"]
    },
    {
      "before": ["z", "\""],
      "after": ["c", "\"", "\"", "<Esc>", "h", "p"]
    },
    {
      "before": ["z", ">"],
      "after": ["c", "<", ">", "<Esc>", "h", "p"]
    },
    {
      "before": ["z", "`"],
      "after": ["c", "`", "`", "<Esc>", "h", "p"]
    }
  ],
  "vim.handleKeys": {
    "<C-a>": false,
    "<C-f>": false
  }
}
```
