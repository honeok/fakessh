# FakeSSH

[![Build Status](https://github.com/honeok/fakessh/actions/workflows/publish.yml/badge.svg)](https://github.com/honeok/fakessh/actions/workflows/publish.yml/badge.svg)
[![License](https://img.shields.io/github/license/honeok/fakessh.svg?style=flat)](./LICENSE)
[![Commit](https://img.shields.io/github/last-commit/honeok/fakessh)](https://github.com/honeok/fakessh)
[![Commit Activity](https://img.shields.io/github/commit-activity/m/honeok/fakessh.svg)](https://github.com/honeok/fakessh)
[![Docker Pulls](https://img.shields.io/docker/pulls/honeok/fakessh.svg)](https://hub.docker.com/r/honeok/fakessh)
[![Issues](https://img.shields.io/github/issues/honeok/fakessh.svg)](https://img.shields.io/github/issues/honeok/fakessh.svg)
[![Stars](https://img.shields.io/github/stars/honeok/fakessh.svg)](https://img.shields.io/github/stars/honeok/fakessh.svg)

<br>
这是一个使用 Go 语言编写的 Docker 化蜜罐 SSH 服务器，旨在记录登录尝试的情况。
所有密码验证都会失败，因此攻击者无法获得终端访问权限。
</br>
<br>
This is a Dockerized honeypot SSH server written in Go, designed to log login attempts. 
All password verifications will fail, preventing attackers from gaining terminal access.
Let me know if you'd like to adjust or expand it further!
</br>
****

## 快速入门

```shell
docker run -it --rm -p 22:22 honeok/fakessh
```
OR
```shell
docker run -d --restart=unless-stopped -p 22:22 --name fakessh honeok/fakessh
docker logs -f fakessh
```
OR
```yaml
services:
  fakessh:
    image: honeok/fakessh
    container_name: fakessh
    restart: unless-stopped
    ports:
      - 22:22
    command: /log
    volumes:
      - ./log:/log
    networks:
      - fakessh

networks:
  fakessh:
    driver: bridge
```

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=honeok/fakessh&type=Date)](https://star-history.com/#honeok/fakessh&Date)

### See also

* [jaksi/sshesame](https://github.com/jaksi/sshesame) - A fake SSH server that lets everyone in and logs their activity.
* [shazow/ssh-chat](https://github.com/shazow/ssh-chat) - Custom SSH server written in Go. Instead of a shell, you get a chat prompt.
* [gliderlabs/ssh](https://github.com/gliderlabs/ssh) - Easy SSH servers in Golang.
* [gliderlabs/sshfront](https://github.com/gliderlabs/sshfront) - Programmable SSH frontend.
* [desaster/kippo](https://github.com/desaster/kippo) - Kippo - SSH Honeypot.
* [micheloosterhof/cowrie](https://github.com/micheloosterhof/cowrie) - Cowrie SSH/Telnet Honeypot.
* [fzerorubigd/go0r](https://github.com/fzerorubigd/go0r) - A simple ssh honeypot in golang.
* [droberson/ssh-honeypot](https://github.com/droberson/ssh-honeypot) - Fake sshd that logs ip addresses, usernames, and passwords.
* [x0rz/ssh-honeypot](https://github.com/x0rz/ssh-honeypot) - Fake sshd that logs ip addresses, usernames, and passwords.
* [tnich/honssh](https://github.com/tnich/honssh) - HonSSH is designed to log all SSH communications between a client and server.
* [Learn from your attackers - SSH HoneyPot](https://www.robertputt.co.uk/learn-from-your-attackers-ssh-honeypot.html)
* [cowrie](https://github.com/cowrie/cowrie) - Cowrie SSH/Telnet Honeypot.
