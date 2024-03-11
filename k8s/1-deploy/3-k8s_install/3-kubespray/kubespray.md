# 通过 kubespray 创建集群
## 拉取镜像
```shell
docker pull quay.io/kubespray/kubespray:v2.24.1
```

## 下载代码
```shell
git clone https://github.com/kubernetes-sigs/kubespray.git
cd kubespray
git checkout v2.24.1
```

## 启动 kubespray 容器
```shell
docker run --net host --rm -it --mount type=bind,source="$(pwd)"/inventory/sample,dst=/inventory \
  --mount type=bind,source="${HOME}"/.ssh/id_rsa,dst=/root/.ssh/id_rsa \
  quay.io/kubespray/kubespray:v2.24.1 bash
```

### 免密登录
```shell
ssh-keygen -t rsa
ssh-copy-id -i ~/.ssh/id_rsa.pub root@172.28.58.199
```

### 修改并生成配置
```shell
cp -r inventory/sample inventory/mycluster
declare -a IPS=(172.28.58.199)

# 多 ip 情况
# declare -a IPS=(192.168.10.11 192.168.10.12)
CONFIG_FILE=inventory/mycluster/hosts.yml python3 contrib/inventory_builder/inventory.py ${IPS[@]}
```

### 修改镜像源
```shell
cat > inventory/mycluster/group_vars/k8s_cluster/vars.yml << EOF
gcr_image_repo: "registry.aliyuncs.com/google_containers"
kube_image_repo: "registry.aliyuncs.com/google_containers"
etcd_download_url: "https://mirror.ghproxy.com/https://github.com/coreos/etcd/releases/download/{{ etcd_version }}/etcd-{{ etcd_version }}-linux-{{ image_arch }}.tar.gz"
cni_download_url: "https://mirror.ghproxy.com/https://github.com/containernetworking/plugins/releases/download/{{ cni_version }}/cni-plugins-linux-{{ image_arch }}-{{ cni_version }}.tgz"
calicoctl_download_url: "https://mirror.ghproxy.com/https://github.com/projectcalico/calico/releases/download/{{ calico_ctl_version }}/calicoctl-linux-{{ image_arch }}"
calico_crds_download_url: "https://mirror.ghproxy.com/https://github.com/projectcalico/calico/archive/{{ calico_version }}.tar.gz"
crictl_download_url: "https://mirror.ghproxy.com/https://github.com/kubernetes-sigs/cri-tools/releases/download/{{ crictl_version }}/crictl-{{ crictl_version }}-{{ ansible_system | lower }}-{{ image_arch }}.tar.gz"
runc_download_url: "https://mirror.ghproxy.com/https://github.com/opencontainers/runc/releases/download/{{ runc_version }}/runc.{{ image_arch }}"
nerdctl_download_url: "https://mirror.ghproxy.com/https://github.com/containerd/nerdctl/releases/download/v{{ nerdctl_version }}/nerdctl-{{ nerdctl_version }}-{{ ansible_system | lower }}-{{ image_arch }}.tar.gz"
containerd_download_url: "https://mirror.ghproxy.com/https://github.com/containerd/containerd/releases/download/v{{ containerd_version }}/containerd-{{ containerd_version }}-linux-{{ image_arch }}.tar.gz"
nodelocaldns_image_repo: "registry.lank8s.cn/dns/k8s-dns-node-cache"
dnsautoscaler_image_repo: "registry.lank8s.cn/cpa/cluster-proportional-autoscaler"
EOF
```

### 如果需要指定用户（非必要）
```shell
vim ansible.cfg
add remote_user=cadmin to [default] section
```

### 部署
```shell
ansible-playbook -i inventory/mycluster/hosts.yml cluster.yml -b -vv \
  --private-key=~/.ssh/id_rsa
```