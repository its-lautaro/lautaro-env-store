from .config import ENTRY_RE, SAVED_RE, STORE


def read_entries():
    """Lista de (nombre, fecha) en el orden del archivo; la fecha puede ser None en entradas viejas."""
    if not STORE.exists():
        return []
    entries = []
    for line in STORE.read_text().splitlines():
        m = ENTRY_RE.match(line)
        if m:
            saved = SAVED_RE.search(line)
            entries.append((m.group("name"), saved.group(1) if saved else None))
    return entries
