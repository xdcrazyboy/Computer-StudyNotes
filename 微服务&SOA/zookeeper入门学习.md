# Zookeeper入门学习

# 它是什么？
Zookeeper（**动物园管理员**，hadoop中的大多是动物名字）是一个开放源码的分布式服务协调组件，是Google Chubby的开源实现。是一个高性能的分布式数据一致性解决方案。他将那些复杂的、容易出错的分布式一致性服务封装起来，构成一个高效可靠的原语集，并提供一系列简单易用的接口给用户使用。

## 功能特性
- **数据存储**
  - zookeeper提供了类似Linux文件系统一样的数据结构。每一个节点对应一个Znode节点，每一个Znode节点都可以存储1MB（默认）的数据。
  - 客户端对zk的操作就是对Znode节点的操作。
    - Znode:包含ACL权限控制、修改/访问时间、最后一次操作的事务Id(zxid)等等
    * 说有数据存储在内存中，在内存中维护这么一颗树。
    * 每次对Znode节点修改都是**保证顺序**和原子性的操作。**写**操作是**原子性**操作。
  * 每一个Znode节点又根据节点的生命周期与类型分为4种节点。
    * 持久节点（带不带顺序编号）、临时节点（带不带顺序编号）
    * 生命周期：当客户端会话结束的时候，是否清理掉这个会话创建的节点。持久-不清理，临时-清理。
    * 类型：每一个会话，创建单独的节点（例子：正常节点：rudytan,顺序编号节点：rudytan001,rudytan002等等）
* **监听机制**
  * zookeeper除了提供对Znode节点的处理能力，还提供了对节点的变更进行监听通知的能力。
  * 监听机制的步骤如下：
    * 任何session(session1,session2)都可以对自己感兴趣的znode监听。
    * 当znode通过session1对节点进行了修改。
    * session1,session2都会收到znode的变更事件通知。
  * 节点常见的事件通知有：
    * session建立成功事件
    * 节点添加
    * 节点删除
    * 节点变更
    * 子节点列表变化
  >**特别注意**:一次监听事件，只会被触发一次，如果想要监听到znode的第二次变更，需要重新注册监听。

## 应用场景        
- **注册中心**
  - 依赖于**临时节点**
    * 消费者启动的时候，会先去注册中心中全量拉取服务的注册列表。
    * 当某个服务节点有变化的时候，通过监听机制做数据更新。
    * zookeeper挂了，不影响消费者的服务调用。

**注册中心对比**: `Eureka` VS `Zookeeper`
  - Eureka中的节点每一个节点对等。Eureka是个AP系统，而不是zk的CP系统。
  - 在注册中心的应用场景下，相对于与强数据一致性，**更加关心可用性**。

- **分布式锁**
  - 依赖于**临时顺序节点**
  * 判断当前client的**顺序号是否是最小**的，如果是获取到锁。
  * 没有获取到锁的节点**监听最小节点的删除事件**（比如lock_key_001）
  * 锁释放，最小节点删除，剩余节点重新开始获取锁。
  * 重复步骤二到四。

**分布式锁对比**： `redis` vs `数据库` vs `zk`
  - 从`理解`的难易程度角度（从低到高）
      > 数据库 > 缓存（Redis） > Zookeepe
  - 从实现的`复杂性`角度（从低到高）
      > Zookeeper >= 缓存（Redis） > 数据库
  - 从`性能`角度（从高到低）
      > 缓存（Redis） > Zookeeper >= 数据库
  - 从`可靠性`角度（从高到低）
      > Zookeeper > 缓存（Redis） > 数据库

- **集群管理与master选举**
  * 依赖于临时节点
  * zookeeper保证无法重复创建一个已存在的数据节点，创建成功的client为master。
  * 非master，在已经创建的节点上注册节点删除事件监听。
  * 当master挂掉后，其他集群节点收到节点删除事件，进行重新选举
  * 重复步骤二到四

>有人说，zookeeper可以做分布式配置中心、分布式消息队列，看到这里的小伙伴们，你们觉得合适么？

## 高性能高可用强一致性保障


