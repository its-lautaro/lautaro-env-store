# lautaro-env-store

CLI mínima para guardar variables de entorno (credenciales, tokens) en `~/.lautaro-env` sin tocar el `.zshrc` a mano. Copia al clipboard el nombre de las variables creadas y lo imprime por stdout.

El archivo `~/.lautaro-env` se crea con permisos `600` y se agrega una línea `source` a tu `~/.zshrc` (o `~/.bashrc`) la primera vez, así cada terminal nueva ya trae las variables.

## Instalación

```bash
go install github.com/its-lautaro/lautaro-env-store@latest && lautaro-env-store --init
```

Requiere Go 1.22+ para instalarlo. Para copiar al clipboard usa `pbcopy` (macOS) o `xclip` / `xsel` / `wl-copy` (Linux); sin ellos igual funciona, solo avisa que no copió.

## Uso

```
🔐 lautaro-env-store — guarda variables de entorno en /Users/vos/.lautaro-env y las deja listas para el shell

Uso:
  lautaro-env-store NOMBRE=valor [NOMBRE=valor ...]
  lautaro-env-store <comando> [args]

Ejemplos:
  lautaro-env-store API_USER=usuario
  lautaro-env-store API_USER=usuario API_PASSWORD=secreto
  lautaro-env-store API_USER="Un Usuario"  # espacios o símbolos: comillas del shell

Comandos:
  --list, ls                 Lista los nombres guardados y su fecha (nunca los valores) y los copia al clipboard
  --remove, rm NOMBRE...     Borra esas variables
  --clear [--yes]            Borra todas; pide confirmación salvo --yes o -y
  --init                     Crea el store y agrega el source al rc del shell
  --help, -h                 Muestra esta ayuda

Notas:
  Los nombres de las variables guardadas se copian al clipboard; los valores nunca se imprimen.
  El store se crea con permisos 600 y ~/.zshrc (o ~/.bashrc) lo carga en cada terminal nueva.
  LAUTARO_ENV_FILE cambia la ruta del store.
```

## Licencia

MIT
