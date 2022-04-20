docker-containerd-ctr、ctr、crictl、docker-runc 等命令的区别与使用 



##### 1.ctr 命令的使用

ctr 命令使用方法与 docker-containerd-ctr 一致，仅是别名关系。

```
# docker-containerd-ctr -a /var/run/docker/containerd/docker-containerd.sock -n moby c ls

# ctr -a /run/containerd/containerd.sock -n moby c ls
# ctr -a /run/containerd/containerd.sock -n moby c  info 9b657f145bdb7a2162ade865f8f0372a2226414180337cea53c5a2f49b52ad0e


# docker-containerd-ctr -a /var/run/docker/containerd/docker-containerd.sock -n moby c ls | grep -v IMAGE | awk '{print $1}' | xargs -ti docker-containerd-ctr -a /var/run/docker/containerd/docker-containerd.sock -n moby t ps {}

# ctr -namespace moby shim --id 2b5d28320969e2ab962d83990a17f08895e2dc731ee8953406cc9ce5ed7b0bfd  state
```



##### 2.runc 命令使用

```
// 查看容器的状态
$ runc --root /var/run/docker/runtime-runc/moby/ list

// 在容器内执行命令
$ runc --root /var/run/docker/runtime-runc/moby/ exec  29f6af06322ce8efd65095bcebca51f5086572083c76020005058cc0e15c3ad4 ls
```



##### 3.docker 接口使用：

```
$ curl -XGET --unix-socket /var/run/docker.sock http://localhost/v1.20/_ping
 
$ curl -XGET --unix-socket /var/run/docker.sock http://localhost/v1.19/containers/json?all=1
 
$ curl -v -XPOST -H "Content-Type: application/json" --unix-socket /var/run/docker.sock http://localhost/v1.19/containers/create -d@c.json

$ curl --unix-socket /var/run/docker.sock "http://1.38/containers/create?name=test01" -v -X POST -H "Content-Type: application/json" -d '{"Image": "nginx:latest"}'
```
