# 简介
## 什么是 MyBatis？
官方中文文档解释： MyBatis 是一款优秀的持久层框架，它支持定制化 SQL、存储过程以及高级映射。MyBatis 避免了几乎所有的 JDBC 代码和手动设置参数以及获取结果集。MyBatis 可以使用简单的 XML 或注解来配置和映射原生类型、接口和 Java 的 POJO（Plain Old Java Objects，普通老式 Java 对象）为数据库中的记录。

mapper文件夹可以等价于dao；
entity文件夹可以等价于bean；

# 入门
## 安装
1. 使用Maven构建项目时，只需要在pom.xml文件中添加：
   ``` xml
    <dependency>
        <groupId>org.mybatis</groupId>
        <artifactId>mybatis</artifactId>
        <version>x.x.x</version>
    </dependency>
   ```

## SqlSessionFactory
SqlSessionFactory是MyBatis的关键对象,它是个单个数据库映射关系经过编译后的内存镜像.SqlSessionFactory对象的实例可以通过SqlSessionFactoryBuilder对象类获得,而SqlSessionFactoryBuilder则可以从XML配置文件或一个预先定制的Configuration的实例构建出SqlSessionFactory的实例.

每一个MyBatis的应用程序都以一个SqlSessionFactory对象的实例为核心.同时SqlSessionFactory也是线程安全的,SqlSessionFactory一旦被创建,应该在应用执行期间都存在.在应用运行期间不要重复创建多次,建议使用单例模式.SqlSessionFactory是创建SqlSession的工厂.

## SqlSession
SqlSession是MyBatis的关键对象,类似于JDBC中的Connection.它是应用程序与持久层之间执行交互操作的一个单线程对象,也是MyBatis执行持久化操作的关键对象.SqlSession对象完全包含以数据库为背景的所有执行SQL操作的方法,它的底层封装了JDBC连接,可以用SqlSession实例来直接执行被映射的SQL语句.

每个线程都应该有它自己的SqlSession实例.SqlSession的实例不能被共享,同时SqlSession也是线程不安全的,绝对不能将SqlSeesion实例的引用放在一个类的静态字段甚至是实例字段中.也绝不能将SqlSession实例的引用放在任何类型的管理范围中,比如Servlet当中的HttpSession对象中.使用完SqlSeesion之后关闭Session很重要,应该确保使用finally块来关闭它.

## SqlSessionFactory和SqlSession实现过程
mybatis框架主要是围绕着SqlSessionFactory进行的，创建过程大概如下：
1. 定义一个Configuration对象，其中包含数据源、事务、mapper文件资源以及影响数据库行为属性设置settings
2. 通过配置对象，则可以创建一个SqlSessionFactoryBuilder对象
3. 通过 SqlSessionFactoryBuilder 获得SqlSessionFactory 的实例。
4. SqlSessionFactory 的实例可以获得操作数据的SqlSession实例，通过这个实例对数据库进行操作

<!-- --------------------- 
作者：可乐丶 
来源：CSDN 
原文：https://blog.csdn.net/u013412772/article/details/73648537  -->

## MyBatis-Spring-Boot-Starter
MyBatis-Spring-Boot-Starter依赖将会提供如下操作：

1. 自动检测现有的DataSource
2. 将创建并注册SqlSessionFactory的实例，该实例使用SqlSessionFactoryBean将该DataSource作为输入进行传递
3. 将创建并注册从SqlSessionFactory中获取的SqlSessionTemplate的实例。
4. 自动扫描您的mappers，将它们链接到SqlSessionTemplate并将其注册到Spring上下文，以便将它们注入到您的bean中。

<!-- 作者：嘟嘟MD
链接：https://juejin.im/post/58fcdcc861ff4b006668f79b
来源：掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。 -->


