### 查看当前系统的 namespace
```shell
lsns -t <type>
```


### 查看某进程的 namespace
```shell
ls -la /proc/<pid>/ns/
```


### 进入某 namespace
```shell
nsenter -t <pid> -n ip a
```


### 在新的 network namespace 中执行 sleep 命令
```shell
unshare -fn sleep 60
```