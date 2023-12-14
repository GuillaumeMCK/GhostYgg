# GhostYgg

[![Release](https://img.shields.io/github/release/GuillaumeMCK/GhostYgg.svg)](https://github.com/GuillaumeMCK/GhostYgg/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/GuillaumeMCK/GhostYgg)](https://goreportcard.com/report/github.com/GuillaumeMCK/GhostYgg)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

<div align="center">
    <img src="https://raw.githubusercontent.com/GuillaumeMCK/GhostYgg/main/.github/banner.png">
</div>

> Made with ☕ for fun.
> GhostYgg is a simple command-line tool for downloading torrents without seeding & increasing download rate.

## Usage

To use GhostYgg, open a terminal or command prompt and execute the following command:

```bash
$ GhostYgg file1.torrent file2.torrent ... [options]
```

### Options:

- `-o`: Specifies the download directory.
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
| <kbd>enter</kbd>                 | Select      |
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

If you using Go1.20 or higher, you can install GhostYgg using the following command:

```bash
$ go install -v github.com/GuillaumeMCK/GhostYgg@latest
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
   $ go build GhostYgg.go
   ```

3. Run the executable:

   ```bash
   $ ./GhostYgg file1.torrent file2.torrent -o .
   ```

## Disclaimer

> [!CAUTION]
> This tool is for educational purposes only. You are responsible for what you do with this tool. 