
# 参考资料

- 官网：http://kylin.apache.org/cn/
  - 官方文档：http://kylin.apache.org/cn/docs/
  - 下载地址：
- github：
  - 官方：
  - [大佬积累](hhh)
- 国内网址：
  - [某网址](hhh)
- 国外网址
  - [某网站](hhh)
- 博客
  - [博可以](jh)
- 问题解答汇总
  1. [为什么？](hhh)

# 简介

## 概况
- Apache Kylin™是一个开源的、**分布式的分析型数据仓库**（最初由 eBay 开发并贡献至开源社区）
  - 提供**Hadoop/Spark 之上**的 **SQL 查询接口**
  - 及**多维分析**（OLAP）能力以支持超大规模数据，
  - 它能在**亚秒内查询巨大**的表。
- Apache Kylin™ 令使用者仅需三步，即可实现超大数据集上的亚秒级查询。
  1. 定义数据集上的一个星形或雪花形模型
  2. 在定义的数据表上构建cube
  3. 使用标准 SQL 通过 ODBC、JDBC 或 RESTFUL API 进行查询，仅需亚秒级响应时间即可获得查询结果

## 特征 & 优势

- 可扩展超快的基于大数据的分析型数据仓库
  >Kylin 是为减少在 Hadoop/Spark 上百亿规模数据查询延迟而设计
- Hadoop ANSI SQL 接口
  >作为一个分析型数据仓库(也是 OLAP 引擎)，Kylin 为 Hadoop 提供标准 SQL 支持大部分查询功能
- 交互式查询能力:
  >通过 Kylin，用户可以与 Hadoop 数据进行亚秒级交互，在同样的数据集上提供比 Hive 更好的性能
- 多维立方体（MOLAP Cube）:
  >用户能够在 Kylin 里为百亿以上数据集定义数据模型并构建立方体
- 实时 OLAP：
  >Kylin 可以在数据产生时进行实时处理，用户可以在秒级延迟下进行实时数据的多维分析。
- 与BI工具无缝整合:
  >Kylin 提供与 BI 工具的整合能力，如Tableau，PowerBI/Excel，MSTR，QlikSense，Hue 和 SuperSet


# 安装实践

## 安装

- 硬件要求
- 系统要求
- 环境准备
  - Hadoop：Kylin 依赖于 Hadoop 集群处理大量的数据集。您需要准备一个配置好 HDFS，YARN，MapReduce，Hive， HBase，Zookeeper 和其他服务的 Hadoop 集群供 Kylin 运行。