#!/bin/bash

## Configuring OpenVPN
ADAPTER="${OPENVPN_NET_ADAPTER:=eth0}"
source /boot/functions.sh

mkdir -p /dev/net

if [ ! -c /dev/net/tun ]; then
    echo "$(datef) Creating tun/tap device."
    mknod /dev/net/tun c 10 200
fi

iptables -A INPUT -i $ADAPTER -p udp -m state --state NEW,ESTABLISHED --dport 1194 -j ACCEPT
iptables -A OUTPUT -o $ADAPTER -p udp -m state --state ESTABLISHED --sport 1194 -j ACCEPT

iptables -A INPUT -i tun0 -j ACCEPT
iptables -A FORWARD -i tun0 -j ACCEPT
iptables -A OUTPUT -o tun0 -j ACCEPT

iptables -A FORWARD -i tun0 -o $ADAPTER -s 10.8.0.0/24 -j ACCEPT
iptables -A FORWARD -m state --state ESTABLISHED,RELATED -j ACCEPT

iptables -t nat -A POSTROUTING -s 10.8.0.0/24 -o $ADAPTER -j MASQUERADE

cd "$OPENVPN_PERSIST_DIR"
easyrsa build-ca nopass << EOF

EOF
easyrsa gen-req MyReq nopass << EOF2

EOF2
easyrsa sign-req server MyReq << EOF3
yes
EOF3
openvpn --genkey --secret ta.key << EOF4
yes
EOF4
easyrsa gen-crl

cp pki/ca.crt pki/issued/MyReq.crt pki/private/MyReq.key pki/crl.pem ta.key /etc/openvpn
##

# Running all services
openvpn --config /etc/openvpn/server.conf &

cd /manager
./app &
tail -f /dev/null
##