# etcd 命令行示例
## 写数据
```shell
# https
etcdctl --endpoints="https://127.0.0.1:12379" \
    --cert="/tmp/etcd-certs/certs/127.0.0.1.pem" \
    --key="/tmp/etcd-certs/certs/127.0.0.1-key.pem" \
    --cacert="/tmp/etcd-certs/certs/ca.pem" \
     put /a A
     
# http
etcdctl --endpoints="http://127.0.0.1:12379" \
    put /a A
```

## 查数据
```shell
# 获取以 / 开头的
etcdctl get --prefix /

# 获取数据细节
etcdctl get /a -wjson
```

## 删数据
```shell
etcdctl --endpoints="https://127.0.0.1:12379" --cert="/tmp/etcd-certs/certs/127.0.0.1.pem" --key="/tmp/etcd-certs/certs/127.0.0.1-key.pem" --cacert="/tmp/etcd-certs/certs/ca.pem" del /a
```