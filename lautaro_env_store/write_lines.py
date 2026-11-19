import os

from .config import STORE


def write_lines(lines):
    STORE.touch(mode=0o600, exist_ok=True)
    os.chmod(STORE, 0o600)
    STORE.write_text("\n".join(lines) + ("\n" if lines else ""))
