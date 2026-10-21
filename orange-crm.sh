#!/bin/bash

echo "Start git pull."
git pull

echo "Start building for Go Linux."

# 检查是否存在orange-crm-v2 文件并删除
if [ -e orange-crm-v2 ]; then
  echo "Old files exist."
  rm -f orange-crm-v2
  echo "The old file has been deleted."
else
  echo "The old file does not exist."
fi

# 设置 Go 环境变量
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

# 执行 Go 编译
go build -o orange-crm-v2 main.go

echo "Compilation completed."


echo "Check the existing orange-crm-v2 process."
# shellcheck disable=SC2009
ps -ef | grep orange-crm-v2

# 删除后台进程
echo "Execute the deletion of the orange-crm-v2 process."
sudo killall -9 orange-crm-v2

# 启动进程
echo "Start the orange-crm-v2 process."
nohup ./orange-crm-v2 &
# 断开进程与当前终端的关联
disown

# shellcheck disable=SC2009
ps -ef | grep orange-crm-v2
echo "End."