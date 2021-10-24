FROM ubuntu:20.04

# Install pulumi and its dependencies
RUN apt-get update -y -q && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y \
    build-essential \
    ca-certificates \
    wget \
    tar \
    unzip \
    gzip \
    curl \
    gcc \
    make \
    python3 \
    python3-pip && \
    curl -fsSL https://get.pulumi.com | bash -s -- --version 3.15.0

# Install awscli
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \
    ./aws/install && \
    rm -rf awscliv2.zip aws && \
    aws --version

# Install golang
ENV GOLANG_VERSION 1.17.2

RUN wget -c https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz -O - | tar -xz -C /usr/local

WORKDIR /infrastructure

ENV PATH=$PATH:/root/.pulumi/bin:/usr/local/go/bin

# Install python dependencies
RUN pip3 install awscli-local && \
    pip3 install pulumi-local
