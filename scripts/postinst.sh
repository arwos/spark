#!/bin/bash


if [ -f "/etc/systemd/system/spark.service" ]; then
    systemctl start spark
    systemctl enable spark
    systemctl daemon-reload
fi
