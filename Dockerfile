FROM ubuntu:22.04

RUN apt-get update

RUN apt-get install -y git wget curl

RUN curl -LO https://go.dev/dl/go1.22.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz && \
    rm go1.22.2.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV PATH="${GOPATH}/bin:${PATH}"

WORKDIR /app



RUN apt-get install -y wget git make libnl-3-dev libnet-dev libbsd-dev runc libcap-dev \
    libgpgme-dev btrfs-progs libbtrfs-dev libseccomp-dev libapparmor-dev libprotobuf-dev \
    libprotobuf-c-dev protobuf-c-compiler protobuf-compiler python3-protobuf \
    software-properties-common zip build-essential pkg-config


RUN git clone https://github.com/cedana/cedana.git

WORKDIR /app/cedana

RUN git fetch &&\
    git checkout main

RUN CGO_ENABLED=1 go build -o cedana


COPY build-start-daemon.sh /usr/local/bin/
RUN cp cedana /usr/local/bin/

ENV USER="root"

# Set entrypoint
ENTRYPOINT ["/bin/bash"]
