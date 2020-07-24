+++
title = "vimrc"
author = ["rx-78-kum0"]
description = "vim配置"
date = 2019-06-04
lastmod = 2020-07-02T14:01:16+08:00
tags = ["vim"]
categories = ["configuration"]
draft = false
+++

```yaml
" .vimrc

" Configuration file for vim
set modelines=0
" Don't write backup file if vim is being called by "crontab -e"
" au BufWrite /private/tmp/crontab.* set nowritebackup nobackup
" Don't write backup file if vim is being called by "chpass"
" au BufWrite /private/etc/pw.* set nowritebackup nobackup
set nobackup
set nowritebackup

let skip_defaults_vim=1
" <leader>
let g:mapleader=","

" 去掉有关vi一致性模式，避免以前版本的bug和局限
set nocompatible
" 设置退格键可用
set backspace=2
" utf-8
set encoding=UTF-8
" 主题
" set background=dark
" colorscheme dracula
" 函数方法名加粗
let g:enable_bold_font = 1
" 注释斜体
let g:enable_italic_font = 1
" 透明背景
let g:hybrid_transparent_background = 1
" airline_theme
let g:airline_theme = "hybrid"
" 行号
set nu!
" 高亮显示寻找匹配
set hls
" 允许用指定语法高亮配色方案替换默认方案
syntax on dracula
" 开启语法高亮
syntax enable

" 设置匹配模式 (当输入一个左括号时会匹配相应的右括号)
set showmatch
" 显示当前光标位置
set ruler

" 使用系统剪切板
set clipboard=unnamed
" 设置格式化时代码缩进为2个空格
set shiftwidth=2
" tab键缩进为4格子
set tabstop=2
"  把连续数量的空格视为一个制表符
set softtabstop=2
" 禁止折行
set nowrap
" tab键转换为空格
set expandtab
" 智能缩进
set smartindent
" 开启实时搜索功能
set incsearch
" 搜索时大小写不敏感
set ignorecase
" vim 自身命令行模式智能补全
set wildmenu
" 开启文件类型侦测
filetype on
" 根据侦测到的不同类型加载对应的插件
filetype plugin on
" 自适应不同语言的智能缩进
filetype indent on
" 将制表符扩展为空格
set nofoldenable
" 基于缩进或语法进行代码折叠
set foldmethod=syntax
" 启动 vim 时关闭折叠代码
set nofoldenable
" 禁止光标闪烁
set gcr=a:block-blinkon0
" 禁止显示滚动条
set guioptions-=l
set guioptions-=L
set guioptions-=r
set guioptions-=R
" 高亮显示当前行/列
set cursorline
set cursorcolumn

" 让配置变更立即生效
autocmd BufWritePost $MYVIMRC source $MYVIMRC

" normal 模式 keymap
nnoremap x "_x
nnoremap X "_X
nnoremap d "_d
nnoremap dd "_dd
nnoremap D "_D
nnoremap s "_s
nnoremap S "_S
nnoremap c "_c
nnoremap C "_C

" insert 模式
inoremap $( ()<esc>i
inoremap $< <><esc>i
inoremap ${ {}<esc>i
inoremap $[ []<esc>i
inoremap $' ''<esc>i
inoremap $" ""<esc>i

" visual 模式
vnoremap s "_s
vnoremap S "_S
vnoremap z) c()<esc>hp
vnoremap z} c{}<esc>hp
vnoremap z] c[]<esc>hp
vnoremap z' c''<esc>hp
vnoremap z" c""<esc>hp
vnoremap z> c<><esc>hp
vnoremap z` c``<esc>hp

" NERDTree
let g:NERDTreeIndicatorMapCustom = {
    \ "Modified"  : "✹",
    \ "Staged"    : "✚",
    \ "Untracked" : "✭",
    \ "Renamed"   : "➜",
    \ "Unmerged"  : "═",
    \ "Deleted"   : "✖",
    \ "Dirty"     : "✗",
    \ "Clean"     : "✔︎",
    \ "Unknown"   : "?"
    \ }
" 自动开启NERDTree
" autocmd vimenter * NERDTree

" UltiSnips
let g:UltiSnipsExpandTrigger="<Leader><TAB>"
let g:UltiSnipsJumpForwardTrigger="<c-f>"
let g:UltiSnipsJumpBackwardTrigger="<c-b>"


"au BufRead,BufNewFile *.go set filetype=go

" ycm
" let g:ycm_key_list_select_completion=['<c-n>']
" let g:ycm_key_list_previous_completion=['<c-p>']
" let g:ycm_key_invoke_completion = '<C-Space>'
" " 关闭加载.ycm_extra_conf.py提示
" let g:ycm_confirm_extra_conf=0
" " 开启 YCM 基于标签引擎
" let g:ycm_collect_identifiers_from_tags_files=1
" " 从第2个键入字符就开始罗列匹配项
" let g:ycm_min_num_of_chars_for_completion=1
" " 禁止缓存匹配项,每次都重新生成匹配项
" let g:ycm_cache_omnifunc=0
" " 语法关键字补全
" let g:ycm_seed_identifiers_with_syntax=1
" " 设置在下面几种格式的文件上屏蔽ycm
" let g:ycm_filetype_blacklist = {
"       \ 'typescript.tsx' : 1,
"       \ 'typescript' : 1,
"       \}
" " 注释和字符串中的文字也会被收入补全
" let g:ycm_collect_identifiers_from_comments_and_strings = 0
" " 输入第2个字符开始补全
" let g:ycm_min_num_of_chars_for_completion=2


