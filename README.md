# unshorten

A simple, verbose URL unshortener.

## Usage

```text
unshorten [OPTIONS] URL
```

### Options

- `-max-redirects`: Maximum number of redirects to follow. (default 10)
- `-quiet`: Run quietly; only display the final URL.
- `-version`: Print version and exit.

## Installation

### macOS via Homebrew

```shell
brew install cdzombak/oss/unshorten
```

### Debian via apt repository

Install my Debian repository if you haven't already:

```shell
sudo apt-get install ca-certificates curl gnupg
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://dist.cdzombak.net/deb.key | sudo gpg --dearmor -o /etc/apt/keyrings/dist-cdzombak-net.gpg
sudo chmod 0644 /etc/apt/keyrings/dist-cdzombak-net.gpg
echo -e "deb [signed-by=/etc/apt/keyrings/dist-cdzombak-net.gpg] https://dist.cdzombak.net/deb/oss any oss\n" | sudo tee -a /etc/apt/sources.list.d/dist-cdzombak-net.list > /dev/null
sudo apt-get update
```

Then install `unshorten` via `apt-get`:

```shell
sudo apt-get install unshorten
```

### Manual installation from build artifacts

Pre-built binaries for Linux and macOS on various architectures are downloadable from each [GitHub Release](https://github.com/cdzombak/unshorten/releases). Debian packages for each release are available as well.

### Build and install locally

```shell
git clone https://github.com/cdzombak/unshorten.git
cd unshorten
make build

cp out/unshorten $INSTALL_DIR
```

## Docker images

Docker images are available for a variety of Linux architectures from [Docker Hub](https://hub.docker.com/r/cdzombak/unshorten) and [GHCR](https://github.com/cdzombak/unshorten/pkgs/container/unshorten). Images are based on the `scratch` image and are as small as possible.

Run them via, for example:

```shell
docker run --rm cdzombak/unshorten:1 https://dzdz.cz
docker run --rm ghcr.io/cdzombak/unshorten:1 https://dzdz.cz
```

## About

- Issues: [github.com/cdzombak/unshorten/issues]https://github.com/cdzombak/unshorten/issues
- Author: [Chris Dzombak](https://www.dzombak.com)
  - [GitHub: @cdzombak](https://www.github.com/cdzombak)

## License

MIT; see `LICENSE` in this repository.
