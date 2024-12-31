## 1. 创建 event 类型的函数
```shell
func create -l go -t cloudevents event-fn
```

## 2. 设置 func.yaml 配置文件
```yaml
specVersion: 0.36.0
name: http-fn
runtime: go
registry: core.harbor.domain/knative
image: core.harbor.domain/knative/http-fn:v1
created: 2024-08-16T10:13:23.213293953+08:00
build:
  builder: pack
  buildEnvs:
    - name: GOPROXY
      value: https://goproxy.cn,direct
run:
  volumes: []
  envs: []
deploy:
  namespace: default
  annotations: {}
  options: {}
  labels: []
  healthEndpoints:
    liveness: /health/liveness
    readiness: /health/readiness
```


## 3. 部署
```shell
func deploy
```

## 4. 创建 broker 和 trigger
```shell
kubectl apply -f broker_trigger_channel.yaml
```


## 5. 查看相关对象
```shell

```


## 4. 部署 event display
```shell
kubectl apply -f event_display.yaml
```


