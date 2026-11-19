#!/usr/bin/env bash
# Instala lautaro-env-store: el código en ~/.local/lib/lautaro-env-store y un link en ~/.local/bin.
set -euo pipefail
PREFIX="${PREFIX:-$HOME/.local/bin}"
LIB_DIR="${LIB_DIR:-$HOME/.local/lib/lautaro-env-store}"
REPO_TARBALL="https://github.com/its-lautaro/lautaro-env-store/archive/refs/heads/main.tar.gz"

SRC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]:-$0}")" 2>/dev/null && pwd || true)"
if [ -n "$SRC_DIR" ] && [ -d "$SRC_DIR/lautaro_env_store" ]; then
  SOURCE="$SRC_DIR"
else
  TMP="$(mktemp -d)"
  trap 'rm -rf "$TMP"' EXIT
  curl -fsSL "$REPO_TARBALL" | tar -xz -C "$TMP"
  SOURCE="$TMP/lautaro-env-store-main"
fi

rm -rf "$LIB_DIR"
mkdir -p "$LIB_DIR" "$PREFIX"
cp -R "$SOURCE/bin" "$SOURCE/lautaro_env_store" "$LIB_DIR/"
find "$LIB_DIR" -name '__pycache__' -type d -exec rm -rf {} +
chmod +x "$LIB_DIR/bin/lautaro-env-store"
ln -sf "$LIB_DIR/bin/lautaro-env-store" "$PREFIX/lautaro-env-store"

"$PREFIX/lautaro-env-store" --init
echo "✅ Instalado en $PREFIX/lautaro-env-store (código en $LIB_DIR)"
case ":$PATH:" in
  *":$PREFIX:"*) ;;
  *) echo "⚠️  $PREFIX no está en tu PATH. Agregá a tu ~/.zshrc:  export PATH=\"$PREFIX:\$PATH\"" ;;
esac
