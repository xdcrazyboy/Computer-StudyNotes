# Spring从入门到放弃

# 目录

[TOC]

# 前言

## Spring简介
Spring是一个轻量级的Java开发框架，为了解决企业开发的复杂性而创建的。核心是**控制反转**（IoC）和面向切面编程（AOP）。

主要作用：**为代码解耦**，降低代码间的耦合度。
>根据功能不同，可以将一个系统中的代码为：**主业务逻辑**和**系统业务逻辑**。 主业务逻辑代码逻辑联紧密，耦合度相对较高，复杂性较低。 系统级业务，相对功能独立，耦合度低， 主要是为主业务提供系统级服务，比如日志、安全、事务等，复用性强。

Spring使用IoC降低业务逻辑之间的耦合度，使得主业务在相互调用过程中，不用再自己维护关系，不用自己创建要使用的的对象。而是通过Spring容器统一管理，自动注入。

使用AOP使得系统级服务得到了最大复用，不需要程序员手工将系统级服务“混杂”到主业务逻辑中，而是由Spring容器统一完成“织入”，x相当于对主业务和系统服务之间进行解耦。 


# 1. Spring概述

## 1.1 Spring体系结构
Spring由20多个模块组成，分为：数据访问/集成(Data Access/Intergration)、AOP（Aspect Oriented Programming） 模块、Aspects模块、Instrumentation模块、Messaging模块、**Core Container**模块和Test模块。

1. **数据处理模块**（**Data Access**） 
   * `JDBC`模块提供了不需要编写冗长的JDBC代码和解析数据库厂商 特有的错误代码的JDBC-抽象层。 
   * `Transactions`模块支持编程和声明式事务管理。 
   * `ORM`模块提供了流行的Object-Relational Mapping（对象-关系映射）API集成层，包含JPA、JDO和Hibernate等ORM框架。
    >Spring对 ORM的支持和封装主要体现在三方面：
    >1. 一致的异常处理体系结构，对第三方ORM框架抛出的专有异常进行了包装；
    >2. 一致的DAO抽象的支 ，为每个框架都提供了模板类来简化和封装常用操作，例如 JdbcSupport、HibernateTemplate等；
    >3. Spring的事务管理机制，为所有数据访问都提供了一致的事务管理。 
     * OXM模块提供抽象层，用于支持Object/XML mapping（对 象/XML映射）的实现，例如JAXB、Castor、XMLBeans、JiBX和 XStream等。 
     * JMS模块（Java Messaging Service）包含生产和消费信息的功能。

2. **Web模块**
   * `Web`模块提供了面向Web开发的集成功能 
   * `WebSocket`模块提供了面向WebSocket开发的集成功能。  
   * `Servlet` 模块（也被称为SpringMVC 模块）包含 Spring 的 ModelView-Controller（模型-视图-控制器，简称MVC）和REST Web Services 实现的Web应用程序。Spring MVC框架使Domain Model（领域模型）代码和Web Form（网页）代码实现了完全分离，并且集成了Spring Framework的所有功能 
   * Portlet模块（也被称为Portlet MVC 模块）是基于Web和Servlet模 块的MVC实现。Portlet和Servlet的最大区别是对请求的处理分为Action 阶段和Render阶段。在处理一次 HTTP请求时:
     * 在 Action阶段处理业务 逻辑响应并且当前逻辑处理只被执行一次；
     * 而在Render阶段随着业务的 定制，当前处理逻辑会被执行多次，
        >这样就保证了业务系统在处理同一个业务逻辑时能够进行定制性响应页面模版渲染。 

3. **AOP模块**
4. Aspects模块
5. Instrumentation模块
6. Messaging模块
7. **Core Container（Spring核心容器模块）**
此模块是Spring的根基。
   - Beans模块和Core模块提供框架的基础部分，包含IoC（Inversion of Control，控制反转）和 DI（Dependency Injection，依赖注入）功能，使用 BeanFactory 基本概念来实现容器对Bean的管理，是所有 Spring应用的核心。Spring本身的运行都是由这种Bean的核心模型进行加载和执行的，是Spring其他模块的核心支撑，是运行的根本保证。 
   - Context（包含 Spring-Context和 Spring-Context-Support两个子模块）模块建立在Core模块和 Beans模块的坚实基础之上，并且集成了 Beans模块的特征，增加了对国际化的支持，也支持Java EE特征。 ApplicationContext接口是Context模块的焦点。Spring-Context-Support模 块支持集成第三方常用库到Spring应用上下文中，例如缓存 （EhCache、Guava）、调度Scheduling框架（CommonJ、Quartz）及模 板引擎（FreeMarker、Velocity）。 
   - SpEL模块（Spring-Expression Language）提供了强大的表达式语 言来查询和操作运行时的对象。 
## 1.2 如何安装和使用


## 1.3 特点

### 1.3.1 非侵入式

### 1.3.2 容器
Spring作为一个容器，可以管理对象的什么周期、对象与对象的

### 1.3.3 IoC

### 1.3.4 AOP 