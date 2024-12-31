#!/bin/bash

set -x

image_file="images.txt"
platform_prefix=""
max_count=72
ALIYUN_REGISTRY="registry.cn-hangzhou.aliyuncs.com"
ALIYUN_NAME_SPACE="xxx"

# 确保文件存在
if [ ! -f "$image_file" ]; then
    echo "文件 $image_file 不存在"
    exit 1
fi

# 函数：裁剪字符串，直到长度小于或等于最大长度
trim_string() {
    local str="$1"
    # 计算字符串长度
    local len=${#str}

    # 只要字符串长度大于最大长度，就继续裁剪
    while [ "$len" -gt $max_count ]; do
        # 去除第一个'/'及其后的所有字符
        str=${str#*/}
        # 更新字符串长度
        len=${#str}
    done

    # 返回裁剪后的字符串
    echo "$str"
}

while IFS= read -r line || [ -n "$line" ]; do
    # 忽略空行与注释
    [[ -z "$line" ]] && continue
    if echo "$line" | grep -q '^\s*#'; then
        continue
    fi

    # 去掉末尾 \r
    line="${line%$'\r'}" 

    # 获取镜像的完整名称，例如kasmweb/nginx:1.25.3（命名空间/镜像名:版本号）
    image=$(trim_string "$line")

    # 获取 镜像名:版本号  例如1.25.3
    image_tag=$(echo "$image" | awk -F":" '{print $2}')

    # 将@sha256:等字符删除
    image_name_tag="${image_tag%%@*}"

    repo=$(echo "$image" | cut -d':' -f1)
    # 替换/为_，这里使用sed命令
    replaced_repo=$(echo "$repo" | sed 's/\//_/g')
    new_image="$ALIYUN_REGISTRY/$ALIYUN_NAME_SPACE/$platform_prefix$replaced_repo:$image_name_tag"

    echo "正在拉取镜像: $new_image"
    echo "nerdctl -n k8s.io image pull $new_image"
    nerdctl -n k8s.io image pull "$new_image"

    # 检查命令是否成功执行
    if [ $? -eq 0 ]; then
        echo "成功拉取: $new_image"
    else
        echo "拉取失败: $new_image"
        exit
    fi

    echo "nerdctl -n k8s.io image tag $new_image $line"
    nerdctl -n k8s.io image tag "$new_image" "$line"

done < images.txt
