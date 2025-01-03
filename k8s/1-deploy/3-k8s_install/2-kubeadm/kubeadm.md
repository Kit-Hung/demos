# 通过 kubeadm 创建集群
### check port

```shell
nc 127.0.0.1 6443
```

### Letting iptables see bridged traffic

```shell
$ cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF

$ cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
$ sudo sysctl --system
```

### ipvs
```shell
apt install ipset ipvsadm -y
# 重启后永久生效
tee /etc/modules-load.d/k8s.conf <<'EOF'
# netfilter
br_netfilter
# containerd.
overlay
# ipvs
ip_vs
ip_vs_rr
ip_vs_wrr
ip_vs_sh
nf_conntrack
EOF

# 临时生效
mkdir -vp /etc/modules.d/
cat > /etc/modules.d/k8s.modules <<EOF
#!/bin/bash
# 允许 iptables 检查桥接流量
modprobe -- br_netfilter
# containerd.
modprobe -- overlay
# ipvs
modprobe -- ip_vs
modprobe -- ip_vs_rr
modprobe -- ip_vs_wrr
modprobe -- ip_vs_sh
modprobe -- nf_conntrack
EOF
chmod 755 /etc/modules.d/k8s.modules && bash /etc/modules.d/k8s.modules && lsmod | grep -e ip_vs -e nf_conntrack
sysctl --system
reboot 
```

### Update the apt package index and install packages needed to use the Kubernetes apt repository:

https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/

```shell
$ sudo apt-get update
$ sudo apt-get install -y apt-transport-https ca-certificates curl gpg
```

### Install kubeadm

```shell
$ sudo curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo apt-key add -
$ sudo curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.29/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
$ echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.29/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list
```

### Add the Kubernetes apt repository

```shell
$ sudo tee /etc/apt/sources.list.d/kubernetes.list <<-'EOF'
deb https://mirrors.aliyun.com/kubernetes/apt kubernetes-xenial main
EOF
```

### Update apt package index, install kubelet, kubeadm and kubectl

```shell
$ sudo apt-get update
$ sudo apt-get install -y kubelet kubeadm kubectl
$ sudo apt-mark hold kubelet kubeadm kubectl
```

### Stop the firewall
```shell
systemctl stop ufw
```

### Config containerd
```shell
containerd config default > /etc/containerd/config.toml

vi /etc/containerd/config.toml
# 约125行，[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]段落
# 默认：
# SystemdCgroup = false
# 改为：
# SystemdCgroup = true

# 约61行，[plugins."io.containerd.grpc.v1.cri"]段落
# 默认：
# sandbox_image = "registry.k8s.io/pause:3.6"
# 改为：
# sandbox_image = "registry.aliyuncs.com/google_containers/pause:3.10"


# version = 2

# [plugins."io.containerd.grpc.v1.cri".registry]
#    config_path = "/etc/containerd/certs.d"

### cat /etc/containerd/certs.d/k8s.gcr.io/host.toml 
# server = "https://k8s.gcr.io"

# [host."https://registry.aliyuncs.com/google_containers"]
#   capabilities = ["pull", "resolve"]

### cat /etc/containerd/certs.d/registry.k8s.io/host.toml 
# server = "https://registry.k8s.io"

# [host."https://registry.aliyuncs.com/google_containers"]
#   capabilities = ["pull", "resolve"]

### cat /etc/containerd/certs.d/core.harbor.domain/host.toml
# server = "https:///core.harbor.domain"

# [host."https:///core.harbor.domain"]
#   capabilities = ["pull", "resolve", "push"]
#   ca = "harbor.crt"



crictl config runtime-endpoint unix:///run/containerd/containerd.sock
systemctl restart containerd
```


```shell
$ kubeadm init \
 --image-repository registry.aliyuncs.com/google_containers \
 --pod-network-cidr=192.168.0.0/16
```

### Copy kubeconfig

```shell
$ mkdir -p $HOME/.kube
$ sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
$ sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

### Untaint master

```shell
$ kubectl taint nodes --all node-role.kubernetes.io/control-plane-
$ kubectl taint nodes --all node-role.kubernetes.io/master-
```

### join node
```shell
kubeadm join xx.xx.xx.xx:6443 --token <token> \
	--discovery-token-ca-cert-hash sha256:<sha256>
```


## Install calico cni plugin

https://docs.projectcalico.org/getting-started/kubernetes/quickstart

```shell
$ kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.27.0/manifests/tigera-operator.yaml
$ kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.27.0/manifests/custom-resources.yaml
```