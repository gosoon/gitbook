#### 1、镜像预热

自研 CRD 批量分发 job 到节点拉取镜像。

#### 2、镜像瘦身，去掉冗余文件

#### 3、镜像压缩解压缩优化

gzip/gunzip 是单线程的压缩/解压工具，可考虑采用 pigz/unpigz 进行多线程的压缩/解压，充分利用多核优势，解压效率能提高20%左右。

containerd 从 1.2 版本开始支持 pigz，节点上安装 unpigz 工具后，会优先用其进行解压。通过这种方法，可通过节点多核能力提升镜像解压效率。

这个过程也需要关注 下载/上传 的并发度问题，docker daemon 提供了两个参数来控制并发度，控制并行处理的镜像层的数量，--max-concurrent-downloads 和 --max-concurrent-uploads。默认情况下，下载的并发度是 3，上传的并发度是 5，可根据测试结果调整到合适的值。

#### 4、改造 docker 使用非压缩镜像

如果内网带宽足够大，可以去掉压缩环境，docker push 时默认会压缩镜像（gunzip 和 gzip），去掉压缩后 docker push 的时候无需执行压缩操作，docker pull 的时候也无需进行解压缩操作，不过可能需要修改 docker 源码实现。

不过目前的场景来说不太合适，磁盘占用量就非常大了。

#### 5、按需加载镜像

容器启动时其实需要镜像中很少的数据，可以对镜像进行分层然后按需加载，即在拉取镜像过程中就直接启动容器

#### 6、Nydus镜像懒加载

[Nydus](https://github.com/dragonflyoss/image-service)是蚂蚁开源的容器镜像加速项目，它在最新的OCI Image-Spec基础之上，重新设计了镜像格式和底层文件系统，从而加速容器启动速度，提高大规模集群中的容器启动效率。

#### 7、原地升级，patch image

#### 8、镜像P2P分发



参考：

[Serverless 场景下 Pod 创建效率优化](https://mp.weixin.qq.com/s/0OLdyVwg4Nsw0Xvvg8if5w)
[镜像搬运工 skopeo 初体验](https://blog.k8s.li/skopeo.html)
