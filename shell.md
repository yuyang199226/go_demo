# shell 常用命令

## find 命令
1. 按文件名过滤文件并cp 到某个目录下

`sudo find . -name '*202006031[0-2]*pb' -type f -exec cp {} ./10-12 \;`

## vim 
1. 查找替换
https://harttle.land/2016/08/08/vim-search-in-file.html

```
:{作用范围}s/{目标}/{替换}/{替换标志}
```
例如:%s/foo/bar/g会在全局范围(%)查找foo并替换为bar，所有出现都会被替换（g）。
作用范围
范围氛围当前行，全文，选区
当前行
`:s/foo/bar/g`
全文
`:%s/foo/bar/g`

# 查找当前目录包含某个字符串 所在的文件
`grep -rn 'FINDSTRING'` .




