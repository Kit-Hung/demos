# 灰度发布
## 1. 检查版本信息
```shell
kubectl get rev

NAME            CONFIG NAME   GENERATION   READY   REASON   ACTUAL REPLICAS   DESIRED REPLICAS
http-fn-00001   http-fn       1            True             0                 0
http-fn-00002   http-fn       2            True             0                 0
http-fn-00003   http-fn       3            True             0                 0
http-fn-00004   http-fn       4            True             0                 0

```


## 2. 检查流量配置
```shell
kubectl get ksvc http-fn -o yaml

traffic:
  - latestRevision: true
    percent: 100
```


## 3. 更新流量规则
canary-ksvc.yaml

```yaml
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: http-fn
spec:
  traffic:
    - revisionName: http-fn-00002
      percent: 50
    - revisionName: http-fn-00003
      percent: 50
```

```shell
kubectl apply -f canary-ksvc.yaml
```


## 4. 测试
```shell
kubectl get pods -w

curl http://http-fn.default.svc.cluster.local
```