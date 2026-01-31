# Lazy CLI

`lazy` is a lightweight CLI that picks the right command to **view**, **open**, or **execute** files based on file extension or MIME type.

## Features
- **VIEW**: Show file content with your preferred viewer.
- **OPEN**: Launch files with your default app.
- **EXEC**: Execute scripts with the matching interpreter.

## Installation
```bash
git clone https://github.com/zetatez/lazy.git
cd lazy
make
make install
```

Uninstall:
```bash
make uninstall
```

## Configuration
Configuration is compiled into the binary. Edit `config.go` and rebuild to customize defaults.

## Usage
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

## Dependencies
- [fd](https://github.com/sharkdp/fd)
- [fzf](https://github.com/junegunn/fzf)
- nvim
- ...

## Best Practice
- dwm: `st -e lazy_open_search_file_of_dir /`
- zsh:
```bash
bindkey -s '^F' 'lazy_open_search_file_of_dir ./\n'
bindkey -s '^X' 'lazy_exec_search_file_of_dir ./\n'
```

## License
MIT.
