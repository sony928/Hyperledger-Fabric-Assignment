#!/bin/bash

if [ "$1" == "stop" ]; then
    echo "Stopping Fabric network..."
    docker-compose down
else
    echo "Starting Fabric network..."
    docker-compose up -d
fi