- **高性能-分布式集群**：
    >高性能，我们通常想到的是通过集群部署来突破单机的性能瓶颈。对于zk来说，就是通过部署多个节点共同对外提供服务，来提供读的高性能。
  * Master/Slave模式。
  * 在zookeeper中部署多台节点对外提供服务，客户端可以连接到任意一个节点。
  * 每个节点的数据都是一样的。
  * 节点根据角色分为Leader节点与Learner节点（包括Follower节点与Observer节点）。
  * 集群中，只有一个Leader节点，完成所有的写请求处理。
  * 每次写请求都会生成一个全局的唯一的64位整型的事务ID(可以理解为全局的数据的版本号)。
  * Learner节点可以有很多，每个Leaner可以独自处理读请求，转写请求到Leader节点。
  * 当Leader节点挂掉后，会从Follower节点中通过选举方式选出一个Leader提供对外服务。
  * Follower节点与Observer节点区别在于不参与选举和提议的事务过半处理。
  * 集群通常是按照奇数个节点进行部署（偶数对容灾没啥影响，浪费机器）。

- **数据一致性**（zab协议-原子广播协议）

# 常用命令

## 删除
```

delete path [version]

```
删除指定节点数据，其version参数的作用于set指定一致

>注意：delete只能删除不包含子节点的节点，如果要删除的节点包含子节点，使用rmr命令

```s
rmr /node_1
```

# Zookeeper Curator 事件监听

## Curator 事件监听

Curator 事件有两种模式：
1. 标准的监听模式是使用Watcher 监听器。
   - 比较简单，只有一种，且**此监听操作只能监听一次**，如果要反复使用，就需要反复的使用usingWatcher提前注册。
2. 第二种缓存监听模式引入了一种本地缓存视图的Cache机制，来实现对Zookeeper服务端事件监听。
   - Cache事件监听的种类有3种**Path Cache，Node Cache，Tree Cache**
     - Node Cache节点缓存可以用于**ZNode节点的监听**
     - Path Cache子节点缓存用于ZNode的**子节点的监听**
     - ree Cache树缓存是Path Cache的增强，不光**能监听子节点，也能监听ZNode节点**自身
   - Cache事件监听可以理解为一个本地缓存视图与远程Zookeeper视图的对比过程。
   - Cache提供了反复注册的功能。
   - Cache是一种缓存机制，可以借助Cache实现监听。
   >简单来说，Cache在客户端缓存了znode的各种状态，当感知到zk集群的znode状态变化，会触发event事件，注册的监听器会处理这些事件。

### Watcher 标准的事件处理器
接口类Watcher用于表示一个标准的事件处理器，其定义了事件通知相关的逻辑，包含KeeperState和EventType两个枚举类，分别代表了通知状态和事件类型。

Watcher接口定义了事件的回调方法：process（WatchedEvent event）。

定义一个Watcher，使用Watcher监听器实例的方式，在Curator的调用链上，加上usingWatcher方法即可，代码如下：
```java
Watcher w = new Watcher() {
    @Override
    public void process(WatchedEvent watchedEvent) {
        log.info("监听器watchedEvent：" + watchedEvent);
    }
};

byte[] content = client.getData()
        .usingWatcher(w).forPath(workerPath);
```

一个Watcher监听器在向服务端完成注册后
1. 当服务端的一些事件触发了这个Watcher，那么就会向指定客户端发送一个事件通知，来实现分布式的通知功能。
2. 客户收到服务器的通知后，Curator 会封装一个WatchedEvent 事件实例，传递给监听器的回调方法process（WatchedEvent event）。

WatchedEvent包含了三个基本属性：

（1）通知状态（keeperState）

（2）事件类型（EventType）

（3）节点路径（path）

适用于一些特殊的场景，比如：会话超时、授权失败等。


Curator引入了Cache来监听ZooKeeper服务端的事件。
Cache对ZooKeeper事件监听进行了封装，能够自动处理反复注册监听

### NodeCache 节点缓存的监听

Node Cache，可以用于监控本节点的新增，删除，更新。

