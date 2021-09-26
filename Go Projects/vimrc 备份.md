



``` shell
git clone https://github.com/fatih/vim-go.git ~/.vim/pack/plugins/start/vim-go
git clone https://github.com/fatih/vim-go.git ~/.vim/pack/plugins/start/vim-go
g:seoul256_background




-- VISUAL --
 1 .vimrc                                                                                         X
  24 set vb t_vb=                   " 当vim进行编辑时，如果命令错误，会发出警报，该设置去掉警报
  25 set ruler                      " 在编辑过程中，在右下角显示光标位置的状态行
  26 set incsearch
  27 set mouse=a                    " 鼠标控制光标
  28
  29 source ~/.vim_runtime/vimrcs/basic.vim
  30 source ~/.vim_runtime/vimrcs/filetypes.vim
  31 source ~/.vim_runtime/vimrcs/plugins_config.vim
  32 source ~/.vim_runtime/vimrcs/extended.vim
  33
  34 try
  35   source ~/.vim_runtime/my_configs.vim
  36 catch
  37   endtry
  38
  39 call plug#begin('~/.vim/plugged')   " 设置完:PlugInstall 参考 https://github.com/junegunn/vim-      plug
  40
  41 call plug#end()
  42
  43 let g:seoul256_background = 236     " theme
  44 silent! colo seoul256
  45
  
  
  
  autocmd
  
  
   1 .vimrc                                                                                         X
   1 set runtimepath+=~/.vim_runtime
   2
   3 source ~/.vim_runtime/vimrcs/basic.vim
   4 source ~/.vim_runtime/vimrcs/filetypes.vim
   5 source ~/.vim_runtime/vimrcs/plugins_config.vim
   6 source ~/.vim_runtime/vimrcs/extended.vim
   7
   8 try
   9 source ~/.vim_runtime/my_configs.vim
  10 catch
  11 endtry
  12
  13 " UTF-8 编码
  14 set encoding=UTF-8
  15 " 搜索忽略大小写
  16 set hlsearch
  17 set incsearch
  18 set ignorecase
  19 " 语法高亮
  20 syntax enable
  21 " 允许backspace和光标键跨越行边界
  22 set whichwrap+=<,>,h,l
  23 " 在被分割的窗口间显示空白，便于阅读
  24 set fillchars=vert:\ ,stl:\ ,stlnc:\
 NORMAL  .vimrc                               vim  utf-8[unix]  3% :11/326☰ :6  ☲ [32]tra… 

  
  
  
  
  
```

