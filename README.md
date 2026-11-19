# lautaro-env-store

CLI mínima para guardar variables de entorno (credenciales, tokens) en `~/.lautaro-env` sin tocar el `.zshrc` a mano. Copia al clipboard el nombre de las variables creadas y lo imprime por stdout.

El archivo `~/.lautaro-env` se crea con permisos `600` y se agrega una línea `source` a tu `~/.zshrc` (o `~/.bashrc`) la primera vez, así cada terminal nueva ya trae las variables.

## Instalación

Con `curl`:

```bash
curl -fsSL https://raw.githubusercontent.com/its-lautaro/lautaro-env-store/main/install.sh | bash
```

O clonando:

```bash
git clone https://github.com/its-lautaro/lautaro-env-store.git
cd lautaro-env-store && ./install.sh
```

Se instala en `~/.local/bin`. Si esa carpeta no está en tu `PATH`, agregá a tu `~/.zshrc`:

```bash
export PATH="$HOME/.local/bin:$PATH"
```

Requiere Python 3.8+ (el de macOS alcanza) y `pbcopy` (macOS) o `xclip` / `xsel` / `wl-copy` (Linux) para el clipboard.

## Uso

Varias variables (la lista va entre comillas simples para que el shell no interprete los paréntesis):

```bash
lautaro-env-store '[("API_USER","usuario"),("API_PASSWORD","secreto")]'
```

Una sola variable, en cualquiera de estas formas:

```bash
lautaro-env-store "API_USER","usuario"
lautaro-env-store '("API_USER","usuario")'
lautaro-env-store API_USER usuario
lautaro-env-store API_USER=usuario API_PASSWORD=secreto
```

Salida:

```
🔐 Variables guardadas en /Users/vos/.lautaro-env:
   ✅ API_USER
   ✅ API_PASSWORD
📋 Nombres copiados al clipboard: API_USER API_PASSWORD
🚀 Para usarlas en esta terminal: source /Users/vos/.lautaro-env  (las nuevas terminales ya las cargan solas)
```

Si la variable ya existía, se reemplaza su valor. Los valores nunca se imprimen ni se copian al clipboard, solo los nombres.

El instalador corre `lautaro-env-store --init`, que crea el store vacío y agrega el `source` al rc del shell, así cualquier terminal (o herramienta que cargue tu `~/.zshrc`) ya ve las variables apenas las guardás.

Para usar otro archivo: `LAUTARO_ENV_FILE=/ruta/otro.env lautaro-env-store ...`.

## Licencia

MIT
