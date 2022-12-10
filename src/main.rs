/*
Author: Lorenzo
Email : zetatez@icloud.com

*/

use serde::{Deserialize, Serialize};
use std::io::Write;

#[derive(Debug, Serialize, Deserialize)]
struct Config {
    view_rule: std::collections::HashMap<String, String>,
    open_rule: std::collections::HashMap<String, String>,
    exec_rule: std::collections::HashMap<String, String>,
}

fn load_config() -> Config {
    let config_file_path = std::env::var("XDG_CONFIG_HOME").unwrap() + "/lazy/lazy.yaml";
    if !std::path::Path::new(&config_file_path).exists() {
        panic!("config file `{}` not found", config_file_path)
    }
    let yaml = std::fs::read_to_string(&config_file_path).unwrap();
    let cfg: Config = serde_yaml::from_str(&yaml).expect("load config file failed");
    cfg
}

struct Lazy {
    cfg: Config,
    file_path: String,
}

impl Lazy {
    fn new(cfg: Config, file_path: String) -> Lazy {
        Lazy { cfg, file_path }
    }

    fn preview(&self) {
        let mime_type = self._get_mime_type();
        match self.cfg.view_rule.get(&mime_type).as_ref() {
            Some(&cmd) => {
                let fullcmd: String = {
                    if cmd.contains("{}") {
                        let fullcmd = format!("{}", cmd.replace("{}", &self.file_path));
                        fullcmd
                    } else {
                        let fullcmd = format!("{} '{}'", cmd, self.file_path);
                        fullcmd
                    }
                };
                println!("bash -c \"{}\"", fullcmd);
                std::process::Command::new("bash")
                    .arg("-c")
                    .arg(fullcmd)
                    .spawn()
                    .unwrap()
                    .wait()
                    .expect("unexpected error occurred");
            }
            _ => println!("mime type: {} is not supported", mime_type),
        }
    }

    fn open(&self) {
        let mime_type = self._get_mime_type();
        match self.cfg.open_rule.get(&mime_type).as_ref() {
            Some(&cmd) => {
                let fullcmd: String = {
                    if cmd.contains("{}") {
                        let fullcmd = format!("{}", cmd.replace("{}", &self.file_path));
                        fullcmd
                    } else {
                        let fullcmd = format!("{} '{}'", cmd, self.file_path);
                        fullcmd
                    }
                };
                println!("bash -c \"{}\"", fullcmd);
                std::process::Command::new("bash")
                    .arg("-c")
                    .arg(fullcmd)
                    .spawn()
                    .unwrap()
                    .wait()
                    .expect("unexpected error occurred");
            }
            _ => println!("mime type: {} is not supported", mime_type),
        }
    }

    fn exec(&self) {
        let extension = std::path::Path::new(&self.file_path)
            .extension()
            .unwrap()
            .to_str()
            .unwrap()
            .to_string();
        match self.cfg.exec_rule.get(&extension).as_ref() {
            Some(&cmd) => {
                let fullcmd: String = {
                    if cmd.contains("{}") {
                        let fullcmd = format!("{}", cmd.replace("{}", &self.file_path));
                        fullcmd
                    } else {
                        let fullcmd = format!("{}", cmd);
                        fullcmd
                    }
                };
                println!("bash -c \"{}\"", fullcmd);
                std::process::Command::new("bash")
                    .arg("-c")
                    .arg(fullcmd)
                    .spawn()
                    .unwrap()
                    .wait()
                    .expect("unexpected error occurred");
            }
            _ => println!("extension type: {} is not supported", extension),
        }
    }

    fn copy(&self) {
        let parent = self._get_parent();
        print!("cp -f {} {}/", self.file_path, parent);
        std::io::stdout().flush().unwrap();
        let dst_file_path = self._input_file_path();
        std::fs::copy(&self.file_path, dst_file_path).expect("copy file failed");
    }

    fn rename(&self) {
        let parent = self._get_parent();
        print!("mv -f {} {}/", self.file_path, parent);
        std::io::stdout().flush().unwrap();
        let dst_file_path = self._input_file_path();
        std::fs::rename(&self.file_path, dst_file_path).expect("rename file failed");
    }

    fn delete(&self) {
        println!("rm -f {}", self.file_path);
        std::fs::remove_file(&self.file_path).expect("remove file failed");
    }

    fn _input_file_path(&self) -> String {
        let parent = self._get_parent();
        let mut file_name_new = String::new();
        std::io::stdin().read_line(&mut file_name_new).unwrap();
        parent + "/" + file_name_new.trim()
    }

    fn _get_mime_type(&self) -> String {
        let output = std::process::Command::new("file")
            .arg("--mime-type")
            .arg("--brief")
            .arg(&self.file_path)
            .output()
            .expect("get mime type failed");
        String::from_utf8(output.stdout).unwrap().trim().to_string()
    }

    fn _get_parent(&self) -> String {
        std::path::Path::new(&self.file_path)
            .parent()
            .unwrap()
            .to_str()
            .unwrap()
            .to_owned()
    }

    fn version() {
        println!("chopin-1.0");
    }

    fn help() {
        let help_str = "
    NAME
        lazy - A cli tool that greatly improves your work efficiency.

    SYNOPSIS
        lazy [-vh]
        lazy [-oecrd] file

    DESCRIPTION
        lazy is a tool for cli to open, exec, copy, rename, delete file automatically.

    OPTIONS
        -v     prints version information to cli and exit.
        -p     preview a file with your default settings automatically.
        -o     open a file with your default settings automatically.
        -e     exec a script with your default settings automatically.
        -c     copy a file.
        -r     rename a file.
        -d     delete a file.

    BUGS
        Send all bug reports with a patch to zetatez@icloud.com.
    ";

        println!("{}", help_str)
    }
}

fn main() {
    let args: Vec<String> = std::env::args().collect();
    if args.len() == 2 {
        if args[1] == "-v" {
            Lazy::version();
            return;
        }
        if args[1] == "-h" {
            Lazy::help();
            return;
        }
    }

    if args.len() != 3 {
        Lazy::help();
        return;
    }

    let option = &args[1];
    let file_path = &args[2];

    let lazy: Lazy = Lazy::new(load_config(), file_path.to_string());
    match option.as_str() {
        "-v" => Lazy::version(),
        "-p" => lazy.preview(),
        "-o" => lazy.open(),
        "-e" => lazy.exec(),
        "-c" => lazy.copy(),
        "-r" => lazy.rename(),
        "-d" => lazy.delete(),
        _ => Lazy::help(),
    }
}
