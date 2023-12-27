#!/bin/bash


if ! [ -d /var/lib/spark/ ]; then
    mkdir /var/lib/spark
fi

if [ -f "/etc/systemd/system/spark.service" ]; then
    systemctl stop spark
    systemctl disable spark
    systemctl daemon-reload
fi
