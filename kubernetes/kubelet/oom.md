1. k8s 中容器在 oom 时会被内核使用 kill -9 强行杀掉，可以看到容器的退出原因是OOMKilled，退出码是137,退出码为何是137而不是9？
containerd-shim在处理时会将加上一个偏移量128，137 = 9 + 128。


ref:
[kubernetes pod内容器状态OOMKilled和退出码137全流程解析](https://www.jianshu.com/p/0a9718199428)
