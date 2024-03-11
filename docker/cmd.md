### 进入容器 ns
```shell
pid=$(docker inspect --format='{{.State.Pid}}' $container_id)

# 进入容器的网络空间
nsenter -t $pid -n
```