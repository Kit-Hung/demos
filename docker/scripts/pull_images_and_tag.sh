#!/bin/bash

# 检查是否有传入参数
if [ $# -eq 0 ]; then
    echo "错误：没有传入参数。"
    echo "可以传入 -h 或者 --help 查看命令参数示例"
    exit 1
fi

if [[ "$1" == "-h" || "$1" == "--help" ]]; then
    echo "example: image_file.txt registry.cn-hangzhou.aliyuncs.com/xxx/ docker.io/"
    exit
fi

image_file=$1
old_repo=$2
new_repo=$3

# 确保文件存在
if [ ! -f "$image_file" ]; then
    echo "文件 $image_file 不存在"
    exit 1
fi

# 读取文件中的每一行
while IFS= read -r line
do
    echo "正在拉取镜像: $line"
    nerdctl -n k8s.io image pull "$line"
    # 检查命令是否成功执行
    if [ $? -eq 0 ]; then
        echo "成功拉取: $line"
    else
        echo "拉取失败: $line"
        exit
    fi

    new_name=$(echo "$line" | sed "s|^$old_repo|$new_repo|")
    nerdctl -n k8s.io image tag "$line" "$new_name"
done < "$image_file"