# Hibernate 入门

# 目录
[TOC]

# 前言
## 学习资料
1. [官方英文文档](https://docs.jboss.org/hibernate/stable/orm/userguide/html_single/Hibernate_User_Guide.html)
# 环境搭建

## 导包
1. 通过Maven导入
```xml
<dependency>
  <groupId>org.hibernate</groupId>
  <artifactId>hibernate-agroal</artifactId>
  <version>5.3.11.Final</version>
  <type>pom</type>
</dependency>
```

2. [下载](http://hibernate.org/orm/downloads/)然后导入`lib\required`下面 jar 包.

## 配置
为了连接数据库，还需要数据库的驱动包。

核心配置文件以 hibernate.cfg.xml 命名，放在类路径下（Idea 放在 resources 下）

```xml
<hibernate-configuration>
    <!-- SessionFactory，相当于之前学习连接池配置 -->
    <session-factory>
        <!-- 1 基本4项 -->
        <property name="hibernate.connection.driver_class">com.mysql.jdbc.Driver</property>
        <property name="hibernate.connection.url">jdbc:mysql:///db01</property>
        <property name="hibernate.connection.username">root</property>
        <property name="hibernate.connection.password">1234</property>
    
    </session-factory>

</hibernate-configuration>
```

# 实体类映射
1. 首先是，构建实体类放在`po`包内；
2. 然后有两种方式进行映射：通过xml或者注解：
   1. xml映射文件放在实体类同层目录下的`mapping`包里面：名字和实体类相同，扩展名为 `实体类名.hbm.xml`
        >例如： User.hbm.xml
   2. 注解： 
## xml方法举例
1. 构建实体类：IdeaTemplate.java
    ```java
    public class IdeaTemplate {
        
        private Long ideaTemplateId;
        
        private Long accountId;
        
        private Integer ideaType;
        
        private Integer deviceType;
        
        private String content;
        
        private Date createDate;
        
        private Date chgDate;
        //省略set、get、和toString方法
    ```
2. 构建映射xml：
    >本例使用了生成器生成，所以`主键id`属性的generator会是一个生成类，如果是原生的`class="native"`。其他主键生成种类可[点击这里查看](#hibernateKeyStyles)
    
    ```xml
    <?xml version="1.0" encoding="utf-8"?>
    <!DOCTYPE hibernate-mapping PUBLIC "-//Hibernate/Hibernate Mapping DTD 3.0//EN"
    "http://hibernate.sourceforge.net/hibernate-mapping-3.0.dtd">

    <hibernate-mapping>
        <class name="com.xxx.yyy.api.po.IdeaTemplate" 
                table="ideatemplate" schema="yyy">
                
            <id name="ideaTemplateId" type="java.lang.Long">
                <column name="ideatemplateid" precision="10" scale="0" not-null="true" />
                <generator class="com.xxx.yyy.api.dao.MyHibernateIdentifierGenerator">
                    <!-- param上面实例化的参数 MyHibernateIdentifierGenerator(String sequence)-->
                    <param name="sequence">kt_ideatemplate_seq</param>
                </generator>
            </id>
            
            <property name="accountId" type="java.lang.Long">
                <column name="accountid" precision="10" scale="0" not-null="true" />
            </property>      
            
            <property name="ideaType" type="java.lang.Integer">
                <column name="ideatype" precision="10" scale="0" not-null="true" />
            </property>   
            
            <property name="deviceType" type="java.lang.Integer">
                <column name="device_type" precision="10" scale="0" not-null="true" />
            </property>   
            
            <property name="content" type="java.lang.String">
                <column name="content" not-null="true" />
            </property> 
            
            <property name="createDate" type="java.util.Date">
                <column name="createdate" length="7" not-null="true" />
            </property>
            
            <property name="chgDate" type="java.util.Date">
                <column name="chg_date" length="7" not-null="true" />
            </property>
            
        </class>
    </hibernate-mapping>
    ```
    
**一定要在核心配置文件中加上映射文件位置**：
```xml
<mapping resource="com/xxx/yyy/api/mapping/IdeaTemplate.hbm.xml"/>
```
## 注解方法举例
```java
// name 对应表名称
@Entity
@Table(name = "t_user")
public class User {
    // 主键
    @Id
    @GeneratedValue
    private Integer uid;
    private String username;
    private String password;
// 省略 get set 方法
}
```

在核心配置文件中加上映射文件位置
    
    <mapping class="com.xxx.yyy.api.vo.User" />

## 测试
```java
public class HelloWorld {
    @Test
    public void hello() {
        // username, password
        User user = new User("123456", "123");

        // 1.加载配置文件
        Configuration configure = new Configuration().configure();
        // 2.获得session factory对象
        SessionFactory sessionFactory = configure.buildSessionFactory();
        // 3.创建session
        Session session = sessionFactory.openSession();
        // 4.开启事务
        Transaction transaction = session.beginTransaction();
        // 5.保存并提交事务
        session.save(user);
        transaction.commit();
        // 6.释放资源
        session.close();
        sessionFactory.close();
    }
}
```

# 核心配置详解
## 核心api
1. **Configuration**
   常用方法：
    * 构造方法：默认加载 hibernate.properties
    * configure() 方法：默认加载 hibernate.cfg.xml
    * configure(String) 方法：加载指定配置文件
    
    手动添加映射
    ```java
    config.addResource("com/xxx/yyy/api/mapping/User.hbm.xml");

    // 手动加载指定类，对应的映射文件 User--> User.hbm.xml
    config.addClass(User.class);
    ```
2. **SessionFactory**
   * SessionFactory 相当于java web连接池，用于管理所有session
    * 获得方式：config.buildSessionFactory();
    * sessionFactory hibernate缓存配置信息 （数据库配置信息、映射文件，预定义HQL语句 等）
    * SessionFactory线程安全，可以是成员变量，多个线程同时访问时，不会出现线程并发访问问题
    * 开启一个 session：factory.openSession();

    * 获取和当前线程绑定的会话（需要配置）：factory.getCurrentSession();
        ```xml
        <property name="hibernate.current_session_context_class">thread</property>
        ```

3. **Session**
   * Session 相当于 JDBC的 Connection -- 会话
    * 通过session操作PO对象 --增删改查
    * session单线程，线程不安全，不能编写成成员变量。
    * api:
      * `save` 保存
      * `update` 更新
      * `delete` 删除
      * `get` 通过id查询，如果没有 null
      * `load` 通过id查询，如果没有抛异常
      * `createQuery("hql") ` 获得Query对象
      * `createCriteria(Class)` 获得Criteria对象
4. **Transaction**
    * 开启事务 `beginTransaction()`
    * 获得事务 `getTransaction()`
    * 提交事务：`commit()`
    * 回滚事务：`rollback()`
    >和 spring 整合后，无需手动管理
5. **Query**
    * hibernate执行hql语句.
    * hql语句：hibernate提供面向对象查询语句，使用对象（类）和属性进行查询。区分大小写。
    * 获得 `session.createQuery("hql")`
    * 方法：
        * `list()`  查询所有
        * `uniqueResult()` 获得一个结果。如果没有查询到返回null，如果查询多条抛异常。

        * `setFirstResult(int)` 分页，开始索引数startIndex
        * `setMaxResults(int)` 分页，每页显示个数 pageSize
6. **Criteria**
   * QBC（query by criteria），hibernate提供纯面向对象查询语言，提供直接使用PO对象进行操作。
    * 获得方式：`Criteria criteria = session.createCriteria(User.class);`
    * 条件
    ```java
    criteria.add(Restrictions.eq("username", "tom"));
    Restrictions.gt(propertyName, value)        大于
    Restrictions.ge(propertyName, value)    大于等于
    Restrictions.lt(propertyName, value)    小于
    Restrictions.le(propertyName, value)    小于等于
    Restrictions.like(propertyName, value)  模糊查询，注意：模糊查询值需要使用 % _
    ```

## 工具类
```java
public class HibernateUtils {
    private static SessionFactory sessionFactory;
    static {
        Configuration configuration = new Configuration().configure();
        sessionFactory = configuration.buildSessionFactory();

        Runtime.getRuntime().addShutdownHook(new Thread(){
            @Override
            public void run() {
                sessionFactory.close();
            }
        });
    }

    public static Session openSession() {
        return sessionFactory.openSession();
    }

    public static Session getCurrentSession() {
        return sessionFactory.getCurrentSession();
    }

    public static void main(String[] args) {
        Session session = openSession();
        System.out.println(session);
        session.close();

    }
}
```

## 核心配置
### 基本配置
```xml
<!-- SessionFactory，相当于之前学习连接池配置 -->
<session-factory>
    <!-- 1 基本4项 -->
    <property name="hibernate.connection.driver_class">com.mysql.jdbc.Driver</property>
    <property name="hibernate.connection.url">jdbc:mysql:///h_day01_db</property>
    <property name="hibernate.connection.username">root</property>
    <property name="hibernate.connection.password">1234</property>

    <!-- 2 与本地线程绑定 -->
    <property name="hibernate.current_session_context_class">thread</property>

        <!-- 3 方言：为不同的数据库，不同的版本，生成sql语句（DQL查询语句）提供依据 
            * mysql 字符串 varchar
            * orcale 字符串 varchar2
        -->
        <property name="hibernate.dialect">org.hibernate.dialect.MySQL5Dialect</property>
    
    <!-- 4 sql语句 -->
    <!-- 显示sql语句 -->
    <property name="hibernate.show_sql">true</property>
    <property name="hibernate.format_sql">true</property>

    <!-- 5 自动创建表（了解） ，学习中使用，开发不使用的。
        * 开发中DBA 先创建表，之后根据表生产 PO类
        * 取值：
        update：
            如果表不存在，将创建表。
            如果表已经存在，通过hbm映射文件更新表（添加）。（映射文件必须是数据库对应）
                表中的列可以多，不负责删除。
        create ：如果表存在，先删除，再创建。程序结束时，之前创建的表不删除。
        create-drop：与create几乎一样。如果factory.close()执行，将在JVM关闭同时，将创建的表删除了。(测试)
        validate：校验 hbm映射文件 和 表的列是否对应，如果对应正常执行，如果不对应抛出异常。(测试)
    -->
    <property name="hibernate.hbm2ddl.auto">create</property>
    
    <!-- 6 java web 6.0 存放一个问题
        * BeanFactory 空指针异常
            异常提示：org.hibernate.HibernateException: Unable to get the default Bean Validation factory
        * 解决方案：取消bean校验
    -->
    <property name="javax.persistence.validation.mode">none</property>

    <!-- 添加映射文件 
        <mapping >添加映射文件
            resource 设置 xml配置文件 （addResource(xml)）
            class 配置类 (addClass(User.class)) 配置的是全限定类名
    -->
    <mapping resource="com/xxx/yyy/api/mapping/IdeaTemplate.hbm.xml"/>
</session-factory>
```

### 主键种类
* 自然主键: 在业务中,某个属性符合主键的三个要求.那么该属性可以作为主键列.
* 代理主键: 在业务中,不存符合以上3个条件的属性,那么就增加一个没有意义的列.作为主键.

在Hibernate中应使用代理主键。在Hibernate中,Hibernate依靠对象表示来区分不同的持久化，而对象标识符则可以通过Hibernate内置的表示生成器来产生。

### 类型对应
Javas数据类型 | Hibernate数据类型 | SQL数据类型（不太DB有差异）
| --- | --- | --- |
byte、java.lang.Byte  | byte    | TINYINT
java.lang.Short | short   | SMALLINT
java.lang.Integer   | integer | BIGINT
java.lang.Long  | long    | TINYINT
java.math.BigDecimal  | big_decimal    | NUMERIC
java.lang.Character  | character    | 	CHAR(1)
java.lang.Boolean  | boolean    | BIT
java.lang.String  | string    | VARCHAR
java.util.Date  | date    | DATE
java.sql.Time  | time    | TIME
byte[]  | binary    | 	VARBINARY

### 普通属性
- **hibernate-mapping**
  - `package` 用于配置PO类所在包
    >例如： package="com.ittianyu.d_hbm"
  - **class** 配置 PO类 和 表 之间对应关系
    - `name`：PO类全限定类名
      >例如：name="com.ittianyu.d_hbm.Person"如果配置 package，name的取值可以是简单类名 name="Person"
    - `table` : 数据库对应的表名
    - `dynamic-insert="false"` 是否支持动态生成insert语句
    - `dynamic-update="false"` 是否支持动态生成update语句
        >如果设置true，hibernate底层将判断提供数据是否为null，如果为null，insert或update语句将没有此项。
    - **id** [主键](#idProperty)
    - **property** 普通字段
      * `name` : PO类的属性
      * `column` : 表中的列名，默认name的值相同
      * `type`:表中列的类型。默认hibernate自己通过getter获得类型，一般情况不用设置。
        - 取值类型1： hibernate类型
            * string 字符串
            * integer 整形
        - 取值类型2： java类型 （全限定类名）
            * java.lang.String 字符串
        - 取值类型3：数据库类型
            * varchar(长度) 字符串
            * int 整形
            ```xml
            <property name="birthday">
                <column name="birthday" sql-type="datetime"></column>
            </property>
            ```
        - javabean 一般使用类型 java.util.Date
        - jdbc规范提供3中
            * java类型              mysql类型
            * java.sql.Date       date
            * java.sql.time       time
            * java.sql.timestamp  timestamp
            * null                datetime
            >以上三个类型都是java.util.Date子类
                    
      * `length` : 列的长度。默认值：255
      * `not-null` : 是否为null
      * `unique` : 是否唯一
      * `access` ：设置映射使用PO类属性或字段
        * _property_ : 使用PO类属性，必须提供setter、getter方法
        * _field_ : 使用PO类字段，一般很少使用。
      * `insert` 生成insert语句时，是否使用当前字段。
      * `update` 生成update语句时，是否使用当前字段。
        >默认情况：hibernate生成insert或update语句，使用配置文件所有项
>注意：配置文件如果使用关键字，列名必须使用重音符 

### <span id="idProperty">主键属性</span>
- **id** 配置主键
  - name:属性名称
  - access="" 设置使用属性还是字段
  - column=""  表的列名
  - length=""  长度
  - type="" 类型
  - **generator**
    -  `class` 属性用于设置[主键生成策略/类型](#hibernateKeyStyles)

# 缓存

## 对象状态
1. **三种状态**：
    * 瞬时态：transient，session没有缓存对象，数据库也没有对应记录。
        >OID特点：没有值
    * 持久态：persistent，session缓存对象，数据库最终会有记录。（事务没有提交）
        >OID特点：有值
    * 脱管态：detached，session没有缓存对象，数据库有记录。
        >OID特点：有值
2. **状态转换**
   ![三种状态转换图](../../配图\数据库\hibernate状态转换图.png)
## 一级缓存
   **一级缓存** ：又称为session级别的缓存。当获得一次会话（session），hibernate在session中创建多个集合（map），用于存放操作数据（PO对象），为程序优化服务，如果之后需要相应的数据，hibernate优先从session缓存中获取，如果有就使用；如果没有再查询数据库。当session关闭时，一级缓存销毁。
```java
@Test
public void demo02(){
    //证明一级缓存
    Session session = factory.openSession();
    session.beginTransaction();
    
    //1 查询 id = 1
    User user = (User) session.get(User.class, 1);
    System.out.println(user);
    //2 再查询 -- 不执行select语句，将从一级缓存获得
    User user2 = (User) session.get(User.class, 1);
    System.out.println(user2);
    
    session.getTransaction().commit();
    session.close();
}
//可以调用方法清除一级缓存
//清除
//session.clear();
session.evict(user);
```
 **快照**
   与一级缓存一样的存放位置，对一级缓存数据备份。保证数据库的数据与 一级缓存的数据必须一致。如果一级缓存修改了，在执行commit提交时，将自动刷新一级缓存，执行update语句，将一级缓存的数据更新到数据库。
    
    当缓存和数据库数据不一样且在提交之前，可以调用 refresh 强制刷新缓存。

## 二级缓存
sessionFactory 级别缓存，整个应用程序共享一个会话工厂，共享一个二级缓存。

由4部分构成：
* 类级别缓存
* 集合级别缓存
* 时间戳缓存
* 查询缓存(二级缓存的第2大部分,三级缓存)

### 并发访问策略
与隔离对应
* **transactional**
  > 可以防止脏读和不可重复读，性能低
* **read-write **
  >可以防止脏读，更新缓存时锁定缓存数据
* **nonstrict-read-write**
  > 不保证缓存和数据库一致，为缓存设置短暂的过期时间，减少脏读
* **read-only**
  > 适用于不会被修改的数据，并发性能高
### 二级缓存应用场景
- 适合放入二级缓存中的数据:
  - 很少被修改
  - 不是很重要的数据, 允许出现偶尔的并发问题
- 不适合放入二级缓存中的数据:
  - 经常被修改
  - 财务数据, 绝对不允许出现并发问题
  - 与其他应用数据共享的数据

### 开启二级缓存

### 使用二级缓存
1. 类缓存
2. 集合缓存
3. 查询缓存
   1. 配置
   2. 使用

# 关系映射
## 一对一
一对一关系一般是可以整合成一张表，也可以分成两张表。
维护两张表的关系可以选择外键也可以选择让主键同步。
1. 实体类
2. 外键维护关系
3. 主键同步关系

## 一对多
1. 实体类
2. 映射文件


## 多对多
1. 实体类
2. 映射文件

## 级联


# 抓取策略
## 检索方法
* 立即检索：立即查询，在执行查询语句时，立即查询所有的数据。
* 延迟检索：延迟查询，在执行查询语句之后，在需要时在查询。（懒加载）


## 检索策略
* 类级别检索：当前的类的属性获取是否需要延迟。
* 关联级别的检索：当前类 关联 另一个类是否需要延迟。

## 类级别检索
* get：立即检索。get方法一执行，立即查询所有字段的数据。
* load：延迟检索。默认情况，load方法执行后，如果只使用OID的值不进行查询，如果要使用其他属性值将查询。可以配置是否延迟检索：
````xml
<!-- lazy 默认值true，表示延迟检索，如果设置false表示立即检索。 -->
<class  lazy="true | false">
````
## 关联级别检索
容器\<set> 提供两个属性：fetch、lazy，用于控制关联检索。
* fetch：确定使用sql格式
  * join：底层使用迫切左外连接
  * select：使用多个select语句（默认值）
  * subselect：使用子查询


* lazy：关联对象是否延迟。

  * false：立即
  * true：延迟（默认值）
  * extra：极其懒惰，调用 size 时，sql 查询 count。（用于只需要获取个数的时候）


## 批量查询

## 检索总结

# 查询
## HQL

### 查询所有
```java
//1  使用简单类名 ， 存在自动导包
// * Customer.hbm.xml <hibernate-mapping auto-import="true">
//  Query query = session.createQuery("from Customer");
//2 使用全限定类名
Query query = session.createQuery("from com.ittianyu.bean.Customer");
// 获取结果
List<Customer> allCustomer = query.list();
```

### 条件查询
```java
//1 指定数据，cid OID名称
//  Query query = session.createQuery("from Customer where cid = 1");
//2 如果使用id，也可以（了解）
//  Query query = session.createQuery("from Customer where id = 1");
//3 对象别名 ,格式： 类 [as] 别名
//  Query query = session.createQuery("from Customer as c where c.cid = 1");
//4 查询所有项，mysql--> select * from...
Query query = session.createQuery("select c from Customer as c where c.cid = 1");

Customer customer = (Customer) query.uniqueResult();
```

### 投影查询
```java
//1 默认
//如果单列 ，select c.cname from，需要List<Object>
//如果多列，select c.cid,c.cname from ，需要List<Object[]>  ,list存放每行，Object[]多列
//  Query query = session.createQuery("select c.cid,c.cname from Customer c");
//2 将查询部分数据，设置Customer对象中
// * 格式：new Customer(c.cid,c.cname)
// * 注意：Customer必须提供相应的构造方法。
// * 如果投影使用oid，结果脱管态对象。
Query query = session.createQuery("select new Customer(c.cid,c.cname) from Customer c");

List<Customer> allCustomer = query.list();
```

### 排序
```java
Query query = session.createQuery("from Customer order by cid desc");
List<Customer> allCustomer = query.list();
```

### 分页
```java
Query query = session.createQuery("from Customer");
// * 开始索引 , startIndex 算法： startIndex = (pageNum - 1) * pageSize;
// *** pageNum 当前页（之前的 pageCode）
query.setFirstResult(0);
// * 每页显示个数 ， pageSize
query.setMaxResults(2);

List<Customer> allCustomer = query.list();
```

### 绑定参数
```java
Integer cid = 1;

//方式1 索引 从 0 开始
//  Query query = session.createQuery("from Customer where cid = ?");
//  query.setInteger(0, cid);
//方式2 别名引用 (:别名)
Query query = session.createQuery("from Customer where cid = :xxx");
//  query.setInteger("xxx", cid);
query.setParameter("xxx", cid);

Customer customer = (Customer) query.uniqueResult();
```

### 聚合函数和分组
```java
//1 
//  Query query = session.createQuery("select count(*) from Customer");
//2 别名
//  Query query = session.createQuery("select count(c) from Customer c");
//3 oid
Query query = session.createQuery("select count(cid) from Customer");

Long numLong = (Long) query.uniqueResult();
```

### 连接查询
```java
//左外连接
//  List list = session.createQuery("from Customer c left outer join c.orderSet ").list();
//迫切左外链接 (默认数据重复)
//  List list = session.createQuery("from Customer c left outer join fetch c.orderSet ").list();
//迫切左外链接 (去重复)
List list = session.createQuery("select distinct c from Customer c left outer join fetch c.orderSet ").list();
```

### 命名查询
Custom.hbm.xml
```xml
...
<!--局部 命名查询-->
<query name="findAll"><![CDATA[from Customer ]]></query>
</class>

<!--全局 命名查询-->
<query name="findAll"><![CDATA[from Customer ]]></query>
```
测试：
```java
//全局
//List list = session.getNamedQuery("findAll").list();
//局部
List list = session.getNamedQuery("com.ittianyu.a_init.Customer.findAll").list();
```

## QBC

### 查询所有
```java
List<Customer> list = session.createCriteria(Customer.class).list();
```

### 分页查询
```java
Criteria criteria = session.createCriteria(Order.class);
criteria.setFirstResult(10);
criteria.setMaxResults(10);
List<Order> list = criteria.list();
```

### 排序
```java
Criteria criteria = session.createCriteria(Customer.class);
// criteria.addOrder(org.hibernate.criterion.Order.asc("age"));
criteria.addOrder(org.hibernate.criterion.Order.desc("age"));
List<Customer> list = criteria.list();
```

### 条件查询
```java
// 按名称查询:
/*Criteria criteria = session.createCriteria(Customer.class);
criteria.add(Restrictions.eq("cname", "tom"));
List<Customer> list = criteria.list();*/

// 模糊查询;
/*Criteria criteria = session.createCriteria(Customer.class);
criteria.add(Restrictions.like("cname", "t%"));
List<Customer> list = criteria.list();*/

// 条件并列查询
Criteria criteria = session.createCriteria(Customer.class);
criteria.add(Restrictions.like("cname", "t%"));
criteria.add(Restrictions.ge("age", 35));
List<Customer> list = criteria.list();
```

### 离线查询
```java
// service 层 封装与 session 无关的 criteria
DetachedCriteria detachedCriteria = DetachedCriteria.forClass(Customer.class);
detachedCriteria.add(Restrictions.eq("id", 4));

// dao 层
Session session = HibernateUtils.openSession();
Criteria criteria = detachedCriteria.getExecutableCriteria(session);
List list = criteria.list();
```


# 事务
## 隔离级别
* read uncommittd，读未提交。存在3个问题。
* read committed，读已提交。解决：脏读。存在2个问题。
* repeatable read ，可重复读。解决：脏读、不可重复读。存在1个问题。
* serializable，串行化。单事务。没有问题

## hibernate 中配置
```xml
<!-- 1,2,3,4分别对应上面的隔离级别，0表示没有事务级别。 -->
<property name="hibernate.connection.isolation">4</property>
```

## 锁
### 悲观锁
采用数据库锁机制。丢失更新肯定会发生。
1. 读锁：共享锁
   ```sql
   select .... from  ... lock in share mode;
   ```
2. 写锁：排它锁
   ```sql
   select ... from  ....  for update;
   ```
**Hibernate** 中使用
```java
Customer customer = (Customer) session.get(Customer.class, 1 ,LockMode.UPGRADE);

```

### 乐观锁
在表中提供一个字段（版本字段），用于标识记录。如果版本不一致，不允许操作。丢失更新肯定不会发生.

**Hibernate** 中使用
1. 在PO对象（javabean）提供字段，表示版本字段。
    ```java
    ...
    private Integer version;
    ...

    ```
2. 在配置文件中增加 version
    ```xml
    <class ...>
        ...
        <version name="version" />
        ...
    ```
# 其他配置


# 参考资料

[1] IT天宇： http://www.jianshu.com/p/50964e92c5fb
[2] Hibernate主键生成种类： https://blog.51cto.com/aiilive/931190

# 附录

## 附录1：主键生成种类 
<span id="hibernateKeyStyles">Hibernate主键生成种类</span>

1) assigned(手工分配主键ID值)
    主键由外部程序负责生成，无需Hibernate参与。该策略要求程序员必须自己维护和管理主键，当有数据需要存储时，程序员必须自己为该数据分配指定一个主键ID值，如果该数据没有被分配主键ID值或分配的值存在重复，则该数据都将无法被持久化且会引起异常的抛出。示例：
    ```java
    @Id
    @GenericGenerator(name = "idGenerator", strategy = "assigned")
    @GeneratedValue(generator = "idGenerator")
    ```

2) hilo
    通过hi/lo 算法实现的主键生成机制，需要额外的数据库表保存主键生成历史状态。


3) seqhilo
    与hilo 类似，通过hi/lo 算法实现的主键生成机制，只是主键历史状态保存在Sequence中，适用于支持Sequence的数据库，如Oracle。


4) **increment**(自然递增)
    主键按数值顺序递增。此方式的实现机制为在当前应用实例中维持一个变量，以保存着当前的最大值，之后每次需要生成主键的时候将此值加1作为主键。该策略不依赖于底层数据库，而依赖于 hibernate 本身。

    >这种方式可能产生的问题是：如果当前有多个实例访问同一个数据库，那么由于各个实例各自维护主键状态，不同实例可能生成同样的主键，从而造成主键重复异常。因此，如果同一数据库有多个实例访问，此方式必须避免使用，比如集群下避免使用。


5) **identity**(自然递增)
    采用数据库提供的主键生成机制。如DB2、SQL Server、MySQL中的主键生成机制。该策略依赖于底层不同的数据库，与Hibernate 和 程序员无关。


