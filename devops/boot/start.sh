#!/bin/bash

echo "Initialising VPN services..."
echo

echo "Configuring OpenVPN..."
/boot/openvpn-init.sh >> /dev/null
echo "OpenVPN is running"
echo

echo "Starting control manager..."
cd /manager
./app &
tail -f /dev/null