**使用**：
1. 构造一个NodeCache缓存实例， 两个构造方法：
```java
NodeCache(CuratorFramework client, String path) 
/**
* client 传入创建的Curator的框架客户端
* path 监听节点的路径
* dataIsCompressed 是否对数据进行压缩
*/
NodeCache(CuratorFramework client, String path, boolean dataIsCompressed)

```
2. 构造一个NodeCacheListener监听器实例,接口定义如下：

```java
public interface NodeCacheListener {

    void nodeChanged() throws Exception;

}
```
>NodeCacheListener监听器接口，只定义了一个简单的方法 nodeChanged，当节点变化时，这个方法就会被回调到

3. 在创建完NodeCacheListener的实例之后，需要将这个实例注册到NodeCache缓存实例，使用缓存实例的addListener方法。 
4. 然后使用缓存实例nodeCache的start方法，启动节点的事件监听
   
```java
nodeCache.getListenable().addListener(l);

nodeCache.start(); 

//如果设置为true的话，在start启动时立即调用NodeCache的getCurrentData方法就能够得到对应节点的信息ChildData类，如果设置为false的就得不到对应的信息。
 start(boolean buildInitial)  //true代表缓存当前节点

```

实例：
```java
@Test
    public void testNodeCache() {

        //检查节点是否存在，没有则创建
        boolean isExist = ZKclient.instance.isNodeExist(workerPath);
        if (!isExist) {
            ZKclient.instance.createNode(workerPath, null);
        }

        CuratorFramework client = ZKclient.instance.getClient();
        try {
            NodeCache nodeCache =
                    new NodeCache(client, workerPath, false);
            NodeCacheListener l = new NodeCacheListener() {
                @Override
                public void nodeChanged() throws Exception {
                    ChildData childData = nodeCache.getCurrentData();
                    log.info("ZNode节点状态改变, path={}", childData.getPath());
                    log.info("ZNode节点状态改变, data={}", new String(childData.getData(), "Utf-8"));
                    log.info("ZNode节点状态改变, stat={}", childData.getStat());
                }
            };
            nodeCache.getListenable().addListener(l);
            nodeCache.start();

            // 第1次变更节点数据
            client.setData().forPath(workerPath, "第1次更改内容".getBytes());
            Thread.sleep(1000);

            // 第2次变更节点数据
            client.setData().forPath(workerPath, "第2次更改内容".getBytes());

            Thread.sleep(1000);

            // 第3次变更节点数据
            client.setData().forPath(workerPath, "第3次更改内容".getBytes());
            Thread.sleep(1000);

            // 第4次变更节点数据
//            client.delete().forPath(workerPath);
            Thread.sleep(Integer.MAX_VALUE);
        } catch (Exception e) {
            log.error("创建NodeCache监听失败, path={}", workerPath);
        }
    }
```


结论：NodeCashe节点缓存能够重复的进行事件节点
>如果NodeCache监听的节点为空（也就是说传入的路径不存在）。那么如果我们后面创建了对应的节点，也是会触发事件从而回调nodeChanged方法。

### PathChildrenCache 子节点监听
PathChildrenCache子节点缓存用于子节点的监听，监控本节点的子节点被创建、更新或者删除。需要强调两点：

（1）只能监听子节点，监听不到当前节点

（2）不能递归监听，子节点下的子节点不能递归监控


**使用**：
1. 构造一个缓存实例， 多个重载的构造方法：
```java
NodeCache(CuratorFramework client, String path) 
/**
* client 传入创建的Curator的框架客户端
* path 监听节点的路径
* cacheData 是否把节点内容缓存起来
* dataIsCompressed 是否对数据进行压缩
* executorService ，表示通过传入的线程池或者线程工厂，来异步处理监听事件
* threadFactory参数（如果有）表示线程池工厂，当PathChildrenCache内部需要开启新的线程执行时，使用该线程池工厂来创建线程
*/
public PathChildrenCache(CuratorFramework client, String path,boolean cacheData)

public PathChildrenCache(CuratorFramework client, String path,boolean cacheData, 
         boolean dataIsCompressed,final ExecutorService executorService)

public PathChildrenCache(CuratorFramework client, String path,boolean cacheData,
         boolean dataIsCompressed,ThreadFactory threadFactory)

public PathChildrenCache(CuratorFramework client, String path,boolean cacheData,
         ThreadFactory threadFactory)

```
2. 构造一个子节点缓存监听器PathChildrenCacheListener实例,接口定义如下：

