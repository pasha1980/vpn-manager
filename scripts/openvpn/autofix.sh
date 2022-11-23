#!/bin/bash

if ! [[ -L /usr/bin/easyrsa ]]; then
  ln -s /usr/share/easy-rsa/easyrsa /usr/bin/easyrsa
fi

if ! [[ -d ${OPENVPN_PERSIST_DIR} ]]; then
  mkdir -p ${OPENVPN_PERSIST_DIR} ${OPENVPN_PERSIST_DIR}/clients ${OPENVPN_PERSIST_DIR}/removed
  cd ${OPENVPN_PERSIST_DIR}
  easyrsa init-pki
  easyrsa gen-dh
  cp pki/dh.pem /etc/openvpn
  cd ${OPENVPN_INSTALL_PATH}
  cp config/server.conf /etc/openvpn/server.conf
  touch ${OPENVPN_PERSIST_DIR}/init.gen
fi

if ! [[ -f /etc/openvpn/ta.key ]]; then
  easyrsa build-ca nopass << EOF

EOF
  easyrsa gen-req MyReq nopass << EOF2

EOF2
  easyrsa sign-req server MyReq << EOF3
yes
EOF3
  openvpn --genkey secret ta.key << EOF4
yes
EOF4
  easyrsa gen-crl
  cp pki/ca.crt pki/issued/MyReq.crt pki/private/MyReq.key pki/crl.pem ta.key /etc/openvpn
fi

if [[ -z $(ps | grep 'openvpn') ]]; then
  openvpn --config /etc/openvpn/server.conf &
fi
