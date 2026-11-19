from ..config import STORE
from ..ensure_sourced import ensure_sourced


def run(args):
    STORE.touch(mode=0o600, exist_ok=True)
    rc, added = ensure_sourced()
    print(f"🧩 Store listo en {STORE}" + (f"; agregué el source a {rc}" if added else " (el shell ya lo cargaba)"))
