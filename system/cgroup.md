## cpu
### 在 cgroup cpu 子系统目录中
```shell
cd /sys/fs/cgroup/cpu
mkdir cpudemo
cd cpudemo
echo $pid > cgroup.procs
echo 10000 > cpu.cfs_quota_us
```