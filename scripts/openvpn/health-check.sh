#!/bin/bash


if [[ -z $(ps | grep 'openvpn') ]]; then
  exit 1
fi

if ! [[ -L /usr/bin/easyrsa ]]; then
  exit 1
fi

if ! [[ -d ${OPENVPN_PERSIST_DIR} ]]; then
  exit 1
fi

if ! [[ -f ${OPENVPN_PERSIST_DIR}/pki/ca.crt ]]; then
  exit 1
fi

if ! [[ -f /etc/openvpn/ta.key ]]; then
  exit 1
fi

if ! [[ -f ${OPENVPN_INSTALL_PATH}/config/client.ovpn ]]; then
  exit 1
fi

if ! [[ -f /etc/openvpn/crl.pem ]]; then
  exit 1
fi

## to be continued

echo "OK"