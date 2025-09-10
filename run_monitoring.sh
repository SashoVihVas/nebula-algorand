#!/bin/bash

while true; do
    echo "-----------------------------------------------------"
    echo "Starting testnet run at $(date)"
    echo "-----------------------------------------------------"
    python3 monitor_peers.py testnet

    echo ""
    echo "[$(date)] Testnet run finished. Waiting for 3 hours..."
    sleep 10800

    echo "-----------------------------------------------------"
    echo "Starting mainnet run at $(date)"
    echo "-----------------------------------------------------"
    python3 monitor_peers.py mainnet

    echo ""
    echo "[$(date)] Mainnet run finished. Waiting for 3 hours..."
    sleep 10800
done