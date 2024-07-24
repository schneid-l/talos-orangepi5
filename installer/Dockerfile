ARG KERNEL

FROM ${KERNEL} AS kernel

ARG SOURCE_DATE_EPOCH

FROM golang:alpine AS builder

ARG SOURCE_DATE_EPOCH

COPY go.mod go.sum /installer/

ENV GOMODCACHE=/.cache/gomod
WORKDIR /installer
RUN --mount=type=cache,target=/.cache/gomod \
    go mod download && \
    go mod verify

COPY . /installer

ENV GOCACHE=/.cache/go
WORKDIR /installer
RUN --mount=type=cache,target=/.cache/gomod \
    --mount=type=cache,target=/.cache/go \
    CGO_ENABLED=0 go build -o /orangepi-5 .

FROM scratch AS installer

ARG SOURCE_DATE_EPOCH
ARG VERSION
ARG IMAGE_NAME="talos-orangepi5-installer"
ARG IMAGE_TITLE="Talos Orange Pi 5 Installer"
ARG IMAGE_DESCRIPTION=$IMAGE_TITLE
ARG IMAGE_SOURCE="https://github.com/si0ls/talos-orangepi5"
ARG IMAGE_AUTHORS="Louis S. <louis@schne.id>"
ARG IMAGE_VENDOR=$IMAGE_AUTHORS
ARG IMAGE_VERSION=$VERSION

LABEL org.opencontainers.image.name=$IMAGE_NAME
LABEL org.opencontainers.image.title=$IMAGE_TITLE
LABEL org.opencontainers.image.description=$IMAGE_DESCRIPTION
LABEL org.opencontainers.image.source=$IMAGE_SOURCE
LABEL org.opencontainers.image.authors=$IMAGE_AUTHORS
LABEL org.opencontainers.image.vendor=$IMAGE_VENDOR
LABEL org.opencontainers.image.version=$IMAGE_VERSION

COPY --from=kernel /dtb/rockchip /artifacts/dtb
COPY --from=builder /orangepi-5 /installers/orangepi-5
COPY profile.yaml /profiles/orangepi-5.yaml

# Path: installer/Dockerfile
