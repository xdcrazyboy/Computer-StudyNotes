
# Spring核心特性学习笔记

# 一、 框架总览

**课前准备**：
1. 心态：戒躁戒躁、谨慎豁达、如履薄冰
2. 方法：
   1. 基础:夯实基础，了解动态
   2. 思考：保持怀疑，验证一切
   3. 分析： 不拘小节，观其大意
   4. 实践：思辨结合，学以致用
3. 工具
   1. JDK： Oracle JDK 8
   2. Spring FrameWork 5.2.2。RELEASE
   3. IDE
   4. Maven 3.2+

## 总览

**核心特性**：
- IoC容器
- Spring事件 Events
- 资源管理（Resources）
- 国际化（i18n）
- 校验（Validation） BeanValidation
- 数据绑定 Data Binding
- 类型转换 Type Conversion 
- Spring表达式 Spring Express Language
- 面向切面编程 AOP

**数据存储**：
- JDBC
- 事务抽象 Transactions
- DAO支持 DAO Support
- O/R映射 Mapping
- XML编列 Marshlling

**Web技术**
- Web Servlet技术栈
  - Spring MVC
  - WebSocket
  - SockJS
- Web Reactive技术栈
  - Spring WebFlux
  - WebClient
  - WebSocket

**技术整合**
- 远程调用 Remoting
- Java消息服务 JMS
- Java连接架构 JCA 可忽略
- Java管理扩展 JMX
- 本地服务 Tasks
- 本地调度 Scheduling
- 缓存抽象 Caching
- Spring测试 Testing
  - 模拟对象 Mock Objects
  - TestContect框架
  - Spring Mvc测试
  - Web测试客户端


## 版本特性
Soring 4.x -- Java 6+ ：springboot1.0
Spring 5.x -- Java标准版 8+  -- Java企业版 Java EE 7  -- springboot2

## 模块化设计
大概分为了 **20**多个模块：
- aop  切面
- aspects  切面
- beans  与context一起组成ioc
- context-indexer
- contect-support
- context
- core
- expression  表达式语言 EL
- instrument  java装配
- jcl  日志框架
- jdbc
- jms  消息服务
- messaging  统一消息服务
- orm
- oxm  xml编列
- test
- tx  事务
- web  想统一，所以分为多个模块
- webflux
- webmvc
- websocket
  

## 技术整合

### Java语言特性运用

Java 语言变化：
- 5： 枚举、泛型、注解、封箱（解箱） spring1.2 第一版开始支持 2004年
  - 注解、枚举 spring1.2+支持
  - for-each、自动转拆箱、泛型 spring3.0+支持
- 6： @Override  spring4.0
- 7： Diamond语法(泛型<>不需要写)、多Catch、Try语法糖（Resource close）。   spring5+
- 8： Lambda语法、可重复注解（一个方法可以加多个注解）、类型注解
- 9：模块化、接口私有方法
- 10：局部变量类型推断


### JDK API实践
- < Java5: 反射 Reflection、Java Beans、动态代理 
- Java5： 并发框架（J.U.C）、格式化（Formatter）、Java管理扩展（JMX）、Instrumentation、XML处理（DOM、SAX、XPath、XSTL）
- Java6： JDBC 4.0（JSR 221）、JAXB 2.0（JSR 222）、可插拔注解处理API（JSR 269）、Common Annotations（JSR 250）、Java Compiler API（JSR 199）、Scripting in JVM(JSR 223)
- Java7: NIO 2(JSR 203)（PathResource）、Fork/Join框架（JSR 116）、invokedynamic 字节码
- Java8： Stream API、CompletableFuture(J.U.C)、Annotation on Java Types、Date and Time API、可重复Annotations 、JavaScript运行时
- Java9：Reactive Stream Flow API、Process API Updates、Variable Handlers、Method Handles、Spin-Wait Hints、Stack-Walking API

>JSR: Java Specification Requests

