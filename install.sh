#!/usr/bin/env bash
# Instala lautaro-env-store en ~/.local/bin (o en $PREFIX si está definido).
set -euo pipefail
PREFIX="${PREFIX:-$HOME/.local/bin}"
mkdir -p "$PREFIX"
SRC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
if [ -f "$SRC_DIR/lautaro-env-store" ]; then
  cp "$SRC_DIR/lautaro-env-store" "$PREFIX/lautaro-env-store"
else
  curl -fsSL https://raw.githubusercontent.com/its-lautaro/lautaro-env-store/main/lautaro-env-store -o "$PREFIX/lautaro-env-store"
fi
chmod +x "$PREFIX/lautaro-env-store"
"$PREFIX/lautaro-env-store" --init
echo "✅ Instalado en $PREFIX/lautaro-env-store"
case ":$PATH:" in
  *":$PREFIX:"*) ;;
  *) echo "⚠️  $PREFIX no está en tu PATH. Agregá a tu ~/.zshrc:  export PATH=\"$PREFIX:\$PATH\"" ;;
esac
