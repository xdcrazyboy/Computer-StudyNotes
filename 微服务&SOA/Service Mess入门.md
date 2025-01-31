# 什么是Service Mesh（服务网格）
又译作 “服务网格”，作为服务间通信的基础设施层。

服务网格（Service Mesh）是致力于解决**服务间通讯**的基础设施层。它负责在现代**云原生应用程序**的**复杂服务拓扑**中进行**可靠地传递请求**。实际上，Service Mesh 通常是通过一组**轻量级网络代理**（Sidecar proxy），与应用程序代码部署在一起来实现，而无需感知应用程序本身（对应用程序透明）。

## Service Mesh 的特点
* 应用程序间通讯的中间层
  >屏蔽分布式系统通信的复杂性(负载均衡、服务发现、认证授权、监控追踪、流量控制等等)，服务只用关注业务逻辑；
* 轻量级网络代理
* 应用程序无感知，对应用透明，Service Mesh组件可以单独升级。
* 解耦应用程序的重试/超时、监控、追踪和服务发现
* 真正的语言无关，服务可以用任何语言编写，只需和Service Mesh通信即可；

如果用一句话来解释什么是 Service Mesh：可以将它比作是**应用程序或者说微服务间的 TCP/IP**，**负责服务之间的网络调用、限流、熔断和监控**。
>对于编写应用程序来说一般无须关心 TCP/IP 这一层（比如通过 HTTP 协议的 RESTful 应用），同样使用 Service Mesh 也就无须关心服务之间的那些原来是通过应用程序或者其他框架实现的事情，比如 Spring Cloud、OSS，现在只要交给 Service Mesh 就可以了。

## Service Mesh如何工作？
下面以 **Linkerd** 为例讲解 Service Mesh 如何工作：

1. Linkerd **将服务请求路由到目的地址**，根据中的**参数判断**是到**生产**环境、**测试**环境还是 **staging** 环境中的服务（服务**可能同时部署在这三个环境**中），是路由到**本地**环境还是**公有云**环境？所有的这些路由信息可以**动态配置**，可以是**全局**配置也可以为某些服务**单独**配置。
2. 当 Linkerd **确认了目的地址后**，**将流量发送到**相应**服务发现**端点，在 kubernetes 中是 service，然后**service**会将服务转发给后端的**实例**。
3. Linkerd 根据它观测到**最近请求的延迟时间**，**选择出**所有应用程序的实例中**响应最快**的实例。
4. Linkerd 将请求发送给该实例，同时**记录响应类型**和**延迟数据**。
5. 如果该实例挂了、不响应了或者进程不工作了，Linkerd 将把请求发送到**其他实例上重试**。
6. 如果该实例**持续返回 error**，Linkerd 会将该实例从**负载均衡池**中**移除**，稍后再**周期性得重试**。
7. 如果**请求的截止时间**已过，Linkerd **主动失败**该请求，而不是再次尝试添加负载。
8. Linkerd 以 metric 和分布式追踪的形式捕获上述行为的各个方面，这些追踪信息将发送到集中 metric 系统。

## 为何使用 Service Mesh？
- 没有带来新功能，只是在Cloud Native的k8s环境下实现了以前的功能：网络调用、限流、熔断和监控。
- 传统的 MVC 三层 Web 应用程序架构下，服务之间的通讯并不复杂，在应用程序内部自己管理即可，但是在现今的复杂的大型网站情况下，单体应用被分解为众多的微服务，服务之间的依赖和通讯十分复杂，出现了 twitter 开发的 Finagle、Netflix 开发的 Hystrix 和 Google 的 Stubby 这样的 “胖客户端” 库，这些就是早期的 Service Mesh，但是它们都近适用于特定的环境和特定的开发语言，并不能作为平台级的 Service Mesh 支持。

在 Cloud Native 架构下，容器的使用给予了异构应用程序的更多可行性，kubernetes 增强的应用的横向扩容能力，用户可以快速的编排出复杂环境、复杂依赖关系的应用程序，同时开发者又无须过分关心应用程序的监控、扩展性、服务发现和分布式追踪这些繁琐的事情而专注于程序开发，赋予开发者更多的创造性。

**微服务**： (Microservices) 是一种软件架构风格，它是以专注于单一责任与功能的小型功能区块 (Small Building Blocks) 为基础，利用模块化的方式组合出复杂的大型应用程序，各功能区块使用与语言无关 (Language-Independent/Language agnostic) 的 API 集相互通信。

## 第二代微服务出现了问题，由此产生了第一代ServiceMesh
第二代微服务模式看似完美，但开发人员很快又发现，它也存在一些本质问题：

- 其一，虽然框架本身屏蔽了分布式系统通信的一些通用功能实现细节，但开发者却**要花更多精力去**掌握和管理**复杂的框架本身**，在实际应用中，去**追踪和解决框架出现的问题也绝非易事**；

- 其二，开发框架通常**只支持一种或几种特定的语言**，回过头来看文章最开始对微服务的定义，一个重要的特性就是**语言无关**，但那些没有框架支持的语言编写的服务，很难融入面向微服务的架构体系，想因地制宜的用多种语言实现架构体系中的不同模块也很难做到；
- 其三，框架**以lib库的形式和服务联编**，复杂项目依赖时的库版本兼容问题非常棘手，同时，**框架库的升级也无法对服务透明**，服务会**因为和业务无关的lib库升级**而**被迫升级**；

因此，以Linkerd，Envoy，Ngixmesh为代表的代理模式（边车模式）应运而生，这就是第一代Service Mesh，它将分布式服务的通信抽象为单独一层，在这一层中实现负载均衡、服务发现、认证授权、监控追踪、流量控制等分布式系统所需要的功能，作为一个和服务对等的代理服务，和服务部署在一起，接管服务的流量，通过代理之间的通信间接完成服务之间的通信请求，这样上边所说的三个问题也迎刃而解。

**第一代**Service Mesh由一系列独立运行的单机代理服务构成，为了提供统一的上层运维入口，演化出了**集中式的控制面板**，所有的单机代理组件通过**和控制面板交互**进行网络拓扑**策略的更新**和**单机数据的汇报**。这就是以Istio为代表的**第二代**Service Mesh。

## 面临的挑战
* Service Mesh组件以代理模式计算并转发请求，一定程度上会降低通信系统性能，并增加系统资源开销；
* Service Mesh组件接管了网络流量，因此服务的整体稳定性依赖于Service Mesh，同时额外引入的大量Service Mesh服务实例的运维和管理也是一个挑战

- 为了解决**端到端的字节码通信**问题，**TCP**协议诞生，让**多机通信**变得**简单可靠**；
- 微服务时代，Service Mesh应运而生，**屏蔽了分布式系统的诸多复杂性**，让开发者可以**回归业务**，聚焦真正的价值。

# Service Mesh 培训会 12-05

## 落地情况
千万
0.2ms

## 简介
- 数据面
- 控制面

## 解决了什么问题
1. 业务和框架偶合严重 27%
2. 机器资源逐年增加 27%
3. 框架不断升级
4. 流量调度的诉求 27%
5. 客户端中间件版本的统一

Mesh 解决了业务开发和基础团队之间的偶合关系。 基础开发。

## 我们落地了哪些能力
- RPC
- 能力解耦，之前注册发现、限流、序列化——>MOSN
- 对接不同的注册中心？


资源、流量调度。
空闲资源——>保活态？运行态？  在MOSN层面做了一些转移？ 用Go写的，占用资源少。

- 无线网关

可行的：做适配，用主流开源或许更靠谱，他们这个MOSN不一定适合我们
