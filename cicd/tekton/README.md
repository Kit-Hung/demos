## 安装文档
https://tekton.dev/docs/installation/


## 示例代码
### Task
文件清单： 
1. [hello-world.yaml](resources/hello-world.yaml)
2. [hello-world-run.yaml](resources/hello-world-run.yaml)

```shell
kubectl apply -f hello-world.yaml
kubectl apply -f hello-world-run.yaml

kubectl get taskrun hello-task-run

# 查看输出
kubectl logs hello-task-run-pod
tkn taskrun logs -n default hello-task-run
```


### Pipeline
文件清单：
1. [hello-world.yaml](resources/hello-world.yaml)
2. [goodbye-world.yaml](resources/goodbye-world.yaml)
3. [hello-goodbye-pipeline.yaml](resources/hello-goodbye-pipeline.yaml)
4. [hello-goodbye-pipeline-run.yaml](resources/hello-goodbye-pipeline-run.yaml)

```shell
kubectl apply -f hello-world.yaml
kubectl apply -f goodbye-world.yaml
kubectl apply -f hello-goodbye-pipeline.yaml
kubectl apply -f hello-goodbye-pipeline-run.yaml

# 查看输出日志
tkn pipelinerun logs -n default hello-goodbye-run
```


### Trigger
1. [hello-world.yaml](resources/hello-world.yaml)
2. [goodbye-world.yaml](resources/goodbye-world.yaml)
3. [hello-goodbye-pipeline.yaml](resources/hello-goodbye-pipeline.yaml)
4. [trigger-template.yaml](resources/trigger-template.yaml)
5. [trigger-binding.yaml](resources/trigger-binding.yaml)
6. [event-listener.yaml](resources/event-listener.yaml)

```shell
kubectl apply -f hello-world.yaml
kubectl apply -f goodbye-world.yaml
kubectl apply -f hello-goodbye-pipeline.yaml
kubectl apply -f trigger-template.yaml
kubectl apply -f trigger-binding.yaml
kubectl apply -f rbac.yaml
kubectl apply -f event-listener.yaml

kubectl port-forward service/el-hello-listener 8080

curl -v \
   -H 'content-Type: application/json' \
   -d '{"username": "Tekton"}' \
   http://localhost:8080

kubectl get pipelinerun

# 查看输出
tkn pipelinerun logs -f hello-goodbye-run-fwrr8
```


### Tekton Chains
#### 安装
```shell
kubectl apply --filename https://storage.googleapis.com/tekton-releases/chains/latest/release.yaml

kubectl patch configmap chains-config -n tekton-chains \
-p='{"data":{"artifacts.oci.storage": "", "artifacts.taskrun.format":"in-toto", "artifacts.taskrun.storage": "tekton"}}'

cosign generate-key-pair k8s://tekton-chains/signing-secrets
```