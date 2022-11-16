# Manager build
FROM golang:1.19-alpine as manager_builder

RUN apk --no-cache add ca-certificates git

WORKDIR /app

COPY ./manager/go.mod .
COPY ./manager/go.sum .

RUN go mod download

COPY ./manager .

RUN go build -o ./app main.go \
    && chmod +x ./app
##

# VPN
FROM alpine:3.14.1 as main
RUN apk add --no-cache sudo bash dumb-init curl
ENV NET_ADAPTER eth0
ENV HOST_ADDR localhost
ENV HOST_TUN_PORT 1194
ENV HOST_CONF_PORT 80

# OpenVPN part
ENV OPENVPN_APP_NAME openvpn
ENV OPENVPN_INSTALL_PATH /opt/${OPENVPN_APP_NAME}
ENV OPENVPN_PERSIST_DIR /opt/${OPENVPN_APP_NAME}_data

COPY devops/config/openvpn ${OPENVPN_INSTALL_PATH}/config
RUN apk add --no-cache openvpn easy-rsa netcat-openbsd zip && \
    ln -s /usr/share/easy-rsa/easyrsa /usr/bin/easyrsa && \
    mkdir -p ${OPENVPN_PERSIST_DIR} && \
    cd ${OPENVPN_PERSIST_DIR} && \
    easyrsa init-pki && \
    easyrsa gen-dh && \
    cp pki/dh.pem /etc/openvpn && \
    cd ${OPENVPN_INSTALL_PATH} && \
    cp config/server.conf /etc/openvpn/server.conf
EXPOSE 1194/udp
##

# Manager part
RUN mkdir /manager
COPY --from=manager_builder /app/app /manager
EXPOSE 80/tcp
##

VOLUME ["/opt/"]

COPY devops/boot /boot
RUN chmod +x /boot/*
ENTRYPOINT ["dumb-init", "/boot/start.sh"]
CMD [""]
##