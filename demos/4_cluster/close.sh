#!/bin/bash

# 获取当前路径
current_path=$(pwd)

# 定义服务器配置
server_ids=("cluster-1" "cluster-2" "cluster-3")

# 关闭所有服务器
for id in "${server_ids[@]}"; do
    echo "Stopping cluster-demo process with id $id..."
    # 查找在当前路径下运行的进程
    pids=$(pgrep -f "$current_path/cluster-demo.*-id $id")
    if [ -n "$pids" ]; then
        # 如果找到进程，则终止它
        echo "Killing main process with id $id (PIDs: $pids)"
        kill $pids
    else
        echo "No main process with id $id found in current directory"
    fi
done

echo "All servers in current directory stopped."
