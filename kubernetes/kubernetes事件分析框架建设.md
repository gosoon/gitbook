[toc]





主要分三类，实例创建问题（直接包含Ready节点的治理问题）、老实例删除问题、异常实例残留问题、notReady 节点的治理问题



#### 目标

1、感知集群中启动异常的实例，进行自愈处理；

2、感知集群中异常的节点，进行自愈处理；

3、感知集群中的残留实例，进行自愈处理；

4、感知节点的健康度，进行重调度/止损处理；

5、依赖组件的问题（ceph、镜像仓库）；

6、实例生命周期记录（TODO：podlist watch 事件行为）；



做云平台的价值主要有业务创新、成本优化、技术创新，但这些也都要在系统稳定的前提下进行，我们主要通过监听 kubernetes 中核心资源对象以及事件的变化，通过消息触发一些操作，建设事件分析的框架来加固平台的稳定性。



注意：以上所有的事件不仅指 k8s 中的 event，还有Pod状态变化产生的事件。



事件数据的分析，

一些关于k8s事件产生的问题：线上遇到过删除了deployment 但是 deployment 关联的资源没有被级联删除，后来发现是因为 k8s event etcd 出了异常导致部分事件丢失，

GarbageCollectorController 负责回收 kubernetes 中的资源，要回收 kubernetes 中所有资源首先得监控所有资源，GarbageCollectorController 会监听集群中所有可删除资源产生的所有事件，这些事件会被放入到一个队列中，然后 controller 会启动多个 goroutine 处理队列中的事件，若为删除事件会根据对象的删除策略删除关联的对象，对于非删除事件会更新对象之间的依赖关系。



![image-20211228105216678](/Users/feilei/Library/Application Support/typora-user-images/image-20211228105216678.png)



##### 针对实例创建问题的处理



实例创建问题

-1、资源不足 pending 实例的跨资源池与集群重调度：理论上在创建实例的时候上层通过全局资源视图可以看到当前有资源的节点，但在实际下发创建实例的任务时，因资源碎片、机器瞬间故障、并发量高等原因还是会出现资源不够导致实例 pending 问题；



-2、实例分配节点成功但无法创建，Pod 处于 pending 状态，无任何事件	

​		可能的原因： 

​		（1）controller 同步实例状态慢：（需要进一步分析是不是 event etcd 有问题或者新创建实例太多导致）

​		（2）kubelet 状态异常，节点处于Ready状态，Pod调度上去后一直处于 pending 状态

​				- 1、节点上 kubelet 部分 goroutine 卡住，之前有使用 ceph 盘的实例因 ceph 侧挂载异常，执行 df 也会卡住；

​				- 2、节点上有 D 进程导致 kubelet 调用超时，kubelet 某个 goroutine hang 住影响主流程，遇到过 cpu manager hang 住 case，kubelet 无法 receive pod；

​				-3、节点处于Ready状态，实例调度上去后创建失败

​					-1、systemd 组件存在 bug：节点缺少systemd/kubepods.slice目录的表象是节点Ready，但调度器分配到该机器后并不会执行创建pod的操作。实例卡在Pending并且存在已经分配到主机的事件；

​					-2、Dockerd /Containerd 进程异常：Dockerd 僵尸进程（defunct）、Containerd 进程退出；

​	  （3）pod 所需要的 pvc 未创建

​			storageclass 为正常创建

```
error while running "VolumeBinding" filter plugin for pod "cloud-1073396-104778-66408-97-deploy-c-6c67d6c7b5-t7cfq": pod has unbound immediate PersistentVolumeClaims
```





-3、Pod 调度成功，但节点实际资源不足，Pod 处于 RunContainerError、UnexpectedAdmissionError

​	  kubelet 节点资源变化导致	

​	磁盘满导致 pod 启动预检失败 - UnexpectedAdmissionError



-4、实例处于PodInitializing 状态

​	实例在创建过程中，启动initContainer时异常（TODO：寻找异常case）



