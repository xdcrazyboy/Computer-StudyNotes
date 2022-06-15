
[TOC]

# 一、环境准备
>如何准备开发环境、制作 CA 证书，安装和配置用到的数据库、应用，以及 Shell 脚本编写技巧


**项目背景**： 实现一个IAM（Identity and Access Management，身份识别与访问管理）系统。
- 为了保障 Go 应用的安全，我们需要对访问进行认证，对资源进行授权。


如何实现访问认证和资源授权呢?
- 认证功能不复杂，我们可以通过 JWT (JSON Web Token)认证来实现。
- 授权功能的复杂性使得它可以囊括很多 Go 开发技能点。 本专栏学习就是将这两种功能实现升级为IAM系统，讲解它的构建过程。


**创建数据库**

```shell
sudo tee /etc/yum.repos.d/mongodb-org-4.4.repo<<'EOF'
[mongodb-org-4.4]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/$releasever/mongodb-org/4.4/x86_64
gpgcheck=1
enabled=1
gpgkey=https://www.mongodb.org/static/pgp/server-4.4.asc
EOF
```

**创建CA证书**


```shell
tee ca-config.json << EOF
{
    "signing": {
        "default": {
        "expiry": "87600h"
        },
        "profiles": {
        "iam": {
            "usages": [
            "signing",
            "key encipherment",
            "server auth",
            "client auth"
            ],
            "expiry": "876000h"
        }
        } 
    }
} 
EOF
```

```shell
$ tee ca-csr.json << EOF 
{
    "CN": "iam-ca",
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names":[
        {
            "C": "CN",
            "ST": "BeiJing",
            "L": "BeiJing",
            "O": "marmotedu",
            "OU": "iam"
        }
    ],
    "ca": {
        "expiry": "876000h"
    }
}
EOF
    
```

```shell
tee iam-apiserver-csr.json <<EOF
  "CN": "iam-apiserver",
  "key": {
    "algo": "rsa",
    "size": 2048
  },
"names": [ {
      "C": "CN",
      "ST": "BeiJing",
      "L": "BeiJing",
      "O": "marmotedu",
      "OU": "iam-apiserver"
} ],
  "hosts": [
    "127.0.0.1",
    "localhost",
    "iam.api.marmotedu.com"
] }
EOF
```

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2NTQ5MjQyNTgsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2NTQ4Mzc4NTgsInN1YiI6ImFkbWluIn0.NB4jJIfet4lfvfJN6KRwQu56VFajxvgS4cDI9BTfRso

'{"password":"User@2021","metadata":{"name":"colin"},"nickname":"colin","email":"colin@foxmail.com","phone":"1812884xxxx"}'


```shell
 curl -s -XPOST -H'Content-Type: application/json' -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2NTQ5MjQyNTgsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2NTQ4Mzc4NTgsInN1YiI6ImFkbWluIn0.NB4jJIfet4lfvfJN6KRwQu56VFajxvgS4cDI9BTfRso' -d '{"password":"User@2021","metadata":{"name":"colin"},"nickname":"colin","email":"colin@foxmail.com","phone":"1812884xxxx"}' http://127.0.0.1:8080/v1/users

  curl -s -XGET -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2NTQ5MjQyNTgsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2NTQ4Mzc4NTgsInN1YiI6ImFkbWluIn0.NB4jJIfet4lfvfJN6KRwQu56VFajxvgS4cDI9BTfRso' 'http://127.0.0.1:8080/v1/users?offset=0&limit=10'
```

```shell
{
"CN": "admin",
"key": {
  "algo": "rsa",
  "size": 2048
},
"names": [ {
} ],
"C": "CN",
"ST": "BeiJing",
"L": "BeiJing",
"O": "marmotedu",
"OU": "iamctl"
"hosts": []
}
```

cfssl gencert -ca=${IAM_CONFIG_DIR}/cert/ca.pem -ca-key=${IAM_CONFIG_DIR}/cert/ca-key.pem  -config=${IAM_CONFIG_DIR}/cert/ca-config.json -profile=iam admin-csr.json | cfssljson -bare admin



# 二、规范设计
>目录规范、日志规范、错误码规范、Commit规范

### 规范设计

一个项目的规范设计主要包括编码类和非编码类这两类规范。

#### 非编码类规范

