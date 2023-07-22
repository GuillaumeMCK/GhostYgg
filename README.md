[![Build Status](https://travis-ci.com/GuillaumeMCK/GhostYgg.svg?branch=main)](https://travis-ci.com/GuillaumeMCK/GhostYgg)
[![Go Report Card](https://goreportcard.com/badge/github.com/GuillaumeMCK/GhostYgg)](https://goreportcard.com/report/github.com/GuillaumeMCK/GhostYgg)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/release/GuillaumeMCK/GhostYgg.svg)](https://github.com/GuillaumeMCK/GhostYgg/releases)

```
                                                                           ▄▄███████████████████████████████████████▄       
                                                                         ▄█████████████████████████████████████████████▄    
     ▄▄▄▓▄▄▄      ▄███▄                                                 ████▀   ▀██████▀         █▌    ██▀    ▀█▀   █████▄  
   ███████████▓   ████▌                                    ▐████      ▄██████▄    ████                              ▓██████ 
 ▄█████▀▀▀▀█████▄ ████▌▄▄▄▄▄▄      ▄▄▄█▄▄▄      ▄▄▄█▄▄▄▄  ▄██████▄▄  ▐████████▄    ██        █████▄       ▐█████    ▓███████
▐█████      ▀▀▀▀  ████████████▄  ▓██████████  ▓██████████▒████████▌     ▄██████▌             █████▌       ██████    ▓██████ 
▐████░   ███████▌▐█████▀ ▐█████ █████   █████ █████▄▄▓██▀  █████▒      ██████████      ▄█     ▀█▌▀          ██▌     ▓████▀  
▐████▒   ████████░████▌   ▓████▐████▌   ▐████ ▀█████████▄  ▓████     ▄███████████▀    ████▄            █▄           ▓███    
 █████▄    ▄████▌ ████▌   ▓████ █████   █████▒▄▄▄▒▀▀█████▌ ▓████▒    ███████████     █████████████▌    █████████    ████████
  █████████████▌  ████▌   ▓████  ███████████▀ ███████████░ ▐██████▌   █████████     █████     ▀▀▀          ▀▀▀      ███████ 
    ▀████████▀    ▀███▀   ▀███▀   ▐▀██████▀    ▀███████▀     ▀████▀    ▀██████     ███████▄         ▄███▄        ▄████████  
                                                                         ▀██████████████████████████████████████████████▀   
                                                                            ▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀      
```
[//]: # (<div align="center" style="margin-bottom: 20px; flex-direction: row;">)

[//]: # (  <img src="https://i.imgur.com/0Q8Z3ZM.png" alt="GhostYgg screenshots">)

[//]: # (   <p><b></b></p> )

[//]: # (</div>)

> Made with ☕ for fun. GhostYgg is a simple command-line tool for downloading torrents without seeding & increase download rate.

## Usage

To use GhostYgg, open a terminal or command prompt and execute the following command:

```bash
$ GhostYgg file1.client file2.client ... [options]
```

<div style="display: flex; flex-wrap: wrap; justify-content: center; gap: 20px;">

<div>

| Shortcut      | Action      |
|---------------|-------------|
| `[o]`         | Open Folder |
| `[↑]`         | Move Up     |
| `[a]`         | Add Torrent |
| `[↓]`         | Move Down   |
| `[backspace]` | Delete      |
| `[?]`         | Help        |
| `[ ]`         | Pause/Play  |
| `[ctrl+c]`    | Quit        |
| `[q]`         | Quit        |

</div>

<div>

GhostYgg supports the following options:

- `-output`: Specifies the download directory.
- `-help`: Displays the help message.

</div>

**Note**: Dragging and dropping torrent files onto the executable is supported only on Windows for now.

</div>

## Installation

### From Binary

You can download the pre-built binaries for your platform from
the [Releases]("https://github.com/GuillaumeMCK/GhostYgg/releases/")
page. After downloading the binary, make it executable if necessary, and put it into your `$PATH` or `%PATH%`.

### From Source

To get started with GhostYgg, follow these steps:

1. Clone the repository:

   ```bash
   $ git clone https://github.com/GuillaumeMCK/GhostYgg.git
   ```

2. Build the executable:

   ```bash
   $ cd GhostYgg
   $ go build ./src/main.go
   ```

3. Run the executable:

   ```bash
   $ ./GhostYgg file1.client file2.client ...
   ```

## Default Download Folder

If no download directory is specified using the `-output` flag, the tool will use the default download folder of your
operating system.

- On Linux & macOS: `/home/<YourUsername>/Downloads`
- On Windows: `C:\Users\<YourUsername>\Downloads`

## Contributing

Contributions to GhostYgg are welcome! If you have bug fixes, improvements, or new features to add, please feel free to
open a pull request 👍.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.