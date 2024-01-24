# Spring MVC入门汇总

[TOC]

# 开始
## 需要导入jar包
1. Spring-context-support：包含支持UI模板，邮件服务，缓存Cache等方面的类。
2. Spring-webmvc：对SpringMVC的体现。

## 配置文件
### web.xml
1. **注册中央调度器**-DispatcherServlet
   ```xml
     <!-- servlet分发 -->
    <servlet>
        <servlet-name>springMVC</servlet-name>
        <servlet-class>org.springframework.web.servlet.DispatcherServlet</servlet-class>
        <init-param>
            <param-name>contextConfigLocation</param-name>
            <param-value>classpath*:appContext-mvc.xml</param-value>
        </init-param>
        <load-on-startup>1</load-on-startup>
    </servlet>
    <servlet-mapping>
        <servlet-name>springMVC</servlet-name>
        <url-pattern>*.action</url-pattern>
    </servlet-mapping>
   ```
   - 全限定类名：org.springframework.web.servlet.DispatcherServlet
   - \<load-on-startup> : 标记是否在Web服务器启动时就创建这个Servlet实例（执行该Servlet的init()方法），还是在真正访问时才创建。它的值时一个整数。
     - 值大于0时，表示容器启动时就加载并初始化这个Servlet，数值越小，优先级越高，被创建得越早。
     - 0或者没有指定，表示在真正使用时才创建
     - 值相同，容器会自己选择创建的顺序。
   - \<url-pattern> : `*.do`或者`*.action`。


# 我遇到的问题

## SpringMVC对于传入多个对象参数遇到的问题

**解决方法1：** 再用一个大对象进行包装：
- 解析应该没办法解析大对象里面的小对象，可以把小对象json成字符串，用的时候再转回来。

**方法2：** 通过Map<String, Object>
- key 固定为对象的一个别名，比如‘user’，表示User对象，提取方法：
```java
入参：（Map<String, Object> models）
User user = JsonXMLUtils.map2obj((Map<String, Object>) models.get('user'), User.class)
Student student = JsonXMLUtils.map2obj((Map<String, Object>) models.get('student'), Student.class)
```

**方法3**： 使用自定义注解实现json和对象的映射。 其实就是显得优雅一些，编程工作量差不多，甚至可读性还差一点。

**方法4：**扩展spring的HandlerMethodArgumentResolver以支持自定义的数据绑定方式。

### 另外一种理解
第1种方法：表单提交，以字段数组接收；
第2种方法：表单提交，以BeanListModel接收；
第3种方法：将Json对象序列化成Json字符串提交，以List接收；
第4种方法：将表单对象序列化成Json字符串提交，以List接收； 
第4种方法其实是第3种方法的升级，就是将表单转成Json对象，再转成Json字符串提交；