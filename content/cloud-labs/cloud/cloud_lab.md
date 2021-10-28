---
title: 云PaaS平台开发
weight: 2
---

# 云PaaS平台开发

## 实验要求

基于Kubernetes，设计并实现一PaaS平台。该平台的终极目标，用户可以通过该平台，实现从源代码到可访问的服务之间的整个自动化流程。

比如，在传统方式下，用户编写了一个React应用，为了能让其他人通过互联网访问到他的应用，他需要首先在云厂商购买服务器，然后在服务器上安装操作系统和NodeJS，然后手动拉取依赖和编译，然后手动将编译好的静态文件放到合适的地方（配置Nginx等）。

而在通过这个PaaS平台，用户只需要提供自己的源代码，并按照平台的需求，编写一个配置文件（可能会包含一个Dockerfile，也可能是某种模板的形式），剩下的工作可以放心地交给PaaS平台来完成。

那么PaaS平台是如何实现这个功能的呢？首先，它根据用户提供的源代码和配置文件，将源代码编译成一个OCI镜像；然后将该镜像交给一个容器管理平台部署为一个容器组（一般是Kubernetes中的Deployment或StatefulSet）；然后根据用户提供的配置文件，将容器组的特定端口暴露到互联网上（例如使用Service的NodePort或LoadBalancer）。

进一步地，一个真正的生产级别的应用不可能由一个容器组提供的功能支撑（最简单的一个前后端分离的应用，就要包含一个前端、一个后端，可能还有数据库和缓存服务等），为了完成整个应用的部署，用户一般会通过配置文件来指定多个不同的容器。这些容器之间的逻辑是相互联系，平台需要能够将它们在逻辑上与其他的应用容器区分开。

因此，可以看到，一个完整的PaaS平台可以包含：

- **镜像管理** 
    1. 用户可以向平台添加自己需要的镜像，可以直接提供一个Dockerfile，也可以直接提供镜像的压缩包，也可以提供源代码（可以是代码的压缩包，也可以是一个代码仓库地址等等），也可以直接让平台去拉公共的镜像仓库
    2. 用户可以浏览、更改和删除自己添加的镜像
    3. 在镜像管理部分，可以自己在本地通过Docker Image管理，也可以考虑搭建一个私有的[Docker Registry](https://docs.docker.com/registry/)
- **容器管理** 用户可创建和修改容器，并监控容器状态 （考虑使用Kubernetes进行部署和管理）
- **应用部署** 用户可以直观地将若干个逻辑上统一的容器编排成完成的应用，并发布。这里这个“应用”的概念是Kubernetes本身没有的，需要你自己去抽象。比如，可以把不同应用的容器组使用namespace隔离等等。在KubeSphere上，这个概念其实就是一个“项目”）

希望同学们在实现上述功能的过程中加深对Kubernetes各项概念的理解，并体会云计算为应用的发布和运维带来的便利。

开发过程中，需要使用到Kubernetes和Docker Engine对外提供的API，可以直接调用他们的OpenAPI，也可以使用官方或第三方封装好的Client SDK。

[Kubernetes的OpenAPI](https://kubernetes.io/zh/docs/concepts/overview/kubernetes-api/)描述文件可以在[这里](https://github.com/kubernetes/kubernetes/blob/master/api/openapi-spec/swagger.json)找到，这里列出了比较流行的一些Kubernetes[客户端库](https://kubernetes.io/zh/docs/reference/using-api/client-libraries/)。

Docker的OpenAPI描述文件可以在[这里](https://docs.docker.com/engine/api/v1.41/#)找到，这里列出了比较流行的一些Docker[客户端库](https://docs.docker.com/engine/api/sdk/)。

## 实验代码管理与部署

实验代码请托管到软院代码托管平台[BuGit](https://git.scs.buaa.edu.cn)上。

{{< hint warning >}}

首次使用代码托管平台时需要激活账户。激活账户时，请注意邮箱的正确性，并牢记密码。

{{< /hint >}}

系统开发将分小组进行，需要小组在[BuGit](https://git.scs.buaa.edu.cn)上创建项目，并邀请所有小组成员加入。

## 可以使用的资源

1. [KubeSphere](https://kube.scs.buaa.edu.cn)，该平台的初始账号密码与BuGit相同，并且其上的项目与BuGit同步。在BiGit上创建项目后，可在KubeSphere对应的项目中部署容器。

2. [Harbor](https://harbor.scs.buaa.edu.cn)，该平台的初始账号密码与BuGit相同，并且其上的项目与BuGit同步。在BiGit上创建进行代码仓库的构建后，可在Harbor对应的项目中查看到创建的镜像。

3. 校内的Docker Hub镜像地址：`10.251.0.37:5000`。