-5、拉取镜像失败，实例处于 ImageInspectError、ErrPullImage、ImagePullBackOff

镜像层损坏问题修复方案：

```
背景介绍
在POD/容器创建时，经常遇到如下两个错误：

InspectFailed
Failed to inspect image "registry.corp.com/ksp/flow-runner:1.0.2462-master-duanhexin-5769959-container-snowman": rpc error: code = Unknown desc = Error response from daemon: readlink /media/disk1/docker/overlay2/l: invalid argument

FailedCreatePodSandBox
Failed to create pod sandbox: rpc error: code = Unknown desc = failed to inspect sandbox image "registry.corp.com/cloud_admin/pause:3.1": Error response from daemon: layer does not exist	

镜像损坏
预备知识
在讨论镜像层损坏前，先解释几个关键文件的用途
● committed  用于标识改层镜像commit层，具备父层，无实际内容
● link  用于记录该层镜像的短链接，避免mount参数过长
● lower  用户记录改层镜像的父层信息，为父层join链接而成

有些上面文件，特别是link和lower文件，overlay2镜像顶层视图正是通过这些信息，将多层联合，形成一个联合文件系统供容器使用。一旦某个镜像层的link和lowner文件有问题，必然影响到相关镜像的使用。

修复方式
业界做法
该问题17年支持overlay2开始就有提出，一直未解决，NPD原生支持检测，GKE通过检测和重建节点解决，更大众化的一般是删除整个overlay2目录，重新初始化docker，成本大。

内部方式
检测容易，修复难。由于镜像层的组织关系复杂，并且很多元数据信息维护在dockerd内存中，如果靠外挂检测直接删除问题镜像层后，会造成数据不一致问题。所以只能再dockerd内部打补丁。

具体实现方式为：
在镜像下载时，会检查各个镜像层是否存在，如果已经存在，跳过下载。在这个逻辑后，再补充一个镜像层的检查逻辑：如果镜像层损坏，删除镜像，并返回镜像层不存在，新下载镜像层。

同样，对于docker inspect而言，如果识别到镜像层损坏，删除该镜像层即可。
```



```
ErrPullImage：（1）镜像不存在；（2）镜像仓库镜像元数据损坏，导致manifest获取异常；
ImageInspectError：（1）镜像不存在；（2）本地镜像元数据损坏；
```



```
InspectFailed
Failed to inspect image "registry.corp.com/ksp/flow-runner:1.0.2462-master-duanhexin-5769959-container-snowman": rpc error: code = Unknown desc = Error response from daemon: readlink /media/disk1/docker/overlay2/l: invalid argument

FailedCreatePodSandBox
Failed to create pod sandbox: rpc error: code = Unknown desc = failed to inspect sandbox image "registry.corp.com/cloud_admin/pause:3.1": Error response from daemon: layer does not exist	
```





-6、实例处于 ContainerCreating

​	- 节点 Ready，但拉取镜像慢：kubelet 串行拉镜像源码分析（TODO）；

​	- 节点 Ready 但是实例创建失败问题，见下文对 node 的分析

​				(1) 节点上docker 异常，调用 docker 接口失败；

​				(2) 节点有高 IO，调用 docker 接口失败；

​				(3) 节点上 kubelet cgroup 配置与 docker 不一致；

​				(4) 节点在 Ready 与 NotReady 之间反复切换：systemd 组件存在 bug，节点缺少hugetlb/kubepods.slice目录的表象是kubelet反复重启，节点在Ready和NotReady之间反复跳动，实例卡在ContainerCreating

​			（5）低版本docker存在 bug，docker init D进程等问题导致节点 pleg nothealth，节点状态反复切换；

​				业务或者 runc 进程出现 D 状态时，会存在 kubelet pleg not health 问题，出现 pleg not health 问题节点会变成 notready ？？？ 需要分析代码

