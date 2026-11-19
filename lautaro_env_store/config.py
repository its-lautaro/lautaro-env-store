import os
import re
from pathlib import Path

STORE = Path(os.environ.get("LAUTARO_ENV_FILE", Path.home() / ".lautaro-env"))
RC_FILES = [Path.home() / ".zshrc", Path.home() / ".bashrc"]
SOURCE_LINE = f'[ -f "{STORE}" ] && source "{STORE}"  # lautaro-env-store'
NAME_RE = re.compile(r"^[A-Za-z_][A-Za-z0-9_]*$")
ENTRY_RE = re.compile(r"^export\s+(?P<name>[A-Za-z_][A-Za-z0-9_]*)=")
SAVED_RE = re.compile(r"# saved (\d{4}-\d{2}-\d{2} \d{2}:\d{2})$")

USAGE = """lautaro-env-store: guarda variables de entorno en ~/.lautaro-env y las deja listas para el shell.

Formatos aceptados:
  lautaro-env-store API_USER=usuario                          # una
  lautaro-env-store API_USER=usuario API_PASSWORD=secreto     # varias
  lautaro-env-store API_USER="Un Usuario"                     # valores con espacios: comillas del shell
  lautaro-env-store --init                                    # solo crea el store y el source en el rc

Mantenimiento:
  lautaro-env-store --list                     # nombres guardados y fecha (nunca los valores)
  lautaro-env-store --remove API_USER          # borra esas variables
  lautaro-env-store --clear [--yes]            # borra todas"""
