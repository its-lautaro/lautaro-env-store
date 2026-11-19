import sys

from .config import USAGE


def usage(code=1):
    print(USAGE, file=sys.stderr)
    sys.exit(code)
