[TOC]

containerd shim 介绍



```
systemd(1)─┬
           ├─containerd(1694)─┬─{containerd}(1696)
           ├─containerd-shim(20799)─┬─pause(20821)
           │                        ├─sleep(20852)
           │                        ├─{containerd-shim}(20802)
           │                        ├─{containerd-shim}(20803)

containerd-shim  20792  1694     0 /usr/bin/containerd-shim-runc-v2 -namespace k8s.io -address /run/containerd/containerd.sock -publish-binary /usr/bin/containerd -id 193727d35be73014e4adfe4ba93aa20b3834654a65990b9de9fba123592044aa -debug start
containerd-shim  20799  20792    0 /usr/bin/containerd-shim-runc-v2 -namespace k8s.io -id                                                               193727d35be73014e4adfe4ba93aa20b3834654a65990b9de9fba123592044aa -address /run/containerd/containerd.sock
runc             20811  20799    0 /usr/bin/runc --root /run/containerd/runc/k8s.io --log /run/containerd/io.containerd.runtime.v2.task/k8s.io/         193727d35be73014e4adfe4ba93aa20b3834654a65990b9de9fba123592044aa/log.json --log-format json --systemd-cgroup create --bundle /run/containerd/io.        containerd.runtime.v2.task/k8s.io/193727d35be73014e4adfe4ba93aa20b3834654a65990b9de9fba123592044aa --pid-file /run/containerd/io.containerd.runtime.v2. task/k8s.io/193727d35be73014e4adfe4ba93aa20b3834654a65990b9de9fba123592044aa/init.pid 193727d35be73014e4adfe4ba93aa20b3834654a65990b9de9fba123592044aa/ run/containerd/io.containerd.runtime.v2.task/k8s.io/193727d35be73014e4adfe4ba93aa20b3834654a65990b9de9fba123592044aa --pid-file /run/containerd/io.     containerd.runtime.v2.task/k8s.io/193727d35be73014e4adfe4ba93aa20b3834654a65990b9de9fba123592044aa/init.pid                                             193727d35be73014e4adfe4ba93aa20b3834654a65990b9de9fba123592044aa
runc             20827  20799    0 /usr/bin/runc --root /run/containerd/runc/k8s.io --log /run/containerd/io.containerd.runtime.v2.task/k8s.io/         193727d35be73014e4adfe4ba93aa20b3834654a65990b9de9fba123592044aa/log.json --log-format json --systemd-cgroup start                                      193727d35be73014e4adfe4ba93aa20b3834654a65990b9de9fba123592044aa
pause            20821  20799    0 /pause
containerd-shim  20834  1694     0 /usr/bin/containerd-shim-runc-v2 -namespace k8s.io -address /run/containerd/containerd.sock -publish-binary /usr/bin/containerd -id f2ba685582ea81fb6568da6cec0ee3086dd8e15db11214ddfed19f4ded41725b -debug start
runc             20841  20799    0 /usr/bin/runc --root /run/containerd/runc/k8s.io --log /run/containerd/io.containerd.runtime.v2.task/k8s.io/         f2ba685582ea81fb6568da6cec0ee3086dd8e15db11214ddfed19f4ded41725b/log.json --log-format json --systemd-cgroup create --bundle /run/containerd/io.        containerd.runtime.v2.task/k8s.io/f2ba685582ea81fb6568da6cec0ee3086dd8e15db11214ddfed19f4ded41725b --pid-file /run/containerd/io.containerd.runtime.v2. task/k8s.io/f2ba685582ea81fb6568da6cec0ee3086dd8e15db11214ddfed19f4ded41725b/init.pid f2ba685582ea81fb6568da6cec0ee3086dd8e15db11214ddfed19f4ded41725b  runc             20859  20799    0 /usr/bin/runc --root /run/containerd/runc/k8s.io --log /run/containerd/io.containerd.runtime.v2.task/k8s.io/         f2ba685582ea81fb6568da6cec0ee3086dd8e15db11214ddfed19f4ded41725b/log.json --log-format json --systemd-cgroup start                                      f2ba685582ea81fb6568da6cec0ee3086dd8e15db11214ddfed19f4ded41725b
sleep            20852  20799    0 /bin/sleep 30000000
```