​	docker hang问题排查：[https://www.likakuli.com/posts/docker-hang/	](https://www.likakuli.com/posts/docker-hang/)

​	https://plpan.github.io/docker-hang-%E6%AD%BB%E9%98%BB%E5%A1%9E-kubelet-%E5%88%9D%E5%A7%8B%E5%8C%96%E6%B5%81%E7%A8%8B/

FROZEON问题定位以及修复方式 ：

```
背景
近期pump大数据业务和搜索均有反馈部分容器进程D状态，容器cgroup freezer子系统的freezer.state处于FROZEN状态，但容器又处于runing状态，不符合预期。

总结
问题源于低版本runc缺陷，runc update操作涉及FROZEN和THAWED，在设置前，会先将字段填充到runc容器的数据结构，如果设置失败，会回滚到之前的状态，容器最终会保存状态到runc容器的状态的state.json文件中。

runc update的FROZEN和THAWED操作没必要修改容器数据结构，如果设置失败，反而会影响到容器的状态。该问题社区已经在2021.07.09修复合入主干，runc 1.0.1分支采用并发布。
```





```
unable to ensure pod container exists: failed to create container for [kubepods podbd0cd7bc-5c7d-4956-8d56-727eda14f314] : dbus: connection closed by user
```

其他case：

```
Error: failed to start container "1218805-hb1-zw-staging-l28632-1": Error response from daemon: OCI runtime create failed: container_linux.go:364: creating new parent process caused: container_linux.go:2005: running lstat on namespace path "/proc/2575743/ns/ipc" caused: lstat /proc/2575743/ns/ipc: no such file or directory: unknown
```







##### 老实例删除问题

-1、terminating 实例卡住问题



Terminating 残留实例，因 docker 老本版 bug/性能问题 产生 init 进程 D 状态或者其他问题出现 Terminating 实例无法删除问题，首先会从集群层面删除该 Pod，然后会从节点层面清理该实例（自愈+kill 进程）；



针对 Terminating 实例处理时，首先按实例优雅退出时间添加定时器，如果超过优雅退出时间直接将实例从 k8s 侧强制删除，直接从 k8s 侧删除实例会存在节点上有容器残留的情况，我们在 npd 中集成了插件来检测节点上容器数与 pod 数量不符合的 case，若该容器通过各种约束条件判断为残留容器时会直接被stop掉（见 npd 优化一节）。

原因：

​	（1）节点负载高 docker 接口超时，删除时出现 FailedKillPod

​	（2）容器 runc init D 进程

优化：

   异步挂载docker各层镜像

##### 异常实例残留问题

-1、节点资源不足时，实例被驱逐，实例处于 failed、evicted等状态

​		节点异常产生的残留实例，比如磁盘满导致的 Evicted、Failed 实例，节点 notready 出现的 Unknown 实例，节点资源变化出现的 RunContainerError、UnexpectedAdmissionError（主要是混部相关的）

​	（TODO：RunContainerError、UnexpectedAdmissionError源码分析）

pod Failed 原因分析：blog



-2、节点异常，实例处于Unknown状态



##### 针对 NotReady 节点的处理

节点 NotReady 的原因

```
1、磁盘满；
2、节点 OOM；
3、老网卡问题导致节点异常重启；
4、/sys/fs/cgroup/hugetlb/kubepods.slice/kubepods-besteffort.slice 文件不存在，kubelet 启动异常
5、kubepods.slice 目录缺失，kubelet 启动失败
I1220 21:40:47.320303 3391776 cgroup_manager_linux.go:273] The Cgroup [kubepods] has some missing paths: [/sys/fs/cgroup/systemd/kubepods.slice /sys/fs/cgroup/cpu,cpuacct/kubepods.slice /sys/fs/cgroup/memory/kubepods.slice /sys/fs/cgroup/cpuset/kubepods.slice /sys/fs/cgroup/pids/kubepods.slice /sys/fs/cgroup/cpu,cpuacct/kubepods.slice]
Dec 20 21:40:47 sd-bjpg-rs113.yz02 kubelet[3391776]: E1220 21:40:47.322050 3391776 node_container_manager_linux.go:50] Failed to create ["kubepods"] cgroup
Dec 20 21:40:47 sd-bjpg-rs113.yz02 kubelet[3391776]: F1220 21:40:47.322063 3391776 kubelet.go:1400] Failed to start ContainerManager Cannot set property TasksAccounting, or unknown property.
Dec 20 21:40:47 sd-bjpg-rs113.yz02 kubelet[3391776]: goroutine 439 [running]:


6、GPU 机器掉卡问题：
Dec 21 14:18:48 public-kce-node-v100-gpu67.txsgp1.sgp.kwaidc.com kubelet[4101185]: W1221 14:18:48.629173 4101185 container.go:526] Failed to update stats for container "/kubepods/burstable/pod8ce5ef84-c24a-41e5-979d-936e29863442/97d660e7b73685e97cbbf1fd5fd5b4264e7ce2e8a2f5b239570002c73c937f63": error while getting gpu utilization: nvml: Not Found


7、pleg not health
（1）容器 freezer 子进程为 THAWED，docker inspect 卡住
# cat freezer.state
THAWED


Dec 21 14:02:34 public-bjzey-rs13-kce-node32.idczw.hb1.kwaidc.com kubelet[3247450]: E1221 14:02:34.364611 3247450 kubelet.go:1926] skipping pod synchronization - PLEG is not healthy: pleg was last seen active 3m0.048796938s ago; threshold is 3m0s
。。。。。。
Dec 21 14:02:37 public-bjzey-rs13-kce-node32.idczw.hb1.kwaidc.com kubelet[3247450]: I1221 14:02:37.903884 3247450 kubelet_node_status.go:494] Recording NodeNotReady event message for node public-bjzey-rs13-kce-node32.idczw.hb1.kwaidc.com
Dec 21 14:02:37 public-bjzey-rs13-kce-node32.idczw.hb1.kwaidc.com kubelet[3247450]: I1221 14:02:37.903896 3247450 setters.go:544] Node became not ready: {Type:Ready Status:False LastHeartbeatTime:2021-12-21 14:02:37.90387372 +0800 CST m=+409436.437901516 LastTransitionTime:2021-12-21 14:02:37.90387372 +0800 CST m=+409436.437901516 Reason:KubeletNotReady Message:PLEG is not healthy: pleg was last seen active 3m3.588072819s ago; threshold is 3m0s}

8、博通网卡问题导致节点反复重启，cgroup子目录丢失，实例创建卡住。

9、docker inspect细分原因非常多，常见的有：管道容量不够（低版本docker），业务僵尸进程或者D状态，runc update/init/ps/kill以及shim  reaper持锁卡住等，进而docker inspect无法获取到容器锁，由于docker inspect没有超时机制，也会卡住，最终导致kubelet PLEG卡住，节点NotReady。
修复方式
容器互斥锁优化成Try锁，以一定频率100ms不断尝试上锁，超时时间为60s，超时退出。具体实现如下：
```



containerd  中有很多地方没有设置超时时间会导致某些操作死锁或者超时问题



runc ：

https://github.com/opencontainers/runc/issues/2865

https://github.com/opencontainers/runc/issues/2530（runc init hang）



#### kubernetes 资源对象变化以及事件监听机制

目前有很多的开源项目都支持采集 k8s 中的 event 然后对接到其他组件，例如：

EventRouter: *https://github.com/heptiolabs/eventrouter*

Kubernetes Event Exporter: *https://github.com/caicloud/event_exporter*

https://github.com/AliyunContainerService/kube-eventer



k8s 中的 event 采集不复杂，只需要数十行代码即可搞定，采集到事件后要进行分析与存储，可以 用来做一些针对 pod 的自愈操作以及回溯 pod 的生命周期。



针对所需要对象的变化以及集群的事件，通过 informer watch pod 和 event 对象。

以下是针对 pod 资源对象变化的监听：

```
func startListWatchPods(clientset *kubernetes.Clientset) corev1listers.PodLister {
	podListWatcher := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), "pods", v1.NamespaceAll, fields.Everything())
	indexer, informer := cache.NewIndexerInformer(podListWatcher, &v1.Pod{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(*v1.Pod)

			// generate pod message and send to kafka
			generatePodMessageToSend()

			glog.Infof("Receive add pod event. Key: %s/%s UID: %s ResourceVersion: %s.", pod.Namespace, pod.Name, string(pod.UID), pod.ResourceVersion)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			pod := newObj.(*v1.Pod)

			// generate pod message and send to kafka
			generatePodMessageToSend()

			glog.Infof("Receive update pod event. Key: %s/%s UID: %s ResourceVersion: %s.", pod.Namespace, pod.Name, string(pod.UID), pod.ResourceVersion)
		},
		DeleteFunc: func(obj interface{}) {
			pod, ok := obj.(*v1.Pod)

			if !ok {
				tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
				if !ok {
					glog.Errorf("couldn't get object from tombstone %#v", obj)
					return
				}

				pod, ok = tombstone.Obj.(*v1.Pod)
				if !ok {
					glog.Errorf("tombstone contained object that is not a pod %#v", obj)
					return
				}
			}

			// generate pod message and send to kafka
			generatePodMessageToSend()

			glog.Infof("Receive delete pod event. Key: %s/%s UID: %s ResourceVersion: %s.", pod.Namespace, pod.Name, string(pod.UID), pod.ResourceVersion)
		},
	}, cache.Indexers{})

	informer.Run(wait.NeverStop)
	// Wait for all involved caches to be synced, before processing items from the queue is started
	if !cache.WaitForCacheSync(wait.NeverStop, informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return nil
	}

	// TODO: check storage and informer have diff periodicity

	return corev1listers.NewPodLister(indexer)
}
```





##### NPD 的使用

针对异常节点以及部分异常实例主要使用 NPD 来进行自愈/止损操作。

主要使用 npd 的 CustomPluginMonitor，plugin 会上报关联节点的事件，在收到对应的事件后会进行一些自愈操作。



##### 对 NPD 的优化

NPD 只缓存了节点信息，没有缓存节点 pod 等信息，对 npd 源码就行了修改可以直接 list-watch 该节点上的实例列表。



```
func NewLister(clientset *kubernetes.Clientset, nodeName string) corelisters.PodLister {
	// create the pod watcher
	podListWatcher := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), "pods", v1.NamespaceDefault, fields.OneTermEqualSelector("spec.nodeName", nodeName))

	// create the workqueue
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	// Bind the workqueue to a cache with the help of an informer. This way we make sure that
	// whenever the cache is updated, the pod key is added to the workqueue.
	// Note that when we finally process the item from the workqueue, we might see a newer version
	// of the Pod than the version which was responsible for triggering the update.
	indexer, informer := cache.NewIndexerInformer(podListWatcher, &v1.Pod{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(old interface{}, new interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(new)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			// IndexerInformer uses a delta queue, therefore for deletes we have to use this
			// key function.
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
	}, cache.Indexers{})

	go informer.Run(wait.NeverStop)

	// Wait for all involved caches to be synced, before processing items from the queue is started
	if !cache.WaitForCacheSync(wait.NeverStop, informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return nil
	}

	podLister := corelisters.NewPodLister(indexer)
	return podLister
}
```



NPD 需要 list-watch master 获取当前节点上的实例列表，NPD 以 DS 的方式部署，如果在一个超大集群中， NPD 关联的所有 pod 启动时都要连接 master，此时对 master 会有一个较大的压力，极端情况下可能将 master 打挂（OOM？连接数打满？），此时设想方案是在每个 NPD 实例连接 master 时 sleep 一个 10 分钟内的随机数来避免热点问题。





镜像压缩以及镜像预热。



**热加载配置**



**事件与自愈操作分离**

在某些特殊场景下可以从外部将自愈的措施关闭。



**docker 版本升级后**

如果 npd 中挂载了节点上的 docker.sock 文件，当节点上 docker 版本升级后，npd 中的 docker.sock 会无法使用，需要重启 npd 实例重新挂载 docker.sock。



TODO：

NPD 的配置文件和脚本目前是打包在镜像里（主要是 CustomPluginMonitor），如果配置或脚本有调整，需要重新构建镜像并上线，流程上比较慢，此处是想将脚本放到数据库中，NPD 在执行过程中可以直接从数据获取脚本并进行热加载，



##### Docker 问题总结

```
// docker cgroup泄漏异常
DOCKER_CGROUP_LET_OUT_ERR1="/sys/fs/cgroup/memory/"
DOCKER_CGROUP_LET_OUT_ERR2="no space left on device"

// docker创建实例名称冲突
DOCKER_CREATE_POD_CONFLICT_ERR="You have to remove (or rename) that container to be able to reuse that name"

// docker无法下载镜像异常
DOCKER_LAYER_NOTEXIST_ERR="layer does not exist"

// docker19版本前的mount异常
DOCKER_DEVICE_BUSY_ERROR="merged: device or resource busy"

// docker19.08 版本后的mount异常
DOCKER_MOUNT_ERROR1="Error response from daemon: OCI runtime create failed: container_linux.go"

DOCKER_MOUNT_ERROR2="running exec setns process for init caused"

// 
KUBELET_KMEM_ERROR="PLEG is not healthy"

# docker sock error
DOCKER_WRITE_BROKEN_ERROR="write unix /var/run/docker.sock->@: write: broken pipe"

# docker_nvidia_plugin_error
DOCKER_NVIDIA_PLUGIN_ERROR="listAndWatch ended unexpectedly for device plugin nvidia.com/gpu with error rpc error: code = Unavailable   desc = transport is closing"

# docker DOCKER_GPU_NVIDIA_ERROR
DOCKER_GPU_NVIDIA_ERROR="nvidia-container-cli: initialization error"

# docker timeout error 
DOCKER_CREATE_ERR="failed to create a sandbox for pod"
DOCKER_TIMEOUT_ERR="operation timeout: context deadline exceeded"

# docker stop error 
DOCKER_STOP_ERR="StopContainer"

# kernel xfs deadlock error
DOCKER_XFS_DEADLOCK_ERR="possible memory allocation deadlock"
```





##### docker hang 死问题

1、docker exec 

runc init阻塞在写pipe操作 -- 》 进而导致 docker stats 卡住 --》 kubelet pleg nothealth



小结一下：containerd-shim启动runc exec去容器内执行用户命令，而runc exec启动runc init进入容器时，由于往2号FD写数据超过限制大小而被阻塞。当最底层的runc init被阻塞时，造成了调用链路上所有进程都被阻塞：

```
runc init → runc exec → containerd-shim exec → containerd exec → dockerd exec
```



##### docker 补丁

解决：1.镜像cache元数据损坏inspect invalid问题；2.docker inspect业务容器卡住PLEG  NotReady问题；3.解决containerd-shim异常退出后docker容器泄漏问题





参考：https://plpan.github.io/docker-hang-%E6%AD%BB%E6%8E%92%E6%9F%A5%E4%B9%8B%E6%97%85/

2、







TODO ： 

1、架构图

2、runc D 状态和进程D状态下，哪些case 会导致 kubelet 出现 NotReady，哪些case 不会导致 kubelet 出现 NotReady



Pod Pending 问题分析：

1、节点上部分容器中有 D 进程，kubelet 在某些操作获取容器信息时会卡住，导致主进程卡住，目前发现有 houseKeeper 和 cpumanager 状态同步卡住问题，ref：docs；

2、







Kubernetes Pod 关闭流程：https://blog.csdn.net/qq_35753140/article/details/105945335



参考：

https://mp.weixin.qq.com/s/csY8T02Ck8bnE3vVcZxVjQ






