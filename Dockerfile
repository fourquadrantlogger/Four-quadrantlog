FROM ubuntu

RUN apt update -y &&\
    DEBIAN_FRONTEND=noninteractive apt install  --no-install-recommends -y \
    wget \
    nginx \
    curl \
    net-tools inetutils-ping\
    zip unzip \
    tar \
    vim \
    ca-certificates \
    xz-utils \
    openssh-server \
    git \
    apt-transport-https \
    gnupg \
    rsync
COPY deploy/nginx-linux.conf /etc/nginx/nginx.conf
RUN v=1.18 && wget  "https://dl.google.com/go/go${v}.linux-amd64.tar.gz"  --progress=bar:force 2>&1 &&tar xzvf "go${v}.linux-amd64.tar.gz" && pwd

ENV NODE_VERSION=16.x
RUN curl -fsSL https://deb.nodesource.com/setup_${NODE_VERSION} | bash -
RUN apt-get install -y nodejs && npm i -g corepack && corepack enable

WORKDIR /opt/project
COPY backend backend
ENV GOPROXY=https://proxy.golang.com.cn,direct
RUN pwd && cd backend &&\
    /go/bin/go build  -o fourquadrantlog app.go

COPY vite-proj vite-proj
RUN pwd && cd vite-proj &&\
    npm install
COPY vue-web vue-web
RUN cd vue-web &&\
    npm install
COPY deploy/run.sh deploy/run.sh
RUN chmod +x deploy/run.sh
CMD sh deploy/run.sh
