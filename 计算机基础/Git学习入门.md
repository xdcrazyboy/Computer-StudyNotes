[TOC]
# Git 入门 

## 一、 常用命令
- 拉取仓库：`git clone`
- 更新：`git pull` 
- 提交：
    1. `git add .` 将本地**工作区**修改更新到本地的**暂存区**,注意有个.，表示添加当前项目的所有文件
    2. `git commit -m` "这里写提交说明"
        >前两步可以一步到位： git commit -am "a代表add，m代表注释说明，这个选项合二为一"
    3. `git push` 正式提交
- `git checkout master` 切换分支，这里是切换到master主干
- `git status` 命令用于查看项目的当前状态。

## 二、 登录

>这些配置我们也可以在 ~/.gitconfig 或 /etc/gitconfig 看到
- `git config --list`
- `git config --global user.name "fengziboboy"`
- `git config --global user.email "fengziboboy@126.com"`
>不需要全局，可以去掉global，注意 -- 两个杠


**克隆私人仓库**：
>git clone https://github-username:github-password@github.com/github-username/github-template-name.git

## 三、 基本概念

### 3.1 工作区、暂存区、版本区
- 工作区：就是你在电脑里能看到的目录。

- 暂存区：英文叫stage, 或index。一般存放在 ".git目录下" 下的index文件（.git/index）中，所以我们把暂存区有时也叫作索引（index）。

- 版本库：工作区有一个隐藏目录.git，这个不算工作区，而是Git的版本库。

![工作区图]()

* 当对工作区修改（或新增）的文件执行 "git add" 命令时，**暂存区的目录树被更新**，同时**工作区修改**（或新增）的文件内容被**写入到对象库中的一个新的对象中**，而该对象的**ID被记录在暂存区的文件索引**中。

* 当执行提交操作（git commit）时，**暂存区的目录树写到版本库**（对象库）中，**master 分支会做相应的更新**。即 master 指向的目录树就是提交时暂存区的目录树。

* 当执行 "git reset HEAD" 命令时，**暂存区的目录树会被重写**，被 **master 分支指向的目录树所替换**，但是工作区不受影响。

* 当执行 "git rm --cached <file>" 命令时，会直接从**暂存区删除文件**，工作区则不做出改变。

* 当执行 "git checkout ." 或者 "git checkout -- <file>" 命令时，会用暂存区全部或指定的文件替换工作区的文件。这个操作**很危险**，会**清除工作区中未添加到暂存区的改动**。

* 当执行 "git checkout HEAD ." 或者 "git checkout HEAD <file>" 命令时，会用 HEAD 指向的 master 分支中的全部或者部分文件替换暂存区和以及工作区中的文件。这个命令也是**极具危险性**的，因为不但会**清除工作区中未提交**的改动，也会**清除暂存区中未提交**的改动。


## 四、命令详解

### 4.1 git status 和 git diff

**git status**：查看在你上次提交之后是否有修改。

- "AM" 状态的意思是，这个文件在我们将它添加到缓存之后又有改动。

- 加 -s 参数，以获得简短的结果输出，如果没加该参数会详细输出内容。

**git diff**：查看执行 git status 的结果的详细信息。显示**已写入缓存**与**已修改但尚未写入缓存**的改动的区别。

git diff 有两个主要的**应用场景**。

- 尚未缓存的改动：`git diff`

- 查看已缓存(也就是add的改动)的改动： `git diff --cached` 

- 查看已缓存的与未缓存的所有改动：`git diff HEAD`

- 显示摘要而非整个 diff：`git diff --stat`

### 4.2 git add 、 git commit
- `git add ./点或者文件名` 命令将想要快照的内容写入缓存区, 
- `git reset HEAD -- 文件名` 用于取消已缓存的内容
- `git rm` 会将条目从缓存区中移除。
    >这与 git reset HEAD 将条目取消缓存是有区别的。 "取消缓存"的意思就是将缓存区恢复为我们做出修改之前的样子。
    - 默认情况下，git rm file 会将文件从**缓存区**和你的硬盘中（**工作目录**）删除。
    - 如果你要在工作目录中留着该文件，可以使用 `git rm --cached file`：只删除缓存区的文件
- `git commit` 将缓存区内容添加到仓库中。

####   git add 添加 多余文件 
这样的错误是由于， 有的时候 可能: git add . （空格+ 点） 表示当前目录所有文件，不小心就会提交其他文件

git add 如果添加了错误的文件的话

- 撤销操作
  * git status 先看一下add 中的文件 
  * git reset HEAD 如果后面什么都不跟的话 就是上一次add 里面的全部撤销了 
  * git reset HEAD XXX/XXX/XXX.java 就是对某个文件进行撤销了

###  4.3 git log 
**查看提交历史**： `git log`
  - 看简洁版本：--oneline 
  - 查看历史中什么时候出现了分支、合并：--graph 
  - 逆向显示所有日志：--reverse
  -  只想查找指定用户的提交日志可以使用命令：git log --author=zhaojinbo
  -  限定时间： --before={3.weeks.ago} --after={2010-04-18}
  -  更多git log相关请看：http://git-scm.com/docs/git-log

## 五、 分支管理
- **创建**分支命令： `git branch (branchname)` 
    >创建完，你当前还是没到分支，进行的操作还是在原来位置。
- **切换**分支命令: `git checkout (branchname)` 
    >当你切换分支的时候，Git 会用该分支的最后提交的快照替换你的工作目录的内容， 所以多个分支不需要多个目录。

- **合并**分支命令: `git merge`
    >将指定的分支合并到当前工作的分支
- **列出**分支基本命令： `git branch` : * master 
    >表示：我们有一个叫做"master"的分支，并且该分支是当前分支。
    >当你执行 `git init` 的时候，默认情况下 Git 就会为你创建"master"分支。如果我们要手动创建一个分支。执行 git branch (branchname) 即可。

