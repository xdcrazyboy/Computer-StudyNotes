# SQL必知必会
[TOC]

# 一、《SQL必知必会》读书笔记

# 二、SQL一些技巧

# 三、SQL优化

# 四、SQL练习

## 4.1 子查询、连接查询
[参考资料1](https://www.cnblogs.com/chiangchou/p/mysql-3.html)

### 子查询(where、from、exists)
1. where型子查询：把内层查询的结果作为外层查询的比较条件。
   >```sql
   >SELECT c.cpcid FROM cpc_0101 as c WHERE c.cpcid = (SELECT MAX(cpcid) FROM cpc_0101)
   >```
2. from型子查询：把内层的查询结果当成临时表，供外层sql再次查询。查询结果集可以当成表看待。临时表要使用一个别名。
3. exists型子查询：把外层sql的结果，拿到内层sql去测试，如果内层的sql成立，则该行取出。内层查询是exists后的查询。
   - 从类别表中取出其类别下有商品的类别(如果该类别下没有商品，则不取出)，[使用where子查询]
        >```sql
        > SELECT c.cat_id,c.cat_name FROM category c WHERE c.cat_id IN (SELECT g.cat_id FROM goods g GROUP BY g.cat_id);
        >```
   - 从类别表中取出其类别下有商品的类别(如果该类别下没有商品，则不取出)，[使用exists子查询]
        >```sql
        > SELECT c.cat_id,c.cat_name FROM category c WHERE EXISTS (SELECT 1 FROM goods g WHERE g.cat_id = c.cat_id);
        >```
     exists子查询，如果exists后的内层查询能查出数据，则表示存在；为空则不存在。
4. any, in 子查询
   - 使用 any 查出类别大于任何一个num值的类别
        >ANY关键词必须后面接一个比较操作符。ANY关键词的意思是“对于在子查询返回的列中的任一数值，如果比较结果为TRUE的话，则返回TRUE”。
        ```sql
        SELECT cat_id,cat_name FROM category WHERE cat_id > ANY (SELECT num FROM nums);
        ```
    - 使用 in 查出cat_id等于num的类别
        ```sql
        SELECT cat_id,cat_name FROM category WHERE cat_id IN (SELECT num FROM nums);
        ```
    - `in` 的效果 跟 `=any` 的效果是一样的, 是对方的别名。
        ```sql
        SELECT cat_id,cat_name FROM category WHERE cat_id IN (SELECT num FROM nums);
        <!-- 等效于下面的 -->
        SELECT cat_id,cat_name FROM category WHERE cat_id = ANY (SELECT num FROM nums);
        ```
    - 使用 all 查询
        >词语ALL必须接在一个**比较操作符的后面**。ALL的意思是“对于子查询返回的列中的所有值，如果比较结果为TRUE，则返回TRUE。”
        ```sql
        SELECT cat_id,cat_name FROM category WHERE cat_id > ALL (SELECT num FROM nums);
        ```
    - `not in` 和 `<> any` 的效果是一样的,`NOT IN`不是`<> ANY`的别名，但是`<> ALL`的别名，也就是说这三个效果都一样。
        ```sql
        SELECT cat_id,cat_name FROM category WHERE cat_id not in (SELECT num FROM nums);
        <!-- 等效于下面的 -->
        SELECT cat_id,cat_name FROM category WHERE cat_id <> ALL (SELECT num FROM nums);
        ```
5. 优化子查询
   - 用子查询替换联合。例如：
        ```sql
        SELECT DISTINCT column1 FROM t1 WHERE t1.column1 IN (SELECT column1 FROM t2);
            <!-- 代替这个： -->
        SELECT DISTINCT t1.column1 FROM t1,t2 WHERE t1.column1 = t2.column1;
        ```
### 连接查询(left join、right join、inner join、union join）

虽然基本不会关注下面两个点，但还是回顾下吧：
1. 笛卡尔积：在数据库中，一张表就是一个集合，每一行就是集合中的一个元素。表之间作联合查询即是作笛卡尔乘积，比如A表有5条数据，B表有8条数据，如果不作条件筛选，那么两表查询就有 5 X 8 = 40 条数据。
    ```sql
    SELECT goods_id,goods_name,cat_name FROM mingoods,category;
    ```
2. 全相乘，虽然带条件了，两表关联了，查出来的不是笛卡尔积式的全部组合，而是想要的数据，但是这个还是全相乘，效率低，全相乘会在内存中生成一个非常大的数据(临时表)，因为有很多不必要的数据。 而且临时表没有索引，用不到索引大大降低查询效率。

实际场景，我们都是使用Join连接：
1. 左连接查询 left join ... on ...
   - 语法：`select A.filed, [A.filed2, .... ,] B.filed, [B.filed4...,] from <left table> as A  left join <right table> as B on <expression>`
   - 表示左边的A表不动（完全保留），右边B表去匹配A表，ON的条件匹配成功的B表的行可以被挑选出来。
   - **左连接**，其实就可以看成**左表是主表**，**右表是从表**。
   - 当ON条件为1时，总记录数还是跟全相乘一样，只是左表不变，右表匹配。
   - 根据cat_id使两表关联行。不会有其它冗余数据，查询速度快，消耗内存小，而且使用了索引。
        ```sql
        SELECT g.goods_name,g.cat_id,c.cat_id,c.cat_name FROM mingoods g LEFT JOIN category c ON g.cat_id = c.cat_id;
        ```
    - 对于左连接查询，如果右表中没有满足条件的行，则默认填充NULL。

2. 右连接查询 right join ... on ...
   - 语法：`select A.field1,A.field2,..., B.field3,B.field4  from <left table> A right join <right table> B on <expression>`
   - 右连接是以右表为主表，会将右表所有数据查询出来，而左表则根据条件去匹配，如果左表没有满足条件的行，则左边默认显示NULL。
        ```sql
        SELECT g.goods_name,g.cat_id AS g_cat_id,  c.cat_id AS c_cat_id,c.cat_name FROM mingoods g RIGHT JOIN mincategory c ON g.cat_id = c.cat_id;
        ```

3. 内连接 inner join ... on ...
   - 语法：`select A.field1,A.field2,.., B.field3, B.field4 from <left table> A inner join <right table> B on <expression>`
   - 内连接查询，就是取左连接和右连接的交集，如果两边不能匹配条件，则都不取出。如果匹配怎会出现两列一样的值。
        ```sql
        SELECT g.goods_name,g.cat_id, c.* from mingoods g INNER JOIN mincategory c ON g.cat_id = c.cat_id;
        ```
4. 全连接 full join ... on ...
   - 语法：`select ... from <left table> full join <right table> on <expression>`
   - 全连接会将两个表的所有数据查询出来，不满足条件的为NULL。
   - 全连接查询跟全相乘查询的区别在于，如果某个项不匹配，全相乘不会查出来，全连接会查出来，而连接的另一边则为NULL。

5. 联合查询 union
   - 语法：`select A.field1 as f1, A.field2 as f2 from <table1> A union (select B.field3 as f1, field4 as f2 from <table2> B)`
   - union是求两个查询的并集。union合并的是结果集，不区分来自于哪一张表，所以可以合并多张表查询出来的数据。
        ```sql
        SELECT id, content, user FROM comment UNION (SELECT id, msg AS content, user FROM feedback);
        ```
    - union查询，列名不一致时，以第一条sql语句的列名对齐.
    - 使用union查询会将重复的行过滤掉.
    - 使用union all查询所有，重复的行不会被过滤.
    - union查询，如果列数不相等，会报列数不相等错误.
    - union 后的结果集还可以再做筛选
        ```sql
        SELECT id,content,user FROM comment UNION ALL (SELECT id, msg, user FROM feedback) ORDER BY id DESC; 
        ```
         >- union查询时，order by放在内层sql中是不起作用的；因为union查出来的结果集再排序，内层的排序就没有意义了；因此，内层的order by排序，在执行期间，被mysql的代码分析器给优化掉了。
        > - order by 如果和limit一起使用，就显得有意义了，就不会被优化掉

总结：
1. 左右连接既然可以互换，出于移植兼容性方面的考虑，尽量使用左连接。
2. `LEFT JOIN` 是 `LEFT OUTER JOIN` 的缩写，同理，`RIGHT JOIN` 是 `RIGHT OUTER JOIN` 的缩写；`JOIN` 是 `INNER JOIN` 的缩写。