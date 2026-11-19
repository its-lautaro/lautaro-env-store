from .fail import fail
from .usage import usage


def parse(args):
    """Devuelve una lista de (nombre, valor) a partir de argumentos NOMBRE=valor."""
    if not args:
        usage()
    bad = [a for a in args if "=" not in a]
    if bad:
        fail(f"cada argumento debe ser NOMBRE=valor; no entiendo: {' '.join(bad)}")
    return [tuple(a.split("=", 1)) for a in args]
