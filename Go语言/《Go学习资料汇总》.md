#  Go 成神之路

# 那些优质且权威的Go语言学习资料 （Tony Bai）


## 书籍

1. No.5 [《The Way To Go》](https://github.com/Unknwon/the-way-to-go_ZH_CN)- Ivo Balbaert  Go 语言百科全书
   - 为什么学习 Go 以及 Go 环境安装入门；
   - Go 语言核心语法；
   - Go 高级用法（I/O 读写、错误处理、单元测试、并发编程、socket 与 web 编程等)；
   - Go 应用（常见陷阱、语言应用模式、从性能考量的代码编写建议、现实中的 Go 应用等）。
2. No.4 [《Go 101》](https://github.com/golang101/golang101)- Go 语言参考手册
   - Go 语法基础；
   - Go 类型系统与运行时实现；
   - 以专题（topic）形式阐述的 Go 特性、技巧与实践模式。
3. No.3 《Go 语言学习笔记》- Go 源码剖析与实现原理探索 。 这本书整体上分为两大部分：
   - **Go 语言详解**：以短平快、“堆干货”的风格对 Go 语言语法做了说明，能用示例说明的，绝不用文字做过多修饰；
   - **Go 源码剖析**：这是这本书的精华，也是最受 Gopher 们关注的部分。这部分对 Go 运行时神秘的内存分配、垃圾回收、并发调度、channel 和 defer 的实现原理、sync.Pool 的实现原理都做了细致的源码剖析与原理总结。
4. No.2 《Go 语言实战》- 实战系列经典之作，紧扣 Go 语言的精华。 这本书的结构框架：
   - 入门：快速上手搭建、编写、运行一个 Go 程序；
   - 语法：数组（作为一个类型而存在）、切片和 map；
   - Go 类型系统的与众不同：方法、接口、嵌入类型；
   - Go 的拿手好戏：并发及并发模式；
   - 标准库常用包：log、marshal/unmarshal、io（Reader 和 Writer）；
   - 原生支持的测试。
5. No.1 《Go 程序设计语言》- 人手一本的 Go 语言“圣经”


## 文档

- Go官方文档——最权威的Go语言资料： https://go.dev/doc/


每个 Gopher **必看的内容**：
Go 官方文档中的[Go 语言规范](https://go.dev/ref/spec)、[Go module 参考文档](https://go.dev/ref/mod)、[Go 命令参考手册](https://go.dev/doc/cmd)、[Effective Go](https://go.dev/doc/effective_go)、[Go 标准库包参考手册](https://pkg.go.dev/std)以及[Go 常见问答](https://go.dev/doc/faq)等都是。

>我强烈建议你一定要抽出时间来仔细阅读这些文档。

## 博客、文章、日/周报、邮件列表

**博客**


- [Go语言官方博客](https://go.dev/blog/)
- [Go 核心团队技术负责人 Russ Cox 的个人博客](https://research.swtch.com/)
- [Go 核心开发者 Josh Bleecher Snyder 的个人博客](https://commaok.xyz/)；
- [Go 核心团队前成员 Jaana Dogan 的个人博客](https://rakyll.org/)；
- [Go 鼓吹者 Dave Cheney 的个人博客](https://dave.cheney.net/)；
- [Go 语言培训机构 Ardan Labs 的博客](https://www.ardanlabs.com/blog)；
- [GoCN 社区](https://gocn.vip/)；
- [Go 语言百科全书：由欧长坤维护的 Go 语言百科全书网站](https://golang.design/)。


**Go日报/周刊邮件列表**


- [Go 语言爱好者周刊](https://studygolang.com/go/weekly)，由 Go 语言中文网维护；
- [Gopher 日报](https://github.com/bigwhite/gopherdaily)，由我本人维护的 Gopher 日报项目，创立于 2019 年 9 月。


## 开源项目


## 技术演讲、大会、PPT
>关于 Go 技术演讲，我个人建议以各大洲举办的 GopherCon 技术大会为主，这些已经基本可以涵盖每年 Go 语言领域的最新发展。

- [Go 官方的技术演讲归档](https://go.dev/talks/)，这个文档我强烈建议你按时间顺序看一下，通过这些 Go 核心团队的演讲资料，我们可以清晰地了解 Go 的演化历程；
- [GopherCon 技术大会](https://www.youtube.com/c/GopherAcademy/playlists)，这是 Go 语言领域规模最大的技术盛会，也是 Go 官方技术大会；
- [GopherCon Europe](https://www.youtube.com/c/GopherConEurope/playlists) 技术大会-欧洲分会；
- [GopherConUK 技术大会](https://www.youtube.com/c/GopherConUK/playlists)；
- [GoLab 技术大会](https://www.youtube.com/channel/UCMEvzoHTIdZI7IM8LoRbLsQ/playlists)；
- [Go Devroom@FOSDEM](https://www.youtube.com/user/fosdemtalks/playlists)；
- [GopherChina 技术大会](https://space.bilibili.com/436361287)，这是中国大陆地区规模最大的 Go 语言技术大会，由 GoCN 社区主办。


## 其他


### 高级冷门

- [Go 语言项目的官方 issue 列表](https://github.com/golang/go/issues)
  - 通过这个 issue 列表，我们可以实时看到 Go 项目演进状态，及时看到 Go 社区提交的各种 bug。
  - 同时，我们通过挖掘该列表，还可以了解某个 Go 特性的来龙去脉，这对深入理解 Go 特性很有帮助。
- [Go 项目的代码 review 站点](https://go-review.googlesource.com/q/status:open+-is:wip)
  - 通过阅读 Go 核心团队 review 代码的过程与评审意见，我们可以看到 Go 核心团队是如何使用 Go 进行编码的，能够学习到很多 Go 编码原则以及地道的 Go 语言惯用法，对更深入地理解 Go 语言设计哲学，形成 Go 语言编程思维有很大帮助。


# 从小白到“老鸟”，我的Go语言进阶之路（孔令飞）
>个人在 Go 语言进阶过程中的一些经验、心得的分享。希望通过这些分享，能帮助到渴望在 Go 研发之路上走的更远的你。

## Go 语言能力级别划分

- **初级**：已经学习完 Go 基础语法课程，能够编写一些简单 Go 代码段，或者借助于 Google/Baidu 能够编写相对复杂的 Go 代码段；这个阶段的你基本具备阅读 Go 项目代码的能力；

- **中级**：能够独立编写完整的 Go 程序，例如功能简单的 Go 工具等等，或者借助于 Google/Baidu 能够开发一个完整、简单的 Go 项目。此外，对于项目中涉及到的其他组件，我们也要知道怎么使用 Go 语言进行交互。在这个阶段，开发者也能够二次开发一个相对复杂的 Go 项目；

- **高级**：不仅能够熟练掌握 Go 基础语法，还能使用 Go 语言高级特性，例如 channel、interface、并发编程等，也能使用面向对象的编程思想去开发一个相对复杂的 Go 项目；

- **资深**：熟练掌握 Go 语言编程技能与编程哲学，能够独立编写符合 Go 编程哲学的复杂项目。同时，你需要对 Go 语言生态也有比较好的掌握，具备较好的软件架构能力；

- **专家**：精通 Go 语言及其生态，能够独立开发大型、高质量的 Go 项目，编程过程中较少依赖 Google/ 百度等搜索工具，且对 Go 语言编程有自己的理解和方法论。除此之外，还要具有优秀的软件架构能力，能够设计、并部署一套高可用、可伸缩的 Go 应用。这个级别的开发者应该是团队的技术领军人物，能够把控技术方向、攻克技术难点，解决各种疑难杂症。

>初级、中级、高级 Go 语言工程师的关注点主要还是使用 Go 语言开发一个实现某种业务场景的应用，但是资深和专家级别的 Go 语言工程师，除了要具有优秀的 Go 语言编程能力之外，还需要具备一些**架构能力**.


## 进阶之路

### 开发者阶段

从初级入门 --->  高级

#### 初级

1. **学习基础语法** 
   1. 一般学习一门编程语言，都会快速阅读两本经典的、讲基础语法的书。《Go 程序设计语言》、《Go 语言编程》。
   2. 还有余力的话，再看两本关于场景化编程的书籍：《Go 并发编程实战》（第 2 版）和《Go Web 编程》。
2. **实战**，主要是通过编码实战加深对 Go 语法知识的理解和掌握。 
   1. 那么具体应该**如何实战呢？**， 核心就是**抄和改**。下面重点说一下：
   2. **找需求**：可以研究优秀的（开源）项目 
      1. 以需求为驱动，找到一个合理的需求，然后实现它。
      2. 需求来源于工作。这些需求可以是产品经理交给你的某一个具体的产品需求，也可以是能够帮助团队 / 自己提高工作效率的工具。
      3. 总之，如果有明确的工作需求最好，如果没有明确的需求，我们就要创造需求。
   3. **如何找优秀开源项目？**。 （以要开发一个版本发布系统为例）
      1. 在 GitHub 上找到一个优秀的版本发布系统，并基于这个系统进行二次开发。通过这种方式，我不仅学习到了一个优秀开源项目的设计和实现，还以最快的速度完成了版本发布系统的开发。
      2. 搜索框： `language:go 版本发布`，按照star排序
      3. 看描述，开code
        >研究完 GitHub 上的开源项目，这时候我还建议你通过[libs.garden](https://libs.garden/go)，再查找一些开源项目。libs.garden 的主要功能是库（在 Go 中叫包）和应用的评分网站，是按不同维度去评分的. 再就是到 GitHub 上的 [awesome-go](https://github.com/avelino/awesome-go) 项目也根据分类记录了很多包和工具。
    3. **如何进行二次开发？**
       1. 手动编译、部署这个开源项目。
       2. 阅读项目的 README 文档，跟着 README 文档使用这个开源项目，至少运行一遍核心功能。
       3. 阅读核心逻辑源码，在不清楚的地方，可以添加一些 fmt.Printf 函数，来协助你理解代码。
       4. 在你理解了项目的核心逻辑或者架构之后，就可以尝试添加 / 修改一些匹配自己项目需求的功能，添加后编译、部署，并调试。
       5. 二次开发完之后，你还需要思考下后续要不要同步社区的代码，如果需要，如何同步代码。
    >在你通过“抄”和“改”完成需求之后，记得还要编写文档，并找个合适的时机在团队中分享你的收获和产出。这点很重要，可以将你的学习输入变成工作产出。


**这样调研很多项目，花费这么多时间的好处是什么？**
1. 最优解：你可以很有底气地跟老板说，这个方案在这个类别就是业界 No.1（开源的）。
2. 高效：基于已有项目进行二次开发，可以提高开发和学习效率。
3. 产出：在学习的过程中，也有工作产出。个人成长、工作贡献可以一起获得。
4. 知识积累：为今后的开发生涯积累项目库和代码库。

#### 中级/高级工程师阶段
中级 / 高级工程师阶段，其实就是不断地利用所学的 Go 基础知识，去编程实践。


这个阶段提升 Go 研发能力的思路也跟前面是一样的：
工作中发现需求 -> 调研优秀的开源项目 -> 二次开发 -> 团队内分享。

>这个阶段，可以刻意地减少对 Google/Baidu 的依赖，尝试自己编码解决问题、实现需求。

在需要的时候，我也会使用 Go 的高级语法 channel、interface 等，结合面向对象编程的思想去开发项目或者改造开源项目。


**案例一**：开发了 HTTP 文件服务器
- 需求来源：因为经常需要将同一个二进制文件部署到不同的机器上，为了便于分发文件，我开发了一个[HTTP 服务器](https://github.com/alex8866/grapehttp)；
- 调研项目：使用了开源的[gohttpserver](https://github.com/codeskyblue/gohttpserver)；
- 效果：开发完成后，在团队中推广，有不少同事使用，显著提高了文件的分发效率。


**案例二**：命令行模板
- 需求来源：因为经常需要编写一些命令行工具，所以我每次都要重复开发一些命令行工具的基础功能，例如命令行参数解析、子命令等。为了避免重复开发这些基础功能，提高工具开发效率和易用度，我开发了一个命令行框架，[cmdctl](https://github.com/lexfei/cmdctl)；
- 调研项目：参考了 Kubernetes 的[kubectl命令行工具](https://github.com/kubernetes/kubernetes/tree/master/cmd/kubectl)的实现；
- 效果：在工作中很多需要自动化的工作，都以命令行工具的形式，添加在了 cmdctl 命令框架中，大大提高了我的开发效率。


**研究过的其他有趣的开源项目**
- [elvish](https://github.com/elves/elvish)：Go 语言编写的 Linux Shell；
- [machinery](https://github.com/RichardKnop/machinery)：Go 语言编写的分布式异步作业系统；
- [gopub](https://github.com/lisijie/gopub)：Go 语言编写的的版本发布系统；
- [crawlab](https://github.com/crawlab-team/crawlab)：Go 语言编写的分布式爬虫管理平台；


### 架构师阶段


### 资深工程师
架构师有很多方向，在云原生技术时代，转型为**云原生架构师**对学 Go 语言的我们来说是一个不错的选择。

要成为云原生架构师，首先要学习云原生技术。云原生技术有很多，我推荐的学习路线如下图所示：
![云原生架构师技术路线图](../../Computer-StudyNotes/img/《Go语言项目开发实战》/cloudTec.jpg)

>如果你还有精力，还可以再学习下 TKEStask、Consul、Cilium、OpenShift 这些项目。

[参考资料](https://github.com/marmotedu/awesome-books)：

- 微服务：《微服务设计》  [英] Sam Newman
- Docker：《Docker 技术入门与实战》（第 3 版）杨保华 / 戴王剑 / 曹亚仑、《Docker ——容器与容器云》（第 2 版）浙江大学SEL实验室
- Kubernetes ： 《Kubernetes 权威指南：从 Docker 到 Kubernetes 实践全接触》（第 4 版）龚正 / 吴治辉 / 崔秀龙 / 闫健勇、《基于 Kubernetes 的容器云平台实战》
- Knative：[Knative Documentation](https://knative.dev/docs/)
- Prometheus：[Prometheus Documentation](https://prometheus.io/docs/introduction/overview/)
- Jaeger ：Jaeger Documentation
- KVM：《KVM 虚拟化技术 : 实战与原理解析》
- Istio：《云原生服务网格 Istio：原理、实践、架构与源码解析》张超盟，章鑫，徐中虎，徐飞
- Kafka：《Apache Kafka 实战》胡夕、《Apache Kafka 源码剖析》徐郡明
- Etcd：etcd 实战课
- Tyk：Tyk Open Source
- TKEStask：TKEStack Documentation
- Consul：Consul Documentation
- Cilium：Cilium Documentation
- OpenShift ：《开源容器云 OpenShift：构建基于 Kubernetes 的企业应用云平台》



### 专家工程师

增强自己架构能力，而不是深入具体细节：

- 调研竞品，了解竞品的架构设计和实现方式。
- 参加技术峰会，学习其他企业的优秀架构设计，例如 ArchSummit 全球架构师峰会、QCon 等。
- 参加公司内外组织的技术分享，了解最前沿的技术、架构和解决方案。
- 关注一些优秀的技术公众号，学习其中高质量的技术文章。
- 作为一名创造者，通过积极思考，设计出符合当前业务的优秀架构。


在架构师阶段你仍然是一名技术开发者，一定不能脱离代码。你可以通过下面这几个方法，让自己保持 Code 能力：- 以 Coder 的身份参与一些核心代码的研发。
以 Reviewer 的身份 Review 成员提交的 PR。
工作之余，阅读项目其他成员开发的源码。
关注一些优秀的开源项目，调研、部署并试用。


- 研发层面和架构层面 走得更远。

- 能够兼具一个 Creator 的角色，能够从 0 到 1，构建满足业务需求的优秀软件系统，甚至能够独立开发一款备受欢迎的开源项目。

## 进阶之路的心得分享

1. 第一点：尽快打怪升级。
2. 第二点：找对方法很重要。
   1. 工作中发现需求 -> 调研优秀的开源项目 -> 二次开发 -> 团队内分享。
   2. 以工作需求为驱动，一方面可以让你有较强的学习动力、学习目标，
   3. 另一方面可以使你在学习的过程中，也能在工作中有所产出，工作产出和学习两不误。
3. 第三点：学架构，先学习当前业务的架构，再学习云原生架构。


毕业 3~5 年的程序员可能是性价比最高的，要时间有时间，要经验有经验，并且当前所积累的研发技能，已经能或者通过后期的学习能够满足公司业务开发需求了。


如何判断一个程序员的性价比呢？就是你的能力要跑赢你当前的年龄和薪资。想跑赢当前的年龄和薪资，需要你尽快地打怪练级，提升自己。