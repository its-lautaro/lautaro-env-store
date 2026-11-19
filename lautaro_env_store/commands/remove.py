from ..fail import fail
from ..remove import remove


def run(names):
    if not names:
        fail("--remove necesita al menos un nombre")
    missing = remove(names)
    removed = [n for n in names if n not in missing]
    for n in removed:
        print(f"🗑  {n} eliminada")
    for n in missing:
        print(f"🤷 {n} no estaba guardada")
    if removed:
        print(f"🚀 En esta terminal además: unset {' '.join(removed)}")
