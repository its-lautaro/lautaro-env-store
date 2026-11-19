import datetime as dt

from .config import ENTRY_RE, STORE
from .read_entries import read_entries
from .shell_quote import shell_quote
from .write_lines import write_lines


def store(pairs):
    """Guarda los pares y devuelve los nombres que ya existían (se reemplazaron)."""
    existing = {name for name, _ in read_entries()}
    names = {name for name, _ in pairs}
    lines = STORE.read_text().splitlines() if STORE.exists() else []
    kept = [l for l in lines if not (ENTRY_RE.match(l) and ENTRY_RE.match(l).group("name") in names)]
    stamp = dt.datetime.now().strftime("%Y-%m-%d %H:%M")
    kept += [f"export {name}={shell_quote(value)}  # saved {stamp}" for name, value in pairs]
    write_lines(kept)
    return names & existing
