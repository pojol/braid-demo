#!/bin/bash


# 获取当前路径
current_path=$(pwd)

# 定义服务器配置
server_ids=("cluster-1" "cluster-2" "cluster-3")
server_ports=("8001" "8002" "8003")

# 启动所有服务器
for i in "${!server_ids[@]}"; do
    id="${server_ids[$i]}"
    port="${server_ports[$i]}"
    echo "Starting main with id $id on port $port..."
    nohup "$current_path/main" -id "$id" -port "$port" > "${current_path}/main_${id}.log" 2>&1 &
    sleep 1  # 等待1秒，确保服务器顺利启动
done


echo "All servers started."