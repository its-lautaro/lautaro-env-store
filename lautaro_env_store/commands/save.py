from ..config import NAME_RE, STORE
from ..copy_to_clipboard import copy_to_clipboard
from ..ensure_sourced import ensure_sourced
from ..fail import fail
from ..parse import parse
from ..store import store


def run(args):
    pairs = parse(args)
    for name, _ in pairs:
        if not NAME_RE.match(name):
            fail(f"'{name}' no es un nombre válido de variable de entorno")

    replaced = store(pairs)
    rc, added = ensure_sourced()
    names = [name for name, _ in pairs]
    copied = copy_to_clipboard(" ".join(names))

    plural = "s" if len(names) > 1 else ""
    print(f"🔐 Variable{plural} guardada{plural}:")
    for name in names:
        print(f"   {'♻️  ' + name + ' (reemplazada)' if name in replaced else '✅ ' + name}")
    if copied:
        print(f"📋 Nombre{plural} copiado{plural} al clipboard")
    else:
        print("⚠️  No encontré un comando de clipboard (pbcopy/xclip/wl-copy); los nombres quedan arriba.")
    if added:
        print(f"🧩 Agregué el source de {STORE} a {rc}")
