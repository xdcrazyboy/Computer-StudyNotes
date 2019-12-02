# 《Hive入门指南》

[TOC]

参考文献：https://cwiki.apache.org/confluence/display/Hive/Tutorial
# Hive是什么？
Hive是一个基于Apache Hadoop的数据仓库。对于数据存储与处理，Hadoop提供了主要的扩展和容错能力。

Hive设计的初衷是：对于大量的数据，使得数据汇总，查询和分析更加简单。它提供了SQL，允许用户更加简单地进行查询，汇总和数据分析。同时，Hive的SQL给予了用户多种方式来集成自己的功能，然后做定制化的查询，例如用户自定义函数（User Defined Functions，UDFs).

## Hive不适合做什么
Hive不是为在线事务处理而设计。它最适合用于传统的数据仓库任务。

## 数据单元
- 数据库
  >命名空间功能。 为了避免表，视图，分区，列等等的命名冲突；也可以用于加强用户或用户组的安全。
  - 表
    >相同数据库的同类数据单元，跟正常数据库的一样。
  - 分区
    >命名空间功能，为了避免表，视图，分区，列等等的命名冲突。数据库也可以用于加强用户或用户组的安全。
  - 桶
    >每个分区的数据，基于表的一些列的哈希函数值，又被分割成桶

## 数据类型
### 原始类型
- 整型 Integers
  - TINYINT － 1位的整型
  - SMALLINT － 2位的整型
  - INT － 4位的整型
  - BIGINT － 8位的整型
* 布尔类型
  * BOOLEAN － TRUE/FALSE
* 浮点数
  * FLOAT － 单精度
  * DOUBLE － 双精度
* 定点数
  * -DECIMAL － 用户可以指定范围和小数点位数
* 字符串
  * -STRING － 在特定的字符集中的一个字符串序列
  * -VARCHAR － 在特定的字符集中的一个有最大长度限制的字符串序列
  * -CHAR － 在特定的字符集中的一个指定长度的字符串序列
* 日期和时间
  * -TIMESTAMP － 一个特定的时间点，精确到纳秒。
  * -DATE － 一个日期
* 二进制
  * -BINARY － 一个二进制位序列  

### 复杂类型
复杂类型可以由原始类型和其他组合类型构建：
- 结构体类型（Stuct): 使用点（.)来访问类型内部的元素。例如，有一列`c`，它是一个结构体类型`{a INT; b INT}`，字段a可以使用表达式`c.a`来访问。
- Map(key-value键值对)：使用`['元素名']`来访问元素。例如，有一个MapM，包含`'group'->gid`的映射，则gid的值可以使用`M['group']`来访问。
- 数组：数组中的元素是相同的类型。可以使用``[n]``来访问数组元素，`n`是数组下标，以0开始。例如有一个数组A，有元素`['a','b','c']`，则`A[1]`返回`'b'`。

## 内置运算符合函数
>Hive所有关键词的**大小写都不敏感**，包括Hive运算符和函数的名字

# 数据操作使用
http://svn.apache.org/viewvc/hive/trunk/ql/src/test/queries/clientpositive/
## 1. 管理表
1. 创建表
   ```sql
    CREATE TABLE country_list (name STRING);
   ```
2. 查看表信息
   ```sql
    DESCRIBE EXTENDED country_list;
   ```
3. 将数据加载到表中
   我们希望将数据加载到表中，因此需要将数据上传到 HDFS。
   ```sql
    
   ```

4. 查询表
   ```sql
    SELECT * FROM country_list;
   ```
5. 删除表
   ```sql
    DROP TABLE country_list;
   ```
   表和 HDFS 目录及文件将被删除。
6. 创建外部表
   ```sql
    CREATE EXTERNAL TABLE country_list (name STRING);
   ```
    所有的东西看起来都和前面的例子一样，也就是说，这个表在 Beeswax 界面中看起来是一样的，并且它的行为也和前面的一个一样。 只有详细的描述显示了不同:
   `DESCRIBE EXTENDED country_list;`
   表格描述现在提到了 tableType: EXTERNAL table，以前是 tableType: MANAGED table。 托管表意味着 Hive 在删除表时将删除数据——对表进行管理。 外部表类型意味着 Hive 在删除一个表时只会删除模式信息——它从 Hive 中消失，但 HDFS 上的数据仍然保留—— Hive 认为它是外部的。

7. 浏览表和分区
   `SHOW TABLES 'page.*';` 这样将会列出以page开头的表，模式遵循Java正则表达式语法。

## 2. 查询
1. 最简单的：
   ```sql
    SELECT * FROM wdi;
   ```
2. SELECT ... **WHERE** ...：
   下面的查询返回名为“ Trade (% of GDP)”的指标的所有行:
   ```sql
    SELECT * FROM wdi
    WHERE indicator_name = 'Trade (% of GDP)';

    <!-- 可以进一步将结果限制为只返回国家名称和2011年的指标结果: -->
    SELECT `country_name`, `2011` AS trade_2011 FROM wdi
    WHERE indicator_name = 'Trade (% of GDP)';

    <!-- 还可以将结果限制在2011年的数据上: -->
    SELECT `country_name`, `2011` AS trade_2011 FROM wdi WHERE
    indicator_name = 'Trade (% of GDP)' AND
    `2011` IS NOT NULL;

   ```
3. SELECT ... ORDER BY ...:
   ```sql
    SELECT `country_name`, `2011` AS trade_2011 FROM wdi WHERE
    indicator_name = 'Trade (% of GDP)' AND
    `2011` IS NOT NULL
    ORDER BY trade_2011 DESC;
   ```
   使用 `order by` 进行全局排序的缺点是，它使用单个减速器实现的，对大量数据进行排序可能需要很长的时间。这种情况下，您只需要接近顺序并调查可以使用 `sort by `语句的数据。 它通过 `reducer` 对数据进行排序，而不是全局排序，这对于大型数据集来说要快得多。
4. SELECT ... SORT BY ...：
   ```sql
    SELECT `country_name`, `2011` AS trade_2011 FROM wdi WHERE
    indicator_name = 'Trade (% of GDP)' AND
    `2011` IS NOT NULL
    SORT BY trade_2011 DESC;
   ```
5. SELECT ... DISTRIBUTE BY ...：
   Distributed BY 告诉 Hive 当数据发送到还原器时，应该通过哪个列来组织数据。 我们可以不使用前面示例中的 CLUSTER BY，而是使用 DISTRIBUTE BY 来确保每个reducer获得每个indicator的所有数据。
   ```sql
    SELECT country_name, indicator_name, `2011` AS trade_2011 FROM wdi WHERE
    (indicator_name = 'Trade (% of GDP)' OR
    indicator_name = 'Broad money (% of GDP)') AND
    `2011` IS NOT NULL
    DISTRIBUTE BY indicator_name;
   ```
6. SELECT ... SORT BY ...：
   ```sql
    SELECT `country_name`, `2011` AS trade_2011 FROM wdi WHERE
  indicator_name = 'Trade (% of GDP)' AND
  `2011` IS NOT NULL
  SORT BY trade_2011 DESC;
   ```