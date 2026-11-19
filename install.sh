#!/usr/bin/env bash
# Instala lautaro-env-store en ~/.local/bin (o en $PREFIX). Requiere Go 1.22+.
set -euo pipefail
PREFIX="${PREFIX:-$HOME/.local/bin}"
MODULE="github.com/its-lautaro/lautaro-env-store"

if ! command -v go >/dev/null 2>&1; then
  echo "❌ Necesito Go (https://go.dev/dl/) para compilar lautaro-env-store" >&2
  exit 1
fi

mkdir -p "$PREFIX"
SRC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]:-$0}")" 2>/dev/null && pwd || true)"
if [ -n "$SRC_DIR" ] && [ -f "$SRC_DIR/go.mod" ]; then
  (cd "$SRC_DIR" && go build -o "$PREFIX/lautaro-env-store" .)
else
  GOBIN="$PREFIX" go install "$MODULE@latest"
fi

"$PREFIX/lautaro-env-store" --init
echo "✅ Instalado en $PREFIX/lautaro-env-store"
case ":$PATH:" in
  *":$PREFIX:"*) ;;
  *) echo "⚠️  $PREFIX no está en tu PATH. Agregá a tu ~/.zshrc:  export PATH=\"$PREFIX:\$PATH\"" ;;
esac
