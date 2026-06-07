#!/bin/bash

set -e

REPO="Jonath-z/commit-ai"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="commit-ai"

case "$(uname -s)" in
  Linux)  OS=linux ;;
  Darwin) OS=darwin ;;
  *)
    echo "Unsupported OS: $(uname -s)" >&2
    echo "Try: go install github.com/${REPO}@latest" >&2
    exit 1
    ;;
esac

case "$(uname -m)" in
  x86_64|amd64)   ARCH=amd64 ;;
  arm64|aarch64)  ARCH=arm64 ;;
  *)
    echo "Unsupported architecture: $(uname -m)" >&2
    echo "Try: go install github.com/${REPO}@latest" >&2
    exit 1
    ;;
esac

ASSET="commit-ai-${OS}-${ARCH}"
URL="https://github.com/${REPO}/releases/latest/download/${ASSET}"

echo "Fetching ${ASSET} from latest release..."
curl -fL -o "$BINARY_NAME" "$URL"

echo "Making $BINARY_NAME executable..."
chmod +x "$BINARY_NAME"

echo "Installing $BINARY_NAME to $INSTALL_DIR..."
sudo mv "$BINARY_NAME" "$INSTALL_DIR/"

if command -v "$BINARY_NAME" &> /dev/null; then
  echo "$BINARY_NAME installed successfully!"
else
  echo "Installation failed!" >&2
  exit 1
fi
