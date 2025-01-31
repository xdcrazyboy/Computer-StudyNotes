# 1 引言
Spring Cloud 是一系列框架的有序集合。它利用 Spring Boot 的开发便利性巧妙地简化了分布式系统基础设施的开发，如服务发现注册、配置中心、消息总线、负载均衡、断路器、数据监控等，都可以用 Spring Boot 的开发风格做到一键启动和部署。
## 参考文献
--------------------- 
- [方志朋Github ](https://github.com/forezp/SpringCloudLearning?utm_source=gold_browser_extension)，强烈推荐大佬的博客，我这基本算是读书笔记。

## 注意事项
1. cloud版本选择的时1.5那个，用spring.io自动生成一定要时刻注意选择统一版本。
2. 依赖默认都用了netflix的，有些教程时去掉这个的，区别大家可以自行百度：
```xml
<dependency>
	<groupId>org.springframework.cloud</groupId>
	<artifactId>spring-cloud-starter-netflix-zuul</artifactId>
</dependency>
		```
# 2 Eureka 服务注册发现

## 2.1 引言
Eureka 是 Netflix 的子模块，它是一个基于 REST 的服务，用于定位服务，以实现云端中间层服务发现和故障转移。

Eureka 采用了 CS 的设计架构。Eureka Server 作为服务注册功能的服务端，它是服务注册中心。而系统中其他微服务则使用 Eureka Client连接到 Eureka Server 并维持心跳连接（服务提供者和消费者都属于Client）。

## 2.2 Eureka Server
Eureka Server 提供服务的注册服务。各个服务节点启动后会在 Eureka Server 中注册服务，Eureka Server 中的服务注册表会存储所有可用的服务节点信息。

## 2.3 Eureka Client
Eureka Client 是一个 Java 客户端，用于简化 Eureka Server 的交互，客户端同时也具备一个内置的、使用轮询负载算法的负载均衡器。在应用启动后，向 Eureka Server 发送心跳（默认周期 30 秒）。如果 Eureka Server 在多个心跳周期内没有接收到某个节点的心跳，Eureka Server 会从服务注册表中将该服务节点信息移除。

## 2.4 搭建注册中心
使用Spriong boot的start.io创建工程，模板选择Cloud Discovery——>Eureka Server。
### 2.4.1 添加依赖
``` xml
<dependencyManagement>
  	<dependencies>
  		<dependency>
			<groupId>org.springframework.cloud</groupId>
			<artifactId>spring-cloud-dependencies</artifactId>
			<version>Greenwich.RELEASE</version>
			<type>pom</type>
			<scope>import</scope>
		</dependency>
		
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-parent</artifactId>
			<version>2.1.3.RELEASE</version>
			<type>pom</type>
			<scope>import</scope>
		</dependency>
  	</dependencies>
</dependencyManagement>
  
<dependencies>
	<!-- eureka 服务端 -->
	<dependency>
		<groupId>org.springframework.cloud</groupId>
		<artifactId>spring-cloud-starter-netflix-eureka-server</artifactId>
	</dependency>
</dependencies>
```
还有一种服务器依赖：
```xml
<!--eureka server -->
<dependency>
	<groupId>org.springframework.cloud</groupId>
	<artifactId>spring-cloud-starter-eureka-server</artifactId>
</dependency>
```

### 2.4.2 配置参数apllication.yml或者application.properties
eureka是一个高可用的组件，它没有后端缓存，每一个实例注册之后需要向注册中心发送心跳（因此可以在内存中完成），在默认情况下erureka server也是一个eureka client ,必须要指定一个 server。eureka server的配置文件application.yml：

```yml
server:
    port: 9500

spring:
  application:
    name: eureka-server
    
eureka:
    instance:
        hostname: localhost   # eureka 实例名称
    client:
        register-with-eureka: false # 不向注册中心注册自己
        fetch-registry: false       # 是否检索服务
        service-url:
            defaultZone: http://${eureka.instance.hostname}:${server.port}/eureka/  # 注册中心访问地址
```

### 2.4.3 开启注册中心
在启动类上添加`@EnableEurekaServer`注解
```java
@EnableEurekaServer
@SpringBootApplication
public class EurekaServerApplication {

	public static void main(String[] args) {
		SpringApplication.run(EurekaServerApplication.class, args);
	}
}
```

然后就启动项目，访问http://localhost:9500,可以查看Eureka服务监控界面。
>http://localhost:9500 是 Eureka 监管界面访问地址，而 http://localhost:9500/eureka/ Eureka 注册服务的地址。

## 2.5 创建一个服务提供者（Eureka Client）
当client向server注册时，它会提供一些元数据，例如主机和端口，URL，主页等。Eureka server 从每个client实例接收心跳消息。 如果心跳超时，则通常将该实例从注册server中删除。

1. 创建过程同server类似,创建完pom.xml如下：
```xml
<!-- 其主要区别在于下面这两个依赖 -->
<!--eureka server -->
<dependency>
	<groupId>org.springframework.cloud</groupId>
	<artifactId>spring-cloud-starter-eureka-server</artifactId>
</dependency>
<!--client端 -->

<dependency>
	<groupId>org.springframework.cloud</groupId>
	<artifactId>spring-cloud-starter-eureka</artifactId>
</dependency>
```

2. 注解改为`@EnableEurekaClient`表明自己是一个client。

3. 需要在配置文件中注明自己的服务注册中心的地址，application.yml配置文件如下：
```yml
server:
  port: 9400
spring:
  application:
	## 这个很重要，以后的服务与服务之间相互调用一般都是根据这个name 。
    name: eureka-provider
eureka:
  client:
    serviceUrl:
      defaultZone: http://localhost:9500/eureka/
```
controller包中的类：
```java
@RestController
public class MessageController {

    @Autowired
    MessageService messageService;

    @Value("${server.port}")
    String port;

    @RequestMapping("/get")
    public String getMessage(String name){
		//getMessage(name) ->This meassage is from ${name}!
        return "provider提供信息："+ messageService.getMessage(name) + "port:" + port;
    }

}
```

## 2.6 高可用的服务注册中心
服务注册中心Eureka Server，是一个实例，当成千上万个服务向它注册的时候，它的负载是非常高的，这在生产环境上是不太合适的，这篇文章主要介绍怎么将Eureka Server集群化。

当有服务注册时，两个Eureka-eserver是对等的，它们都存有相同的信息，这就是**通过服务器的冗余来增加可靠性**，当有一台服务器宕机了，服务并不会终止，因为另一台服务存有相同的数据。


# 3. 服务消费者 Ribbon （负载均衡、服务调用）
Spring cloud有两种服务调用方式，一种是ribbon+restTemplate，另一种是feign，Feign默认集成了ribbon。

## 搭建小集群
启动eureka-server 工程；启动service-hi工程，它的端口为8762；将service-hi的配置文件的端口改为8763,并启动，这时你会发现：service-hi在eureka-server注册了2个实例，这就相当于一个小的集群。例如改成9401,在启动按钮哪里修改Edit Configurations进去找到：选择该boot的启动类，右上角勾选Allow Parallel，出来后，启动按钮会重新出现，启动一下还有一个小数字（Idea）。

然后访问server地址，会看到两个ip同一个服务名称。

## 建立一个服务消费者
它的pom.xml文件分别引入起步依赖spring-cloud-starter-eureka、spring-cloud-starter-ribbon、spring-boot-starter-web：
```xml
<properties>
        <java.version>1.8</java.version>
        <spring-cloud.version>Edgware.SR6</spring-cloud.version>
    </properties>

    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-ribbon</artifactId>
        </dependency>

        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
    </dependencies>

    <dependencyManagement>
        <dependencies>
            <dependency>
                <groupId>org.springframework.cloud</groupId>
                <artifactId>spring-cloud-dependencies</artifactId>
                <version>${spring-cloud.version}</version>
                <type>pom</type>
                <scope>import</scope>
            </dependency>
        </dependencies>
    </dependencyManagement>

    <build>
        <plugins>
            <plugin>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-maven-plugin</artifactId>
            </plugin>
        </plugins>
    </build>
```
一些坑：
	- 注意cloud的版本号！

application.yml中端口换一个，名字取一个，服务器注册中心地址保持一致：
```yml
server:
  port: 8764
spring:
  application:
    name: service-ribbon
```

## 启动类中
在工程的启动类中,通过@EnableDiscoveryClient向服务中心注册；并且向程序的ioc注入一个bean: restTemplate;并通过@LoadBalanced注解表明这个restRemplate开启负载均衡的功能。
```java
@SpringBootApplication
@EnableDiscoveryClient
public class ServiceRibbonApplication {

	public static void main(String[] args) {
		SpringApplication.run(ServiceRibbonApplication.class, args);
	}

	@Bean
	@LoadBalanced
	RestTemplate restTemplate() {
		return new RestTemplate();
	}

}

```

## 创建Service和Controller去调用服务
主要核心代码是：
service包中：
```java
@Service
public class HelloService {

    @Autowired
    RestTemplate restTemplate;

    public String hiService(String name) {
        return restTemplate.getForObject("http:///eureka-provider/get?name="+ name, String.class);
    }

}
```
controller包中：
```java
@RestController
public class HelloControler {

    @Autowired
    HelloService helloService;

    @RequestMapping(value = "/hi")
    public String hi(@RequestParam String name){
        return helloService.hiService(name);
    }
}
```
在浏览器上多次访问http://localhost:9600/hi?name=bobo4，浏览器交替显示：
```
provider提供信息：This meassage is from bobo4!port:9400
provider提供信息：This meassage is from bobo4!port:9401
```

# 服务消费者 Feign
## 引言-简介
Feign是一个声明式的伪Http客户端，它使得写Http客户端变得更简单。使用Feign，只需要创建一个接口并注解。它具有可插拔的注解特性，可使用Feign 注解和JAX-RS注解。Feign支持可插拔的编码器和解码器。Feign默认集成了Ribbon，并和Eureka结合，默认实现了负载均衡的效果。

## 创建一个feign服务
1. 新建一个spring-boot工程，pom：
```xml
		<dependency>
			<groupId>org.springframework.cloud</groupId>
			<artifactId>spring-cloud-starter-feign</artifactId>
		</dependency>
```
2. application.yml：
```yml
eureka:
  client:
    serviceUrl:
      defaultZone: http://localhost:9500/eureka/
server:
  port: 9700
spring:
  application:
    name: service-feign

```

3. 启动类ServiceFeignApplication ，加上@EnableFeignClients注解开启Feign的功能:
```java
@SpringBootApplication
@EnableDiscoveryClient
@EnableFeignClients
public class ServiceFeignApplication {

	public static void main(String[] args) {
		SpringApplication.run(ServiceFeignApplication.class, args);
	}
}
```

4. 定义一个feign接口，通过@ FeignClient（“服务名”），来指定调用哪个服务。比如在代码中调用了service-hi服务的“/hi”接口，代码如下：
```java
//这个value的值时所要调用的服务配置文字里的名称
@FeignClient(value = "service-hi")
public interface SchedualServiceHi {
	//这个hi就是所要调用的映射的词，/hi接口
    @RequestMapping(value = "/hi",method = RequestMethod.GET)
    String sayHiFromClientOne(@RequestParam(value = "name") String name);
}

```

5. 最核心的不同：在Web层的controller层，对外暴露一个"/hi"的API接口，通过上面定义的Feign客户端SchedualServiceHi 来消费服务。
```java
@RestController
public class HiController {

    @Autowired
    SchedualServiceHi schedualServiceHi;

    @RequestMapping(value = "/hi",method = RequestMethod.GET)
    public String sayHi(@RequestParam String name){
        return schedualServiceHi.sayHiFromClientOne(name);
    }
}

```
6. 启动即可

# 断路器（Hystrix）解决服务雪崩现象（通过断路，降级）
一般一个业务需要调用很多服务，如果某个服务崩了（比如积分增加服务），如果不处理可能就会让整个订单服务卡在那。

较底层的服务如果出现故障，会导致连锁故障。当对特定的服务的调用的不可用达到一个阀值（Hystric 是5秒20次） 断路器将会被打开。断路打开后，可用避免连锁故障，fallback方法可以直接返回一个固定值。

## 在ribbon中使用断路器
1. 改造serice-ribbon 工程的代码，首先在pox.xml文件中加入spring-cloud-starter-hystrix的起步依赖：
```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-hystrix</artifactId>
</dependency>

```
1. 启动类：加@EnableHystrix注解开启Hystrix。@EnableDiscoveryClient也一直要有。
   
2. 改造HelloService类，在hiService**方法上**加上@HystrixCommand注解。该注解对该方法创建了熔断器的功能，并指定了fallbackMethod熔断方法，熔断方法直接返回了一个字符串，字符串为"hi,"+name+",sorry,error!"，代码如下：
```java
@Service
public class HelloService {

    @Autowired
    RestTemplate restTemplate;

    @HystrixCommand(fallbackMethod = "hiError")
    public String hiService(String name) {
        return restTemplate.getForObject("http://eureka-provider/get?name="+name,String.class);
    }

    public String hiError(String name) {
        return "hi,"+name+",sorry,error!";
    }
}

```
1. 测试：启动：service-ribbon 工程，当我们访问http://localhost:9600/hi?name=bobo,浏览器显示：
>provider提供信息：This meassage is from bobo!port:9401

此时关闭 service-hi 工程，当我们再访问http://localhost:9500/hi?name=bobo，浏览器会显示：
>hi ,forezp,orry,error!

## 在Feign中使用断路器
1. Feign是自带断路器的，在D版本的Spring Cloud中，它没有默认打开。需要在配置文件中配置打开它，在配置文件加以下代码：
>feign.hystrix.enabled=true
别忘记再pom.xml文件添加依赖：
```xml
<dependency>
   <groupId>org.springframework.cloud</groupId>
   <artifactId>spring-cloud-starter-hystrix</artifactId>
</dependency>
```

2. 基于service-feign工程进行改造，只需要在FeignClient的SchedualServiceHi接口的注解中加上fallback的指定类就行了：
```java
@FeignClient(value = "service-hi",fallback = SchedualServiceHiHystric.class)
public interface SchedualServiceHi {
    @RequestMapping(value = "/hi",method = RequestMethod.GET)
    String sayHiFromClientOne(@RequestParam(value = "name") String name);
}

```
3. SchedualServiceHiHystric需要实现SchedualServiceHi 接口，并注入到Ioc容器中，代码如下：
```java
@Component
public class SchedualServiceHiHystric implements SchedualServiceHi {
    @Override
    public String sayHiFromClientOne(String name) {
        return "sorry "+name;
    }
}
```

## Hystrix Dashboard (断路器：Hystrix 仪表盘)
基于service-ribbon 改造，Feign的改造和这一样。

1. 首选在pom.xml引入spring-cloud-starter-hystrix-dashboard的起步依赖：
```xml
<dependency>
	<groupId>org.springframework.boot</groupId>
	<artifactId>spring-boot-starter-actuator</artifactId>
</dependency>

<dependency>
	<groupId>org.springframework.cloud</groupId>
	<artifactId>spring-cloud-starter-hystrix-dashboard</artifactId>
</dependency>

```
2. 主程序启动类中加入@EnableHystrixDashboard注解，开启hystrixDashboard：
```java
@SpringBootApplication
@EnableDiscoveryClient
@EnableHystrix
@EnableHystrixDashboard
public class ServiceRibbonApplication {

	public static void main(String[] args) {
		SpringApplication.run(ServiceRibbonApplication.class, args);
	}

	@Bean
	@LoadBalanced
	RestTemplate restTemplate() {
		return new RestTemplate();
	}

}
```

3. 打开浏览器：访问http://localhost:9500/hystrix
4. 点击monitor stream，进入下一个界面，访问：http://localhost:9500/hi?name=bobo


# Zuul 路由转发和过滤器
## 引言
Zuul的主要功能是路由转发和过滤器。路由功能是微服务的一部分，比如／api/user转发到到user服务，/api/shop转发到到shop服务。zuul默认和Ribbon结合实现了负载均衡的功能。

## 搭建测试路由功能
1. 在原有的工程上，创建一个新的工程。
2. zuul需要增加一个依赖：
```xml
<dependency>
	<groupId>org.springframework.cloud</groupId>
	<artifactId>spring-cloud-starter-zuul</artifactId>
</dependency>
```

3. 其入口applicaton类加上注解@EnableZuulProxy，开启zuul的功能：

4. 配置文件application.yml加上以下的配置代码：
```xml
eureka:
  client:
    serviceUrl:
      defaultZone: http://localhost:9500/eureka/
server:
  port: 9800
spring:
  application:
    name: service-zuul
zuul:
  routes:
    api-a:
      path: /api-a/**
      serviceId: service-ribbon
    api-b:
      path: /api-b/**
      serviceId: service-feign
```

5. 测试:以/api-a/ 开头的请求都转发给service-ribbon服务；以/api-b/开头的请求都转发给service-feign服务；
依次运行这五个工程;打开浏览器访问：http://localhost:9800/api-a/hi?name=bobo ;浏览器显示：
>provider提供信息：This meassage is from bobo!port:9400
打开浏览器访问：http://localhost:9800/api-b/hi?name=bobo ;浏览器显示：
>provider提供信息：This meassage is from bobo!port:9400
这说明zuul起到了路由的作用,地址不同访问的服务就不同，不过这里两个服务内容一致所以一样。可以试着变动一下输出内容。

## 服务过滤
过滤，做一些安全验证

在service或者启动类的文件夹创建一个过滤器：
```java
@Component
public class MyFilter extends ZuulFilter{

    private static Logger log = LoggerFactory.getLogger(MyFilter.class);
    @Override
    public String filterType() {
        return "pre";
    }

    @Override
    public int filterOrder() {
        return 0;
    }

    @Override
    public boolean shouldFilter() {
        return true;
    }

    @Override
    public Object run() {
        RequestContext ctx = RequestContext.getCurrentContext();
        HttpServletRequest request = ctx.getRequest();
        log.info(String.format("%s >>> %s", request.getMethod(), request.getRequestURL().toString()));
        Object accessToken = request.getParameter("token");
        if(accessToken == null) {
            log.warn("token is empty");
            ctx.setSendZuulResponse(false);
            ctx.setResponseStatusCode(401);
            try {
                ctx.getResponse().getWriter().write("token is empty");
            }catch (Exception e){}

            return null;
        }
        log.info("ok");
        return null;
    }
}
- filterType：返回一个字符串代表过滤器的类型，在zuul中定义了四种不同生命周期的过滤器类型，具体如下：
  - pre：路由之前
  - routing：路由之时
  - post： 路由之后
  - error：发送错误调用
- filterOrder：过滤的顺序
- shouldFilter：这里可以写逻辑判断，是否要过滤，本文true,永远过滤。
- run：过滤器的具体逻辑。可用很复杂，包括查sql，nosql去判断该请求到底有没有权限访问。

这时访问：http://localhost:9800/api-a/hi?name=bobo ；网页显示：
>token is empty
访问 http://localhost:9800/api-a/hi?name=bobo&token=12 ；网页显示：
>provider提供信息：This meassage is from bobo!port:9400
```

# 分布式配置中心(Spring Cloud Config)
## 引言
在分布式系统中，由于服务数量巨多，为了方便服务配置文件统一管理，实时更新，所以需要分布式配置中心组件。在Spring Cloud中，有分布式配置中心组件spring cloud config ，它支持配置服务放在配置服务的内存中（即本地），也支持放在远程Git仓库中。在spring cloud config 组件中，分两个角色，一是config server，二是config client。

## 如何配置

##  高可用的分布式配置中心(Spring Cloud Config)
搭建一个Eureka服务注册中心，专门用于对配置中心，配置client进行发现管理。


# 其他
## 消息总线(Spring Cloud Bus)

## 服务链路追踪(Spring Cloud Sleuth)

