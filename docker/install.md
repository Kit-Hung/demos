### Install docker

```shell
apt install docker.io
```


### Update cgroupdriver to systemd

```shell
vi /etc/docker/daemon.json
{
  "exec-opts": ["native.cgroupdriver=systemd"]
}
systemctl daemon-reload
systemctl restart docker
```
