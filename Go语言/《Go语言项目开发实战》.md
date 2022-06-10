
[TOC]

# 一、环境准备
>如何准备开发环境、制作 CA 证书，安装和配置用到的数据库、应用，以及 Shell 脚本编写技巧


**项目背景**： 实现一个IAM（Identity and Access Management，身份识别与访问管理）系统。
- 为了保障 Go 应用的安全，我们需要对访问进行认证，对资源进行授权。


如何实现访问认证和资源授权呢?
- 认证功能不复杂，我们可以通过 JWT (JSON Web Token)认证来实现。
- 授权功能的复杂性使得它可以囊括很多 Go 开发技能点。 本专栏学习就是将这两种功能实现升级为IAM系统，讲解它的构建过程。


**创建数据库**

```shell
sudo tee /etc/yum.repos.d/mongodb-org-4.4.repo<<'EOF'
[mongodb-org-4.4]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/$releasever/mongodb-org/4.4/x86_64
gpgcheck=1
enabled=1
gpgkey=https://www.mongodb.org/static/pgp/server-4.4.asc
EOF
```

**创建CA证书**


```shell
tee ca-config.json << EOF
{
    "signing": {
        "default": {
        "expiry": "87600h"
        },
        "profiles": {
        "iam": {
            "usages": [
            "signing",
            "key encipherment",
            "server auth",
            "client auth"
            ],
            "expiry": "876000h"
        }
        } 
    }
} 
EOF
```

```shell
$ tee ca-csr.json << EOF 
{
    "CN": "iam-ca",
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names":[
        {
            "C": "CN",
            "ST": "BeiJing",
            "L": "BeiJing",
            "O": "marmotedu",
            "OU": "iam"
        }
    ],
    "ca": {
        "expiry": "876000h"
    }
}
EOF
    
```

```shell
tee iam-apiserver-csr.json <<EOF
  "CN": "iam-apiserver",
  "key": {
    "algo": "rsa",
    "size": 2048
  },
"names": [ {
      "C": "CN",
      "ST": "BeiJing",
      "L": "BeiJing",
      "O": "marmotedu",
      "OU": "iam-apiserver"
} ],
  "hosts": [
    "127.0.0.1",
    "localhost",
    "iam.api.marmotedu.com"
] }
EOF
```

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2NTQ5MjQyNTgsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2NTQ4Mzc4NTgsInN1YiI6ImFkbWluIn0.NB4jJIfet4lfvfJN6KRwQu56VFajxvgS4cDI9BTfRso

'{"password":"User@2021","metadata":{"name":"colin"},"nickname":"colin","email":"colin@foxmail.com","phone":"1812884xxxx"}'


```shell
 curl -s -XPOST -H'Content-Type: application/json' -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2NTQ5MjQyNTgsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2NTQ4Mzc4NTgsInN1YiI6ImFkbWluIn0.NB4jJIfet4lfvfJN6KRwQu56VFajxvgS4cDI9BTfRso' -d '{"password":"User@2021","metadata":{"name":"colin"},"nickname":"colin","email":"colin@foxmail.com","phone":"1812884xxxx"}' http://127.0.0.1:8080/v1/users

  curl -s -XGET -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2NTQ5MjQyNTgsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2NTQ4Mzc4NTgsInN1YiI6ImFkbWluIn0.NB4jJIfet4lfvfJN6KRwQu56VFajxvgS4cDI9BTfRso' 'http://127.0.0.1:8080/v1/users?offset=0&limit=10'
```

```shell
{
"CN": "admin",
"key": {
  "algo": "rsa",
  "size": 2048
},
"names": [ {
} ],
"C": "CN",
"ST": "BeiJing",
"L": "BeiJing",
"O": "marmotedu",
"OU": "iamctl"
"hosts": []
}
```

cfssl gencert -ca=${IAM_CONFIG_DIR}/cert/ca.pem -ca-key=${IAM_CONFIG_DIR}/cert/ca-key.pem  -config=${IAM_CONFIG_DIR}/cert/ca-config.json -profile=iam admin-csr.json | cfssljson -bare admin



# 二、规范设计\
>目录规范、日志规范、错误码规范、Commit规范

## 代码规范



# 三、基础功能设计或开发
>开发基础功能，如日志包、错误包、错误码

# 四、服务开发
>解析一个企业级的 Go 项目代码，让你学会如何开发 Go 应用. 怎么设计和开发 API 服务、Go SDK、客户端工具


# 五、服务测试
>讲解单元测试、功能测试、性能分析和 性能调优的方法，最终让你交付一个性能和稳定性都经过充分测试的、生产级可用的服 务。


# 六、服务部署
>如何部署一个高可用、安 全、具备容灾能力，又可以轻松水平扩展的企业应用。 传统部署和容器化部署。
