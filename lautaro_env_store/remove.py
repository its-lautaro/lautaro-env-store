from .config import ENTRY_RE, STORE
from .read_entries import read_entries
from .write_lines import write_lines


def remove(names):
    """Borra las variables indicadas; devuelve las que no existían."""
    lines = STORE.read_text().splitlines() if STORE.exists() else []
    present = {name for name, _ in read_entries()}
    kept = [l for l in lines if not (ENTRY_RE.match(l) and ENTRY_RE.match(l).group("name") in names)]
    write_lines(kept)
    return set(names) - present
