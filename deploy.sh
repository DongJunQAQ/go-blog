#!/bin/bash
bak_image_dir=/image_backup
docker rm -f blog
mkdir $bak_image_dir
docker save -o $bak_image_dir/go-blog_"$(date +"%Y%m%d%H%M%S")".tar 192.168.246.152/blog/go-blog:1.1
docker rmi 192.168.246.152/blog/go-blog:1.1
echo "$1" | docker login --username=dongjun --password-stdin 192.168.246.152
docker run -d --name blog -p:8080:8080 --restart always 192.168.246.152/blog/go-blog:1.1