1. 新开发的项目最好按照开源标准来规范，以驱动其成为一个高质量的项目。
2. 开发之前，最好提前规范好文档目录，并选择一种合适的方式来编写 API 文档。
3. 版本规范


**API文档规范**


一个规范的 API 接口文档，通常需要包含一个完整的 API 接口介绍文档、API 接口变更 历史文档、通用说明、数据结构说明、错误码描述和 API 接口使用文档。


**版本号规范**

- 到底该如何确定版本号呢?
  * 第一，在实际开发的时候，我建议你使用 0.1.0 作为第一个开发版本号，并在后续的每次 发行时递增次版本号。
  * 第二，当我们的版本是一个稳定的版本，并且第一次对外发布时，版本号可以定为 1.0.0。
  * 第三，当我们严格按照 Angular commit message 规范提交代码时，版本号可以这么来确定:
    * fix 类型的 commit 可以将修订号 +1。
    * feat 类型的 commit 可以将次版本号 +1。
    * 带有 BREAKING CHANGE 的 commit 可以将主版本号 +1。


**Commit 规范**


- Commit Message 包含三个部分，分别是 Header、Body 和Footer，格式如下:
```c
 <type>[optional scope]: <description> 
 // 空行
 [optional body]
 // 空行
 [optional footer(s)]
```
- 每行50/72/100字比较合理。
- Type分为两大类
  - Production： 修改会影响最终用户和生成环境的代码，这类改动要做好充分的测试。
    - feat： 新增功能
    - fix： Bug修复
    - perf： 提高代码性能的变更
    - refactor： 其他代码类变更， 不属于feat、fix、perf、style的。
  - Development： 一些项目管理类的变更，不影响生产环境，可以免测发布，比如CI流程、构建方式
    - style： 代码格式类变更。 比如删除空行、gofmt格式化代码。
    - docs： 文档的更新
    - test： 新增测试用例或者更新现有测试用例
    - chore： 其他不影响生产环境的变更。 比如 构建流程、依赖管理或者辅助工具的变更。


**如何归类？**
- 如果**变更了应用代码**： 在代码类中，有 4 种具有明确变更意图的类型:feat、fix、perf 和 style;如果我们的代码变更不 属于这 4 类，那就全都归为 refactor 类，也就是优化代码。
- 如果我们**变更了非应用代码**： 例如更改了文档，那它属于非代码类。在非代码类中，有 3 种具有明确变更意图的类型:test、ci、docs;如果我们的非代码变更不属于这 3 类，那 就全部归入到 chore 类。


- **scope**： 说明 commit 的影响范围的，它必须是名词。
  - 主要是根据组件名和功能来设置的。例如，支持 apiserver、 authzserver、user 这些 scope。
  - 不适合设置太具体的值。
- **subject**： 简短描述
  - 必须以动词开头，使用现在时。
  - subject 的 结尾不能加英文句号。
- **body**： 更详细的变更说明，可选
  - 动词开头，必须包括修改的动机，以及和上一个版本的改动点。
    >例如： The body is mandatory for all commits except for those of scope "docs". When t。。。
- **footer**： 说明本次commit导致的后果。
  - 在实际应用中，Footer 通常用来说明不兼容的改动和关闭的 Issue 列表
  - 不兼容的改动: 如果当前代码跟上一个版本不兼容，需要在 Footer 部分，以 `BREAKING CHANG`: 开头，后面跟上不兼容改动的摘要。
  - 关闭的 Issue 列表:关闭的 Bug 需要在 Footer 部分新建一行，并以 Closes 开头列 出，例如:Closes #123。关闭多个用逗号分隔。
- 还原： revert: 开头，后跟还原的 commit 的 Header。
  - 而且，在 Body 中必须写成 This reverts commit <hash> ，其中 hash 是 要还原的 commit 的 SHA 标识。

>为了更好地遵循 Angular 规范，建议你在提交代码时养成不用 git commit -m，即不用 -m 选项的习惯，而是直接用 git commit 或者 git commit -a 进入交互界面编辑 Commit Message。


 **提交频率**


>随意 commit 不仅会让 Commit Message 变得难以理解，还会让其他研发同事觉 得你不专业
  - 开发完一个完整的功能，测试通过后就提交.
  - 规定一个时间，定期提交。这里我建议代码下班前固定提交一次，并且要确保本地未提交的代码，延期不超过 1 天。


 **rebase**


