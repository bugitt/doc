---
title: 虚拟机使用说明
weight: 100
---

# 虚拟机使用说明

## 连接虚拟机

### Linux系统

首先从云平台中获取虚拟机的IP和登录名，之后即可在本地通过任意ssh客户端登录。

{{< tabs "uu1" >}}
{{< tab "MacOS" >}}

使用系统自带的Terminal.app登录即可。

为了更好的使用体验，推荐使用[iterm2](https://iterm2.com/)登录。

当然，你也可以使用[termius](https://termius.com/)进行多个ssh连接的管理。

{{< /tab >}}

{{< tab "Linux" >}}

如果你是Linux Desktop用户，那么你肯定已经有了自己喜爱的终端模拟器，此处不再赘述。

{{< /tab >}}

{{< tab "Windows" >}}

一般来讲，Windows 10（及以上）自带的cmd.exe都自带ssh client，打开cmd后直接`ssh foo@x.x.x.x`即可登录。

为了更好的使用体验，推荐下载使用[Windows Terminal](https://aka.ms/terminal)。

当然，你也可以使用[termius](https://termius.com/)或者其他工具（如 Xshell等）进行多个ssh连接的管理。

{{< /tab >}}

{{< tab "校外跳板" >}}

通过 d.buaa.edu.cn 跳转登录即可。

{{< /tab >}}

{{< /tabs >}}

## 联网

### Linux系统

推荐使用[Dr-Bluemond/srun](https://github.com/Dr-Bluemond/srun)提供的工具。云平台提供已经编译好的Linux64位版本。可以这样获取：

```bash
wget http://10.251.253.10/scsos/tools/linux/buaalogin
chmod +x ./buaalogin
```

使用前请使用`config`命令配置一下校园网用户名和密码（注意，如果用户名中有英文的话，请大小写都尝试一下）：

```bash
./buaalogin config
```

配置完成后，使用`login`命令登录即可：

```bash
./buaalogin login
```

或，直接：

```bash
./buaaloign
```

如果想作为系统命令使用的话（注意替换合适的安装路径）：

```bash
sudo install ./buaalogin /usr/local/bin
```

## 传输文件

### Linux系统

{{< tabs "uu2" >}}
{{< tab "MacOS & Linux" >}}

可以使用使用SCP命令进行服务器与本地之间的文件交换。

{{< /tab >}}

{{< tab "Windows" >}}

除了在终端中使用SCP命令外，

还可以使用[WinSCP](https://winscp.net/eng/download.php)进行图形化的文件管理。

{{< /tab >}}

{{< tab "校外跳板" >}}

d.buaa.edu.cn 的Linux界面已经提供了比较完善的文件管理工具。

{{< /tab >}}

{{< /tabs >}}