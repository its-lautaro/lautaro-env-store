from ..config import STORE
from ..read_entries import read_entries
from ..write_lines import write_lines


def run(args):
    entries = read_entries()
    if not entries:
        print("📭 No había nada que borrar")
        return
    if "--yes" not in args and "-y" not in args:
        answer = input(f"⚠️  Vas a borrar {len(entries)} variable(s) de {STORE}. ¿Seguro? [y/N] ").strip().lower()
        if answer not in ("y", "yes", "s", "si", "sí"):
            print("👌 Cancelado")
            return
    write_lines([])
    print(f"🧹 Store vacío. En esta terminal además: unset {' '.join(n for n, _ in entries)}")
