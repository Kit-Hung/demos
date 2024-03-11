# 通过 cluster api 安装集群

https://cluster-api.sigs.k8s.io/user/quick-start.html


## 通过 kind 安装管理集群，使用 docker 作为基础设施提供者
### 生成配置文件
```shell
cat > kind-cluster-with-extramounts.yaml <<EOF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  ipFamily: dual
nodes:
- role: control-plane
  extraMounts:
    - hostPath: /var/run/docker.sock
      containerPath: /var/run/docker.sock
EOF
```

### 创建管理集群
```shell
kind create cluster --config kind-cluster-with-extramounts.yaml
```



## 下载二进制
```shell
curl -L https://github.com/kubernetes-sigs/cluster-api/releases/download/v1.6.2/clusterctl-linux-amd64 -o clusterctl
```

## 初始化 provider
```shell
# Enable the experimental Cluster topology feature.
export CLUSTER_TOPOLOGY=true

# Enable the experimental Machine Pool feature
export EXP_MACHINE_POOL=true

# Initialize the management cluster
clusterctl init --infrastructure docker
```

## 生成集群配置
```shell
# The list of service CIDR, default ["10.128.0.0/12"]
export SERVICE_CIDR=["10.96.0.0/12"]

# The list of pod CIDR, default ["192.168.0.0/16"]
export POD_CIDR=["192.168.0.0/16"]

# The service domain, default "cluster.local"
export SERVICE_DOMAIN="k8s.test"

clusterctl generate cluster capi-quickstart --flavor development \
  --kubernetes-version v1.29.2 \
  --control-plane-machine-count=1 \
  --worker-machine-count=1 \
  > capi-quickstart.yaml
```

## 创建工作集群
```shell
kubectl apply -f capi-quickstart.yaml
```

## 获取工作集群的 kubeconfig
```shell
clusterctl get kubeconfig capi-quickstart > capi-quickstart.kubeconfig
```

## 访问工作集群
```shell
docker ps | grep lb
# 01072213195c   kindest/haproxy:v20230510-486859a6   "haproxy -W -db -f /…"   29 minutes ago   Up 29 minutes   0/tcp, 0.0.0.0:32768->6443/tcp     capi-quickstart-lb

kubectl get nodes --kubeconfig capi-quickstart.kubeconfig --server https://127.0.0.1:32768
# NAME                                     STATUS     ROLES           AGE     VERSION
# capi-quickstart-md-0-blg4k-f82dl-rxl9x   NotReady   <none>          6m3s    v1.29.2
# capi-quickstart-nsxtv-pwdnw              NotReady   control-plane   8m12s   v1.29.2
# capi-quickstart-worker-o3fucl            NotReady   <none>          6m3s    v1.29.2

```

## 安装网络插件
```shell
kubectl --kubeconfig capi-quickstart.kubeconfig --server https://127.0.0.1:32768 apply -f https://github.com/projectcalico/calico/blob/v3.27.2/manifests/calico.yaml

kubectl --kubeconfig capi-quickstart.kubeconfig --server https://127.0.0.1:32768 get nodes
# NAME                                     STATUS   ROLES           AGE   VERSION
# capi-quickstart-md-0-blg4k-f82dl-rxl9x   Ready    <none>          11m   v1.29.2
# capi-quickstart-nsxtv-pwdnw              Ready    control-plane   13m   v1.29.2
# capi-quickstart-worker-o3fucl            Ready    <none>          11m   v1.29.2

```
