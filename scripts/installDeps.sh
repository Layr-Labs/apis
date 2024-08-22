#!/usr/bin/env bash

export BIN="/usr/local/bin"
export VERSION="1.33.0"

curl -sSL "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" -o "${BIN}/buf"

chmod +x "${BIN}/buf"
