#!/bin/bash

echo "Initialising environment..."

if [[ -z $HOST_ADDR ]]; then
  export HOST_ADDR=$(curl ifconfig.me) >> /dev/null
fi

if [[ -z $HOST_URL ]]; then
  export HOST_URL="http://$HOST_ADDR:8080"
fi

echo

echo "Initialising VPN services..."
echo

echo "Configuring OpenVPN..."
/boot/openvpn-init.sh >> /dev/null
echo "OpenVPN is running"
echo

echo "Starting control manager..."
cd /manager
./app