# 创建并部署函数
```shell
# 创建函数 fn
func create -l go fn
cd fn

# 编译
func build --registry core.harbor.domain/knative

# 部署
func deploy --build=false
```

func.yaml
```yaml
specVersion: 0.36.0
name: fn
runtime: go
registry: core.harbor.domain/knative
created: 2024-08-07T18:01:38.522683932+08:00
build:
  builder: pack
  buildEnvs:
  - name: GOPROXY    # 编译需要代理
    value: https://goproxy.cn,direct

```

# 部署后查看状态
```shell
kubectl get ksvc

NAME   URL                                   LATESTCREATED   LATESTREADY   READY   REASON
fn     http://fn.default.svc.cluster.local   fn-00001        fn-00001      True    

```

## 访问服务
```shell
curl http://fn.default.svc.cluster.local
```