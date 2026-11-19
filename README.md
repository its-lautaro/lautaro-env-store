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

Cada argumento es `NOMBRE=valor`, una o varias:

```bash
lautaro-env-store API_USER=usuario
lautaro-env-store API_USER=usuario API_PASSWORD=secreto
```

Si el valor tiene espacios o caracteres especiales, usá las comillas del shell. Se guardan sin las comillas:

```bash
lautaro-env-store API_USER="Un Usuario"
echo "$API_USER"   # Un Usuario
```

Salida:

```
🔐 Variables guardadas:
   ✅ API_USER
   ✅ API_PASSWORD
📋 Nombres copiados al clipboard
```

Si la variable ya existía, se reemplaza su valor y la salida lo marca con ♻️. Los valores nunca se imprimen ni se copian al clipboard, solo los nombres.

## Mantenimiento

```bash
lautaro-env-store --list                  # nombres y fecha de guardado (alias: ls)
lautaro-env-store --remove API_USER API_PASSWORD   # borra esas variables (alias: rm)
lautaro-env-store --clear                 # borra todas, pide confirmación (--yes para saltearla)
```

Borrar del store no afecta la terminal ya abierta: la salida te deja el `unset` listo para pegar.

El instalador corre `lautaro-env-store --init`, que crea el store vacío y agrega el `source` al rc del shell, así cualquier terminal (o herramienta que cargue tu `~/.zshrc`) ya ve las variables apenas las guardás.

Para usar otro archivo: `LAUTARO_ENV_FILE=/ruta/otro.env lautaro-env-store ...`.

## Licencia

MIT
