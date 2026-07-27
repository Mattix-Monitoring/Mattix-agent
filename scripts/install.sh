#!/usr/bin/env bash

set -euo pipefail

REPO="Mattix-Monitoring/Mattix-agent"
BINARY="mattix"

if [[ "$EUID" -eq 0 ]]; then
  SUDO=""
else
  SUDO="sudo"
fi

echo "==> Detecting architecture..."

case "$(uname -m)" in
x86_64)
  ARCH="amd64"
  ;;
*)
  echo "Unsupported architecture: $(uname -m)"
  exit 1
  ;;
esac

TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

echo "==> Downloading latest release..."

VERSION=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" |
  grep '"tag_name":' |
  sed -E 's/.*"([^"]+)".*/\1/')

URL="https://github.com/$REPO/releases/download/${VERSION}/${BINARY}-linux-${ARCH}.tar.gz"

curl -fL "$URL" -o "$TMP/mattix-agent.tar.gz"

echo "==> Extracting..."

tar -xzf "$TMP/mattix-agent.tar.gz" -C "$TMP"

echo "==> Installing binary..."

$SUDO install -Dm755 \
  "$TMP/mattix-agent" \
  "/usr/local/bin/$BINARY"

echo
echo "✔ Mattix Agent installed successfully!"
echo
echo "Version : $VERSION"
echo "Binary  : /usr/local/bin/$BINARY"
echo "Config  : ~/.config/mattix/config.yaml"
echo
echo "Useful command:"
echo "  mattix"
