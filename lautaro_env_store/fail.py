import sys

from .usage import usage


def fail(message):
    print(f"❌ {message}", file=sys.stderr)
    usage()
