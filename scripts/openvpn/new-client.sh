#!/bin/bash

cd "$OPENVPN_PERSIST_DIR"
CLIENT_ID=$1
CLIENT_PATH="$OPENVPN_PERSIST_DIR/clients/$CLIENT_ID"
easyrsa build-client-full "$CLIENT_ID" nopass &> /dev/null
mkdir -p $CLIENT_PATH
cp "pki/private/$CLIENT_ID.key" "pki/issued/$CLIENT_ID.crt" pki/ca.crt /etc/openvpn/ta.key $CLIENT_PATH
cd "$OPENVPN_INSTALL_PATH"
cp config/client.ovpn $CLIENT_PATH/

echo -e "\nremote $HOST_ADDR $HOST_TUN_PORT" >> "$CLIENT_PATH/client.ovpn"

cat <(echo -e '<ca>') \
  "$CLIENT_PATH/ca.crt" <(echo -e '</ca>\n<cert>') \
  "$CLIENT_PATH/$CLIENT_ID.crt" <(echo -e '</cert>\n<key>') \
  "$CLIENT_PATH/$CLIENT_ID.key" <(echo -e '</key>\n<tls-auth>') \
  "$CLIENT_PATH/ta.key" <(echo -e '</tls-auth>') \
     >> "$CLIENT_PATH/client.ovpn"

echo ";client-id $CLIENT_ID" >> "$CLIENT_PATH/client.ovpn"

echo "$CLIENT_PATH/client.ovpn"