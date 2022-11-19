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

ARG HOST_URL="localhost:8080"

RUN apk add --no-cache sudo bash dumb-init curl
ENV NET_ADAPTER eth0
ENV HOST_ADDR localhost
ENV HOST_TUN_PORT 1194
ENV HOST_CONF_PORT 80

# Manager part
RUN mkdir /manager
COPY --from=manager_builder /app/app /manager

ENV HTTP_ADDRESS "0.0.0.0:80"
ENV OPERATOR_URL "https://operator.tgvpn.dev2.feraru.org/manager/hook"
EXPOSE 80
##

# OpenVPN part
ENV OPENVPN_APP_NAME openvpn
ENV OPENVPN_INSTALL_PATH /opt/${OPENVPN_APP_NAME}
ENV OPENVPN_PERSIST_DIR /opt/${OPENVPN_APP_NAME}_data
ENV OPENVPN_SCRIPT_DIR /manager/${OPENVPN_APP_NAME}

COPY devops/config/openvpn ${OPENVPN_INSTALL_PATH}/config
EXPOSE 1194/udp

COPY scripts/openvpn ${OPENVPN_SCRIPT_DIR}
RUN chmod +x ${OPENVPN_SCRIPT_DIR}/*
##

COPY devops/boot /boot
RUN chmod +x /boot/*
ENTRYPOINT ["dumb-init", "/boot/start.sh"]
CMD [""]
##