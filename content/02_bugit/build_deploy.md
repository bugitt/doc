---
title: 自动构建与部署
weight: 100
---

# 自动构建与部署

## 原理

本质上，自动构建与部署的过程是，系统根据用户提供的包含在代码仓库中的配置文件，将代码编译成一个OCI规范的镜像，然后上传到镜像中心，最后通知Kubernetes集群拉取镜像运行，以对外提供可用服务。

自动构建与部署的几个步骤：

1. 用户编写配置合法的配置文件，并包含在代码仓库的根目录中。

2. 用户触发自动构建与部署。目前支持自动监听代码推送（git push）动作，和在前端手动点击按钮触发。

3. 系统拉取用户指定的仓库的指定的某次提交的代码，并根据指定的Dockerfile进行镜像构建。

4. 系统将构建完成的镜像将推送到镜像中心（harbor.scs.buaa.edu.cn）。

5. 系统通知Kubernetes拉取镜像，并部署之。

## 配置文件

配置文件一共包含两个：`Dockerfile` 和 `.bugit.yaml`。

### Dockerfile

Dockerfile用来描述该代码仓库希望被如何编译和打包成一个OCI镜像。具体的编写规则可以参考[Dockerfile Reference](https://docs.docker.com/engine/reference/builder/)。

### .bugit.yaml

`.bugit.yaml`是一个YAML文件（名称`.bugit.yml`也是合法的，并且请注意文件名最前面那个`.`）。它是对整个构建和部署过程的描述。

下面是一个`.bugit.yaml`文件支持的全部指令的示例（请注意缩进）。


{{< hint danger >}}

下方示例中，提到的非必需字段，都可以在`.bugit.yaml`省略不写。

{{< /hint >}}


```yaml
# 必需字段。表示当前的.bugit.yaml 所适用的构建与部署流程的版本号，目前仅支持 0.0.1
version: 0.0.1

# on 字段中的内容用来表示在哪个分支发生什么事件时，自动启动构建与部署流程
# 该字段中可以包含若干组内容，每一组的 key （比如，下方示例中的 main 和 master） 都是分支名称，其 value （比如下方示例中的 [push]）是一个数组，表示希望系统监听哪些事件的发生
# 比如下面的示例就表示，希望系统在远程仓库的 main 分支和 master 分支发生代码推送事件（git push）时，自动启动构建与部署流程
# 如果希望开启“自动”构建与部署的功能，那么该字段是必需的
on: 
  main: ["push"]
  master: ["push"]

# 必需字段。build 字段用来描述如何系统如何构建OCI镜像
build:
  name: build-1  # 必需。名称标识，目前没有太大意义。可以是任意字符串，但请不要带空格
  type: docker   # 必需。构建的类型，目前仅支持docker
  docker_tag: simple        # 非必需。表示希望给构建好的字段加的额外tag
  dockerfile: ./Dockerfile  # 必需。表示使用的 Dockerfile 与代码仓库根目录的相对路径

# 非必需字段。build 字段用来描述如何构建好的镜像将会被如何部署。如果不希望使用部署功能的话，该字段可以忽略。
deploy:
  # 必需。其值为一个列表。表示希望在哪几个分支的代码中开启部署功能。（“部署”是一个非常重的操作，需要用户明确确认）
  on: ["main", "master"]
  # 必需(至少包含一个端口）。表示运行起来的容器将向外暴露哪些端口
  ports:
    - name: name-1    # 名称，该端口的一个标识。必须为小写的英文字母和数字组合，可以包含短横线。但数字不能作为开头，短横线不能作为结尾。
      protocol: tcp   # 使用的协议，支持tcp和udp，默认是tcp
      port: 80        # 容器向外暴露的端口
    - name: name-2
      protocol: udp
      port: 9934
     
  # 非必需。envs 表示服务部署时需要使用的环境变量。key 和 value 一一对应。
  envs:         
    SOME_ENV_1: "some_env_1"
    SOME_ENV_2: "some_env_2"
  # 非必需。stateful 表示该服务是有状态的还是无状态的。该字段默认为false，即默认无状态。
  stateful: false
  # 非必需。work_dir 表示容器开始运行时，执行的命令所在的目录。如果该字段为空，默认使用镜像中指定的 workDir
  work_dir: /path/to/work_dir
  # 非必需。cmd表示容器启动时指定的命令。其又分为两部分（两个列表），其中，command用来指定命令， args 用来指定命令需要使用的参数
  # 如下方的示例，对应我们常见的命令形式就是 java -jar awesome.jar，表示使用java命令运行一个jar
  cmd:
    command: ["java"]
    args: ["-jar", "awesome.jar"]
  # 非必需。cpu限额，表示该服务最多使用多少CPU资源。默认为250m（四分之一个CPU核心）
  cpu: 250m     # 可以直接使用数字，如 3、100，分别表示使用3个CPU核心、100个CPU核心；也可以使用m作为单位，一个CPU核心是1000m，那么250m就表示使用四分之一个CPU核心
  # 非必需。内存限额，表示该服务最多使用多少内存资源。默认为0.5G
  memory: 512Mi # 请使用单位 Mi，Gi，或 M，G，如 0.5G，512Mi 等

```

{{<hint warning>}}

BuGit平台的每个个人和项目的资源配额为2核4G。所以请合理调配每个代码仓库的资源配额。

{{</hint>}}

{{<hint warning>}}

对于典型的资源消耗大户（如Java应用对内存的消耗），需要注意在容器的启动命令中手动限制资源配额（如手动指定是jvm的内存参数等），防止容器因为资源有限而使服务启动失败。

{{</hint>}}

## 示例

BuGit平台中的[test-project项目](https://git.scs.buaa.edu.cn/test-project)包含的每个项目都进行了自动构建和部署的配置，都可以作为参考。

特别地，下面给出了一些典型的示例。

### Static Web

适用于纯静态文件的网站部署（如，仅包含html，css，js等文件）。

可参考项目 [static-web](https://git.scs.buaa.edu.cn/test-project/static-web)。

#### Dockerfile

```dockerfile
FROM nginx

# 下面的 . 表示使用的是当前这个Dockerfile所在的目录作为网站的根目录
# 如果你的 index.html 所在的位置与此不同，请根据实际情况修改
COPY . /usr/share/nginx/html
```

#### .bugit.yaml

```yaml
version: 0.0.1

on:
  master: [push]
  main: [push]

build:
  name: build-static-nginx
  type: docker
  docker_tag: web
  dockerfile: ./Dockerfile

deploy:
  on: [main, master]
  ports:
    - name: web
      protocol: tcp
      port: 80
```