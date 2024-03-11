# 使用 kind 创建集群
## 获取 kind 二进制文件
```shell
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.22.0/kind-linux-amd64
```

## 通过 kind 创建集群
```shell
kind create cluster
```

## 获取集群
```shell
kind get clusters
```