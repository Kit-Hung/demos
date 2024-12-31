## 1. 更新 kpa 的 burst capacity
```shell
kubectl annotate kpa http-fn-00003 autoscaling.knative.dev/targetBurstCapacity=1
```

## 2. 访问函数服务
```shell
curl http://http-fn.default.svc.cluster.local -I
```

## 3. 查看 ep
```shell
kubectl get ep http-fn-00003 http-fn-00003-private

NAME                    ENDPOINTS                                                                    AGE
http-fn-00003           192.168.235.240:8012,192.168.235.240:8112                                    5h49m
http-fn-00003-private   192.168.235.221:9091,192.168.235.221:8012,192.168.235.221:8022 + 3 more...   5h49m
```

## 4. 查看 serverlessService
```shell
kubectl get sks -w

NAME            MODE    ACTIVATORS   SERVICENAME     PRIVATESERVICENAME      READY     REASON
http-fn-00003   Proxy   2            http-fn-00003   http-fn-00003-private   True      
http-fn-00003   Serve   2            http-fn-00003   http-fn-00003-private   True      
http-fn-00003   Serve   2            http-fn-00003   http-fn-00003-private   True      
http-fn-00003   Proxy   2            http-fn-00003   http-fn-00003-private   True      
http-fn-00003   Proxy   2            http-fn-00003   http-fn-00003-private   True
```