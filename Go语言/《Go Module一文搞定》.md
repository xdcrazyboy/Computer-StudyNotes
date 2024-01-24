[TOC]
# Go Module 那些道道

## 导入本地 module
- 可以借助 go.mod 的 replace 指示符，来解决这个问题。
  - 首先，我们需要在 module a 的 go.mod 中的 require 块中，手工加上这一条:
  ```go  
    //这里的 v1.0.0 版本号是一个“假版本号”，目的是满足 go.mod 中 require 块的语法要求。
    require github.com/user/b v1.0.0 
  ```
  - 然后，我们再在 module a 的 go.mod 中使用 replace，将上面对 module b v1.0.0 的依赖，替换为本地路径上的 module b:
  ```go
  replace github.com/user/b v1.0.0 => module b的本地源码路径    
  ```


## 拉取私有 module 的需求与参考方案

- 配置公共 GOPROXY 服务拉取公共 Go Module，同时再把私有仓库配置到 GOPRIVATE 环境变量，就可以了。

- 这样，所有私有 module 的拉取，都会直连代码托管服务器，不会走 GOPROXY 代理服务，也不会去 GOSUMDB 服务器做 Go 包的 hash 值校验。


更多的公司 / 组织，可能会将私有 Go Module 放在公司 / 组织内部的 vcs（代码版本控制）服务器上。 一般有两个方案：
- 第一个方案，通过直连组织公司内部的私有 Go Module 服务器拉取。
  >十分适合内部有完备 IT 基础设施的公司。这类型的公司内部的 vcs 服务器都可以通过域名访问（比如 git.yourcompany.com/user/repo），因此，公司内部员工可以像访问公共 vcs 服务那样，访问内部 vcs 服务器上的私有 Go Module。
- 第二种方案，将外部 Go Module 与私有 Go Module 都交给内部统一的 GOPROXY 服务去处理：
  >可以将所有复杂性都交给 in-house goproxy 这个节点，开发人员可以无差别地拉取公共 module 与私有 module，心智负担降到最低。

## 如何升级修改版本

- **查看版本**： `o list -m -versions github.com/sirupsen/logrus`
- **修改版本**：
  - 方法1： 以在项目的 module 根目录下，执行带有版本号的 go get 命令：`go get github.com/sirupsen/logrus@v1.7.0`
  - 方法2： 修改go.mod文件，然后tidy一下：
    - `go mod edit -require=github.com/sirupsen/logrus@v1.7.0`
    - `go mod tidy`


### 如何添加主版本号大于1的依赖（在代码中import）
在 Go Module 构建模式下，当依赖的主版本号为 0 或 1 的时候，我们在 Go 源码中导入依赖包，**不需要在包的导入路径上增加版本号**，也就是：

```go
import github.com/user/repo/v0 等价于 import github.com/user/repo
import github.com/user/repo/v1 等价于 import github.com/user/repo
```

- 如果新旧版本的包使用相同的导入路径，那么新包与旧包是兼容的。 反过来说，如果不兼容，那剧需要采用不同的导入路径。


如果引入的主版本大于1的依赖（比如v2.0.0），那么就不能直接使用`github.com/user/repo`,因为这是默认0/1，这个与2是不兼容的。 需要向下面这样导入：
```go
import github.com/user/repo/v2/xxx
```
 - 也就是在声明它的导入路径的基础上，加上版本号信息。
 - 然后要从新下载最新的：`go get github.com/go-redis/redis/v7`

### 升级依赖版本到一个不兼容版本

跟上面类似，修改版本号，然后重新下载。
```go

import (
  _ "github.com/go-redis/redis/v8"
  "github.com/google/uuid"
  "github.com/sirupsen/logrus"
)

//
$go get github.com/go-redis/redis/v8
```

### 移除一个依赖

- 在业务代码中删除依赖后，直接build不会删除不用的依赖，因为如果源码满足成功构建的条件，go build 命令是不会“多管闲事”地清理 go.mod 中多余的依赖项的。
- 运行下 `go mod tidy`就行， go mod tidy 会自动分析源码依赖，而且将不再使用的依赖从 go.mod 和 go.sum 中移除。


### 特殊情况：使用vendor

**什么情况下还需要用vendor？**


- 在一些不方便访问外部网络，并且对 Go 应用构建性能敏感的环境，比如在一些内部的持续集成或持续交付环境（CI/CD）中，使用 vendor 机制可以实现与 Go Module 等价的构建。


**怎么用mod模式下用vendor？**

- Go Module 构建模式下，我们再也无需手动维护 vendor 目录下的依赖包了，Go 提供了可以快速建立和更新 vendor 的命令：
  - `go mod vendor` 项目根目录，创建vendor目录。
    - go mod vendor 命令在 vendor 目录下，创建了一份这个项目的依赖包的副本
    - 并且通过 vendor/modules.txt 记录了 vendor 下的 module 以及版本。
- 如果我们要基于 vendor 构建，而不是基于本地缓存的 Go Module 构建，我们需要在 go build 后面加上 `-mod=vendor` 参数。
- 在 Go 1.14 及以后版本中，如果 Go 项目的顶层目录下存在 vendor 目录，那么 go build **默认也会优先基于 vendor 构建**，除非你给 go build 传入 -mod=mod 的参数。


>go get会下载gotests下面所有的包，如果gotests是一个可执行文件的项目（带有main包main函数）. go get会在下载包之后构建这个项目并把可执行文件放入$GOPATH/bin下。


## 作为module维护者，你需要知道的事情

