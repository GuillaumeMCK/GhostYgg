# GhostYgg

[![Build Status](https://travis-ci.com/GuillaumeMCK/GhostYgg.svg?branch=main)](https://travis-ci.com/GuillaumeMCK/GhostYgg)
[![Go Report Card](https://goreportcard.com/badge/github.com/GuillaumeMCK/GhostYgg)](https://goreportcard.com/report/github.com/GuillaumeMCK/GhostYgg)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/release/GuillaumeMCK/GhostYgg.svg)](https://github.com/GuillaumeMCK/GhostYgg/releases)

<div align="center">

```
╭───────────────────────────────────────────────────────────────────────────────────╮
│ Name                             Progress     Seeder…  Leech…  Download …  ETA    │
│───────────────────────────────────────────────────────────────────────────────────│
│ Dave the Diver                   0.3/1.8GB    21       0       86.97MB/s   00:01… │
│ Kali Linux 2020.4                                                          ✓      │
│                                                                                   │
│                                                                                   │
│                                                                                   │
│                                                                                   │
│                                                                                   │
│                                                                                   │
│                                                                                   │
│                                                                                   │
│                                                                                   │
╰───────────────────────────────────────────────────────────────────────────────────╯
 [?] help • [ctrl+c] [q] quit • [a] add torrent • [backspace] delete • [ ] pause/play 
```

</div>

> Made with ☕ for fun. <br>
> GhostYgg is a simple command-line tool for downloading torrents without seeding & increase download rate.

## Usage

To use GhostYgg, open a terminal or command prompt and execute the following command:

```bash
$ GhostYgg file1.client file2.client ... [options]
```

> **Note**: Dragging and dropping torrent files onto the executable is supported only on Windows for now.

### Options:

- `-output`: Specifies the download directory.
- `-help`: Displays the help message.

If no download directory is specified using the `-output` flag, the tool will use the default download folder of your
operating system.

- On Linux & macOS: `/home/<YourUsername>/Downloads`
- On Windows: `C:\Users\<YourUsername>\Downloads`

### Controls:

| Shortcut                       | Action      |
|--------------------------------|-------------|
| <kbd>o</kbd>                   | Open Folder |
| <kbd>a</kbd>                   | Add Torrent |
| <kbd>space</kbd>               | Pause/Play  |
| <kbd>↑</kbd>                   | Move Up     |
| <kbd>↓</kbd>                   | Move Down   |
| <kbd>backspace</kbd>           | Delete      |
| <kbd>?</kbd>                   | Help        |
| <kbd>ctrl+c</kbd> <kbd>q</kbd> | Quit        |

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

## Contributing

Contributions to GhostYgg are welcome! If you have bug fixes, improvements, or new features to add, please feel free to
open a pull request 👍.
