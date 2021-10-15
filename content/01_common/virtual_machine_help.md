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

登录虚拟机后，可以执行下面命令进行联网，请注意替换学号和密码：

```bash
/usr/bin/curl -k -d "action=login&username=学号&password=密码&type=2&n=117&ac_id=1" "https://gw.buaa.edu.cn/cgi-bin/srun_portal"
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