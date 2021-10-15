---
title: 附录：创建Kubernetes集群
weight: 10
---

# 创建Kubernetes集群

Kubernetes生态发展至今已非常完善，部署一个Kubernetes集群已经不再是一件非常繁琐和困难的事情。社区有大量简单可靠的解决方案。

下面给出几种可选的方案，根据自己的实际情况，**选择其一即可**。

## kubeadm（不推荐）

Kubernetes官方推荐使用[kubeadm](https://kubernetes.io/zh/docs/reference/setup-tools/kubeadm/) 来初始化一个Kubernetes集群。通过它，用户可以获得是一个相对“纯净”的Kubernetes集群。但该方法相对繁琐，而且对国内用户非常不友好，因此不推荐这种方式部署。

有兴趣的同学可以尝试。

## KubeKey（推荐）

[KubeKey](https://github.com/kubesphere/kubekey) 是由[Kubesphere](https://kubesphere.io/zh/) （一个国内公司主导的开源的Kubernetes管理平台）开源的Kubernetes和Kubesphere部署工具。

KubeKey使用声明式的配置方式，用户只需要通过一个YAML配置文件给出所需集群的相关配置，即可通过KubeKey创建集群或修改集群的状态。更详细内容和使用方式请参考[KubeKey的文档](https://github.com/kubesphere/kubekey/blob/master/README.md)。

### 创建集群

在本次实验分配的虚拟机中，已经提前完成了KubeKey的部分配置，只需按照下面几步操作即可完成Kubernetes集群的创建。

{{< hint danger >}}

**注意**  

为了减少不必要的麻烦，请**直接使用root账户**登录虚拟机，并完成相关操作。

{{< /hint >}}

1. 进入`/root/kubesphere`目录

    ```bash
    cd /root/kubesphere
    ```

2. 修改配置文件中的IP和登录密码。可以按需修改主机的hostname。如果修改了hostname，也需要同步修改`roleGroups`中的值。

    ```yaml
    apiVersion: kubekey.kubesphere.io/v1alpha1
    kind: Cluster
    metadata:
    name: main-cluster
    spec:
    hosts:
    - {name: node1, address: 1.1.1.1, internalAddress: 1.1.1.1, user: root, password: '&&shieshuyuan21'}
    - {name: node2, address: 1.1.1.1, internalAddress: 1.1.1.1, user: root, password: '&&shieshuyuan21'}
    roleGroups:
        etcd:
        - node1
        master:
        - node1
        worker:
        - node1
        - node2
    controlPlaneEndpoint:
        domain: lb.kubesphere.local
        address: ""
        port: 6443
    kubernetes:
        version: v1.20.4
        imageRepo: kubesphere
        clusterName: cluster.local
    network:
        plugin: calico
        kubePodsCIDR: 172.20.0.0/16
        kubeServiceCIDR: 172.21.0.0/16
    registry:
        registryMirrors: []
        insecureRegistries: []
    addons: []
    ```

3. 保存文件并退出后，执行以下命令即可：

    ```bash
    ./kk create cluster -f config.yaml
    ```

4. 等待完成即可。完成后，可以直接在机器中使用`kubectl get node`来验证。

### 删除集群

如果需要删除创建好的集群，只需要执行下述删除命令即可：

```bash
cd /root/kubesphere
./kk delete cluster -f config.yaml
```

## RKE（一般推荐）

{{< hint danger >}}

使用RKE时，需要向目标机器传递公钥。请**不要**使用虚拟机默认生成的公钥。

因为所有实验机器密钥对都相同，当你向目标机器传递了公钥后，本次实验的其他同学将可以无障碍登录你的目标机器。

{{< /hint >}}

[RKE（Rancher Kubernetes Engine）](https://rancher.com/products/rke/)是rancher提供的一个Kubernetes管理工具。

与KubeKey一样，RKE同样提供了声明式的配置方式，你可以在RKE的引导下，创建配置文件，并以此创建集群和管理集群的状态。

本次实验分配的虚拟机中，提供了rke的可执行文件，可以直接使用。对此感兴趣的同学可以参考[rke的安装说明文档](https://rancher.com/docs/rke/latest/en/installation/)。

## k3s（一般推荐）

k3s是一个非常轻量的Kubernetes发行版。其安装和配置方法在其[官方文档](https://rancher.com/docs/k3s/latest/en/quick-start/)中写的很详细。

## 其他（单机推荐）

如果需要在本地创建Kubernetes集群，可以选择使用Minikube、Docker Desktop等。

