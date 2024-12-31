## 环境选择
- 内核至少需要 4.9 或者更新的版本，推荐使用更新的 5.x 内核
- 在使用 CO-RE 之前，内核需要开启 CONFIG_DEBUG_INFO_BTF=y 和 CONFIG_DEBUG_INFO=y 这两个编译选项
    - 已经默认开启这些编译选项的发行版
        - Ubuntu 20.10+
        - Fedora 31+
        - RHEL 8.2+
        - Debian 11+


## 工具安装
```shell
# For Ubuntu20.10+
sudo apt-get install -y  make clang llvm libelf-dev libbpf-dev bpfcc-tools libbpfcc-dev linux-tools-$(uname -r) linux-headers-$(uname -r)

# For RHEL8.2+
sudo yum install libbpf-devel make clang llvm elfutils-libelf-devel bpftool bcc-tools bcc-devel
```