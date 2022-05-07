kubelet 使用 --container-runtime 参数支持配置不同的 runtime：

```
--container-runtime=remote --container-runtime-endpoint=unix:///run/containerd/containerd.sock
```



#### ctr vs crictl

ctr是containerd本身的CLI，根据CRI规范，Kubernetes社区也定义了专门CLI工具crictl（https://github.com/kubernetes-sigs/cri-tools）



#### containerd VS cri-o 

从**功能性**来讲，containerd 和 cri-o 都符合 CRI 和 OCI 的标准。从**稳定性**来说，单独使用 containerd 和 cri-o 都没有足够的生产环境经验。但庆幸的是，containerd 一直在 docker 里使用，而 docker 的生产环境经验可以说比较充足。**可见在稳定性上 containerd 略胜一筹。所以我们最终选用了 containerd**。





参考：

[容器运行时从 docker 到 containerd 的迁移](https://www.infoq.cn/article/odslclsjvo8bnx*mbrbk)

[容器中的 Shim 到底是个什么鬼？](https://mp.weixin.qq.com/s/Dr6851XnkNLVFHaj1b13RQ)
