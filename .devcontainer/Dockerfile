FROM mcr.microsoft.com/devcontainers/base:bookworm

RUN apt-get update \
    && apt-get install -y --no-install-recommends gnupg curl ca-certificates python3 socat \
    && rm -rf /var/lib/apt/lists/*

ARG GO_VERSION=1.26.5
RUN set -eux; \
    arch="$(dpkg --print-architecture)"; \
    case "$arch" in \
      amd64) goarch=amd64 ;; \
      arm64) goarch=arm64 ;; \
      *) echo "unsupported arch: $arch" >&2; exit 1 ;; \
    esac; \
    curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-${goarch}.tar.gz" -o /tmp/go.tgz; \
    tar -C /usr/local -xzf /tmp/go.tgz; \
    rm /tmp/go.tgz
ENV GOPATH="/go"
ENV PATH="/usr/local/go/bin:/go/bin:${PATH}"

RUN mkdir -p /go && chown -R vscode:vscode /go
RUN sudo -u vscode env "PATH=$PATH" GOPATH=/go go install github.com/goreleaser/goreleaser/v2@latest
RUN sudo -u vscode env "PATH=$PATH" GOPATH=/go go install github.com/go-delve/delve/cmd/dlv@latest

ENV VERSION="1.165.1"
RUN curl -fsSL https://raw.githubusercontent.com/speakeasy-api/speakeasy/main/install.sh | sh
