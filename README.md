# lazy
A cli tool that will greatly improve your working efficiency.

## Dependencies
- [fd](https://github.com/sharkdp/fd)
- [fzf](https://github.com/junegunn/fzf)
- nvim
- ...

## Installation
```bash
git clone https://github.com/zetatez/lazy.git
cd lazy

make install
```

## Uninstall
```bash
make uninstall
```

## Configuration
See *etc*
```
lazy
├── etc
```

## Usage
```bash
lazy -h
```

## Best Practice

- Use lazy in dwm:
    - Try it yourself. Hint: `st -e lazy-open-search-file-of-dir /`

- Use lazy in zsh:
    ```bash
    bindkey -s '^F' 'lazy-open-search-file-of-dir ./\n'
    bindkey -s '^X' 'lazy-exec-search-file-of-dir ./\n'
    bindkey -s '^N' 'lazy-copy-search-file-of-dir ./\n'
    bindkey -s '^V' 'lazy-rename-search-file-of-dir ./\n'
    bindkey -s '^Z' 'lazy-delete-search-file-of-dir ./\n'
    ```

## LICENSE

MIT.
