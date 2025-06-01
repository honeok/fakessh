## FakeSSH

<p align="center">
<img src="https://img.shields.io/github/license/honeok/fakessh.svg?style=flat" alt="License" />
<img src="https://img.shields.io/github/last-commit/honeok/fakessh?style=flat" alt="Last Commit" />
<img src="https://img.shields.io/github/commit-activity/m/honeok/fakessh.svg?style=flat" alt="Commit Activity" />
</p>

A Docker containerized honeypot SSH server written in Go that logs login attempts.

Password authentication always fails, so the attacker cannot gain terminal access. 💀

[![](https://dockerico.blankenship.io/image/honeok/fakessh)](https://hub.docker.com/r/honeok/fakessh)

### Quick Start

```shell
docker run -it --rm -p 22:22 --name fakessh honeok/fakessh
```
or
```shell
docker run -d --restart=unless-stopped -p 22:22 --name fakessh honeok/fakessh
docker logs -f fakessh
```

### Local build

```shell
git clone https://github.com/honeok/fakessh.git
cd fakessh
CGO_ENABLED=0 go build -v -trimpath -ldflags="-s -w -buildid=" -o ./fakessh fakessh.go
```

### See also

* [fffaraz/fakessh](https://github.com/fffaraz/fakessh) - A dockerized fake SSH server honeypot written in Go that logs login attempts.