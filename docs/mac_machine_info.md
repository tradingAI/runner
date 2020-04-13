runner使用[gopsutil](https://github.com/shirou/gopsutil)获取机器信息

[两种常见的获取宿主信息的方式](https://blog.csdn.net/wufolangren/java/article/details/86702274):
- 在启动容器的时候，挂载/proc/目录到指定目录，注意容器容也有proc目录，容器中不可再用/proc目录，然后根据cpuinfo或者其他文件自己提取数据再算出来，比如:
```
docker run -it -v /proc:/hostinfo/proc:ro alpine sh
```
- 在容器中使用ssh连接到主机获取主机信息

```
docker run -it alpine sh
# apk update   //进入容器后首先更新
# apk add openssh-client   //安装ssh客户端，只需要客户端就可以连接其他主机，不需要安装server
# apk add sshpass    //安装sshpass，ssh不能指定登录密码，需要sshpass协助
# sshpass -p "123456" ssh root@192.168.0.61 "df -h"  //-p后面是ssh登录密码，“df -h“是需要执行的命令
```

Mac 系统下是通过sysctl命令解释输出信息的， container要想获取宿主的信息，需要ssh
登录到宿主主机上，配置相对烦琐
