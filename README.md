# Lazy CLI

`lazy` is a lightweight command-line tool that automatically selects the appropriate way to **view, open, or execute files** based on their type, improving your workflow efficiency.

---

## Features

- **VIEW**: View file content using the default configured method.
- **OPEN**: Open files with their default applications.
- **EXEC**: Execute scripts or executable files.

`lazy` selects commands based on file extension or MIME type, and falls back to default commands if no match is found.

---

## Installation

1. Clone the repository

```bash
git clone https://github.com/zetatez/lazy.git
cd lazy
```

2. Build the binary
```bash
make
```

3. Install the binary
```bash
make install
```

4. Uninstall the binary
```bash
make uninstall
```

---

## Dependencies
- [fd](https://github.com/sharkdp/fd)
- [fzf](https://github.com/junegunn/fzf)
- nvim
- ...

---

## Configuration
Configuration is compiled into the binary. Edit `config.go` to customize default commands.

---

## Usage Examples
```bash
# Show help
lazy -h

# View a file
lazy -o view -f /path/to/file

# Open a file
lazy -o open -f /path/to/file

# Execute a script
lazy -o exec -f /path/to/script.sh
```

## Best Practice

- Use lazy in dwm:
    - Hint: `st -e lazy_open_search_file_of_dir /`

- Use lazy in zsh:
    ```bash
    bindkey -s '^F' 'lazy_open_search_file_of_dir ./\n'
    bindkey -s '^X' 'lazy_exec_search_file_of_dir ./\n'
    ```

## LICENSE

MIT.
