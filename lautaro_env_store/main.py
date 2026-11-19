import sys

from .commands import clear, init, remove, save
from .commands import list as list_cmd
from .usage import usage

COMMANDS = {
    "--list": list_cmd.run,
    "ls": list_cmd.run,
    "--remove": remove.run,
    "rm": remove.run,
    "--clear": clear.run,
    "--init": init.run,
}


def main():
    args = sys.argv[1:]
    if args and args[0] in ("-h", "--help"):
        usage(0)
    if args and args[0] in COMMANDS:
        COMMANDS[args[0]](args[1:])
        return
    save.run(args)
