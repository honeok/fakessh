# FakeSSH

这是一个使用 Go 语言编写的 Docker 化蜜罐 SSH 服务器，旨在记录登录尝试的情况。
所有密码验证都会失败，因此攻击者无法获得终端访问权限。

[![Build Status](https://github.com/honeok/fakessh/actions/workflows/build-image.yml/badge.svg)](https://github.com/honeok/fakessh/actions/workflows/build-image.yml/badge.svg)
[![](https://dockeri.co/image/honeok/fakessh)](https://hub.docker.com/r/honeok/fakessh)

## 快速入门

```shell
go install github.com/fffaraz/fakessh@latest
sudo setcap 'cap_net_bind_service=+ep' ~/go/bin/fakessh
fakessh [optional-log-directory]
```
OR
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