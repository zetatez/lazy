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

mkidr -p ~/.config/lazy
cp -f lazy.yaml ~/.config/lazy/
cargo install --path .
```

## Uninstall
```bash
rm -f ~/.config/lazy/lazy.yaml
cargo uninstall --bin lazy
```

## Configuration
See [*lazy.yaml*](https://github.com/zetatez/lazy/blob/main/lazy.yaml)

HINTS:
- To get mime-type of a file:
    ```bash
    file --dereference --brief --mime-type filename
    ```

## Usage

## Best Practice

- Use lazy in dwm:
    ```c
    #define SUPKEY Mod4Mask
    #define TM(cmd)    { "st", "-e", "/bin/sh", "-c", cmd, NULL }

    static const char *lazy_open[]       = TM("lazy -o \"\$(fd --type f --hidden --exclude .git . '/home/dionysus'|fzf --prompt='open>' --preview 'lazy -p {}' --select-1 --exit-0)\"");
    static const char *lazy_exec[]       = TM("lazy -e \"\$(fd -e sh -e jl -e py -e tex -e c -e cpp -e go -e scala -e java -e rs -e sql --exclude .git . '/home/dionysus'|fzf --prompt='exec>' --preview 'lazy -p {}' --select-1 --exit-0)");
    static const char *lazy_copy[]       = TM("lazy -c \"\$(fd --type f --hidden --exclude .git . '/home/dionysus'|fzf --prompt='copy>' --preview 'lazy -p {}' --select-1 --exit-0)\"");
    static const char *lazy_rename[]     = TM("lazy -r \"\$(fd --type f --hidden --exclude .git . '/home/dionysus'|fzf --prompt='rename>' --preview 'lazy -p {}' --select-1 --exit-0)\"");
    static const char *lazy_delete[]     = TM("lazy -d \"\$(fd --type f --hidden --exclude .git . '/home/dionysus'|fzf --prompt='delete>' --preview 'lazy -p {}' --select-1 --exit-0)\"");
    static const char *lazy_open_wiki[]  = TM("lazy -o \"\$(fd --type f --hidden --exclude .git . '/home/dionysus/obsidian/wiki'|fzf --prompt='wikis>' --preview 'lazy -p {}' --select-1 --exit-0)\"");
    static const char *lazy_open_book[]  = TM("lazy -o \"\$(fd -e pdf -e epub -e djvu -e mobi --exclude .git . '/home/dionysus/obsidian/library'|fzf --prompt='books>' --preview 'lazy -p {}' --reverse --select-1 --exit-0)\"");
    static const char *lazy_open_media[] = TM("lazy -o \"\$(fd -e jpg -e jpeg -e png -e gif -e bmp -e tiff -e mp3 -e flac -e mkv -e avi -e mp4 --exclude .git . '/home/dionysus'|fzf --prompt='medias>' --preview 'lazy -p {}' --reverse --select-1 --exit-0)\"");

    static Key keys[] = {
		// ...
        { SUPKEY,                       XK_f,          spawn,             {.v = lazy_open       } },
        { SUPKEY,                       XK_x,          spawn,             {.v = lazy_exec       } },
        { SUPKEY,                       XK_n,          spawn,             {.v = lazy_copy       } },
        { SUPKEY,                       XK_v,          spawn,             {.v = lazy_rename     } },
        { SUPKEY,                       XK_z,          spawn,             {.v = lazy_delete     } },
        { SUPKEY,                       XK_w,          spawn,             {.v = lazy_open_wiki  } },
        { SUPKEY,                       XK_p,          spawn,             {.v = lazy_open_book  } },
        { SUPKEY,                       XK_a,          spawn,             {.v = lazy_open_media } },
		// ...
    };
    ```

- Use lazy in zsh:
    ```bash
    alias lazy-open="lazy -o \"\$(fd --type f --hidden --exclude .git . './'|fzf --prompt='open>' --preview 'lazy -p {}' --select-1 --exit-0)\""
    alias lazy-exec="lazy -e \"\$(fd -e sh -e jl -e py -e tex -e c -e cpp -e go -e scala -e java -e rs -e sql --exclude .git . './'|fzf --prompt='exec>' --preview 'lazy -p {}' --select-1 --exit-0)\""
    alias lazy-copy="lazy -c \"\$(fd --type f --hidden --exclude .git . './'|fzf --prompt='copy>' --preview 'lazy -p {}' --select-1 --exit-0)\""
    alias lazy-rename="lazy -r \"\$(fd --type f --hidden --exclude .git . './'|fzf --prompt='rename>' --preview 'lazy -p {}' --select-1 --exit-0)\""
    alias lazy-delete="lazy -d \"\$(fd --type f --hidden --exclude .git . './'|fzf --prompt='delete>' --preview 'lazy -p {}' --select-1 --exit-0)\""
    alias lazy-open-wiki="lazy -o \"\$(fd --type f --hidden --exclude .git . '$HOME/obsidian/wiki'|fzf --prompt='wikis>' --preview 'lazy -p {}' --select-1 --exit-0)\""
    alias lazy-open-book="lazy -o \"\$(fd -e pdf -e epub -e djvu -e mobi --exclude .git . '$HOME/obsidian/library'|fzf --prompt='books>' --preview 'lazy -p {}' --reverse --select-1 --exit-0)\""
    alias lazy-open-media="lazy -o \"\$(fd -e jpg -e jpeg -e png -e gif -e bmp -e tiff -e mp3 -e flac -e mkv -e avi -e mp4 --exclude .git . '$HOME'|fzf --prompt='medias>' --preview 'lazy -p {}' --reverse --select-1 --exit-0)\""

    bindkey -s '^F' 'lazy-open\n'
    bindkey -s '^X' 'lazy-exec\n'
    bindkey -s "^N" 'lazy-copy\n'
    bindkey -s "^V" 'lazy-rename\n'
    bindkey -s "^Z" 'lazy-delete\n'
    bindkey -s '^W' 'lazy-open-wiki\n'
    bindkey -s '^P' 'lazy-open-book\n'
    bindkey -s '^A' 'lazy-open-media\n'
    ```

## LICENSE

MIT.