# 使用
## 注解的方式
在dao里写Mapper接口，例如：
```java
@Component
@Mapper
public interface UserMapper {
    @Insert("Insert into user(user_name,user_password) values(#{userName},#{userPassword})")
    int add(User user);

    @Update("update user set user_name=#{userName},user_password=#{userPassword} where id=#{id}")
    int update(User user);

    @Delete("delete from user where id=#{id}")
    int deleteById(long id);

    @Delete("delete from user where user_name=#{userName")
    int deleteByName(String userName);

    @Select("select id,user_name from user where user_name=#{userName}")
    @Results(id = "user",value= {
            @Result(property = "user_name", column = "user_name", javaType = String.class),
            @Result(property = "id", column = "id", javaType = Long.class)
    })
    User queryUserByName(String userName);

}
```

## xml的方式

1. 在application.yml中添加xml文件的路径，还有mybatis的配置文件（可选）：
```yml
mybatis:
  ##检查 mybatis 配置是否存在，一般命名为 mybatis-config.xml
  check-config-location: true
  ##配置文件位置
  config-location: classpath:mybatis/mybatis-config.xml
  ## mapper xml 文件地址
  # type-aliases扫描路径
  # type-aliases-package:
  # mapper xml实现扫描路径
  mapper-locations: classpath:mapper/*Mapper.xml
```
   
   注意：如果设置了xml路径，程序会优先去寻找这个文件，如果文件不存在会直接报错，而不是转向使用注解。所以两种方式请二选一。
2. 然后再mapper文件夹中添加UserMapper.xml等文件，做为映射
   

# XML 映射文件
MyBatis 的真正强大在于它的映射语句，这是它的魔力所在。由于它的异常强大，映射器的 XML 文件就显得相对简单。如果拿它跟具有相同功能的 JDBC 代码进行对比，你会立即发现省掉了将近 95% 的代码。MyBatis 为聚焦于 SQL 而构建，以尽可能地为你减少麻烦。

SQL 映射文件只有很少的几个顶级元素（按照应被定义的顺序列出）：

- cache – 对给定命名空间的缓存配置。
- cache-ref – 对其他命名空间缓存配置的引用。
- resultMap – 是最复杂也是最强大的元素，用来描述如何从数据库结果集中来加载对象。
- sql – 可被其他语句引用的可重用语句块。
- insert – 映射插入语句
- update – 映射更新语句
- delete – 映射删除语句
- select – 映射查询语句

## select
查询语句是 MyBatis 中最常用的元素之一，光能把数据存到数据库中价值并不大，只有还能重新取出来才有用，多数应用也都是查询比修改要频繁。对每个插入、更新或删除操作，通常间隔多个查询操作。
``` xml
<select id="selectPerson" parameterType="int" resultType="hashmap">
  SELECT * FROM PERSON WHERE ID = #{id}
</select>
```
这个语句被称作 selectPerson，接受一个 int（或 Integer）类型的参数，并返回一个 HashMap 类型的对象，其中的键是列名，值便是结果行中的对应值。

select 元素允许你配置很多属性来配置每条语句的作用细节。
```xml
<select
  id="selectPerson"    
  parameterType="int"
  parameterMap="deprecated"
  resultType="hashmap"
  resultMap="personResultMap"
  flushCache="false"
  useCache="true"
  timeout="10"
  fetchSize="256"
  statementType="PREPARED"
  resultSetType="FORWARD_ONLY">
```
Select元素的属性
属性 | 描述
---  | ---
id | 在命名空间中唯一的标识符，可以被用来引用这条语句。
parameterType | 将会传入这条语句的参数类的完全限定名或别名。可选，默认值为未设置（unset）。
resultType | 	从这条语句中返回的期望类型的类的完全限定名或别名。 注意如果返回的是集合，那应该设置为集合包含的类型，而不是集合本身。可以使用 resultType 或 resultMap，但不能同时使用。
resultMap | 外部 resultMap 的命名引用。结果集的映射是 MyBatis 最强大的特性，如果你对其理解透彻，许多复杂映射的情形都能迎刃而解。
timeout	| 这个设置是在抛出异常之前，驱动程序等待数据库返回请求结果的秒数。默认值为未设置（unset）（依赖驱动）。

