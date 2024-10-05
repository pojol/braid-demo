#!/bin/bash


# 获取当前路径
current_path=$(pwd)

# 定义服务器配置
server_ids=("cluster-1" "cluster-2" "cluster-3")
server_external_ports=("8001" "8002" "8003")
server_internal_ports=("7001" "7002" "7003")

# 启动所有服务器
for i in "${!server_ids[@]}"; do
    id="${server_ids[$i]}"
    eport="${server_external_ports[$i]}"
    iport="${server_internal_ports[$i]}"
    echo "Starting cluster-demo with id $id on port $eport..."
    nohup "$current_path/cluster-demo" -id "$id" -eport "$eport" -iport "$iport" > "${current_path}/main_${id}.log" 2>&1 &
    sleep 1  # 等待1秒，确保服务器顺利启动
done


echo "All servers started."