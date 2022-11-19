#!/bin/bash

# Move file to ${OPENVPN_PERSIST_DIR}/removed
CLIENT_ID=$1

cd ${OPENVPN_PERSIST_DIR}
easyrsa --batch revoke "$CLIENT_ID"
easyrsa --batch --days=3650 gen-crl
rm -f /etc/openvpn/crl.pem
cp ${OPENVPN_PERSIST_DIR}/pki/crl.pem /etc/openvpn/crl.pem
mv clients/${CLIENT_ID} removed/