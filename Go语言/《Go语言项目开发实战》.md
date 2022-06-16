
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

规范设计

一个项目的规范设计主要包括编码类和非编码类这两类规范。

## 非编码类规范

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

**如何规范目录?**

>目录规范，通常是指我们的项目由哪些目录组成，每个目录下存放什么文件、实现什么功能，以及各个目录间的依赖关系是什么等。

- 命名清晰： 不长不短，最好用单数。
- 功能明确： 当需要新增一个功能时，能够清楚知道放哪个目录。
- 全面性： 尽可能全面地包含研发过程中需要的功能，例如文档、脚本、源码管理、API实现、工具、第三方包、测试、编译产物。
- 可预测性： 能够在项目变大时，仍然保持之前的目录结构。
- 可扩展性： 存同类功能，项目变大时，还可以存更多？ **莫名其妙**，感觉像是再说，子目录不要取名字太宽泛，避免目录太深。


根据功能，我们可以将目录结构分为结构化目录结构和平铺式目录结构两种。
- 结构化目录一般用在Go应用中，相对复杂。
- 平铺式目录一般用在Go包中，相对简单。
  - 引用路径长度明显减少


应用目录结构分为3大部分：Go 应用 、项目管理和文档。
- Go应用主要存放前后端代码
  - /web  前端代码，静态资源、服务端模板、单页应用
  - /cmd  组件
    - 每个组件的目录名应该跟你期望的可执行文件名是一致的。
    - 这里要保证 /cmd/<组件名> 目 录下不要存放太多的代码。
    - cmd/<component-name>/main.go
  - /internal 私有应用和库代码，不希望在其他应用和库中被导入的代码。
    - 可以通过 Go 语言本身的机制来约束其他项目 import 项目内部的包。
    - /internal/apiserver: 该目录中存放真实的应用代码。
    - /internal/pkg ： 存放项目内可共享，项目外不共享的包。 校验代码、code码
    - 建议：一开始将所有的共享代码存放在 /internal/pkg 目录下，当该共享代码做 好了对外开发的准备后，再转存到/pkg目录下。
  - /pkg ： 存放可以被外部应用使用的代码库，其他项目可以直接通过 import 导入这里的代码。
  - /vendor  项目依赖，可通过 go mod vendor 创建。需要注意的是，如果是一个 Go 库，不要提交 vendor 依赖包。
  - /third_party 外部帮助工具，分支代码或其他第三方应用(例如 Swagger UI)
    - 比如我们 fork 了一个 第三方 go 包，并做了一些小的改动，我们可以放在目录 /third_party/forked 下。
  - /test  用于存放其他外部测试应用和测试数据。
    - Go 也会忽略以“.”或 “_” 开头的目录或文件。这样在命名测试数据目 录方面，可以具有更大的灵活性。
  - /configs  配置文件模板或默认配置，敏感信息用占位符取代，不要放在配置代码中。 
  - /deployments  用来存放 Iaas、PaaS 系统和容器编排部署配置和模板
  - /init  存放初始化系统(systemd，upstart，sysv)和进程管理配置文件(runit， supervisord)。比如 sysemd 的 unit 文件。这类文件，在非容器化部署的项目中会用到。
- 项目管理类
  - /Makefile 一个很老的项目管理工具，通 常用来执行静态代码检查、单元测试、编译等功能。
    - 执行顺序建议： 首先生成代码 gen -> format -> lint -> test -> build。
    - 在实际开发中，我们可以将一些重复性的工作自动化，并添加到 Makefile 文件中统一管 理。
  - /script  存放脚本文件，实现构建、安装、分析等不同功能。
    - /scripts/make-rules:用来存放 makefile 文件，实现 /Makefile 文件中的各个功能。 Makefile 有很多功能.
    - /scripts/lib:shell 库，用来存放 shell 脚本。
      - shell 脚本中的函数名，建议采用语义化的命名方式，例如 iam::log::info 这种 语义化的命名方式，可以使调用者轻松的辨别出函数的功能类别，便于函数的管理和引 用。
    - /scripts/install:如果项目支持自动化部署，可以将自动化部署脚本放在此目录下。
  - /build  存放安装包和持续集成相关的文件
    - /build/package:存放容器(Docker)、系统(deb, rpm, pkg)的包配置和脚本。 
    - /build/ci:存放 CI(travis，circle，drone)的配置文件和脚本。 
    - /build/docker:存放子项目各个组件的 Dockerfile 文件。
  - /tools 存放这个项目的支持工具。这些工具可导入来自 /pkg 和 /internal 目录的代码.
  - /assets  其他资源 (图片、CSS、JavaScript 等)。
  - /website  如果你不使用 GitHub 页面，那么可以在这里放置项目网站相关的数据。