>从 Go Module 的作者或维护者的视角，来聊聊在规划、发布和维护 Go Module 时需要考虑和注意什么事情，包括 go 项目仓库布局、Go Module 的**发布**、**升级 module 主版本号**、作废特定版本的 module.


**仓库布局：是单module耗时多module**

- 能单就单。 管理方便，导入时也方便。
  - 然后我们对仓库打 tag，这个 tag 也会成为 Go Module 的版本号，这样，对仓库的版本管理其实就是对 Go Module 的版本管理。


如果组织层面要求采用单一仓库（monorepo）模式，也就是所有 Go Module 都必须放在一个 repo 下，那我们只能使用单 repo 下管理多个 Go Module 的方法了。 就是所谓的 **大仓**。

例如：
```go
./srmm
├── module1
│   ├── go.mod
│   └── pkg1
│       └── pkg1.go
└── module2
    ├── go.mod
    └── pkg2
        └── pkg2.go
```
- 这种情况下，module 的 path 也不能随意指定，必须包含子目录的名字。
- 如果我们要发布 module1 的 v1.0.0 版本，我们不能通过给仓库打 v1.0.0 这个 tag 号来发布 module1 的 v1.0.0 版本，正确的作法应该是打 module1/v1.0.0 这个 tag 号。


**发布Go Module**


- 发布的步骤也十分简单，就是为 repo 打上 tag 并推送到代码服务器上就好了。
  - 单module，给 repo 打的 tag 就是 module 的版本。
  - 多module，在 tag 中加上各个 module 的子目录名，这样才能起到发布某个 module 版本的作用，否则 module 的用户通过 go get xxx@latest 也无法看到新发布的 module 版本。


**作废特定版本的Go Module**


- 修复 broken 的 module 版本并重新发布。
  - m1 的作者只需要删除掉远程的 tag: v1.0.2，
  - 在本地 fix 掉问题，
  - 然后重新 tag v1.0.2 并 push 发布到 bitbucket 上的仓库中就可以了。
- 如果m1所有的消费者，都是通过m1所在代码托管服务器来获取m1的特定版本，那么只要清理掉本地缓存module cache（`go clean -modcache`），然后再重新构建就可以了.
- 但现实的情况时，现在大家都是通过 Goproxy 服务来获取 module 的。
  - 当某个消费者通过他配置的 goproxy 获取这个版本时，这个版本就会在被缓存在对应的代理服务器上。
  - 后续 m1 的消费者通过这个 goproxy 服务器获取那个版本的 m1 时，请求不会再回到 m1 所在的源代码托管服务器。
- 如果 m1 的作者删除了 bitbucket 上的 v1.0.2 这个发布版本，各大 goproxy 服务器上的 broken v1.0.2 版本是否也会被同步删除呢？ **不会**。


**那怎么解决？**
>Go 社区更为常见的解决方式就是**发布 module 的新 patch 版本**.


现在我们废除掉 v1.0.2，在本地修正问题后，直接打 v1.0.3 标签，并发布 push 到远程代码服务器上。

- 重新拉取最新的会获得v1.0.3， 但是对于之前曾依赖 v1.0.2 版本的消费者 c2 来说，这个时候他们需要手工介入才能解决问题。


从 Go 1.16 版本开始，Go Module 作者还可以在 go.mod 中使用新增加的retract 指示符，标识出哪些版本是作废的且不推荐使用的。retract 的语法形式如下：
```go

// go.mod
retract v1.0.0           // 作废v1.0.0版本
retract [v1.1.0, v1.2.0] // 作废v1.1.0和v1.2.0两个版本
```


如果要提示用某个 module 的某个大版本整个作废，我们用 Go 1.17 版本引入的 Deprecated 注释行更适合。下面是使用 Deprecated 注释行的例子：
```go
// Deprecated: use bitbucket.org/bigwhite/m1/v2 instead.
module bitbucket.org/bigwhite/m1
```

### 升级 module 的 major 版本号

- 在同一个 repo 下，不同 major 号的 module 就是完全不同的 module，甚至同一 repo 下，不同 major 号的 module 可以相互导入。
- 这意味着高版本的代码要与低版本的代码彻底分开维护，通常 Go 社区会采用为新的 major 版本建立新的 major 分支的方式，来将不同 major 版本的代码分离开。


以将 bitbucket.org/bigwhite/m1 的 major 版本号升级到 v2 为例看看。

- 首先，我们要建立 v2 代码分支并切换到 v2 分支上操作
- 然后修改 go.mod 文件中的 module path，增加 v2 后缀：
```go
//go.mod
module bitbucket.org/bigwhite/m1/v2

go 1.17
```
  - 如果module内部包间有互相导入，那么在升级major号时，这些包的 import 路径上也要增加 v2。 否则就会出现高major号的module引用低module。
- 使用者：需要在这个依赖 module 的 import 路径的后面，增加 /vN 就可以了（这里是 /v2），当然代码中也要针对不兼容的部分进行修改，然后 go 工具就会自动下载相关 module。


**多module的情况下升级major版本号？**
分两种情况：

- 第一种情况：repo 下的所有 module 统一进行版本发布。
  - 建立 vN 版本分支，在 vN 分支上对 repo 下所有 module 进行演进，统一打 tag 并发布。
  - 当然 tag 要采用带有 module 子目录名的那种方式，比如：module1/v2.0.0。
- 第二个情况：repo 下的 module 各自独立进行版本发布。
  - 需要建立 major 分支矩阵。
  - 假设我们的一个 repo 下管理了多个 module，从 m1 到 mN，那么 major 号需要升级时，我们就需要将 major 版本号与 module 做一个组合，形成下面的分支矩阵： v2_m1/v2_m2/v3_m1

