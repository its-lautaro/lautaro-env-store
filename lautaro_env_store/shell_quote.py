def shell_quote(value):
    return "'" + value.replace("'", "'\"'\"'") + "'"
