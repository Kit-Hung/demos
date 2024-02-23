# 安装 etcd

## 手工安装
### 安装 etcd 二进制
```shell
tar -zxvf etcd-v3.5.12-linux-amd64.tar.gz
cp etcd-v3.5.12-linux-amd64/etcd* /usr/local/bin/
```

### 生成证书
```shell
# 安装 cfssl
apt install golang-cfssl

# 生成证书
git clone https://github.com/etcd-io/etcd.git
cd etcd/hack/tls-setup/

# 只保留需要的 host，如 127.0.0.1 和 localhost
vim config/req-csr.json
make

# 把生成的证书移动到所需目录，如： /tmp
mkdir /tmp/etcd-certs
mv certs/ /tmp/etcd-certs/
```

### 启动命令
#### 单节点启动
##### 不带证书
```shell
etcd --listen-client-urls 'http://localhost:12379' \
 --advertise-client-urls 'http://localhost:12379' \
 --listen-peer-urls 'http://localhost:12380' \
 --initial-advertise-peer-urls 'http://localhost:12380' \
 --initial-cluster 'default=http://localhost:12380'
```


##### 带证书
```shell
etcd --listen-client-urls 'https://localhost:12379' \
 --advertise-client-urls 'https://localhost:12379' \
 --listen-peer-urls 'https://localhost:12380' \
 --initial-advertise-peer-urls 'https://localhost:12380' \
 --initial-cluster 'default=https://localhost:12380' \
 --client-cert-auth --trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
 --cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
 --key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem \
 --peer-client-cert-auth --peer-trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
 --peer-cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
 --peer-key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem
```


#### 集群启动
```shell
nohup etcd --name infra0 \
--data-dir=/tmp/etcd/infra0 \
--listen-peer-urls https://127.0.0.1:3380 \
--initial-advertise-peer-urls https://127.0.0.1:3380 \
--listen-client-urls https://127.0.0.1:3379 \
--advertise-client-urls https://127.0.0.1:3379 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster infra0=https://127.0.0.1:3380,infra1=https://127.0.0.1:4380,infra2=https://127.0.0.1:5380 \
--initial-cluster-state new \
--client-cert-auth --trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem \
--peer-client-cert-auth --peer-trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--peer-cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--peer-key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem 2>&1 > /var/log/infra0.log &

nohup etcd --name infra1 \
--data-dir=/tmp/etcd/infra1 \
--listen-peer-urls https://127.0.0.1:4380 \
--initial-advertise-peer-urls https://127.0.0.1:4380 \
--listen-client-urls https://127.0.0.1:4379 \
--advertise-client-urls https://127.0.0.1:4379 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster infra0=https://127.0.0.1:3380,infra1=https://127.0.0.1:4380,infra2=https://127.0.0.1:5380 \
--initial-cluster-state new \
--client-cert-auth --trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem \
--peer-client-cert-auth --peer-trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--peer-cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--peer-key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem 2>&1 > /var/log/infra1.log &

nohup etcd --name infra2 \
--data-dir=/tmp/etcd/infra2 \
--listen-peer-urls https://127.0.0.1:5380 \
--initial-advertise-peer-urls https://127.0.0.1:5380 \
--listen-client-urls https://127.0.0.1:5379 \
--advertise-client-urls https://127.0.0.1:5379 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster infra0=https://127.0.0.1:3380,infra1=https://127.0.0.1:4380,infra2=https://127.0.0.1:5380 \
--initial-cluster-state new \
--client-cert-auth --trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem \
--peer-client-cert-auth --peer-trusted-ca-file=/tmp/etcd-certs/certs/ca.pem \
--peer-cert-file=/tmp/etcd-certs/certs/127.0.0.1.pem \
--peer-key-file=/tmp/etcd-certs/certs/127.0.0.1-key.pem 2>&1 > /var/log/infra2.log &
```


## helm 安装
```shell
helm repo add bitnami https://charts.bitnami.com/bitnami
helm pull bitnami/etcd
tar -zxvf etcd-9.14.2.tgz

# 修改相关属性
vim etcd/values.yaml

# 安装到 k8s 集群
helm install my-etcd ./etcd

# 运行 client pod
kubectl run my-etcd-client --restart='Never' --image docker.io/bitnami/etcd:3.5.12-debian-12-r7 --env ROOT_PASSWORD=$(kubectl get secret --namespace default my-etcd -o jsonpath="{.data.etcd-root-password}" | base64 -d) --env ETCDCTL_ENDPOINTS="my-etcd.default.svc.cluster.local:2379" --namespace default --command -- sleep infinity

# 通过 client 进行数据操作
kubectl exec --namespace default -it my-etcd-client -- bash
etcdctl --user root:$ROOT_PASSWORD put /message Hello
etcdctl --user root:$ROOT_PASSWORD get /message

# 提供外部访问
 kubectl port-forward --namespace default svc/my-etcd 2379:2379 &
    echo "etcd URL: http://127.0.0.1:2379"
    
# 获取密码
export ETCD_ROOT_PASSWORD=$(kubectl get secret --namespace default my-etcd -o jsonpath="{.data.etcd-root-password}" | base64 -d)
```