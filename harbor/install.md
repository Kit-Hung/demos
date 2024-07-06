# Harbor 安装
## 下载 helm chart
```shell
helm repo add harbor https://helm.goharbor.io
helm fetch harbor/harbor --untar
```

## 修改配置
```shell
vim harbor/values.yaml

# 测试修改的配置
expose:
  type: nodePort
tls:
  commonName: 'core.harbor.domain'

persistence: false
```

## 安装 Harbor
```shell
helm install harbor ./harbor -n harbor
```

## 获取证书
```shell
# 页面
https://10.105.111.219/harbor/configs/setting

# 接口
https://10.105.111.219/api/v2.0/systeminfo/getcert
```

## 配置 docker
```shell
mkdir -p /etc/docker/certs.d/core.harbor.domain
cp ca.crt /etc/docker/certs.d/core.harbor.domain
systemctl restart docker

docker login -u admin -p Harbor12345 core.harbor.domain
```

## 配置 containerd
```shell
mkdir -p /etc/containerd/certs.d/core.harbor.domain
cp ca.crt /etc/containerd/certs.d/core.harbor.domain/

vim /etc/containerd/certs.d/core.harbor.domain/hosts.toml
server = "https://core.harbor.domain"

[host."http://core.harbor.domain"]
  capabilities = ["pull", "resolve", "push"]
  skip_verify = true
  ca = "ca.crt"

systemctl restart containerd
```

## 修改 host
```shell
# harbor svc 的 cluster ip
10.105.111.219 core.harbor.domain
```

## 查看 repositories 和 blobs
```shell
kubectl -n harbor exec -it harbor-registry-7886456f94-vkfv5 -- bash

ls -la /storage/docker/registry/v2/repositories/
ls -la /storage/docker/registry/v2/blobs/
```

## 查看数据库数据
```shell
kubectl exec -it harbor-database-0 -- bash

psql -U postgres -d postgres -h 127.0.0.1 -p 5432
\c registry
select * from harbor_user;
```