### JavaEE API整合

**Java EE Web技术相关**：
- Servlet + JSP--- DispatcherServlet
- JSTL  -- JstlView
- JSF（JavaServer Faces） -- 
- Portlet
- SOAP 简单对象访问协议
- WebServices
- WebSocket

**数据存储相关**：
- JDO -- JdoTemplate  5.0之后不支持了
- JTA 事务
- JPA  注解方式支持
- Java Caching API 

**Bean技术相关**
- JMS  -- JmsTemplate
- EJB
- Dependency Injection For Java 依赖注入 spring2.5+
- Bean Validation 3.0+

**相关资源**
-[小马哥JSR收藏](https://github.com/mercyblitz/jsr)
- [JSR官方网站](https://jcp/org)


## 编程模型

### 面向对象编程

- 契约接口
- 设计模式： 
  - 观察者模式 EventObject  
  - 组合模式  Composite
  - 模板模式  Template  JdbcTemplate
  - 对象继承  Abstract Application...

### 面向切面编程

#### 动态代理

JdkDynamic
CglibAopProxy 

#### 字节码提升

### 面向元编程

#### 配置元信息
Environment.java
PropertySource.java

#### 注解
模式注解
- @Component
- @Repository
- @Service

#### 泛型

GenericTypeResolver
TypeResolver
ResolvableType

### 面向模块编程

#### Maven Artifacts

#### Java 9 Automatic Modules

#### Spring @Enable*注解
@EnableCache
@EnableMvc
@Enable   激活

### 面向函编程

- FunctionalInterface.java
- 函数接口：ApplicationEventPublisher
- Reactive： Spring WebFlux

#### Lambada

#### Reactive
异步非阻塞

## Spring 核心价值

- 生态系统
  - Spring Boot
  - Spring Cloud
  - Spring Security
  - Spring Data
  - 其他
- API抽象
  - AOP抽象
  - 事务抽象
  - Environment抽象
  - 生命周期
- 编程模型
  - 面向对象
    - 契约接口
  - 面向切面
    - 动态代理
    - 字节码提升
  - 面向元
    - 配置元信息
    - 注解
  - 面向函数
    - Lambda
    - Reactive
  - 面向模块
- 设计思想
  - OOP
  - IoC/DI
  - DDD
  - TDD
  - EDP
  - FP
- 设计模式
  - 专属模式
    - 前缀模式
      - Enable模式
      - Configurable
    - 后缀模式
      - 处理模式
        - Processor
        - Resolver
        - Handler
      - 意识模式
        - Aware
      - 配置器模式
        - Configuror
      - 选择器模式
        - ImprotSelector
  - 传统GoF23
    - 创建模式
    - 结构模式
    - 行为模式
      - 责任链
- 用户基础
  - Spring用户
    - Spring Framework
    - Spring Boot
    - Spring CLoud
  - 传统用户
    - Java SE
    - Java EE

## 面试题

- 沙雕面试题
  - 什么是Spring Framework
    - 企业应用
    - 易用
    - 提供很多好特性
- 996面试题
  - Spring有哪些重要的模块
    - core:资源管理、泛型处理
    - beans： 依赖查找、依赖注入
    - aop： 动态代理、AOP字节码提升
    - context： 事件驱动、注解驱动、模块驱动
    - expression： 表达式语言模块
- 劝退题
  - Spring Framework的优势和不足是什么？ 贯穿整个系列，慢慢补充

---

# 二、 IoC容器

## 重新认识IoC

### 1. IoC发展简介
- 1983年 好莱坞原则： 演员不用去找导演，导演会联系演员
- 1988 控制反转
- 1996 控制反转命名为好莱坞原则
- 2004 Martin Fowler 提出了自己对IoC和DI的一些理解
- 2005 Martin Fowler 对IoC做出进一步说明

### 2. IoC主要实现策略
- service locator pattern
- 依赖查找  lookup
- 依赖注入  injection
  - 构造器注入
  - 参数注入
  - Setter injection
  - 接口注入
-  contextualized lookup
- 模板方法设计模式 template
- 策略模式


### 3. IoC容器的职责
IoC遵循下面几个原则：
- 实现与任务运行之间解耦
- 关注设计目标模块而不是实现
- To prevent side effects when replacing a module.
- To free modules from assumptions about how other systems do what they do and instead rely on contracts.
>"好莱坞原则"： 不要打给我们，我们会打给你。   我们：需要的东西和资源    你：系统、模块

**职责**：
- 通用职责
- 依赖处理
  - 依赖查找
  - 依赖注入
- 生命周期管理
  - 容器
  - 托管资源（Java Beans 或其他资源）
- 配置
  - 容器
  - 外部化配置
  - 托管的资源Java Beans 或其他资源）

### 4. IoC容器的实现
- Java SE
  - Java Beans
  - Java ServiceLoader SPI
  - JNDI（Java Naming and Directory Interface）
- Java EE
  - EJB(Enterprise Java Beans)
  - Servlet
- 开源
  - Apache Avalon
  - PicoContainer
  - Googel Guice
  - Spring Framework


### 5. 传统IoC容器实现
- Java Beans 作为IoC容器
- 特性
  - 依赖查找
  - 声明周期管理
  - 配置元信息
  - 事件
  - 自定义
  - 资源管理
  - 持久化
- 规范
  - JavaBeans： https://www.oracle.com/technetwork/java/javase/tech。。。
  - BeanContext：

**什么是JavaBeans**：
PropertyEditor


### 6. 轻量级IoC容器
- 管理到我的应用代码，控制启停生命周期
- 快速启动
- 不需要一些特殊配置， 不像EJB需要大量XML
- 轻量级内存占用，少量的API。 EJB需要大量的API
- 提供管控的渠道 


**好处**：
- 执行层面和实现的解耦
- 最大化代码复用
- 更大化的面向对象
- 更大化的产品化
- 更好的可测性

### 7. 依赖查找 VS 依赖注入
**优劣对比**：

类型 | 依赖处理 | 实现便利性 | 代码入侵性 | API依赖性 | 可读性
---------|----------|---------|---------|----------|---------
 依赖查找 | 主动获取 | 相对繁琐 |  侵入业务逻辑 | 依赖容器API | 良好
 依赖注入 | 被动提供 | 相对便利 |  低侵入性 | 不依赖容器API | 一般


### 8. 构造器注入 VS Setter注入
- 鼓励构造器注入。 不变的对象，确保对象不为空。
- Setter注入可选。 让对象更可配

Setter注入优点：
- JavaBeans properties 更好的支持 IDEs
- JavaBeans 属性是自文档的

setter缺点：
- Setter没有顺序
- 不是所有的setter方法都是必须的

构造器优势：
- 字段赋值，鼓励对象是不变的，final修饰。
 
**参考书**： 《Expert One-On-One J2EE Development without EJB》倾向于setter，spring官方文档倾向于 构造器注入

### 9. 面试题精选
- 沙雕面试题： 什么是IoC？
> 反转控制，类似于好莱坞原则，主要有依赖查找，依赖注入（构造器、setter）。 推的模式。

- 996面试题： **依赖查找和依赖注入**的**区别**？
> 依赖查找是主动或手动的依赖查找方式，通常需要依赖容器或标准API实现。 而依赖注入则是手动或自动依赖绑定的方式，无需依赖特性的容器和API。

- 劝退面试题：** Spring作为IoC容器**有什么**优势**？
  - 典型的IoC管理，依赖查找，依赖注入
  - AOP抽象
  - 事务抽象
  - 事件机制
  - SPI扩展
  - 强大的第三方整合
  - 易测试性
  - 更好的面向对象


## Spring IoC容器概述


### 1. Spring IoC依赖查找
- 根据Bean名称查找
  - 实时查找
  - 延迟查找： ObjectBean
- 根据Bean类型查找
- 单个Bean对象

### 2. Spring IoC依赖注入


### Spring IoC依赖来源
- 自定义的Bean： UserRepository
- 容器内建Bean对象： Environment
- 容器内建依赖 ： BeanFacotory

### 3. Spring IoC配置元信息
- Bean定义配置
  - 基于XML文件
  - 基于Properties文件
  - 基于Java注解
  - 基于Java API
- IoC容器配置
  - 基于XML文件
  - 基于Java注解
  - 基于Java API
- 外部化属性配置
  - 基于Java注解 @Value

### 4. Spring IoC容器 底层
**BeanFactory 和ApplicationContext谁才是Spring IoC容器？**
- BeanFactory是底层的IoC容器。
- ApplicationContext通过组合方式引入了BeanFactory的实现，提供更多企业级特性（更好的跟AOP集成，消息资源处理、事件发布），是超集。

### 5. Spring应用上下文
ApplicationContext除了IoC容器角色，还提供：
- AOP
- 配置元信息（Configuration Metadata）
- 资源管理（Resources）
- 事件（Event）
- 国际化
- 注解
- Environment抽象

### 6. 使用Spring IoC容器
- BeanFactory是Spring底层IoC容器
- ApplicationContext是具备应用特性的BeanFactory的超集

### 7. Spring IoC容器生命周期
- 启动 :refresh()
  - prepareRefresh
  - prepareBeanFactory  
- 运行
- 停止

### 8. 面试题
- 沙雕面试题： 什么是Spring IoC容器？
  - DI是IoC实现的一种，原则。依赖查找已经移除了
  - 伴随很多依赖
- 996面试：BeanFactory 和 FactoryBean
  - BeanFactory是IoC底层容器
  - FactoryBean是创建Bean的一种方式，帮助实现复杂的初始化逻辑
- 劝退： Spring IoC容器启动时做了哪些准备？
  - AbstractApplicationContext.java
  - IoC配置元信息读取和解析
  - IoC容器声明周期
  - Spring事件发布
  - 国际化


## Spring IoC依赖查找

### 1. 依赖查找的前世今生

### 2. 单一类型依赖查找

### 3. 集合类型依赖查找

### 4. 层次性依赖查找

### 5. 延迟性依赖查找

### 6. 安全依赖查找
不太理解安全性的意思：
- 根据类型查找，如果有多个同类型的Bean，会报错，这就是不安全？？
- 
  
### 7. 内建可查找的依赖


Bean名称 | Bean实例 | 使用场景
---------|----------|---------
 environment | Environment对象 | 外部化配置以及Profiles
 systemProperties | java.util.Properties对象 | Java系统属性
 systemEnvironment | B3 | C3
 systemEnvironment | B3 | C3
systemEnvironment | B3 | C3
 systemEnvironment | B3 | C3

 **注解驱动Spring应用上下文内建可查找的依赖**

### 8. 依赖查找中的经典异常
**BeanException**子类型

异常类型 | 触发条件 | 场景距离 
---------|----------|---------
NoSuchBeanDefinitionException | B1 | C1
NoUniqueBeanDefinitionException  | B2 | C2
BeanInstantiationException | B3 | C3
BeanCreateException | B3 | C3
 BeanDefinitionStoreException | B3 | C3

### 9. 面试题精选
- 沙雕面试题：ObjectFactory与BeanFactory的区别？
  - 两者都提供依赖查找的能力
  - 不过ObjectFactory仅关注一个或者一种类型的Bean依赖查找，并且自身不具备依赖查找的能力，能力则由BeanFactory输出
  - BeanFactory则提供了单一类型、集合类型以及层次性等多种依赖查找方式
- 996面试题：BeanFactory.getBeab操作是否线程安全？
  - 是线程安全的。举例用DefaultListableBeanFactory
- 劝退面试题：Spring依赖查找与注入在来源上的区别？

## 依赖注入

### 1. 依赖注入的模式和类型
**手动模式**：配置或者编程的方式，提前安排注入规则
- XML 资源配置元信息
- Java注解配置元信息
- API配置元信息

**自动模式**：实现方式提供依赖自动关联的方式，按照内建的注入规则
- Autowiring

**依赖注入类型**：
- Setter方法： <property name="user" ref="userBean">
- 构造器：<constructor-arg name="user" ref="userBean">
- 字段： `@Autorwire User user`
- 方法： `@Autorwire public void user(User user){...}`
- 接口回调： `class MyBean implements BeanFactoryAware{...}`

### 2. 自动绑定（AutoWiring）
**优点**：
- 减少属性、构造器参数的配置
- 更新绑定，引用传递


### 3. 自动绑定模式
模式分类：
- **no**： 默认值，未激活Autowiring， 需要手动指定依赖注入对象
- **byName**： 根据被注入属性的名称作为Bean名称进行依赖查找，并将对象设置到该属性
- **ByType**： 根据被注入属性的类型作为依赖类型查找，并将对象设置到该属性
- **constructor**： 特殊byType类型，用于构造器参数

### 4. 自动绑定限制和不足
看官方文档：
- 精确依赖会覆盖自动
- 不能绑定一些简单的类型，原生类型。可以用@Value
- 
- 不唯一就会报异常，可以加 primary

### 5. Setter方法依赖注入
实现方法：
- 手动模式
  - xml
  - 注解
  - api
- 自动模式
  - byType
  - byName

### 6. 构造器依赖注入
- 手动模式
  - xml： <constructor-arg ref="user">
  - 注解： @Bean new User（xxx）
  - api： 是有顺序的
- 自动模式
  - autowire="constructor"
  

### 7. 字段注入
- 手动模式
  - Java注解配置元信息
  - @Autowire
    - 会忽略掉静态字段，不会注入
  - @Resource
  - @Inject（可选）

### 8. 方法注入
不是注入方法。。。
- 手动模式
  - Java注解配置元信息
  - @Autowire
    - 会忽略掉静态字段，不会注入
  - @Resource
  - @Inject（可选）
  - @Bean

### 9. 回调注入

### 10. 依赖注入类型选择

### 11. 基础类型注入
- 原生类型：int
- 标量类型：enum
- 常规类型：Object、String
- Spring类型： 

### 12. 集合类型注入
- 数组类型（Array）：原生类型、标量类型、常规类型、Spring类型
- 集合类型（collection）：
  - Collection：List、Set（SortedSet、NavigableSet、EnumSet）
  - Map：Properties

### 13. 限定注入
- `@Qualifier`
- 

### 14. 延迟依赖注入

### 15. 依赖处理注入

### 16. @Autowire注入原理

### 17. JSR-330@Inject注入原理

### 18. Java通用注解注入原理

### 19. 自定义依赖注入注解

### 20. 面试题精选

# 三、 Bean

## Spring Bean基础

### 1. 定义Spring Bean 
BeanDefinition: 是Spring Framwork 中定义Bean的 **配置元信息接口**，包含：
- Bean的类名
- Bean行为配置元素，如作用域、自动绑定的模式、生命周期回调等
- 其他Bean引用，又可称为合作者（Collaborators）或者依赖（Dependencies）
- 配置设置，比如Bean配置（Properties）
 
### 2. BeanDefinition元信息
- Class： Bean全类名，必须是具体类， 不能用抽象类或接口
- Name： Bean的名称或者ID
- Scope： Bean的作用域（singleton、prototype）
- Constructor arguments：构造器参数， 用于依赖注入
- Properties： 属性设置，用于依赖注入
- AUtoWiring mode： 自动绑定模式，byName byType
- Lazy initialization mode： 延迟初始化模式
- Initialization method： 初始化回调方法名称
- Destruction method： 销毁回调方法名称


BeanDefinition如何构建？
- 通过BeanDefinitionBuilder
- 通过AbstractBeanDefinition以及派生类


### 3. 命名Spring Bean
Bean的名称：
- 允许出现特殊字符，比如.
- 如果想要引入别名，可以在name属性使用英文逗号或者分号来间隔
- id非必填，留空的话，容器会自动生成唯一的名称


Bean的名称生成器，因为id、name非必填。 注解式的一般都不命名，使用默认生成的。 XML一般都命名居多。

### 4. Spring Bean的别名
别名的价值场景：
- 在不同的系统叫不同的名称

### 5. 注册Spring Bean
- XML配置元信息
  - <bean name=".." ../>
- Java注解配置元信息
  - @Bean
  - @Component
  - @Import()
- Java API配置元信息
  - 命名方式：BeanDefinitionRegistry#registerBeanDefinition(String, BeanDefinition)
  - 非命名方式: BeanDefinitionReaderUtil#registerWithGeneratedName(AbstractBeanDefinition, BeanDifinitionRegistey)
  - 配置类方式： AnnotatedBeanDefinitionReader#register(Class...)

### 6. 实例化Spring Bean
- 常规方式
  - 通过构造器 (配置元信息：XML、Java注解和Java API) 
  - 静态工厂方法(配置元信息：XML、Java API) 
  - Bean工厂方法(配置元信息：XML、Java API) 
  - FactoryBean(配置元信息：XML、Java API) 
- 特殊方式
  - 通过ServiceLoaderFactoryBean(配置元信息：XML、Java注解和Java API) 
  - 通过AutowireCapableBeanFactory#createBean(Class, int, boolean)
  - 通过BeanDefinitionRegistry#registerBeanDefinition(String, BeanDefinition)

### 7. 初始化Spring Bean
Initialization
- @PostConstruct标注方法
- 实现Initialization接口的afterPropertiesSet()方法
- 自定义初始化方法
  - XML配置：<bean init-method="init" ../>
  - Java注解：@Bean(initMethod="init)
  - Java API:AbstractBeanDefinition#setInitMethodName(String)

>如果同时出现，那顺序是 PostConstruct > Initialization > 自定义

### 8. 延迟初始化Spring Bean
Lazy Initialization
- XML配置： <bean lazy-init="true" ../>
- Java注解： @Lazy(true)

>当某个Bean定义为延迟初始化，那么。Spring容器返回的对象与非延迟的对象存在怎样的差异？ 延迟加载是在上下文初始化之后进行加载的， 非延迟加载是在上下文初始化之前加载

### 9. 销毁Spring Bean
Destroy
- @PreDestory标注方法
- 实现DisposableBean接口的destroy()方法
- 自定义初始化方法
  - XML配置：<bean destroy="destroy" ../>
  - Java注解：@Bean(destroy="destroy")
  - Java API:AbstractBeanDefinition#setDestroyMethodName(String)

>如果同时出现，那顺序是 @PreDestory > DisposableBean > 自定义

### 10. 垃圾回收Spring Bean
1. 关闭Spring容器（应用上下文）
2. 执行GC
3. Spring Bean覆盖finalize()方法回调

### 11. 面试题精选
1. 沙雕面试题：如何注册一个Spring Bean
   1. 通过BeanDefinition和外部单体对象注册
2. 996面试题：什么是Spring BeanDefinition？
   1. 有很多属性， scope、role、primary、各种元信息接口
3. 劝退面试题：Spring容器是怎样管理注册Bean？
   1. IoC配置、依赖注册、依赖查找、生命周期等

## Bean实例

## Bean作用域

## Bean生命周期

# 四、 元信息

## 注解

## 配置源信息

## 外部化属性

# 五、 基础设施

## 资源管理

## 类型转换

## 数据绑定

## 校验

## 国际化

## 事件

## 泛型处理