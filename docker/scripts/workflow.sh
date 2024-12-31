#!/bin/bash

# docker login -u "$ALIYUN_REGISTRY_USER" -p "$ALIYUN_REGISTRY_PASSWORD" "$ALIYUN_REGISTRY"
ALIYUN_REGISTRY="registry.cn-hangzhou.aliyuncs.com"
ALIYUN_NAME_SPACE="xxx"

max_count=72

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

    echo "docker pull $line"
    # docker pull "$line"

    platform=$(echo "$line" | awk -F'--platform[ =]' '{if (NF>1) print $2}' | awk '{print $1}')
    echo "platform is $platform"
    # 如果存在架构信息 将架构信息拼到镜像名称前面
    if [ -z "$platform" ]; then
        platform_prefix=""
    else
        platform_prefix="${platform//\//_}_"
    fi
    echo "platform_prefix is $platform_prefix"

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

    echo "docker tag $line $new_image"
    # docker tag "$line" "$new_image"

    echo "docker push $new_image"
    # docker push "$new_image"

    echo "开始清理磁盘空间"
    echo "=============================================================================="

    # df -hT
    echo "=============================================================================="

    # docker rmi "$line"
    # docker rmi "$new_image"

    echo "磁盘空间清理完毕"
    echo "=============================================================================="

    # df -hT
    echo "=============================================================================="

done < images.txt