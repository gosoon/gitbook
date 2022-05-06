Docker 运行过程中的一些异常，因单机或节点上基础组件问题导致的。

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
