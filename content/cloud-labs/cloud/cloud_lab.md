---
title: 云PaaS平台开发
weight: 2
---

# 云PaaS平台开发

## 实验要求

基于Kubernetes，设计并实现一PaaS平台，可以考虑实现如下功能（不限于）：

- 镜像管理
- 容器管理（用户可创建和修改容器，并监控容器状态）
- 应用部署（用户可以直观地将若干个逻辑上统一的容器编排成完成的应用，并发布）

希望同学们在实现上述功能的过程中加深对Kubernetes各项概念的理解，并体会云计算为应用的发布和运维带来的便利。

## 实验代码管理与部署

实验代码请托管到软院代码托管平台[BuGit](https://git.scs.buaa.edu.cn)上。

{{< hint warning >}}

首次使用代码托管平台时需要激活账户。激活账户时，请注意邮箱的正确性，并牢记密码。

{{< /hint >}}

系统开发将分小组进行，需要小组在[BuGit](https://git.scs.buaa.edu.cn)上创建项目，并邀请所有小组成员加入。

## 可以使用的资源

1. [KubeSphere](https://kube.scs.buaa.edu.cn)，该平台的初始账号密码与BuGit相同，并且其上的项目与BuGit同步。在BuGit上创建项目后，可在KubeSphere对应的项目中部署容器。

2. [Harbor](https://harbor.scs.buaa.edu.cn)，该平台的初始账号密码与BuGit相同，并且其上的项目与BuGit同步。在BuGit上创建进行代码仓库的构建后，可在Harbor对应的项目中查看到创建的镜像。

3. 校内的Docker Hub镜像地址：`10.251.0.37:5000`。