## insert, update 和 delete
数据变更语句 insert，update 和 delete 的实现非常接近：
```xml
<insert
  id="insertAuthor"
  parameterType="domain.blog.Author"
  flushCache="true"
  statementType="PREPARED"
  keyProperty=""
  keyColumn=""
  useGeneratedKeys=""
  timeout="20">

<update
  id="updateAuthor"
  parameterType="domain.blog.Author"
  flushCache="true"
  statementType="PREPARED"
  timeout="20">

<delete
  id="deleteAuthor"
  parameterType="domain.blog.Author"
  flushCache="true"
  statementType="PREPARED"
  timeout="20">
```

下面就是 insert，update 和 delete 语句的示例：
```xml
<insert id="insertAuthor">
  insert into Author (id,username,password,email,bio)
  values (#{id},#{username},#{password},#{email},#{bio})
</insert>

<update id="updateAuthor">
  update Author set
    username = #{username},
    password = #{password},
    email = #{email},
    bio = #{bio}
  where id = #{id}
</update>

<delete id="deleteAuthor">
  delete from Author where id = #{id}
</delete>
```
在插入语句里面有一些额外的属性和子元素用来处理主键的生成，而且有多种生成方式。

首先，如果你的数据库支持自动生成主键的字段（比如 MySQL 和 SQL Server），那么你可以设置 `useGeneratedKeys=”true”`，然后再把 `keyProperty` 设置到目标属性上就 OK 了。例如，如果上面的 Author 表已经对 id 使用了自动生成的列类型，那么语句可以修改为：
```xml
<insert id="insertAuthor" useGeneratedKeys="true"
    keyProperty="id">
  insert into Author (username,password,email,bio)
  values (#{username},#{password},#{email},#{bio})
</insert>
```

对于不支持自动生成类型的数据库或可能不支持自动生成主键的 JDBC 驱动，MyBatis 有另外一种方法来生成主键。

这里有一个简单（甚至很傻）的示例，它可以生成一个随机 ID（你最好不要这么做，但这里展示了 MyBatis 处理问题的灵活性及其所关心的广度）：  
>先用selectKey生成一个id，order的before表示在sql语句之前生存； 之后一般是作为有自增的时候，插入后立马读取到新自增的id作为后续使用。
```xml
<insert id="insertAuthor">
  <selectKey keyProperty="id" resultType="int" order="BEFORE">
    select CAST(RANDOM()*1000000 as INTEGER) a from SYSIBM.SYSDUMMY1
  </selectKey>
  insert into Author
    (id, username, password, email,bio, favourite_section)
  values
    (#{id}, #{username}, #{password}, #{email}, #{bio}, #{favouriteSection,jdbcType=VARCHAR})
</insert>
```
在上面的示例中，selectKey 元素中的语句将会首先运行，Author 的 id 会被设置，然后插入语句会被调用。

如果你的数据库还支持多行插入, 你也可以传入一个 Author 数组或集合，并返回自动生成的主键。
```xml
<insert id="insertAuthor" useGeneratedKeys="true"
    keyProperty="id">
  insert into Author (username, password, email, bio) values
  <foreach item="item" collection="list" separator=",">
    (#{item.username}, #{item.password}, #{item.email}, #{item.bio})
  </foreach>
</insert>
```
## sql
这个元素可以被用来定义可重用的 SQL 代码段，这些 SQL 代码可以被包含在其他语句中。它可以（在加载的时候）被静态地设置参数。 在不同的包含语句中可以设置不同的值到参数占位符上。比如：
```xml
<sql id="userColumns"> ${alias}.id,${alias}.username,${alias}.password </sql>
```

这个 SQL 片段可以被包含在其他语句中，例如：
```xml
<select id="selectUsers" resultType="map">
  select
    <include refid="userColumns"><property name="alias" value="t1"/></include>,
    <include refid="userColumns"><property name="alias" value="t2"/></include>
  from some_table t1
    cross join some_table t2
</select>
```

## 参数
参数是 MyBatis 非常强大的元素。对于简单的使用场景，大约 90% 的情况下你都不需要使用复杂的参数，比如：
```xml
<select id="selectUsers" resultType="User">
  select id, username, password
  from users
  where id = #{id}
</select>
```