shim v1 --> containerd-shim（**io.containerd.runtime.v1.linux**）

`io.containerd.runtime.v1.linux` 是最早的 shim 实现，使用shim v1 API，在 Containerd 1.0 之前被设计出来，目前 v1 版 shim API 已被废弃，并将于 Containerd 2.0 被删除。



Shim v2 runc v1 --> containerd-shim-runc-v1（**io.containerd.runc.v1**）

https://github.com/containerd/containerd/pull/2434 增加了shim v2 的支持，新增 containerd-shim-runc-v1 二进制文件，v1.1 版本支持。containerd-shim-{runtime-name}-{version} 是shim v2 的命令约定，后面使用 shim v2 都将使用该命名约定。



Shim v2 runc v2 --> containerd-shim-runc-v2 （**io.containerd.runc.v2**）

https://github.com/containerd/containerd/pull/3004 支持 runc v2，新增 containerd-shim-runc-v2 二进制文件，支持 container group）,v1.3 版本支持

containerd shim v1.4 支持 cgroup v2。





TODO：

1、Containerd shim 进程 PPID 为什么是 1？

containerd-shim 父进程是1号进程的原因:[containerd-shim parent process is not containerd,why separated them？](https://github.com/containerd/containerd/discussions/7458#discussioncomment-3770865)

```
1、为了分离 containerd 和 containerd shim 进程，避免升级 containerd 进程时影响 containerd-shim 进程，containerd 进程退出时如何影响 containerd-shim？？？ 需要分析
2、containerd-shim 进程是通过 os/exec 包中的 Start 方法启动的，最终执行 linux 的 forkAndExecInChild 函数，所以搞了半天最终的系统调用还不是 fork。。。execve 落实下来是一个名为 do_execve 的函数，使用 execve 系统调用创建出来的进程是全新的，不会从原进程复制进程结构。
/usr/bin/containerd-shim-runc-v2 。。。 start 退出后，/usr/bin/containerd-shim-runc-v2 就成了孤儿进程，孤儿进程会被 init 进程也即是 1 号进程（systemd）收养，这就是 containerd-shim 进程的 PPID 全都是 1 的原因。
```

containerd-shim 的父进程如果是 containerd，当 containerd 重启后 containerd-shim 的父进程则会发生变化。

containerd-shim 通过两次进程 fork 创建出来的，使 containerd 的父进程为 PID 1 的进程。



2、containerd shim 支持 group containerd 为什么 docker 不支持？

​	docker 是为容器设计的，和 k8s 的理念不一致，docker 层面感知不到 pod 的存在，所以不支持多个 container 使用同一个 shim 的功能。

3、shim 退出流程分析？





社区在 container shim 这块搞的替代品特别多，v1 版本的 shim 是每个容器需要启动一个 shim，在 19 年 containerd 社区开发除了 v2 版本的 shim io.containerd.runc.v2，可以支持一组 container 共用一个  shim，io.containerd.runc.v2 也是 containerd 当前默认的 shim，一个 containerd shim 大概需要 10M 内存，多个容器共用 shim 之后可以降低额外资源的消耗，此外 containerd 社区还有通过使用 rust 开发了一个运行时资源更小的 shim，以及也有单节点上所有容器共用一个 shim 的组件 [containerd-shim-systemd-v1](https://github.com/cpuguy83/containerd-shim-systemd-v1)。

containerd 默认使用 io.containerd.runc.v2 shim，kubelet 无论使用 dockershim 还是通过直连 containerd 创建出的 pod 都使用该shim，但 kubelet 只有在直连 containerd 时通过其 cri 接口创建出的 pod 下所有容器会公用一个 shim，经过翻阅 docker 的代码发现其在生成 runtime spec 时没有添加 io.containerd.runc.v2 shim 所需要用到的 Annotation 字段所以不支持多个 container 共用 shim 的功能。

如果 config.json 文件中有 annotations.io.kubernetes.cri.sandbox-id 则会支持 pod 下多个 container 共用一个shim的功能。

```
#cat /run/containerd/io.containerd.runtime.v2.task/k8s.io/08b0507c023aa051c92ae2c1b3a6aaa58d274eb754a08616adbbb4f9157541e0
...
  "annotations": {
    "io.kubernetes.cri.container-name": "nginx",
    "io.kubernetes.cri.container-type": "container",
    "io.kubernetes.cri.image-name": "docker.io/library/nginx:latest",
    "io.kubernetes.cri.sandbox-id": "b05389da37c6df685f2ef1c0ace29a7156d5facccc1df16b22afb06b9c7f7502",
    "io.kubernetes.cri.sandbox-name": "nginx-768c6f9745-mn4bb",
    "io.kubernetes.cri.sandbox-namespace": "default"
  },
...
```



4、shim v2 的作用？

shim v2 是为了让 runtime 的作者在 containerd 之外实现自己的 shim，比如 kata-shim。

shim v2 定义了



~~5、shim v1 与 shim v2 的区别：~~

~~shim v1 与 shim v2 最大的区别就是 v2 支持 group containerd~~







TODO：

1、config.json 是怎么生成的

Docker 请求 containerd 时直接传入 container 配置，containerd 使用传入的配置生成了 config.json,

使用 cri 创建容器时，cri server 会配置 container 的spec，然后服务内通过 grpc 调用 container service 创建 container 生成 config.json。



config.json 在 /run/containerd/io.containerd.runtime.v2.task/k8s.io/00d505fc63e71d3e018197877ef9a502b9e6532efd9b6ebbf247d0cabd845dfd 路径下可以看到，如果要使用 container groups per shim 功能，config.json 中 annotations 代码块下需要有 io.kubernetes.cri.sandbox-id label。



```
containerd support container groups per shim 
```





---

TODO：

1、runc v1 和 runc v2 的区别

2、containerd 与 shim 进程是如何关联并进行通信的？





一个 pod 一个 shim，如果是 sandbox 则会启动一个新的shim，如果是container 则会复用 sandbox 的shim。



##### runtime 历史:

containerd/plugin/plugin.go

```
const (
    // RuntimeLinuxV1 is the legacy linux runtime
    RuntimeLinuxV1 = "io.containerd.runtime.v1.linux"
    // RuntimeRuncV1 is the runc runtime that supports a single container
    RuntimeRuncV1 = "io.containerd.runc.v1"
    // RuntimeRuncV2 is the runc runtime that supports multiple containers per shim
    RuntimeRuncV2 = "io.containerd.runc.v2"
)
```


从 v1.3 之前默认的 runtime type 是 io.containerd.runtime.v1.linux，v1.3 开始默认的 runtime type 是 io.containerd.runc.v1，从 v1.4 开始默认的 runtime type 是 io.containerd.runc.v2。





#### shim API v2 propose

- 解决runtime状态管理的问题

- 支持其他运行时与 containerd 集成，比如 kata；

主要功能：

- 支持 `shim start` 子命令；
- 支持 `shim stop` 子命令；
- 使用 ttrpc 进行通信； 



shim v2 是什么？？



shim v2+（ runc v1 or runc v2）， not shim v1+ ？？？

shim-v1+runtime v1 与 v2 (2018-07)

runtime v1 commit（shim-v1）： https://github.com/containerd/containerd/pull/2434



shim v2+（ runc v1 or runc v2）+  container groups per shim ？？？

shim-v2+runtime v2 with container groups (2019-02)

io.containerd.runc.v2 with Container Groups（shim-v2）：https://github.com/containerd/containerd/pull/3004  支持一组容器使用一个shim

<img src="images/image-20220721202851000.png"></img>



io.containerd.runc.v2 和 io.containerd.runc.v1 都使用 shim v2 interface，io.containerd.runc.v1 是一个 shim 对应一个 container，io.containerd.runc.v2 支持一个 shim 对应一组 container。



shim v1 --> containerd-shim

```
shim v1 定义在 containerd/runtime/v1/shim/v1/shim_ttrpc.pb.go 中
```

shim v2 + runc v1 --> containerd-shim-runc-v1

```
shim v2 定义在 
containerd/api/runtime/task/v2/shim_ttrpc.pb.go 中

增加了命令行 containerd-shim-runc-v1 cmd 代码，shim v2 定义在 runtime/v2/shim ，runtime 支持 runc，runc 没有区分 v1 和 v2 版本 
```

shim v2 + runc v2 --> containerd-shim-runc-v2

```
增加了命令行 containerd-shim-runc-v2 cmd 代码，将 runtime/v2/runc 改为 runtime/v2/runc/v1 增加runtime/v2/runc/v2
```









```
查看容器列表
# ctr --namespace k8s.io c list

查看容器详细信息，可以看到同一pod 下不同容器中的 io.kubernetes.cri.sandbox-id 是相同的
# ctr --namespace k8s.io c info de2ef7862b3fea427f759ad6fab8a7349bbb312bc2dc31ec8f02fb535fd25657

        "annotations": {
            "io.kubernetes.cri.container-name": "kflax-core",
            "io.kubernetes.cri.container-type": "container",
            "io.kubernetes.cri.image-name": "registry.corp.com/ksp/public/containernetwork/v2/kflax-core:2.0.9",
            "io.kubernetes.cri.sandbox-id": "2863371c1ac3bababc0e550508342aac07547476e7ceec9c728cc3fb318bc0b9",
            "io.kubernetes.cri.sandbox-name": "kflax-core-5669b77d98-zvnpp",
            "io.kubernetes.cri.sandbox-namespace": "kube-system"
        },
        "linux": {
        
pause 与 业务容器中 io.kubernetes.cri.sandbox-id value 均为 sandbox 的 id。


// 测试
# ctr run --label io.containerd.runc.v2.group=test   --rm --runtime io.containerd.runc.v2  docker.io/library/redis:alpine demo1

# ctr run -t --label io.containerd.runc.v2.group=test   --rm --runtime io.containerd.runc.v2  docker.io/library/redis:alpine demo2 sh

  
# pstree -s -p 963007
systemd(1)───containerd-shim(963007)─┬─redis-server(963027)─┬─{redis-server}(963044)
                                     │                      ├─{redis-server}(963045)
                                     │                      ├─{redis-server}(963046)
                                     │                      └─{redis-server}(963047)
                                     ├─sh(964397)  

```





#### shim 启动流程

首先执行 start 子命令

```
# ./containerd-shim-runc-v2 -namespace k8s.io -id 98491f510c3e30c819a3d1b28407b2e8b16755a9d344466c6edcf07c25b26f09 -address /run/containerd/containerd.sock start
```

start 子命令如果执行成功会返回一个 unix socket 字符串，unix sockek 字符串会写入到标准输出。如果 start 子命令在执行期间有打印日志到到标准输出，则会影响 containerd-shim后续的执行流程。





然后使用没有 start 的子命令启动

```
# ./containerd-shim-runc-v2 -namespace k8s.io -id 98491f510c3e30c819a3d1b28407b2e8b16755a9d344466c6edcf07c25b26f09 -address /run/containerd/containerd.sock
```



调试日志：

<img src="images/image-20220804173502053.png"></img>


/usr/bin/containerd-shim-runc-v2 -namespace k8s.io -address /run/containerd/containerd.sock -publish-binary /usr/bin/containerd -id 08b0507c023aa051c92ae2c1b3a6aaa58d274eb754a08616adbbb4f9157541e0 -debug start

执行结果是返回一个 unix socket 的地址并写到标准输出

<img src="images/image-20220823195121132.png"></img>




TODO：

1、/run/containerd/s/ 下文件过多原因





参考：

[proposal Shim API v2](https://github.com/containerd/containerd/issues/2426)

https://nanikgolang.netlify.app/post/containers/

```
The shim allows for daemonless containers.  It basically sits as the parent of the container's process to facilitate a few things.

First it allows the runtimes, i.e. runc,to exit after it starts the container.  This way we don't have to have the long running runtime processes for containers.  When you start mysql you should only see the mysql process and the shim.

Second it keeps the STDIO and other fds open for the container incase containerd and/or docker both die.  If the shim was not running then the parent side of the pipes or the TTY master would be closed and the container would exit.  

Finally it allows the container's exit status to be reported back to a higher level tool like docker without having the be the actual parent of the container's process and do a wait.  
    
```



```
--- args  [create --bundle /run/containerd/io.containerd.runtime.v2.task/default/u67]
--- cmd  /home/nanik/AndroidProjects/docker/docker/runc --root /run/containerd/runc/default --log /run/containerd/io.containerd.runtime.v2.task/default/u67/log.json --log-format json create
--bundle /run/containerd/io.containerd.runtime.v2.task/default/u67 --pid-file /run/containerd/io.containerd.runtime.v2.task/default/u67/init.pid --console-socket /tmp/pty415594316/pty.sock u
67

The command used to execute 'runc' is as follows
"/home/nanik/AndroidProjects/docker/docker/runc --root /run/containerd/runc/default --log /run/containerd/io.containerd.runtime.v2.task/default/u67/log.json --log-format json create --bundle /run/containerd/io.containerd.runtime.v2.task/default/u67 --pid-file /run/containerd/io.containerd.runtime.v2.task/default/u67/init.pid --console-socket /tmp/pty415594316/pty.sock u67"
     
```

```
sudo ./runc --root /run/docker/runtime-runc/moby  list
 
ID                                                                 PID         STATUS      BUNDLE                                                                                                                 CREATED                          OWNER
f182f95645673b94af95495ea4c2a7c0f58dcce523f3d4e7174d7e482e136e08   12212       running     /run/containerd/io.containerd.runtime.v1.linux/moby/f182f95645673b94af95495ea4c2a7c0f58dcce523f3d4e7174d7e482e136e08   2020-01-05T21:28:05.371871841Z   root
  
```

```
The runtime (runc) uses so-called runtime root directory to store and obtain the information about containers. Under this root directory, runc places sub-directories (one per container), and each of them contains the state.json file, where the container state description resides.

The default location for runtime root directory is either /run/runc (for non-rootless containers) or $XDG_RUNTIME_DIR/runc (for rootless containers) - the latter also usually points to somewhere under /run (e.g. /run/user/$UID/runc).

When the container engine invokes runc, it may override the default runtime root directory and specify the custom one (--root option of runc). Docker uses this possibility, e.g. on my box, it specifies /run/docker/runtime-runc/moby as the runtime root.

That said, to make runc list see your Docker containers, you have to point it to Docker's runtime root directory by specifying --root option. Also, given that Docker containers are not rootless by default, you will need the appropriate privileges to access the runtime root (e.g. with sudo).

    So, that's how this should work:

    $ docker run -d alpine sleep 1000
    4acd4af5ba8da324b7a902618aeb3fd0b8fce39db5285546e1f80169f157fc69

    $ sudo runc --root /run/docker/runtime-runc/moby/ list
    ID                                                                 PID         STATUS      BUNDLE                                                                                                                               CREATED                          OWNER
    4acd4af5ba8da324b7a902618aeb3fd0b8fce39db5285546e1f80169f157fc69   18372       running     /run/docker/containerd/daemon/io.containerd.runtime.v1.linux/moby/4acd4af5ba8da324b7a902618aeb3fd0b8fce39db5285546e1f80169f157fc69   2019-07-12T17:33:23.401746168Z   root

As to images, you can not make runc see them, as it has no notion of image at all - instead, it operates on bundles. Creating the bundle (e.g. based on image) is responsibility of the caller (in your case - containerd).
  
```





### 参考：

[set containerd-shim log output to a specific log file](https://github.com/containerd/containerd/issues/2547)

[LINUX PID 1 和 SYSTEMD](https://coolshell.cn/articles/17998.html)

[What are the differences between runc v1 and runc v2 in containerd-shim-runc-v1 and containerd-shim-runc-v2?](https://github.com/containerd/containerd/discussions/7407)

[容器中的 Shim 到底是个什么鬼？](https://cloud.tencent.com/developer/article/1948155)

为什么要分离containerd-shim 和containerd：https://www.sobyte.net/post/2021-09/containerd-usage/