6) **sequence***(序列)
    采用数据库提供的sequence 机制生成主键。如Oralce 中的Sequence。


7) **native**
    由Hibernate根据底层数据库自行判断采用identity、hilo、sequence其中一种作为主键生成方式。


8) uuid.hex
    由Hibernate基于128 位唯一值产生算法生成16 进制数值（编码后以长度32 的字符串表示）作为主键。



9) uuid.string
    与uuid.hex 类似，只是生成的主键未进行编码（长度16）。在某些数据库中可能出现问题（如PostgreSQL）。能够保证网络环境下的主键唯一性，也就能够保证在不同数据库及不同服务器下主键的唯一性。
    >uuid 最终被编码成一个32位16进制数的字符串，占用的存储空间较大。用于为 String 类型生成唯一标识，适用于所有关系型数据库。


10) foreign
    使用外部表的字段作为主键。一般而言，利用uuid.hex方式生成主键将提供最好的性能和数据库平台适应性。


另外由于常用的数据库，如Oracle、DB2、SQLServer、MySql 等，都提供了易用的主键生成机制（Auto-Increase 字段或者Sequence）。我们可以在数据库提供的主键生成机制上，采用generator-class=native的主键生成方式。

>不过值得注意的是，一些数据库提供的主键生成机制在效率上未必最佳，大量并发insert数据时可能会引起表之间的互锁。


数据库提供的主键生成机制，往往是通过在一个内部表中保存当前主键状态（如对于自增型主键而言，此内部表中就维护着当前的最大值和递增量），之后每次插入数据会读取这个最大值，然后加上递增量作为新记录的主键，之后再把这个新的最大值更新回内部表中，这样，一次Insert操作可能导致数据库内部多次表读写操作，同时伴随的还有数据的加锁解锁操作，这对性能产生了较大影响。

因此，对于并发Insert要求较高的系统，推荐采用uuid.hex 作为主键生成机制。   