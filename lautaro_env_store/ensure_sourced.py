from .config import RC_FILES, SOURCE_LINE


def ensure_sourced():
    """Agrega la línea de source al rc del shell que exista (zsh primero), una sola vez."""
    for rc in RC_FILES:
        if rc.exists():
            content = rc.read_text()
            if "lautaro-env-store" in content:
                return rc, False
            with rc.open("a") as f:
                f.write(f"\n# lautaro-env-store: variables guardadas con `lautaro-env-store`\n{SOURCE_LINE}\n")
            return rc, True
    return None, False