上面的这个示例说明了一个非常简单的命名参数映射。参数类型被设置为 int，这样这个参数就可以被设置成任何内容。原始类型或简单数据类型（比如 Integer 和 String）因为没有相关属性，它会完全用参数值来替代。 然而，如果传入一个复杂的对象，行为就会有一点不同了。比如：
```xml
<insert id="insertUser" parameterType="User">
  insert into users (id, username, password)
  values (#{id}, #{username}, #{password})
</insert>
```
如果**User 类型的参数对象传递到了语句中**，**id、username 和 password 属性将会被查找**，然后将它们的值传入预处理语句的参数中。

大多时候你只须简单地指定属性名，其他的事情 MyBatis 会自己去推断，顶多要为可能为空的列指定 jdbcType(比如HashMap对象)。
```xml
#{firstName}
#{middleInitial,jdbcType=VARCHAR}
#{lastName}
```

## 字符串替换
默认情况下,使用`#{}`格式的语法会导致 MyBatis 创建 PreparedStatement 参数占位符并安全地设置参数（就像使用 ? 一样）。

## 结果映射
resultMap 元素是 MyBatis 中最重要最强大的元素。它可以让你从 90% 的 JDBC ResultSets 数据提取代码中解放出来，并在一些情形下允许你进行一些 JDBC 不支持的操作。

实际上，在为一些比如连接的复杂语句编写映射代码的时候，一份 resultMap 能够代替实现同等功能的长达数千行的代码。

ResultMap 的设计思想是，对于简单的语句根本不需要配置显式的结果映射，而对于复杂一点的语句只需要描述它们的关系就行了。
```xml
<select id="selectUsers" resultType="map">
  select id, username, hashedPassword
  from some_table
  where id = #{id}
</select>
```
上述语句只是简单地将所有的列映射到 HashMap 的键上，这由 resultType 属性指定。虽然在大部分情况下都够用，但是 HashMap 不是一个很好的领域模型。你的程序更可能会使用 JavaBean 或 POJO（Plain Old Java Objects，普通老式 Java 对象）作为领域模型。

```java
package com.someapp.model;
public class User {
  private int id;
  private String username;
  private String hashedPassword;

  //省略get、set语句...
```
基于 JavaBean 的规范，上面这个类有 3 个属性：id，username 和 hashedPassword。这些属性会对应到 select 语句中的列名。

这样的一个 JavaBean 可以被映射到 ResultSet，就像映射到 HashMap 一样简单。

```xml
<select id="selectUsers" resultType="com.someapp.model.User">
  select id, username, hashedPassword
  from some_table
  where id = #{id}
</select>
```
**类型别名**——使用它们，你就可以不用输入类的完全限定名称了。比如：

```xml
<!-- mybatis-config.xml 中 -->
<typeAlias type="com.someapp.model.User" alias="User"/>

<!-- SQL 映射 XML 中 -->
<select id="selectUsers" resultType="User">
  select id, username, hashedPassword
  from some_table
  where id = #{id}
</select>
```
这些情况下，MyBatis 会在幕后自动创建一个 ResultMap，再基于属性名来映射列到 JavaBean 的属性上。如果列名和属性名没有精确匹配，可以在 SELECT 语句中对列使用别名（这是一个基本的 SQL 特性）来匹配标签。比如：

```xml
<select id="selectUsers" resultType="User">
  select
    user_id             as "id",
    user_name           as "userName",
    hashed_password     as "hashedPassword"
  from some_table
  where id = #{id}
</select>
```
如果使用外部的 resultMap 会怎样，这也是解决列名不匹配的另外一种方式。

```xml
<resultMap id="userResultMap" type="User">
  <id property="id" column="user_id" />
  <result property="username" column="user_name"/>
  <result property="password" column="hashed_password"/>
</resultMap>
```
而在引用它的语句中使用 resultMap 属性就行了（注意我们去掉了 resultType 属性）。比如:
```xml
<select id="selectUsers" resultMap="userResultMap">
  select user_id, user_name, hashed_password
  from some_table
  where id = #{id}
</select>
```