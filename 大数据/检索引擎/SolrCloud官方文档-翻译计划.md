# SolrCloud

# 开始使用

SolrCloud 旨在提供一个高可用性、容错性的环境，用于跨多个服务器分发索引内容和查询请求。

在这个系统中，数据被组织成多个片段，或者分片，可以托管在多台机器上，副本为可伸缩性和容错性提供了冗余，而 ZooKeeper 服务器帮助管理整个结构，这样索引和搜索请求都可以正确路由。

本节将详细解释 SolrCloud 及其内部工作原理，但是在深入讨论之前，最好先了解一下您想要实现的目标。

本页面提供了一个简单的教程，以便在 SolrCloud 模式下启动 Solr，这样您就可以开始了解在索引和服务查询期间分片如何相互交互。为此，我们将使用在单台机器上配置 SolrCloud 的简单示例，这显然不是真正的生产环境，其中将包括多个服务器或虚拟机。在实际的生产环境中，您还将使用实际的机器名称，而不是我们在这里使用的“ localhost”。

在本节中，您将学习如何使用启动脚本和特定的配置集启动 SolrCloud 集群。

>注意： 本教程假设您已经熟悉了使用 Solr 的基本知识。如果你需要复习，请参阅[入门部分](https://lucene.apache.org/solr/guide/8_5/getting-started.html#getting-started)，了解 Solr 概念的基础知识。如果作为这个练习的一部分加载文档，那么应该为这些 SolrCloud 教程重新安装 Solr。

## SolrCloud示例

### 交互式启动
Bin/Solr 脚本使您可以很容易地开始使用 SolrCloud，因为它将引导您完成以 SolrCloud 模式启动 Solr 节点并添加集合的过程。首先，只需要输入:
```bash
bin/solr -e cloud
```
这将启动一个交互式会话，带您完成使用嵌入的 ZooKeeper 设置简单的 SolrCloud 集群的步骤。

该脚本首先询问您希望在本地集群中运行多少个 Solr 节点，默认值为2。
```
Welcome to the SolrCloud example!

This interactive session will help you launch a SolrCloud cluster on your local workstation.
To begin, how many Solr nodes would you like to run in your local cluster? (specify 1-4 nodes) [2]
```
该脚本支持最多启动4个节点，但是我们建议在启动时使用默认的2。这些节点将各自存在于一台计算机上，但将使用不同的端口在不同的服务器上模拟操作。

接下来，脚本将提示您将每个 Solr 节点绑定到的端口，例如:
```
Please enter the port for node1 [8983]
```
为每个节点选择任何可用的端口; 第一个节点的默认端口为8983，第二个节点为7574。该脚本将按顺序启动每个节点，并显示用于启动服务器的命令，例如:
```bash
solr start -cloud -s example/cloud/node1/solr -p 8983
```
第一个节点还将启动一个绑定到端口9983的嵌入式 ZooKeeper 服务器。第一个节点的 Solr home位于/cloud/node1/Solr 中，如-s 选项所示。

启动集群中的所有节点后，脚本提示您输入要创建的集合的名称:
```
 Please provide a name for your new collection: [gettingstarted]
```
建议的默认值是“ gettingstarted” ，但是你可能需要为你的特定搜索应用程序选择一个更合适的名称。

接下来，该脚本提示您输入分发集合的分片数量。稍后将更详细地介绍 Sharding，因此如果您不确定，我们建议使用默认值2，这样您就可以看到集合是如何在 SolrCloud 集群中的多个节点之间分布的。

接下来，该脚本将提示您为每个切分创建的副本数量。本指南稍后将更详细地介绍副本，因此如果您不确定，请使用缺省值2，以便查看如何在 SolrCloud 中处理副本。

最后，该脚本将提示您输入集合的配置目录的名称。您可以选择 _ default，或 sample _ techproducts _ confgs。配置目录是从 server/solr/configsets/中提取出来的，如果您愿意，可以事先查看它们。当您仍在为文档设计模式并且在使用 Solr 时需要一定的灵活性时，_ default 配置非常有用，因为它具有无模式功能。但是，在创建您的集合之后，可以禁用无模式功能来锁定模式(以便在这样做之后索引的文档不会改变模式)或自己配置模式。这可以按照以下方式完成(假设您的集合名为 mycollection) :
```bash
# V1 API:
curl http://host:8983/solr/mycollection/config -d '{"set-user-property": {"update.autoCreateFields":"false"}}'
# V2 API SolrCloud 
curl http://host:8983/api/collections/mycollection/config -d '{"set-user-property": {"update.autoCreateFields":"false"}}'
```
此时，您应该在本地 SolrCloud 集群中创建了一个新集合。要验证这一点，您可以运行 status 命令:
```bash
bin/solr status
```
如果在此过程中遇到任何错误，请检查示例/cloud/node1/logs 和示例/cloud/node2/logs 中的 Solr 日志文件。

通过访问 Solr Admin UI: http://localhost:8983/Solr/#/~cloud 集群中的云面板，您可以看到您的集合是如何部署到集群中的。Solr 还提供了一种使用 healthcheck 命令对集合执行基本诊断的方法:
```bash
bin/solr healthcheck -c gettingstarted
```
Healthcheck 命令收集集合中每个副本的基本信息，例如文档的数量、当前状态(active、 down 等)和地址(副本存在于集群中的位置)。

现在可以使用 [Post 工具](https://lucene.apache.org/solr/guide/8_5/post-tool.html#post-tool)将文档添加到 SolrCloud 中。

要在 SolrCloud 模式下停止 Solr，可以使用 bin/Solr 脚本并发出 stop 命令，如下所示:
```bash
bin/solr stop -all
```
### 从-noprompt开始
你也可以使用以下命令使用默认值而不是交互式会话来启动 SolrCloud:
```bash
bin/solr -e cloud -noprompt
```


### 重启节点
可以使用 bin/solr 脚本重新启动 SolrCloud 节点。例如，要重新启动端口8983上运行的 node1(使用嵌入式 ZooKeeper 服务器) ，您可以这样做:
```bash
bin/solr restart -c -p 8983 -s example/cloud/node1/solr
```
要在端口7574上重新启动运行的 node2，可以这样做:
```bash
bin/solr restart -c -p 7574 -z localhost:9983 -s example/cloud/node2/solr
```
>注意，在启动 node2时需要指定 ZooKeeper 地址(- z localhost: 9983) ，以便它可以使用 node1加入集群。


### 向集群中添加节点
向现有集群添加节点有点高级，并且涉及到对 Solr 的更多理解。一旦你使用启动脚本启动了一个 SolrCloud 集群，你可以通过以下方式添加一个新的节点:
```bash
mkdir <solr.home for new Solr node>
cp <existing solr.xml path> <new solr.home>
bin/solr start -cloud -s solr.home/solr -p <port num> -z <zk hosts string>
```
>注意，上面要求您创建 Solr 主目录。您要么需要将 solr.xml 复制到 solr_home 目录，要么集中保存在 ZooKeeper/solr.xml 中。

示例(具有目录结构) ，将一个节点添加到以“ bin/solr-e cloud”开始的示例中:
```bash
mkdir -p example/cloud/node3/solr
cp server/solr/solr.xml example/cloud/node3/solr
bin/solr start -cloud -s example/cloud/node3/solr -p 8987 -z localhost:9983
```
上一个命令将启动端口8987上的另一个 Solr 节点，Solr home 设置为 example/cloud/node3/Solr。新节点将把它的日志文件写到 example/cloud/node3/logs。

一旦您熟悉了 SolrCloud 示例的工作方式，我们建议在生成环境中启动solrCloud节点时使用[Taking Solr to Production](https://lucene.apache.org/solr/guide/8_5/taking-solr-to-production.html#taking-solr-to-production)上描述的方法流程。

---

# SolrCloud是如何运行的
- 关键概念
  - 逻辑概念
  - 物理概念

下面的章节将提供有关各种 SolrCloud 特性如何工作的一般信息。为了理解这些特性，首先理解一些与 SolrCloud 相关的关键概念是很重要的。
- 在SolrCloud中的分片和索引数据
- 分布式请求
- 标准和路由别名

>如果您已经熟悉 SolrCloud 的概念和基本功能，可以跳到介绍 SolrCloud 配置和参数的部分。

## 主要的SolrCloud概念
SolrCloud 集群由一些“逻辑”概念组成，这些概念分层在一些“物理”概念之上。

### 逻辑概念
- 一个集群可以承载多个 Solr 文档集合
- 一个集合可以被划分为多个Shards，其中包含集合中的一个文档子集
- 集合确定的分片数:
  - 集合可包含的文件合理数量的理论限制
  - 单个搜索请求可支持的并行处理量

### 物理概念
- 集群由一个或多个 Solr 节点组成，它们运行 Solr 服务器进程的实例
- 每个节点可以承载多个核心
- 集群中的每个核心都是逻辑分片的物理副本
- 每个副本都使用相同配置，配置由其所在集合决定。Every Replica uses the same configuration specified for the Collection that it is a part of.
- 每个 Shard 拥有的副本数量决定了:
  - 集合中内置的冗余级别以及在某些节点不可用时集群的容错能力
  - 在重负载下可处理的并发搜索请求数量的理论限制

>确保您的集群中的 DNS解析是稳定的，也就是说，对于属于集群的每个存活主机，主机名总是对应于相同的特定 IP 和物理节点。例如，在 AWS 上部署的集群中，这将需要设置`preserve_hostname: true` 在 `/etc/cloud/cloud.cfg`. 更改活动节点的 DNS 解析可能导致意外错误。详情请参阅 SOLR-13159 


## SolrCloud中的分片和索引数据

[小目录]

当您的集合对于一个节点来说太大时，您可以通过创建多个分片将其分解并分段存储。

分片是集合的分区，包含集合中文档的子集，这样集合中的每个文档都恰好包含在一个分片中。集合中包含每个文档的切分取决于该集合的总体“切分”策略。

例如，您可能有一个集合，其中每个文档的“ country”字段确定它属于哪个分片，因此来自同一个国家的文档位于同一个位置。不同的集合可以简单地在每个文档的 uniqueKey 上使用“散列”来确定其分片。

在 SolrCloud 之前，Solr 支持分布式搜索，允许跨多个分片执行一个查询，因此查询是针对整个 Solr 索引执行的，搜索结果中不会遗漏任何文档。因此，跨分片分割索引并不是 SolrCloud 独有的概念。然而，分布式方法存在一些问题，使得有必要用SolrCloud来进行改进:

- 将索引分割成分片多少有点手工操作的味道
- 没有对分布式索引的支持，这意味着您需要显式地将文档发送到特定的分片; Solr 自己无法确定将文档发送到哪些分片
- 没有负载平衡或故障转移，因此，如果收到大量查询，就需要确定将它们发送到哪里，如果一个分片死亡，它就消失了

SolrCloud解决了这些限制。它支持自动分发索引进程和查询，而 ZooKeeper 提供故障转移和负载平衡。此外，每个分片都可以有多个副本，以增强健壮性。

### Leaders和副本

在solrCloud没有主从。相反，每个分片至少由一个物理副本组成，其中恰好有一个是leader。leader是自动选举产生的，起初是先到先得的基础上，然后基于在http://ZooKeeper.apache.org/doc/r3.5.5/recipes.html#sc_leaderelection 文档中描述的 ZooKeeper 流程。

如果一个leader挂掉，其他副本中的一个会自动当选为新的leader。

当文档发送到 Solr 节点进行索引时，系统首先确定该文档属于哪个 Shard，然后确定当前哪个节点承载该 Shard 的leader。然后文档被转发给当前的leader进行索引，leader将更新转发给所有其他副本。

#### 副本的类型

默认情况下，如果他们的领导下台，所有的副本都有资格成为leader。然而，这是有代价的: 如果所有的副本在任何时候都可以成为leader，那么每个副本必须在任何时候都与它的leader保持同步。添加到leader的新文档必须路由到副本，每个副本必须执行提交。如果一个副本挂起，或者暂时不可用，然后重新加入集群，如果错过了大量更新，恢复可能会很慢。

这些问题对于大多数用户来说都不是问题。然而，一些用例如果副本表现得更像前一个模型（ behaved a bit more like the former model），或者不要求实时同步，或者根本没有资格成为leader（by not being eligible to become leaders at all），那么它们的性能会更好。

Solr 通过允许您在创建新集合或添加副本时设置副本类型来实现这一点。可供选择的种类如下:
- **NRT**：默认类型。NRT 副本(NRT = NearRealTime)维护事务日志并在本地将新文档写入其索引。这种类型的任何副本都有资格成为leader。传统上，这是 Solr 支持的唯一类型
  
- **TLOG**：这种类型的副本维护事务日志，但不在本地索引文档更改。这种类型有助于加快索引，因为副本中不需要进行提交。当这种类型的副本需要更新其索引时，它通过复制leader的索引来更新索引。这种类型的副本也有资格成为分片leader; 它将首先处理其事务日志。如果它成为一个leader，它的行为就像它是一个 NRT 类型的副本一样
  
- **PULL**：这种类型的副本不在本地维护事务日志或索引文档更改。它只复制来自分片leader的索引。它不具备成为分片leader的资格，也根本不参与分片leader的选举。

>如果在创建副本时没有指定它的类型，那么它将是 NRT 类型。

#### 在集群中混合副本类型

建议使用三种副本混合类型：
- 全为NRT副本
  >对于中小型集群，甚至是更新(索引)吞吐量不太高的大型集群，可以使用此方法。NRT 是唯一支持软提交的副本类型，所以在需要 NearRealTime 时也可以使用这个组合。
- 全为TLOG副本
  >如果不需要 NearRealTime 并且每个分片的副本数量很高，但是您仍然希望所有副本都能够处理更新请求，那么可以使用这个组合。
- 带PULL副本的TLOG副本
  >如果不需要 NearRealTime，每个分片的副本数量很高，并且您希望在文档更新之上增加搜索查询的可用性，即使这意味着临时提供过时的结果。

不推荐副本类型的其他组合。如果分片中的多个副本正在写入自己的索引，而不是从 NRT 副本复制，那么领导选举会导致分片的所有副本与leader不同步，所有副本都必须复制完整的索引

#### 使用 PULL 副本进行恢复
如果 PULL 副本关闭或离开集群，则需要考虑以下几种情况。

如果 PULL 副本由于Leader已关闭而无法与Leader同步，则不会发生复制。不过，它将继续提供查询服务。一旦它可以再次连接到Leader，复制就会恢复。

如果 PULL 副本不能连接到 ZooKeeper，它将被从集群中删除，并且查询不会从集群中路由到它。

如果 PULL 副本死亡或者由于其他原因无法访问，那么它将不可查询。当它重新加入集群时，它将从Leader进行复制，当复制完成时，它将准备好再次服务查询。

#### 设置首选副本类型的查询
默认情况下，所有副本都服务于查询。有关如何为查询指示首选副本类型的详细信息，请参阅[分片指南·参数](https://lucene.apache.org/solr/guide/8_5/distributed-requests.html#shards-preference-parameter) 一节。

### 文档路由
Solr 通过在[创建集合](https://lucene.apache.org/solr/guide/8_5/collection-management.html#create)时指定 router.name 参数，提供了指定集合使用的路由器实现的能力。

如果您使用 compositeId 路由器(默认值) ，则可以发送文档 ID 中带有前缀的文档，该前缀将用于计算 Solr 使用的散列，以确定文档发送到哪个分片去做索引。前缀可以是自定义任何名称(例如，它不必是分片名) ，但它必须是一致的，这样 Solr 的行为才能一致。

**例如**：
如果希望为客户共同定位文档，可以使用客户名称或 ID 作为前缀。例如，如果您的客户是“ IBM” ，并且有一个 ID 为“12345”的文档，那么您将在文档 ID 字段中插入前缀: “ IBM！“12345”。叹号在这里是至关重要的，因为它区分了用于确定将文档定向到哪个分片的前缀。

然后在查询时，使用 _route _ parameter (即 q = solr & _ route _ = IBM!)将前缀(es)包含到查询中直接查询特定的分片。在某些情况下，这可能会提高查询性能，因为在查询所有分片时，会存在网络延迟。

compositeId 路由器支持包含最多2个路由级别的前缀。例如: 首先按区域，然后按客户: “ usa! ibm！12345”

**另一个用例**：
如果客户“IBM”拥有大量文档，并且您希望将其分散到多个分片上。这种用例的语法是: shard _ key/num! document _ id，其中/num 是在复合散列中使用的 shard 键的位数。

所以 IBM/3！12345将从分片密钥中提取3个位，从唯一的文件id中提取29个位，将租户分散到集合中1/8的分片上。同样，如果num值为2，那么它将把文档分布在分片数的1/4上。在查询时，使用 `_route _` 参数 (即 `q = solr & _ route_ = IBM/3!`)在查询中包含前缀(es)和位数直接查询特定的分片。

如果不希望影响文档的存储方式，就不需要在文档ID中指定前缀。

如果在创建时创建了集合并定义了“隐式”路由器，那么还可以定义 router.field 参数来使用每个文档中的字段来标识文档所属的分片。但是，如果文档中缺少指定的字段，则文档将被拒绝。您还可以使用 _ route _ 参数来命名特定的分片。

### 分片再分割
当您在 SolrCloud 中创建一个集合时，您将决定要使用的初始数字分片。但是很难预先知道需要多少分片，特别是当组织需求可能随时改变时，而且后来发现自己选错的代价可能很高，包括创建新的内核和重新获取所有数据。

分割分片的能力在Collection API中。它目前允许将一个分片分割成两部分。现有的分片保持原样，因此分割操作实际上将数据作为新的分片复制两份。您可以在稍后准备好时删除旧的分片。

关于如何使用分片的更多细节见Collection API的[SPLITSHARD命令](https://lucene.apache.org/solr/guide/8_5/shard-management.html#splitshard)一节。

### 在SolrCloud中忽略客户端应用程序的提交
在大多数情况下，在 SolrCloud 模式下运行时，索引客户机应用程序不应该发送显式的提交请求。相反，您应该使用 openSearcher = false 和 auto soft-commit 来配置自动提交，以使最近的更新在搜索请求中可见。这样可以确保自动提交按照集群中的常规调度进行。

为了执行客户端应用程序不应发送显式提交的策略，您需要更新所有将数据索引到 SolrCloud 中的客户端应用程序。但是，这并不总是可行的，因此 Solr 提供了 IgnoreCommitOptimizeUpdateProcessorFactory，它允许您忽略客户机应用程序的显式提交和/或优化请求，而无需重构客户机应用程序代码。

要激活这个请求处理器，你需要在你的 `solrconfig.xml` 文件中添加以下内容:
```xml
<updateRequestProcessorChain name="ignore-commit-from-client" default="true">
  <processor class="solr.IgnoreCommitOptimizeUpdateProcessorFactory">
    <int name="statusCode">200</int>
  </processor>
  <processor class="solr.LogUpdateProcessorFactory" />
  <processor class="solr.DistributedUpdateProcessorFactory" />
  <processor class="solr.RunUpdateProcessorFactory" />
</updateRequestProcessorChain>
```

如上面的示例所示，处理器将向客户机返回200，但将忽略提交/优化请求。
>请注意，您还需要连接到 SolrCloud 所需的隐式处理器，因为这个自定义链取代了默认链。

在下面的示例中，处理器将引发一个带有自定义错误消息的403代码异常:
```xml
<updateRequestProcessorChain name="ignore-commit-from-client" default="true">
  <processor class="solr.IgnoreCommitOptimizeUpdateProcessorFactory">
    <int name="statusCode">403</int>
    <str name="responseMessage">Thou shall not issue a commit!</str>
  </processor>
  <processor class="solr.LogUpdateProcessorFactory" />
  <processor class="solr.DistributedUpdateProcessorFactory" />
  <processor class="solr.RunUpdateProcessorFactory" />
</updateRequestProcessorChain>
```

最后，你也可以配置它忽略优化，让提交通过:
```xml
<updateRequestProcessorChain name="ignore-optimize-only-from-client-403">
  <processor class="solr.IgnoreCommitOptimizeUpdateProcessorFactory">
    <str name="responseMessage">Thou shall not issue an optimize, but commits are OK!</str>
    <bool name="ignoreOptimizeOnly">true</bool>
  </processor>
  <processor class="solr.RunUpdateProcessorFactory" />
</updateRequestProcessorChain>
```


## 分布式请求

[子目录]

当 Solr 节点接收到搜索请求时，请求将在后台被路由到被检索的集合的一个分片副本上。

选择的副本充当聚合器: 它为集合中的每个分片随机选择的副本创建内部请求，协调响应，根据需要发出任何后续内部请求(例如，细化 facets 值，或请求额外的存储字段) ，并为客户机构造最终响应。

### 限制查询哪些分片
选择的副本充当聚合器: 它为集合中的每个分片随机选择的副本创建内部请求，协调响应，根据需要发出任何后续内部请求(例如，细化 facets 值，或请求额外的存储字段) ，并为客户机构造最终响应。

- 一个集合的所有分片查询只是一个没有定义分片参数的查询:
    ```text
    http://localhost:8983/solr/gettingstarted/select?q=*:*
    ```

- 如果只想搜索一个分片，可以使用分片参数通过其逻辑 ID 指定分片，如下所示:
    ```text
    http://localhost:8983/solr/gettingstarted/select?q=*:*&shards=shard1
    ```

- 如果你想搜索一组碎片，你可以在一个请求中用逗号分隔每个碎片:
    ```text
    http://localhost:8983/solr/gettingstarted/select?q=*:*&shards=shard1,shard2
    ```
    >在上述两个示例中，虽然只查询特定的分片，但是任何分片的随机副本都会获得请求。

- 或者，你可以通过用逗号分隔副本 id 来指定一个副本列表来代替分片 id:
    ```text
    http://localhost:8983/solr/gettingstarted/select?q=*:*&shards=localhost:7574/solr/gettingstarted,localhost:8983/solr/gettingstarted
    ```

- 或者，您可以在不同的副本 id 之间使用管道符号(|)为单个分片指定一个副本列表(用于负载平衡目的) :
    ```text
    http://localhost:8983/solr/gettingstarted/select?q=*:*&shards=localhost:7574/solr/gettingstarted|localhost:7500/solr/gettingstarted
    ```

- 最后，您可以指定一个分片列表(用逗号分隔) ，每个分片由一个副本列表(用管道分隔)定义。

在下面的例子中，查询了2个分片，第一个是 shard1中的随机副本，第二个是显式管道分隔列表中的随机副本:
```text
http://localhost:8983/solr/gettingstarted/select?q=*:*&shards=shard1,localhost:7574/solr/gettingstarted|localhost:7500/solr/gettingstarted
```

### 配置 ShardHandlerFactory
为了实现更细粒度的控制，您可以直接配置和调优 Solr 分布式搜索中使用的并发性和线程池方面。**默认配置倾向于提高吞吐量，而不是控制延迟**。