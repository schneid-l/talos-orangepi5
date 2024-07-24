# syntax=docker/dockerfile:1

FROM alpine AS downloader

ARG SOURCE_DATE_EPOCH

RUN apk --no-cache add \
    curl \
    git \
    patch \
    tar

ARG KERNEL_VERSION=6.10
ARG KERNEL_SOURCE=https://cdn.kernel.org/pub/linux/kernel/v6.x/linux-6.10.tar.gz
ARG KERNEL_PATCHES_DIR=./patches

RUN mkdir -p /kernel/src && \
    mkdir -p /kernel/patches && \
    curl -L $KERNEL_SOURCE | tar -xz -C /kernel/src --strip-components=1

COPY $KERNEL_PATCHES_DIR/* /kernel/patches/

RUN for patch in /kernel/patches/*.patch; do \
        patch -d /kernel/src -p1 < $patch; \
    done

FROM ubuntu:22.04 AS builder

ARG SOURCE_DATE_EPOCH

ENV DEBIAN_FRONTEND=noninteractive

RUN --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    rm -f /etc/apt/apt.conf.d/docker-clean; \
    echo 'Binary::apt::APT::Keep-Downloaded-Packages "true";' > /etc/apt/apt.conf.d/keep-cache && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone && \
    apt-get update && apt-get install --no-install-recommends -y \
    bc \
    bison \
    bsdmainutils \
    build-essential \
    ca-certificates \
    cpio \
    device-tree-compiler \
    dwarves \
    flex \
    git \
    kmod \
    libbpf-dev \
    libelf-dev \
    liblz4-dev \
    liblz4-tool \
    libncurses-dev \
    libssl-dev \
    libzstd-dev \
    python3 \
    u-boot-tools \
    zstd \
    && rm -rf /var/lib/apt/lists/*

COPY --from=downloader /kernel/src /kernel/src

COPY ./certs/* /kernel/certs/
COPY ./config /kernel/.config

ENV KBUILD_OUTPUT=/kernel/build
ENV ARCH=arm64

WORKDIR /kernel
RUN --mount=type=cache,target=/kernel/build \
    cp .config build/.config && \
    cp certs/x509.genkey src/certs/x509.genkey

WORKDIR /kernel/src
RUN --mount=type=cache,target=/kernel/build \
    sed -i 's/KBUILD_CFLAGS   := -Wall/KBUILD_CFLAGS   := -Wno-error/' Makefile && \
    make -j$(nproc) olddefconfig && \
    make -j$(nproc) && \
    make -j$(nproc) modules && \
    make -j$(nproc) dtbs

    RUN --mount=type=cache,target=/kernel/build \
    fdtoverlay -o /kernel/build/arch/arm64/boot/dts/rockchip/rk3588s-orangepi-5-sata.dtb \
        -i /kernel/build/arch/arm64/boot/dts/rockchip/rk3588s-orangepi-5.dtb \
        /kernel/build/arch/arm64/boot/dts/rockchip/overlay/rockchip-rk3588-sata.dtbo && \
    fdtoverlay -o /kernel/build/arch/arm64/boot/dts/rockchip/rk3588-orangepi-5-plus-sata.dtb \
        -i /kernel/build/arch/arm64/boot/dts/rockchip/rk3588-orangepi-5-plus.dtb \
        /kernel/build/arch/arm64/boot/dts/rockchip/overlay/rockchip-rk3588-sata.dtbo

RUN --mount=type=cache,target=/kernel/build \
    mkdir -p /kernel/out/boot && \
    mkdir -p /kernel/out/dtb/rockchip/overlay && \
    mkdir -p /kernel/out/certs && \
    cp -r /kernel/build/certs/signing_key.x509 /kernel/out/certs/signing_key.x509 && \
    cp -r /kernel/build/arch/arm64/boot/Image /kernel/out/boot/vmlinuz && \
    cp -r /kernel/build/arch/arm64/boot/dts/rockchip/rk3588s-orangepi-5.dtb /kernel/out/dtb/rockchip/ && \
    cp -r /kernel/build/arch/arm64/boot/dts/rockchip/rk3588s-orangepi-5-sata.dtb /kernel/out/dtb/rockchip/ && \
    cp -r /kernel/build/arch/arm64/boot/dts/rockchip/rk3588-orangepi-5-plus.dtb /kernel/out/dtb/rockchip/ && \
    cp -r /kernel/build/arch/arm64/boot/dts/rockchip/rk3588-orangepi-5-plus-sata.dtb /kernel/out/dtb/rockchip/ && \
    cp -r /kernel/build/arch/arm64/boot/dts/rockchip/overlay/rockchip-rk3588-sata.dtbo /kernel/out/dtb/rockchip/overlay && \
    export KERNEL_RELEASE=$(cat /kernel/build/include/config/kernel.release) && \
    make -j$(nproc) modules_install INSTALL_MOD_PATH=/kernel/out INSTALL_MOD_STRIP=1 && \
    depmod -b /kernel/out $KERNEL_RELEASE && \
    unlink /kernel/out/lib/modules/$KERNEL_RELEASE/build

FROM scratch AS kernel

ARG SOURCE_DATE_EPOCH
ARG KERNEL_VERSION=6.10
ARG IMAGE_NAME="talos-orangepi5-kernel"
ARG IMAGE_TITLE="Talos Orange Pi 5 Kernel"
ARG IMAGE_DESCRIPTION=$IMAGE_TITLE
ARG IMAGE_SOURCE="https://github.com/si0ls/talos-orangepi5"
ARG IMAGE_AUTHORS="Louis S. <louis@schne.id>"
ARG IMAGE_VENDOR="The Linux Foundation"
ARG IMAGE_VERSION=$KERNEL_VERSION

LABEL org.opencontainers.image.name=$IMAGE_NAME
LABEL org.opencontainers.image.title=$IMAGE_TITLE
LABEL org.opencontainers.image.description=$IMAGE_DESCRIPTION
LABEL org.opencontainers.image.source=$IMAGE_SOURCE
LABEL org.opencontainers.image.authors=$IMAGE_AUTHORS
LABEL org.opencontainers.image.vendor=$IMAGE_VENDOR
LABEL org.opencontainers.image.version=$IMAGE_VERSION

COPY --from=builder /kernel/out/ /

# Path: kernel/Dockerfile
