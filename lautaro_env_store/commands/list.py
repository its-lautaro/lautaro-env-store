from ..config import STORE
from ..copy_to_clipboard import copy_to_clipboard
from ..read_entries import read_entries


def run(args):
    entries = read_entries()
    if not entries:
        print(f"📭 No hay variables guardadas en {STORE}")
        return
    print(f"🗂  {len(entries)} variable{'s' if len(entries) > 1 else ''} en {STORE}:")
    for name, date in entries:
        print(f"   🔑 {name}" + (f"  (guardada {date})" if date else ""))
    copy_to_clipboard(" ".join(name for name, _ in entries))