- **删除**分支：`git branch -d (branchname)`

- `git checkout -b (branchname)` 命令来创建新分支并立即切换到该分支下，从而在该分支中操作。

### 5.1 合并分支
- **合并**分支命令: `git merge`,创建完，你当前还是没到分支，进行的操作还是在原来位置。
- **合并冲突**：可以用git diff 查看：合并冲突就出现了，接下来我们需要手动去修改它。 解决完用 git add 去查看是否修改成功。 （**带完善补充**）

## 六、 不常用的命令

- `git mv a a1` 重命名
- `git tag -a v1.0` : 如果你达到一个重要的阶段，并希望永远记住那个特别的提交快照，你可以使用 git tag 给它打上标签。
    >-a 选项意为"创建一个带注解的标签"。 不用 -a 选项也可以执行的，但它不会记录这标签是啥时候打的，谁打的，也不会让你添加个标签的注解。
    - 追加标签，给过去某次提交：git tag -a v0.9 85fc7e7
    - 查看所有标签：git tag
    - 指定标签信息命令：git tag -a <tagname> -m "w3cschool.cc标签"
    - PGP签名标签命令：git tag -s <tagname> -m "w3cschool.cc标签"

## 七、GitHub
 Github 作为远程仓库，你可以 [Github 简明教程](http://rogerdudler.github.io/git-guide/index.zh.html?spm=5176.10731542.0.0.189e684eQPfjFK)。

 - 添加远程库:可以指定一个简单的名字，以便将来引用,命令格式如下：`git remote add [shortname] [url]`


# 八、 Git 服务器搭建
上一章节中我们远程仓库使用了 Github，Github 公开的项目是免费的，但是如果你不想让其他人看到你的项目就需要收费。

这时我们就需要自己搭建一台Git服务器作为私有仓库使用。

接下来我们将以 Centos 为例搭建 Git 服务器。

1. 安装Git
```s
    $ yum install curl-devel expat-devel gettext-devel openssl-devel zlib-devel perl-devel
    $ yum install git

```
接下来我们 创建一个git用户组和用户，用来运行git服务：
```s
    $ groupadd git
    $ adduser git -g git
```
2. 创建证书登录
收集所有需要登录的用户的公钥，公钥位于id_rsa.pub文件中，把我们的公钥导入到/home/git/.ssh/authorized_keys文件里，一行一个。

如果没有该文件创建它：

```s
    $ cd /home/git/
    $ mkdir .ssh
    $ chmod 700 .ssh
    $ touch .ssh/authorized_keys

```
3. 初始化Git仓库
首先我们选定一个目录作为Git仓库，假定是/home/gitrepo/runoob.git，在/home/gitrepo目录下输入命令：

```s
$ cd /home
$ mkdir gitrepo
$ chown git:git gitrepo/$ cd gitrepo
$ git init --bare runoob.gitInitialized empty Git repository in /home/gitrepo/runoob.git/
```
以上命令Git创建一个空仓库，服务器上的Git仓库通常都以.git结尾。然后，把仓库所属用户改为git：

```s
$ chown -R git:git runoob.git

```
4. 克隆仓库
```s
$ git clone git@192.168.45.4:/home/gitrepo/runoob.git
Cloning into 'runoob'...warning: You appear to have cloned an empty repository.Checking connectivity... done.

```
192.168.45.4 为 Git 所在服务器 ip ，你需要将其修改为你自己的 Git 服务 ip。

这样我们的 Git 服务器安装就完成了，接下来我们可以禁用 git 用户通过shell登录，可以通过编辑/etc/passwd文件完成。找到类似下面的一行：
>`git:x:503:503::/home/git:/bin/bash`
改为：
>`git:x:503:503::/home/git:/sbin/nologin`

# 参考文献
[阿里云开发者社区《学习 Git》](https://developer.aliyun.com/course/489?spm=5176.10731542.0.0.554a684eaXBDC5)

# 问题解决

## remote: Permission to XXXA/xxxx.git denied to XXXB
1. 生成一个新的SSH KEY：
```sh
ssh-keygen  -t rsa –C "youremail@example.com"

```
>如果已经有其他公钥占用了默认名字，可以在`Enter file in which to save the key (): `后面写个名字，比如 xd。

2. 复制xd.pub里面的公钥，到github上的 SSH key，新建复制就好。
3. 修改.git/config 里面的url为 git@github.com:fengziboboy...形式的。 此步是否必须不确定
4. **重要**：Adding your SSH key to the ssh-agent
   1. ensure the ssh-agent is running.
    ```sh
        # start the ssh-agent in the background
        $ eval $(ssh-agent -s)
        > Agent pid 59566
    ```
    2. Add your SSH private key to the ssh-agent.
        >$ ssh-add ~/.ssh/id_rsa

## 撤销某次还未push的commit

> 问题来自于我提交了超过100m的大文件，被拒绝，导致其他commit也都无法push
- 1、git status 查看未被传送到远程代码库的提交次数

2、git cherry -v 查看未被传送到远程代码库的提交描述和说明

3、git reset commit_id 撤销未被传送到远程代码库的提交

做到这里就已经可以重新添加提交了（注意一定要撤销有大文件的提交）

**上面还是存在问题，只能reset到最前面一次，但如果问题就出在那最前面的一次？ 需要重置commit：**
```
git reset --soft HEAD^
```
- 这样就成功的撤销了你的commit, 注意，仅仅是撤回commit操作，您写的代码仍然保留。
- HEAD^的意思是上一个版本，也可以写成HEAD~1, 如果你进行了2次commit，想都撤回，可以使用HE
