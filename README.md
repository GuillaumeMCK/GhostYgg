# GhostYgg

[![Release](https://img.shields.io/github/release/GuillaumeMCK/GhostYgg.svg)](https://github.com/GuillaumeMCK/GhostYgg/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/GuillaumeMCK/GhostYgg)](https://goreportcard.com/report/github.com/GuillaumeMCK/GhostYgg)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

<div align="center">

```
GhostYgg 👻                                                                        
╭───────────────────────────────────────────────────────────────────────────────────╮
│      [?] help • [ctrl+c, esc] exit • [a] add torrent • [backspace] delete…        │
╰───────────────────────────────────────────────────────────────────────────────────╯
╭───────────────────────────────────────────────────────────────────────────────────╮
│ Name                             Progress     Seeder…  Leech…  Speed       ETA    │
│───────────────────────────────────────────────────────────────────────────────────│
│ Dave the Diver                   0.3/1.8GB    21       0       86.97MB/s   00:31… │
│ Kali Linux 2020.4                                                          ✓      │
│                                                                                   │
│                                                                                   │
│                                                                                   │
│                                                                                   │
╰───────────────────────────────────────────────────────────────────────────────────╯
```

</div>

> Made with ☕ for fun.
> GhostYgg is a simple command-line tool for downloading torrents without seeding & increasing download rate.

## Usage

To use GhostYgg, open a terminal or command prompt and execute the following command:

```bash
$ GhostYgg file1.torrent file2.torrent ... [options]
```

> **Note**: Dragging and dropping torrent files onto the executable is supported only on Windows for now.

### Options:

- `-output`: Specifies the download directory.
- `-help`: Displays the help message.

If no download directory is specified using the `-o` flag, the tool will use the default download folder of your
operating system.

- On Linux & macOS: `/home/<YourUsername>/Downloads`
- On Windows: `C:\Users\<YourUsername>\Downloads`

### Controls:

| Shortcut                         | Action      |
|----------------------------------|-------------|
| <kbd>o</kbd>                     | Open Folder |
| <kbd>a</kbd>                     | Add Torrent |
| <kbd>enter</kbd>                 | Validate    |
| <kbd>space</kbd>                 | Pause/Play  |
| <kbd>backspace</kbd>             | Delete      |
| <kbd>↑</kbd>                     | Move Up     |
| <kbd>↓</kbd>                     | Move Down   |
| <kbd>?</kbd>                     | Help        |
| <kbd>ctrl+c</kbd> <kbd>esc</kbd> | Exit        |

## Installation

### From Binary

You can download the pre-built binaries for your platform from
the [Releases]("https://github.com/GuillaumeMCK/GhostYgg/releases/")
page. After downloading the binary, make it executable if necessary, and put it into your `$PATH` or `%PATH%`.

### Using Go

If you have Go installed, you can install GhostYgg from source:

```bash
$ go get -u github.com/GuillaumeMCK/GhostYgg
```

### From Source

To get started with GhostYgg, follow these steps:

1. Clone the repository:

   ```bash
   $ git clone https://github.com/GuillaumeMCK/GhostYgg.git
   ```

2. Build the executable:

   ```bash
   $ cd GhostYgg
   $ go build -o GhostYgg src/main.go
   ```

3. Run the executable:

   ```bash
   $ ./GhostYgg file1.torrent file2.torrent -o .
   ```

## Disclaimer

You are responsible for what you download with GhostYgg.