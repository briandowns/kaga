GO = go

BINDIR := bin
BINARY := kaga

VERSION = v0.1.0
GIT_SHA = $(shell git rev-parse HEAD)

K3S_PKG=vendor/github.com/rancher/k3s
LDFLAGS = -ldflags "-X $(K3S_PKG)/pkg/version.GitCommit=$(GIT_SHA) \
					-X ${K3S_PKG}/pkg/version.Version=${VERSION}   \
					-X ${K3S_PKG}/pkg/version.Program=${PROG}"

$(BINDIR)/$(BINARY): clean
	$(GO) build $(LDFLAGS) -o $@

.PHONY: clean
clean:
	$(GO) clean
	rm -f $(BINARY)
	rm -f $(BINDIR)/*
