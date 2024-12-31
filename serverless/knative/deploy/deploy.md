# 官方安装链接
https://knative.dev/docs/install/


# 安装 knative
## 验证镜像签名
```shell
go install github.com/sigstore/cosign/v2/cmd/cosign@latest
apt-get install jq
```


## 安装 serving
```shell
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.15.1/serving-crds.yaml
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.15.1/serving-core.yaml
```


### 安装网络层（istio）
```shell
kubectl apply -l knative.dev/crd-install=true -f https://github.com/knative/net-istio/releases/download/knative-v1.15.1/istio.yaml
kubectl apply -f https://github.com/knative/net-istio/releases/download/knative-v1.15.1/istio.yaml

kubectl apply -f https://github.com/knative/net-istio/releases/download/knative-v1.15.1/net-istio.yaml

kubectl --namespace istio-system get service istio-ingressgateway
```


### 安装 DNS
Knative provides a Kubernetes Job called default-domain that configures Knative Serving to use sslip.io as the default DNS suffix.

```shell
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.15.1/serving-default-domain.yaml
```

### 安装 hpa
```shell
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.15.1/serving-hpa.yaml
```


## 安装 eventing
```shell
kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v1.15.0/eventing-crds.yaml
kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v1.15.0/eventing-core.yaml
```


### 安装 Kafka
```shell
kubectl apply -f https://github.com/knative-extensions/eventing-kafka-broker/releases/download/knative-v1.15.0/eventing-kafka-controller.yaml
kubectl apply -f https://github.com/knative-extensions/eventing-kafka-broker/releases/download/knative-v1.15.0/eventing-kafka-channel.yaml

```


## 安装 func cli
```shell
https://github.com/knative/func/releases
https://github.com/knative/client/releases

```