### Set no password for sudo
```shell
%sudo ALL=(ALL:ALL) NOPASSWD:ALL
```


### Swap off

```shell
swapoff -a

vi /etc/fstab
# remove the lin with swap
```