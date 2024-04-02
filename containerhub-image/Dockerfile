FROM debian:bookworm-slim

ARG S6_OVERLAY_VERSION=3.1.6.0

RUN useradd -m -s /bin/bash user

USER root
RUN rm -rf /etc/apt/sources.list.d/debian.sources && \
    echo "deb http://mirrors.ustc.edu.cn/debian bookworm main contrib non-free non-free-firmware" >> /etc/apt/sources.list && \
    echo "deb http://mirrors.ustc.edu.cn/debian bookworm-updates main contrib non-free non-free-firmware" >> /etc/apt/sources.list && \
    echo "deb http://mirrors.ustc.edu.cn/debian/ bookworm-backports main contrib non-free non-free-firmware" >> /etc/apt/sources.list && \
    echo "deb http://mirrors.ustc.edu.cn/debian-security/ bookworm-security main contrib non-free non-free-firmware" >> /etc/apt/sources.list && \
    apt-get update && apt-get install -y --no-install-recommends \
    'tmux' \
    'cmake' \
    'build-essential' \
    'openssh-server' \
    'htop' \
    'vim' \
    'curl' \
    'wget' \
    'zip' \
    'unzip' \
    'sudo' \
    'git' \
    'locales' && \
    rm -rf /var/lib/apt/lists/* 

# set locale
RUN localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG=en_US.utf8

# set default password and enable ssh
RUN echo "user:user" | chpasswd && \
    adduser user sudo && \
    ssh-keygen -A && \
    mkdir /run/sshd && \
    echo "TrustedUserCAKeys /etc/ssh/ssh_ca.pub" >>/etc/ssh/sshd_config.d/containerhub.conf && \
    echo "AuthorizedPrincipalsFile /etc/ssh/ssh_principals" >>/etc/ssh/sshd_config.d/containerhub.conf &&\
    echo "PasswordAuthentication no" >>/etc/ssh/sshd_config.d/containerhub.conf

# add s6-overlay
ADD https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-noarch.tar.xz /tmp
RUN tar -C / -Jxpf /tmp/s6-overlay-noarch.tar.xz
ADD https://github.com/just-containers/s6-overlay/releases/download/v${S6_OVERLAY_VERSION}/s6-overlay-x86_64.tar.xz /tmp
RUN tar -C / -Jxpf /tmp/s6-overlay-x86_64.tar.xz

# copy s6 services
COPY services.d/ /etc/services.d/
COPY cont-init.d/ /etc/cont-init.d/

EXPOSE 22

CMD ["/init"]
