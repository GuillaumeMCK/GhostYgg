<div align="center">

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
</div>

>A simple command-line tool for downloading torrents without seeding & increase download rate.

## Usage

```bash
$ GhostYgg file1.torrent file2.torrent ... [options]
```
**Note:** Dragging and dropping torrent files onto the executable is also supported.

## Options

- `-output` : Specifies the download directory.
- `-help` : Displays this help message.

## Prerequisites

Make sure you have Go installed on your system.

## Installation

1. Clone the repository:

   ```bash
   $ git clone https://github.com/GuillaumeMCK/GhostYgg.git
   ```

2. Build the executable:

   ```
   $ cd GhostYgg
   $ go build ./src/main.go
   ```

3. Run the executable:

   ```bash
   $ ./GhostYgg file1.torrent file2.torrent ...
   ```

## Example

To download a torrent file, run the following command:

```bash
$ ./GhostYgg example.torrent
```

## Progress Tracking

The tool will display the download progress for each torrent file, including the following information:

- Status: Downloaded bytes/Total size Percentage
- Seeders: Number of connected seeders
- Name: Torrent file name

Once the download is complete, the tool will display a message indicating the completion.

## Default Download Folder

If no download directory is specified using the `-output` flag, the tool will use the default download folder of your operating system.

On Linux:

```bash
~/Downloads
```

On Windows:

```bash
C:\Users\<YourUsername>\Downloads
```

## Interrupting the Download

To interrupt the download process, press `Ctrl+C`.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.