```java
import org.apache.curator.framework.CuratorFramework;
 
public interface PathChildrenCacheListener {

   void childEvent(CuratorFramework client, PathChildrenCacheEvent e) throws Exception;

}
```
>NodeCacheListener监听器接口，只定义了一个简单的方法 childEvent，当子节点有变化时，这个方法就会被回调到

3. 在创建完PathChildrenCacheListener的实例之后，需要将这个实例注册到PathChildrenCache缓存实例，使用缓存实例的addListener方法。 
4. 然后使用缓存实例nodeCache的start方法，启动节点的事件监听

这里的start方法，需要传入启动的模式。可以传入三种模式：
（1）NORMAL——异步初始化cache， 完成后不会发出通知

（2）BUILD_INITIAL_CACHE——同步初始化cache，以及创建cache后，就从服务器拉取对应的数据

（3）POST_INITIALIZED_EVENT——异步初始化cache，初始化完成触发PathChildrenCacheEvent.Type#INITIALIZED事件，cache中Listener会收到该事件的通知
   

PathChildrenCache来监听节点的事件，完整的实例代码如下
```java
 @Test
    public void testPathChildrenCache() {

        //检查节点是否存在，没有则创建
        boolean isExist = ZKclient.instance.isNodeExist(workerPath);
        if (!isExist) {
            ZKclient.instance.createNode(workerPath, null);
        }

        CuratorFramework client = ZKclient.instance.getClient();

        try {
            PathChildrenCache cache =
                    new PathChildrenCache(client, workerPath, true);
            PathChildrenCacheListener l =
                    new PathChildrenCacheListener() {
                        @Override
                        public void childEvent(CuratorFramework client,
                                               PathChildrenCacheEvent event) {
                            try {
                                ChildData data = event.getData();
                                switch (event.getType()) {
                                    case CHILD_ADDED:

                                        log.info("子节点增加, path={}, data={}",
                                                data.getPath(), new String(data.getData(), "UTF-8"));

                                        break;
                                    case CHILD_UPDATED:
                                        log.info("子节点更新, path={}, data={}",
                                                data.getPath(), new String(data.getData(), "UTF-8"));
                                        break;
                                    case CHILD_REMOVED:
                                        log.info("子节点删除, path={}, data={}",
                                                data.getPath(), new String(data.getData(), "UTF-8"));
                                        break;
                                    default:
                                        break;
                                }

                            } catch (
                                    UnsupportedEncodingException e) {
                                e.printStackTrace();
                            }
                        }
                    };
            cache.getListenable().addListener(l);
            cache.start(PathChildrenCache.StartMode.BUILD_INITIAL_CACHE);
            Thread.sleep(1000);
            for (int i = 0; i < 3; i++) {
                ZKclient.instance.createNode(subWorkerPath + i, null);
            }

            Thread.sleep(1000);
            for (int i = 0; i < 3; i++) {
                ZKclient.instance.deleteNode(subWorkerPath + i);
            }

             } catch (Exception e) {
            log.error("PathCache监听失败, path=", workerPath);
        }

    }

运行的结果如下：

\- 子节点增加, path=/test/listener/node/id-0, data=to set content

\- 子节点增加, path=/test/listener/node/id-2, data=to set content

\- 子节点增加, path=/test/listener/node/id-1, data=to set content

......

\- 子节点删除, path=/test/listener/node/id-2, data=to set content

\- 子节点删除, path=/test/listener/node/id-0, data=to set content

\- 子节点删除, path=/test/listener/node/id-1, data=to set content
```

>Curator的监听原理，无论是PathChildrenCache，还是TreeCache，所谓的监听，都是进行Curator本地缓存视图和ZooKeeper服务器远程的数据节点的对比。


**在什么场景下触发事件呢？** 以节点增加事件NODE_ADDED为例
所在本地缓存视图开始的时候，本地视图为空，在数据同步的时候，本地的监听器就能监听到NODE_ADDED事件。
- 这是因为，刚开始本地缓存并没有内容，
- 然后本地缓存和服务器缓存进行对比，发现ZooKeeper服务器有节点而本地缓存没有，
- 这才将服务器的节点缓存到本地，就会触发本地缓存的NODE_ADDED事件。

