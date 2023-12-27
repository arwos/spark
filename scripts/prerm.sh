#!/bin/bash


if [ -f "/etc/systemd/system/spark.service" ]; then
    systemctl stop spark
    systemctl disable spark
    systemctl daemon-reload
fi