" 注释
" Add spaces after comment delimiters by default
let g:NERDSpaceDelims = 1
" Use compact syntax for prettified multi-line comments
let g:NERDCompactSexyComs = 1
" Align line-wise comment delimiters flush left instead of following code indentation
let g:NERDDefaultAlign = 'left'
" Set a language to use its alternate delimiters by default
let g:NERDAltDelims_java = 1
" Add your own custom formats or override the defaults
let g:NERDCustomDelimiters = { 'c': { 'left': '/**','right': '*/' } }
" Allow commenting and inverting empty lines (useful when commenting a region)
let g:NERDCommentEmptyLines = 1
" Enable trimming of trailing whitespace when uncommenting
let g:NERDTrimTrailingWhitespace = 1
" Enable NERDCommenterToggle to check all selected lines is commented or not
let g:NERDToggleCheckAllLines = 1


" coc
let g:coc_global_extensions = [
  \ 'coc-json',
  \ 'coc-html',
  \ 'coc-css',
  \ ]

set cmdheight=2
set updatetime=300
set shortmess+=c
set signcolumn=yes

au BufNewFile,BufRead *.ts setlocal filetype=typescript
au BufNewFile,BufRead *.tsx setlocal filetype=typescript.tsx

nmap <silent> gd <Plug>(coc-definition)
nmap <silent> gy <Plug>(coc-type-definition)
nmap <silent> gi <Plug>(coc-implementation)
nmap <silent> gr <Plug>(coc-references)
nmap <leader>rn <Plug>(coc-rename)
nmap <silent> [c <Plug>(coc-diagnostic-prev)
nmap <silent> ]c <Plug>(coc-diagnostic-next)
nmap <leader>a  <Plug>(coc-codeaction-selected)
nmap <leader>ac  <Plug>(coc-codeaction)
nmap <leader>qf  <Plug>(coc-fix-current)
nmap <silent> <TAB> <Plug>(coc-range-select)

xmap <silent> <TAB> <Plug>(coc-range-select)
xmap <silent> <S-TAB> <Plug>(coc-range-select-backword)
xmap <leader>a  <Plug>(coc-codeaction-selected)

" 显示文档
nnoremap <silent> K :call <SID>show_documentation()<CR>

" 回车：补全框?确认补全:回车可以撤回
inoremap <expr> <cr> pumvisible() ? "\<C-y>" : "\<C-g>u\<CR>"
" <c-c>触发补全
inoremap <silent><expr> <c-c> coc#refresh()
inoremap <expr><S-TAB> pumvisible() ? "\<C-p>" : "\<C-h>"
inoremap <silent><expr> <TAB>
      \ pumvisible() ? "\<C-n>" :
      \ <SID>check_back_space() ? "\<TAB>" :
      \ coc#refresh()

" Use `:Format` to format current buffer
command! -nargs=0 Format :call CocAction('format')
" Use `:Fold` to fold current buffer
command! -nargs=? Fold :call     CocAction('fold', <f-args>)

" set statusline^=%{coc#status()}%{get(b:,'coc_current_function','')}

" autocmd CursorHold * silent call CocActionAsync('highlight')

function! s:show_documentation()
  if (index(['vim','help'], &filetype) >= 0)
    execute 'h '.expand('<cword>')
  else
    call CocAction('doHover')
  endif
endfunction

function! s:check_back_space() abort
  let col = col('.') - 1
  return !col || getline('.')[col - 1]  =~# '\s'
endfunction



" -------------------- plug ------------------------
call plug#begin('~/.vim/plugins')
Plug 'dracula/vim', { 'as': 'dracula' }
Plug 'mhinz/vim-startify'
Plug 'SirVer/ultisnips'
Plug 'jiangmiao/auto-pairs'
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'scrooloose/nerdcommenter'

" NERDTree
Plug 'scrooloose/nerdtree'
" git
Plug 'Xuyuanp/nerdtree-git-plugin'
" icon
" Plug 'ryanoasis/vim-devicons'

" coc
" Plug 'neoclide/coc.nvim', {'do': 'yarn install --frozen-lockfile' }

" golang
" Plug 'fatih/vim-go'
" typescript
" Plug 'leafgarland/typescript-vim', {'for': ['typescript', 'typescript.tsx', 'js']}

" prettier
"Plug 'prettier/vim-prettier', {
"  \ 'do': 'yarn install',
"  \ 'for': ['javascript', 'typescript', 'css', 'less', 'scss', 'json', 'graphql', 'markdown', 'vue', 'yaml', 'html'] }
"Plug 'Valloric/YouCompleteMe'
call plug#end()


" 替换函数
" 参数说明：
" confirm：是否替换前逐一确认
" wholeword：是否整词匹配
" replace：被替换字符串
function! Replace(confirm, wholeword, replace)
    wa
    let flag = ''
    if a:confirm
        let flag .= 'gec'
    else
        let flag .= 'ge'
    endif
    let search = ''
    if a:wholeword
        let search .= '\<' . escape(expand('<cword>'), '/\.*$^~[') . '\>'
    else
        let search .= expand('<cword>')
    endif
    let replace = escape(a:replace, '/\&~')
    execute 'argdo %s/' . search . '/' . replace . '/' . flag . '| update'
endfunction
" 不确认、非整词
nnoremap <Leader>R :call Replace(0, 0, input('Replace '.expand('<cword>').' with: '))<CR>
" 不确认、整词
nnoremap <Leader>rw :call Replace(0, 1, input('Replace '.expand('<cword>').' with: '))<CR>
" 确认、非整词
nnoremap <Leader>rc :call Replace(1, 0, input('Replace '.expand('<cword>').' with: '))<CR>
" 确认、整词
nnoremap <Leader>rcw :call Replace(1, 1, input('Replace '.expand('<cword>').' with: '))<CR>
nnoremap <Leader>rwc :call Replace(1, 1, input('Replace '.expand('<cword>').' with: '))<CR>
```
