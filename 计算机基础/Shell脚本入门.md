# Shell脚本入门

## 常用命令

## 运行脚本
运行 Shell 脚本有两种方法：
1、作为可执行程序

将上面的代码保存为 test.sh，并 cd 到相应目录：
```bash

chmod +x ./test.sh  #使脚本具有执行权限
./test.sh  #执行脚本

```
>注意，一定要写成 ./test.sh，而不是 test.sh，运行其它二进制的程序也一样，直接写 test.sh，linux 系统会去 PATH 里寻找有没有叫 test.sh 的，而只有 /bin, /sbin, /usr/bin，/usr/sbin 等在 PATH 里，你的当前目录通常不在 PATH 里，所以写成 test.sh 是会找不到命令的，要用 ./test.sh 告诉系统说，就在当前目录找。

2、作为解释器参数

这种运行方式是，直接运行解释器，其参数就是 shell 脚本的文件名，如：

```bash
/bin/sh test.sh
/bin/php test.php

```

这种方式运行的脚本，不需要在第一行指定解释器信息，写了也没用。


## 判断条件

-z 判断 变量的值，是否为空； zero = 0


 - 变量的值，为空，返回0，为true

 - 变量的值，非空，返回1，为false

 -n 判断变量的值，是否为空 name = 名字

 - 变量的值，为空，返回1，为false

 - 变量的值，非空，返回0，为true

 pid="123"

  [ -z "$pid" ] 单对中括号变量必须要加双引号

  [[ -z $pid ]] 双对括号，变量不用加双引号

 


  [ -n "$pid" ] 单对中括号，变量必须要加双引号

  [[ -z $pid ]] 双对中括号，变量不用加双引号

 2、多个条件判断，[] 和 [[]] 的区别？

 2.1：[[ ]] 双对中括号，是不能使用 -a 或者 -o的参数进行比较的；

 && 并且 || 或 -a 并且 -o 或者

 [[ ]] 条件判断 && 并且 || 或

 


 [[ 5 -lt 3 || 3 -gt 6 ]] 一个条件，满足，就成立 或者的关系 

 [[ 5 -lt 3 || 3 -gt 6 ]] 一个条件满足，就成立 或者的关系 

 


 [[ 5 -lt 3 ]] || [[3 -gt 6 ]] 

 [[ 5 -lt 3 ]] || [[3 -gt 6 ]] 写在外面也可以

 


 


 && 必须两个条件同时满足，和上述一样，这里想说明的问题的是：

 


 [[ 5 -lt 3]] -o [[ 3 -gt 6 ]] [[ 5 -lt 3 -o 3 -gt 6 ]] 

 [[ 5 -lt 3 -a 3 -gt 6 ]] [[ 5 -lt 3 -a 3 -gt 6 ]] 

 -a 和 -o就不成立了，是因为，[[]] 双对中括号，不能使用 -o和 -a的参数

 直接报错：


   2.2 [ ] 可以使用 -a -o的参数，但是必须在 [ ] 中括号内，判断条件，例如：
  

   [ 5 -lt 3 -o 3 -gt 2 ] 或者条件成立
  

   [5 -lt 3 ] -o [ 3 -gt 2] 或者条件， 这个不成立，因为必须在中括号内判断
  

   

  

   如果想在中括号外判断两个条件，必须用&& 和 || 比较
  

   [5 -lt 3 ] || [ 3 -gt 2] 
  

   [5 -gt 3 ] && [ 3 -gt 2] 成立
  

   

  

   相对的，|| 和 && 不能在中括号内使用，只能在中括号外使用
  

   3、当判断某个变量的值是否满足正则表达式的时候，必须使用[[ ]] 双对中括号
  

   
   

  

   单对中括号，直接报错：
  


# 案例实践

## 批量解压、分类归档压缩

```shell
#!/bin/sh
# 1. 批量解压
for zipfile in `ls tauriel关键词下载-*.zip`
do
  unzip $zipfile -d temp
done

for file in `ls temp/download_cpc_info_*.csv`
do
  filename=$(basename "$file")
  # 名称download_cpc_info_22525608_1.csv 通过‘_’来分隔，如果想要取账号，就是取第4项
  accountId=`echo $filename | cut -d '_' -f 4`
  echo "Input File: $file"
  echo "Input FileName: $filename"
  echo "Input accountId: $accountId"
  zip $accountId.zip $file
done

```