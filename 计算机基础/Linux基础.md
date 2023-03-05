# 常用操作

## 常用命令
- **ls**
  - `ll` = `ls -al`  别名
- **cd**
  - `cd ~` = `cd $HOME`
  - `cd -`，回到之前的工作目录
- **cp**
  - `cp -r 源 目标`  ：-r，复制该目录下的所有子目录和文件
  - `cp -s  /home/cmd/hello.c softlink`： 创建软链接
- **cat** 小文件
  - `-n` 显示行号
- **more** 大文件
  - 当文件的内容大于一屏时，more命令可以按页来显示，并且支持翻页、直接跳转行 
  - 打开后命令： q退出；回车一行；空格键下一页；ctrl+b 上一页；=显示当前行号；d，f跳很多页； 
  - +n	从第n行开始显示
  - -n	指定每屏显示的行数
  - +/pattern	在每个文档显示前搜寻该字(pattern)，然后从该字串之后开始显示
- **less** 
  - 与more类似，可以前后翻页，用键盘上下键即可。
  - q退出；b 向后翻一页；u 向前滚动半页；
  - /hello：向下搜索字符串“hello”； ?hello：向上搜索字符串“hello”
- **touch**
  - 改变一个文件的时间戳，或者新建一个空文件.
- **mkdir**
  - `-p` 新建一个已存在的文件夹不会报错
  - `-m` 设置目录的读写权限。 
    - 比如新建一个具有读写、执行权限的目录：scripts：`mkdir -m 777 scripts`
- **rm**
  - `rm -r` 命令删除一个**目录**时，该目录下的所有子目录、文件都会被删除。
  - `rm -i test/*`  删除前需要确认。

- **tar** 打包和解压文件
  - 解压： 常用-zxvf
    - -z	通过gzip指令压缩/解压缩文件，文件格式：*.tar.gz
  - 打包：
    - -c	新建打包压缩文件
    - -r	添加文件到压缩文件 
    - 将某一个目录（比如：test）打包成压缩文件包test.tar.gz: `tar cvfz test.tar.gz test/`
  - 通用
    - -x	解压缩打包文件
    - -v	在压缩/解压缩过程中，显示正在处理的文件名或目录
    - -f	（压缩或解压时）指定要处理的压缩文件
    - -C dir	指定压缩/解压缩的目录，若无指定，默认是当前目录
  - —delete	从压缩文件中删除指定的文件

- **chmod** 改变文件或目录的权限
  - 读权限的值为4、写的权限值为2、可执行的权限值为1， 读写6，读执行5；
  - -R参数是可选的，可以进行递归地持续更改，将指定目录下所有的子目录或文件都修改。
- **whereis**: 定位一个文件的存储位置，这个文件可以是二进制文件、源文件或文本文件。
- **whichis**: 在PATH环境变量指定的路径中，搜索某个系统命令的位置，并且返回第一个搜索结果。
- **whatis**: 用一句话介绍命令的功能
- **tee**: 从标准输入设备读取数据，并写到标准输出设备和指定的文件上。
  - `tee input.txt`: 使用tee命令后，shell就会进入输入交互状态，接收用户从标准输入设备（一般是键盘）输入的字符，将用户输入的字符显示到屏幕上，并写到指定的文件input.txt。
  - ` tee input1.txt  input2.txt  input3.txt` 内容一样
  - `tailf all.log | tee test.log` 会把hello.txt的内容 写入test.log
- **wc** : 统计一个文件中的行数、字数、字节数。
  * -w	统计字数：由空白、tab或换行字符分隔的字符串个数
  * -c	统计字节数，单位为Byte
  * -l	统计行数
  * -m	统计字符数
  * -L	打印最长行的长度
* **ifconfig**: 用来查看和配置网络设备
  * -a	查看全部网络接口配置信息
  * -s 简短摘要信息 类似 `netstat -i`

## 查看

- `tree` 显示目录的树状结构。

### 按照资源分类
1. 查**磁盘**使用率
```
df -Th
```

- 查看当前目录大小
- du -h --max-depth=0 
  > --max-depth=n表示只深入到第n层目录，此处设置为0，即表示不深入到子目录。
- du -s * | sort -nr | head 选出排在前面的10个，
- du -s * | sort -nr | tail 选出排在后面的10个。


2. **网络**

```
 ifconfig
```

**端口使用情况**
```s
netstat -an | ag 2181
```


3. **内存**使用情况
```
free -m

以MB为单位显示内存使用情况
```
- 查看java程序设的内存，可以通过 ps -ef | grep jar (如果是resin容器启动，就看resin，设置是conf里面的resin.properties)

jvm_args  : -Xms1024m -Xmx15000m -XX:MaxPermSize=2048m -Xdebug -Xrunjdwp:transport=dt_socket,address=5005,server=y,suspend=n


4. **CPU** 
5. 
top后键入P看一下谁占用最大
```
# top -d 5
```


5. **端口**占用情况：
   1. windows对比
        ```
           netstat -aon|findstr "49157"  查到pid为2720
           tasklist|findstr "2720"  查看经常为  svchost.exe
           打开任务管理器，关掉或者。
           - taskkill /f /t /im svchost.exe  
           - taskkill -f -pid 14128

        ```

## 字符串操作

**查找目录下所有文件中是否包含某个字符串**：
```sh
find .|xargs grep -ri "showIdeaDetailList.action"
```

## 网络，文件
**下载文件**：  `curl http://www.linux.com >> linux.html`