### Tree Cache 节点树缓存
Tree Cache可以看做是上两种的合体，Tree Cache观察的是当前ZNode节点的所有数据。而TreeCache节点树缓存是PathChildrenCache的增强，不光能监听子节点，也能监听节点自身。

**使用**：
1. 构造一个TreeCache缓存实例
```java

//maxDepth表示缓存的层次深度，默认为整数最大值。
//executorService 表示监听的的执行线程池，默认会创建一个单一线程的线程池。
//createParentNodes 表示是否创建父亲节点，默认为false。
   
TreeCache(CuratorFramework client, String path) 
 

TreeCache(CuratorFramework client, String path,
          boolean cacheData, boolean dataIsCompressed, int maxDepth, 
		 ExecutorService executorService, boolean createParentNodes,
		 TreeCacheSelector selector) 
```
2. 构造一个TreeCacheListener监听器实例:
   
```java
   
 import org.apache.curator.framework.CuratorFramework;

public interface TreeCacheListener {
    void childEvent(CuratorFramework var1, TreeCacheEvent var2) throws Exception;

}
```
3. 注册监听器
4. 启动事件监听

完整实例：
```java
@Test
    public void testTreeCache() {

        //检查节点是否存在，没有则创建
        boolean isExist = ZKclient.instance.isNodeExist(workerPath);
        if (!isExist) {
            ZKclient.instance.createNode(workerPath, null);
        }

        CuratorFramework client = ZKclient.instance.getClient();

        try {
            TreeCache treeCache  =
                    new TreeCache(client, workerPath);
            TreeCacheListener l =
                    new TreeCacheListener() {
                        @Override
                        public void childEvent(CuratorFramework client,
                                               TreeCacheEvent event) {
                            try {
                                ChildData data = event.getData();
                                if(data==null)
                                {
                                    log.info("数据为空");
                                    return;
                                }
                                switch (event.getType()) {
                                    case NODE_ADDED:

                                        log.info("[TreeCache]节点增加, path={}, data={}",
                                                data.getPath(), new String(data.getData(), "UTF-8"));

                                        break;
                                    case NODE_UPDATED:
                                        log.info("[TreeCache]节点更新, path={}, data={}",
                                                data.getPath(), new String(data.getData(), "UTF-8"));
                                        break;
                                    case NODE_REMOVED:
                                        log.info("[TreeCache]节点删除, path={}, data={}",
                                                data.getPath(), new String(data.getData(), "UTF-8"));
                                        break;
                                    default:
                                        break;
                                }

                            } catch (
                                    UnsupportedEncodingException e) {
                                e.printStackTrace();
                            }
                        }
                    };
            treeCache.getListenable().addListener(l);
            treeCache.start();
            Thread.sleep(1000);
            for (int i = 0; i < 3; i++) {
                ZKclient.instance.createNode(subWorkerPath + i, null);
            }

            Thread.sleep(1000);
            for (int i = 0; i < 3; i++) {
                ZKclient.instance.deleteNode(subWorkerPath + i);
            }
            Thread.sleep(1000);

            ZKclient.instance.deleteNode(workerPath);

            Thread.sleep(Integer.MAX_VALUE);

        } catch (Exception e) {
            log.error("PathCache监听失败, path=", workerPath);
        }

    }

```

说明下事件的类型，对应于节点的增加、修改、删除，TreeCache 的事件类型为：

（1）NODE_ADDED

（2）NODE_UPDATED

（3）NODE_REMOVED

这一点，与Path Cache 的事件类型不同，与Path Cache 的事件类型为：

（1）CHILD_ADDED

（2）CHILD_UPDATED

（3）CHILD_REMOVED



# 参考资料
- [一篇文章带你深入理解Zookeeper](https://mp.weixin.qq.com/s/38JLeXS54Ji-ozLRQpv0JQ)
- [Zookeeper Curator 事件监听](https://blog.csdn.net/crazymakercircle/article/details/85922561)