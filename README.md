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
See [*lazy.yaml*](https://github.com/zetatez/lazy/blob/master/lazy.yaml)

## Usage
```bash
lazy -h
```

## Best Practice

- Use lazy in dwm:
    ```c
    #define SUPKEY Mod4Mask
    #define TM(cmd)    { "st", "-e", "/bin/sh", "-c", cmd, NULL }

    static const char *cmd_lazy_exec[]            =  ST("lazy -o exec -f \"$(fd --type f -e sh -e jl -e py -e tex -e c -e cpp -e go -e scala -e java -e rs -e sql --exclude .git . '${HOME}'|fzf --prompt='exec>' --preview 'lazy -o view -f {}' --select-1 --exit-0|xargs lazy -e {}\"");
    static const char *cmd_lazy_copy[]            =  ST("lazy -o copy -f \"$(fd --type f --hidden --exclude .git . '/home/dionysus'|fzf --prompt='copy>' --preview 'lazy -o view -f {}' --select-1 --exit-0)\"");
    static const char *cmd_lazy_rename[]          =  ST("lazy -o rename -f \"$(fd --type f --hidden --exclude .git . '/home/dionysus'|fzf --prompt='rename>' --preview 'lazy -o view -f {}' --select-1 --exit-0)\"");
    static const char *cmd_lazy_delete[]          =  ST("lazy -o delete -f \"$(fd --type f --hidden --exclude .git . '/home/dionysus'|fzf --prompt='delete>' --preview 'lazy -o view -f {}' --select-1 --exit-0)\"");
    static const char *cmd_lazy_open[]               =  ST("lazy -o open -f \"$(fd --type f --hidden --exclude .git . '/home/dionysus/'|fzf --prompt='open>' --preview 'lazy -o view -f {}' --select-1 --exit-0)\"");
    static const char *cmd_lazy_open_book[]          =  ST("lazy -o open -f \"$(fd --type f -e pdf -e epub -e djvu -e mobi --exclude .git . '/home/dionysus/my-library'|fzf --prompt='books>' --preview 'lazy -o view -f {}' --reverse --select-1 --exit-0)\"");
    static const char *cmd_lazy_open_media[]         =  ST("lazy -o open -f \"$(fd --type f -e jpg -e jpeg -e png -e gif -e bmp -e tiff -e mp3 -e flac -e mkv -e avi -e mp4 --exclude .git . '/home/dionysus/'|fzf --prompt='medias>' --preview 'lazy -o view -f {}' --reverse --select-1 --exit-0)\"");
    static const char *cmd_lazy_open_wiki[]          =  ST("lazy -o open -f \"$(fd --type f --hidden --exclude .git . '/home/dionysus/my-wiki'|fzf --prompt='wikis>' --preview 'lazy -o view -f {}' --select-1 --exit-0)\"");

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
    alias lazy-open="lazy -o open -f \"\$(fd --type f --hidden --exclude .git . './'|fzf --prompt='open>' --preview 'lazy -o view -f {}' --select-1 --exit-0)\""
    alias lazy-exec="lazy -o exec -f \"\$(fd --type f -e sh -e jl -e py --exclude .git . './'|fzf --prompt='exec>' --preview 'lazy -o view -f {}' --select-1 --exit-0)\""
    alias lazy-copy="lazy -o copy -f \"\$(fd --type f --hidden --exclude .git . './'|fzf --prompt='copy>' --preview 'lazy -o view -f {}' --select-1 --exit-0)\""
    alias lazy-rename="lazy -o rename -f \"\$(fd --type f --hidden --exclude .git . './'|fzf --prompt='rename>' --preview 'lazy -o view -f {}' --select-1 --exit-0)\""
    alias lazy-delete="lazy -o delete -f \"\$(fd --type f --hidden --exclude .git . './'|fzf --prompt='delete>' --preview 'lazy -o view -f {}' --select-1 --exit-0)\""
    alias lazy-open-wiki="lazy -o open -f \"\$(fd --type f --hidden --exclude .git . '$HOME/my-wiki'|fzf --prompt='wikis>' --preview 'lazy -o view -f {}' --select-1 --exit-0)\""
    alias lazy-open-book="lazy -o open -f \"\$(fd --type f -e pdf -e epub -e djvu -e mobi --exclude .git . '$HOME/my-library'|fzf --prompt='books>' --preview 'lazy -o view -f {}' --reverse --select-1 --exit-0)\""
    alias lazy-open-media="lazy -o open -f \"\$(fd --type f -e jpg -e jpeg -e png -e gif -e bmp -e tiff -e mp3 -e flac -e mkv -e avi -e mp4 --exclude .git . '$HOME'|fzf --prompt='medias>' --preview 'lazy -o view -f {}' --reverse --select-1 --exit-0)\""

    bindkey -s '^F' 'lazy-open\n'
    bindkey -s '^X' 'lazy-exec\n'
    bindkey -s '^N' 'lazy-copy\n'
    bindkey -s '^V' 'lazy-rename\n'
    bindkey -s '^Z' 'lazy-delete\n'
    bindkey -s '^W' 'lazy-open-wiki\n'
    bindkey -s '^P' 'lazy-open-book\n'
    bindkey -s '^A' 'lazy-open-media\n'
    ```

## LICENSE

MIT.
