# Port Scanner

A simple concurrent TCP port scanner written in Gooooo...
It scans the ports and gives you information whether they are open or closed.

![GoScan](assets/readme_page.png)

## Features

- TCP port scanning
- Concurrent port scanning using goroutines
- Simple terminal output

## Instalation

Binaries are avaliable on Releases page.

### Linux

[Download for Linux](https://github.com/shantanuparte/Port-Scanner/releases/download/v1.0.0/goport-linux-amd64)

Give the binary permission to run after download:

```bash
chmod +x geoshan-linux-amd64
```

Then run it:

```bash
./geoshan-linux-amd64
```

### Windows ¯\\_(ツ)_/¯(really!!)

[For windows use this](https://github.com/shantanuparte/Port-Scanner/releases/download/v1.0.0/goport-windows-amd64.exe)

## Usage

Provides a host followed by the ports you want to scan

` ./pscan localhost 22 80 443`

You can scan a remote host as well:

`./pscan google.com 443 80 8000`

You can also add a flag names -time :

`./pscan -time 5 localhost 22 80 8000`

The timeout is now 5 second

Also, if a port is wrong

`./pscan localhost 22 80 banana 99999`

Tell which are not real ports

## License

I am using MIT License for this project :D hehehe
