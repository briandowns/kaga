ARG KUBERNETES_VERSION=dev
FROM rancher/hardened-build-base:v1.16.6b7 AS build
RUN apk --no-cache add \
    bash \
    curl \
    file \
    git \
    libseccomp-dev \
    rsync \
    mingw-w64-gcc \
    gcc \
    bsd-compat-headers \
    py-pip \
    pigz