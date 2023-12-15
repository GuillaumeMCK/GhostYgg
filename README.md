# GhostYgg

[![Release](https://img.shields.io/github/release/GuillaumeMCK/GhostYgg.svg)](https://github.com/GuillaumeMCK/GhostYgg/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/GuillaumeMCK/GhostYgg)](https://goreportcard.com/report/github.com/GuillaumeMCK/GhostYgg)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

<div align="center">
    <img src="https://raw.githubusercontent.com/GuillaumeMCK/GhostYgg/main/.github/banner.png" height="450">
</div>

> Made with ☕ for fun.
> GhostYgg is a simple command-line tool for downloading torrents without seeding & increasing download rate.

## Installation

### From Binary

You can download the pre-built binaries for your platform from
the [Releases]("https://github.com/GuillaumeMCK/GhostYgg/releases") page.
page. After downloading the binary, make it executable if necessary, and put it into your `$PATH` or `%PATH%`.

### Using Go

If you using Go1.20 or higher, you can install GhostYgg using the following command:

```bash
$ go install -v github.com/GuillaumeMCK/GhostYgg@latest
```

### From Source

To get started with GhostYgg, follow these steps:

```bash
git clone https://github.com/GuillaumeMCK/GhostYgg.git
cd GhostYgg
go build GhostYgg.go
./GhostYgg
```

## Usage

To use GhostYgg, open a terminal or command prompt and execute the following command:

> [!CAUTION]
> You are responsible for the torrents you download with GhostYgg.

### Syntax

```bash
$ GhostYgg file1.torrent file2.torrent ... [options]
```

### Options

- `-o`: Specifies the download directory.
- `-help`: Displays the help message.

If no download directory is specified using the `-o` flag, the tool will use the default download folder of your
operating system.

- On Linux & macOS: `/home/<YourUsername>/Downloads`
- On Windows: `C:\Users\<YourUsername>\Downloads`

### Keybindings

| Key                              | Action      | Description                            |
|----------------------------------|-------------|----------------------------------------|
| <kbd>o</kbd>                     | Open Folder | Open the download folder               |
| <kbd>a</kbd>                     | Add Torrent | Show the input prompt to add a torrent |
| <kbd>enter</kbd>                 | Select      | Validate the path in the input prompt  |
| <kbd>space</kbd>                 | Pause/Play  | Pause or play the selected torrent     |
| <kbd>backspace</kbd>             | Delete      | Delete the selected torrent            |
| <kbd>↑</kbd>                     | Move Up     | Cursor moves up                        |
| <kbd>↓</kbd>                     | Move Down   | Cursor moves down                      |
| <kbd>?</kbd>                     | Help        | Expand the help menu                   |
| <kbd>ctrl+c</kbd> <kbd>esc</kbd> | Exit        | Exit the program or input prompt       |

<br>
<br>
