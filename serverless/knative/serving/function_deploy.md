## 1. 创建函数，指定为 go 语言
```shell
func create -l go http-fn
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

## 3. 编译
```shell
func build 
```

## 4. 部署
```shell
func deploy --build=false
```

## 5. 测试
```shell
kubectl get ksvc

NAME      URL                                        LATESTCREATED   LATESTREADY     READY   REASON 
http-fn   http://http-fn.default.svc.cluster.local   http-fn-00001   http-fn-00001   True   
```

## mount secret
```shell
kubectl apply -f test-secret.yaml
kn service update http-fn --mount /test-secret=secret:test-secret
```