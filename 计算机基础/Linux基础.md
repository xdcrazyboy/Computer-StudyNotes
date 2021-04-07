# 常用操作

## 查看

1. 查磁盘使用率
```
df -Th
```
    - 查看当前目录大小
     - du -h --max-depth=0 
        > --max-depth=n表示只深入到第n层目录，此处设置为0，即表示不深入到子目录。
    - du -s * | sort -nr | head 选出排在前面的10个，
    - du -s * | sort -nr | tail 选出排在后面的10个。

2. 网络

```
 ifconfig
```

3. 内存使用情况
```
free -m

以MB为单位显示内存使用情况
```
- 查看java程序设的内存，可以通过 ps -ef | grep jar (如果是resin容器启动，就看resin，设置是conf里面的resin.properties)

jvm_args  : -Xms1024m -Xmx15000m -XX:MaxPermSize=2048m -Xdebug -Xrunjdwp:transport=dt_socket,address=5005,server=y,suspend=n

1. CPU 

top后键入P看一下谁占用最大
```
# top -d 5
```

5. 端口占用情况：
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