- 可以在最后合并代码或者提交 Pull Request 前，执行 git rebase -i 合并之前的所有 commit。
  - 如何操作？用rebase。
  - git rebase 的最大作用是它可以重写历史。
  - 通常会通过 git rebase -i <commit ID>使用 git rebase 命令，-i 参数表示交 互(interactive)，该命令会进入到一个交互界面中，其实就是 Vim 编辑器。
    - 首先列出给定<commit ID>之前(不包括，越下面越新)的所有 commit，每个 commit 前面有一个操作命令，默认是 pick。
    - 可以选择不同的 commit，并修改 commit 前面的命令，来对该 commit 执行不同的变更操作。
  - git rebase支持的变更操作：
    - p，pick 不对该commit做任何处理
    - r，reword  保留该commit，但是修改提交信息
    - edit  保留该commit，但是rebase时会暂停，允许你修改这个commit
    - squash 保留该commit，将当前commit与上一个commit合并
    - fixup 与squash相同，但是不会保存当前commit的提交信息
    - exec 执行其他shell命令
    - drop 删除该commit
  - squash 和 fixup 可以用来合并 commit。
    >例如用 squash 来合并， 我们只需要把要合并的 commit 前面的动词，改成 squash(或者 s)即可。

```s
 pick 07c5abd Introduce OpenPGP and teach basic usage 2 s de9b1eb Fix PostChecker::Post#urls
 s 3e7ee36 Hey kids, stop all the highlighting
 pick fa20af3 git interactive rebase, squash, amend

 //合并成2条commit
 # This is a combination of 3 commits.
 # The first commit's message is:
 Introduce OpenPGP and teach basic usage

 # This is the 2ndCommit Message:
 Fix PostChecker::Post#urls

 # This is the 3rdCommit Message:
 Hey kids, stop all the highlighting

```
**注意事项**：
* 删除某个 commit 行，则该 commit 会丢失掉。 
* 删除所有的 commit 行，则 rebase 会被终止掉。 
* 可以对 commits 进行排序，git 会从上到下进行合并。


**步骤**：
1. `git rebase -i <commid ID>`
2. 编辑交互界面，执行squash 操作，在每个提交前面增加 s。
3. 看一下是否合并成功：`git log --oneline`
4. `git checkout master`
5. `git merge feature/user`  可以将 feature 分支 feature/user 的改动合并到主干分支，从而完成新 功能的开发。
6. `git log --oneline`


如果有太多commit需要合并，可以不合并，撤销之前n次，然后再建一个新的。
```s
 $ git reset HEAD~3
 $gitadd.
 $ git commit -am "feat(user): add user resource"
```
>除了 commit 实在太多的时候，一般情况下我不建议用这种方法，有点粗 暴，而且之前提交的 Commit Message 都要重新整理一遍。


**修改 Commit Message**


>遇到提交的 Commit Message 不 符合规范的情况，这个时候就需要我们能够修改之前某次 commit 的 Commit Message。
有两种修改方法，分别对应两种不同情况:
1. git commit --amend:修改最近一次 commit 的 message;
   1. 在当前 Git 仓库下执行命令:git commit --amend，后会进入一个交互界面，在交互界 面中，修改最近一次的 Commit Message，
2. git rebase -i:修改某次 commit 的 message。
   1. 先看当前分支日志记录：`git log --oneline`
   2. 指定想修改的记录上一条commit的 id 的 message。`git rebase -i 55892fa`，
   3. 使用 reword 或者 r，保留倒 数第二次的变更信息，但是修改其 message
   > git rebase -i 会变更父 commit ID 之后所有提交的 commit ID。

>如果当前分支有未 commit 的代码，需要先执行 git stash 将工作状态进行暂存，当 修改完成后再执行 git stash pop 恢复之前的工作状态。


###  目录结构设计:如何组织一个可维护、可扩展的代码目录?


## 代码规范



# 三、基础功能设计或开发
>开发基础功能，如日志包、错误包、错误码

# 四、服务开发
>解析一个企业级的 Go 项目代码，让你学会如何开发 Go 应用. 怎么设计和开发 API 服务、Go SDK、客户端工具


# 五、服务测试
>讲解单元测试、功能测试、性能分析和 性能调优的方法，最终让你交付一个性能和稳定性都经过充分测试的、生产级可用的服 务。


# 六、服务部署
>如何部署一个高可用、安 全、具备容灾能力，又可以轻松水平扩展的企业应用。 传统部署和容器化部署。