- 文档
  - /README.md  包含了项目的介绍、功能、快速安装和使用指引、详细的文档链 接以及开发指引等。
    - 过长需要跳转，需要 添加 markdown toc 索引，可以借助工具  tocenize 来完成索引的添加。
  - /docs 存放设计文档、开发文档和用户文档等(除了 godoc 生成的文档)。
    - /docs/devel/{en-US,zh-CN}:存放开发文档、hack 文档等。
    - /docs/guide/{en-US,zh-CN}: 存放用户手册，安装、quickstart、产品文档等，分为中 文文档和英文文档。
    - /docs/images:存放图片文件。
  - /api 目录中存放的是当前项目对外提供的各种不同类型的 API 接口定义文件.
    - 其中可能包含类似 /api/protobuf-spec、/api/thrift-spec、/api/http-spec、 openapi、swagger 的目录.
  - /CONTRIBUTING.md 开源贡献说明
  - /LICENSE  
    - 常用的开源协议有:Apache 2.0、MIT、 BSD、GPL、Mozilla、LGPL。
    - 可自动化生成，推荐工具:  addlicense 。
  - /CHANGELOG  为了方便了解当前版本的更新内容或者历史更新内容，需要将更新记录 存放到 CHANGELOG 目录
  - /examples  存放应用程序或者公共包的示例代码。


**一些建议**
- 对于小型项目， 可以考虑先包含 cmd、pkg、internal 3 个目录，其他目录后面按需创建。
- 空目录无法提交到git，可以加一个 .keep 文件
- utils, common这类目录不建议用，在Go项目中，每个包名字应该唯一、功能单一、明确。这 类目录存放了很杂的功能，后期维护、查找都很麻烦。

- GO 基于功能划分目录， DDD 基于模型划分目录，如何理解？


### 工作流设计： **如何设计合理的多人开发模式?**

在使用 Git 开发时，有 4 种常用的工作流，也叫开发模式，按演进顺序分为集中式工作流、功能分支工作流、Git Flow 工作流和 Forking 工作流。
- 集中式工作流：
  -  都在主干开发。
  - 适合用在团队 人数少、开发不频繁、不需要同时维护多个版本的小项目中
- 分支工作流：
  - 在功能分支上进行开发，开发完再合并到主干。 不是版本号递增那种，而是自己取名字。 git checkout -b feature/rate-limiting
  - `git merge --no-ff`  ： feature 分支上所有的 commit 都会加到 master 分支上，并且会生成一个 merge commit。
  - `git merge --squash` : 使该 pull request 上的所有 commit 都合并成一个 commit ，然后加到 master 分支上，但**原来的 commit 历史会丢失**。如果开发人员在 feature 分支上提交 的 commit 非常随意，没有规范，那么我们可以选择这种方法来丢弃无意义的 commit。
  - `git rebase`: 将 pull request 上的所有提交历史按照原有顺序依次添加到 master 分支的头部(HEAD)。不熟悉别用。
- Git Flow 工作流
  - Git Flow 中定义了 5 种分支，分别是 master、develop、feature、release 和 hotfix。 其中，master 和 develop 为常驻分支，其他为非常驻分支，不同的研发阶段会用到不同 的分支。
  - 上手难度大，Git Flow 工作流的 每个分支分工明确，这可以最大程度减少它们之间的相互影响。。
  - 比较适合开发团队相对固定，规模较大的项目。
-  Forking 工作流
   - 适用于：开源项目、开发者有衍生出自己的衍生版的需求、开发者不固定。
   - fork到自己账号下，clone到本地，创建功能分支，开发commit，合并commit：git rebase -i origin/master，git rebase -i --autosquash 自动合并commit，push到主干。
   - 在个人远程仓库页面创建 pull request。创建 pull request 时，base 通常选择目标 远程仓库的 master 分支。

### 研发流程设计

#### 如何设计 Go 项目的开发流程?
待看，暂时不需要


#### 如何管理应用的生命周期?
待看


## 编码规范

### 设计方法:怎么写出优雅的 Go 项目?

1. 为什么是 Go 项目，而不是 Go 应用? 
>Go 项目是一个偏工程化的概念，不仅包含了 Go 应用，还包含了项目 管理和项目文档:


2. 一个优雅的 Go 项目具有哪些特点?
>不仅要求我们的 Go 应用是优雅的，还要确保我们的项目管理和文档也是优雅的。
* 符合 Go 编码规范和最佳实践; 
* 易阅读、易理解，易维护; 
* 易测试、易扩展; 
* 代码质量高。


**编写高质量的 Go 应用**


做好5个方面: 代码结构、代码规范、代码质量、编程哲学和软件设计方法.


- **代码结构**
  - 组织一个好的目录结构，看前面那讲。
  - 选择一个好的模块拆分方法。目的就是模块职责分明，高内聚低耦合
    - 按层拆分，比如MVC结构。
      - 问题：相同功能可能在不同层被使用到，而这些功能又分散在不同的层中，很容易造成循环引用。
    - 按功能拆分，比如把user、order、billing拆分为三个模块。
      - 好处1：不同模块，功能单一，可以实现高内聚低耦合的设计哲学。
      - 好处2：因为所有的功能只需要实现一次，引用逻辑清晰，会大大减少出现循环引用的概率。
- **代码规范**
  - 编码规范： 《Uber Go 语言编码规范》比较受欢迎
  - 静态代码检查工具： golangci-lint
  - 最佳实践文章：
    - 《Effective Go》: 高效 Go 编程，由 Golang 官方编写。
    - 《Go Code Review Comments》:Golang 官方编写的 Go 最佳实践，作为 Effective Go 的补充。
    -  Style guideline for Go packages:包含了如何组织 Go 包、如何命名 Go 包、如何 写 Go 包文档的一些建议。
- **代码质量**
  - 写**单测**，mock
    - 为了**提高代码的可测性**，降低单元测试的复杂度，对 function 和 mock 的要求是:
      - 要尽可能减少 function 中的依赖，让 function 只依赖必要的模块。编写一个功能单 一、职责分明的函数，会有利于减少依赖。
      - 依赖模块应该是易 Mock 的。
    - 常用mock工具
      -  golang/mock，是官方提供的 Mock 框架。它实现了基于 interface 的 Mock 功 能，能够与 Golang 内置的 testing 包做很好的集成
      -  sqlmock，可以用来模拟数据库连接。
      -  httpmock，可以用来 Mock HTTP 请求。
      -  bouk/monkey，猴子补丁，能够通过替换函数指针的方式来修改任意函数的实现。  猴子补丁提供了单元测试 Mock 依赖的最终解决方案。
      -  定期检查单元测试覆盖率
         -  `go test -race -cover -coverprofile=./coverage.out -timeout=10m -short -v ./...`
         -  `go tool cover -func ./coverage.out`
    - 提高我们的单元测试覆盖率
      - 使用 gotests 工具自动生成单元测试代码，减少编写单元测试用例的工作量，将你从重 复的劳动中解放出来。

  - **Code Review** 

编写高质量 Go 代码的- 
- 外功: 组织一个合理的代码结构、编写符合 Go 代码规范的代码、保 证代码质量，在我看来都是。
- 内功: 编程哲学和软件设计方法。


**编程哲学**

>面向接口编 程和面向“对象”编程。

Go 接口是一组方法的集合。
任何类型，只要实现了该接口中的方法集，那么就属于这个类型，也称为实现了该接口。
接口的作用，其实就是为不同层级的模块提供一个定义好的中间层。-这样，上游不再需要 依赖下游的具体实现，充分地对上下游进行了解耦。


### Go常用设计模式概述



# 三、基础功能设计或开发
>开发基础功能，如日志包、错误包、错误码

## API 风格

### 如何设计RESTful API?


### RPC API介绍


## 项目管理:如何编写高质量的Makefile?


## 研发流程实战:IAM项目是如何进行研发流程管理的?


## 代码检查:如何进行静态代码检查?

## API 文档:如何生成 Swagger API 文档 ?


## 错误处理

### 如何设计一套科学的错误码?


### 如何设计错误包?


## 日志处理

### 如何设计日志包并记录日志?


### 手把手教你从 0 编写一个日志包


## 应用构建三剑客:Pflag、Viper、Cobra 核心功能介绍


## 应用构建实战:如何构建一个优秀的企业应用框架?

# 四、服务开发
>解析一个企业级的 Go 项目代码，让你学会如何开发 Go 应用. 怎么设计和开发 API 服务、Go SDK、客户端工具


# 五、服务测试
>讲解单元测试、功能测试、性能分析和 性能调优的方法，最终让你交付一个性能和稳定性都经过充分测试的、生产级可用的服 务。


# 六、服务部署
>如何部署一个高可用、安 全、具备容灾能力，又可以轻松水平扩展的企业应用。 传统部署和容器化部署。
