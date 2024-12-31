## 创建证书

```shell
cat >metrics-server-csr.json << EOF
{
	"CN": "system:metrics-server",
	"hosts": [],
	"key": {
		"algo": "rsa",
		"size": 2048
	},
	"names": [{
		"C": "CN",
		"ST": "Beijing",
		"L": "Beijing",
		"O": "k8s",
		"0U": "system"
	}]
}
EOF


```

