

## 如何在开发机上运行多个Elasticsearch实例
bin/elasticsearch -E node.name=node0 -E cluster.name=clusterDemo -E path.data=node0_data -d
bin/elasticsearch -E node.name=node1 -E cluster.name=clusterDemo -E path.data=node1_data -d
bin/elasticsearch -E node.name=node2 -E cluster.name=clusterDemo -E path.data=node2_data -d
bin/elasticsearch -E node.name=node3 -E cluster.name=clusterDemo -E path.data=node4_data -d






# 查询

## 语法示例

### match查询（分词）

- 会对字段值进行分词匹配
  
```json
GET http://ip:prot/textbook/_search
{
  "query": {
    "match": {
      "bookName":"my test"
    }
  }
}
```


**match_phrase查询**

>**不分词？ 默认是要连续匹配**：它不是匹配到某一处分词的结果就算是匹配成功了，而是需要query中所有的词都匹配到，而且相对顺序还要一致，而且默认还是连续的，其实类似精确包含.

```json
GET http://ip:prot/textbook/_search
{
  "query": {
    "match_phrase": {
      "bookName":"is a test"
    }
  }
}
```


**搜索的严格程度：slop**

将slop置为1，然后搜索"is test"，虽然is test中间省略了一个词语"a"，但是在slop为1的情况下是可以容忍你中间省略一个词语的，也可以搜索出来结果。
```json
GET http://ip:prot/textbook/_search
{
  "query": {
    "match_phrase": {
      "bookName":{
        "query":"is test",
        "slop":1
      }
    }
  }
}
```

### multi_match查询

- 类似Or操作，满足一个就返回

```json
GET http://ip:prot/textbook/_search
{
  "query": {
    "multi_match": {
        "query" : "老坛",
        "fields" : ["bookName", "author"]
    }
  }
}
```

### term查询（不分词）

- 它和match的唯一区别就是match需要对query进行分词，而term是不会进行分词的，它会直接拿query整体和原文进行匹配。
- 但是原文指的是被分词后的原文， 在原文被分好的每一个词语里，没有一个词语是："This is a test doc"，那自然是什么都搜不到了。所以在这种情况下就只能用某一个词进行搜索才可以搜到， 比如等于 “test”
```json
GET http://ip:prot/textbook/_search
{
  "query": {
    "term": {
      "bookName": "This is a test doc"
    }
  }
}
```


**terms查询**

- terms查询事实上就是多个term查询取一个交集
- 也就是要满足多个term查询条件匹配出来的结果才可以查到，所以是比单纯的term条件更为严格了：

- 比如这个例子，是要求原文中既有This这个词，又有is这个词才可以被查到，那按照这个规则我们是可以匹配到数据的：
  - 但是如果改成了一个不存在的词便匹配不到了：
  
```json
{
  "query": {
    "terms": {
      "bookName": ["This", "is"]
    }
  }
}
{
  "query": {
    "terms": {
      "bookName": ["This", "my"]
    }
  }
}
```


### fuzzy查询

- fuzzy是ES里面的模糊搜索，它可以借助term查询来进行理解。
- fuzzy和term一样，也**不会将query进行分词**，但是不同的是它在进行匹配时可以容忍你的词语拼写有错误。
- 至于容忍度如何，是根据参数**fuzziness**决定的。
- fuzziness默认是2，也就是在默认情况下，fuzzy查询容忍你有两个字符及以下的拼写错误。无论是错写多写还是少写都是计算在内的。
>即如果你要匹配的词语为test，但是你的query是text，那也可以匹配到。

```json
GET http://ip:prot/textbook/_search
{
  "query": {
    "fuzzy": {
      "bookName":"text"
    }
  }
}

{
  "query": {
    "fuzzy": {
      "bookName":{
        "value":"texts",
        "fuzziness":1
      }
    }
  }
}
```

### range查询

range查询时对于某一个**数值字段**的大小范围查询.

* gte：大于等于
* gt：大于
* lt：小于
* lte：小于等于

```json
GET http://ip:prot/textbook/_search
{ 
  "query": {
    "range": { 
      "num": { 
          "gte":20, 
          "lt":30 
      } 
    }
  } 
}
```

### bool查询
bool查询是上面查询的一个综合，它可以用多个上面的查询去组合出一个大的查询语句，它也有一些关键字：

* must：代表且的关系，也就是必须要满足该条件
* should：代表或的关系，代表符合该条件就可以被查出来
* must_not：代表非的关系，也就是要求不能是符合该条件的数据才能被查出来


例如： 要求must里面的match是必须要符合的，但是should里面的两个条件就可以符合一条即可。

```json
GET http://ip:prot/textbook/_search
{
    "query":{
        "bool":{
            "must":{
                "match":{
                    "bookName":"老坛"
                }
            },
            "should":{
                "term":{
                    "author":"老坛"
                },
                "range":{
                    "num":{
                        "gt":20
                    }
                },
            }
        }
    }
}

```

### 排序和分页

排序和分页也是建立在上述的那些搜索之上的。排序和分页的条件是和query平级去写的。

先举个例子：

```json
GET http://ip:prot/textbook/_search
{
    "query":{
        "match":{
            "bookName":"老坛"
        }
    },
    // 它代表的意思是按照页容量为100进行分页，取第一页​。
    "from":0,
    "size":100,
    "sort":{
        "num":{
            "order":"desc"
        }
    }
}
```