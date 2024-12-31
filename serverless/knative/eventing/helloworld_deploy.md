## 镜像打包
```shell
cd helloworld-go
docker buildx build --platform linux/arm64,linux/amd64 -t "core.harbor.domain/knative/helloworld-go:v1" --push .
```

## 部署 default channel
```shell
kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v1.15.1/in-memory-channel.yaml
```


## 部署应用
```shell
kubectl apply -f broker_trigger.yaml
kn service create helloworld-go --image core.harbor.domain/knative/helloworld-go:v1 namespace: knative-samples
```

## 查看资源
```shell
kubectl -n knative-samples get broker
kubectl -n knative-samples get deployment
kubectl -n knative-samples get svc
kubectl -n knative-samples get trigger
```

## 给 broker 发送 cloudEvent
```shell
# 获取 broker 地址

```