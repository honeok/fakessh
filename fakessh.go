package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	errBadPassword = fmt.Errorf("Permission denied, please try again.")
	serverVersions = []string{
		"SSH-2.0-OpenSSH_9.6p1 Ubuntu-3ubuntu13.5",
	}
)

func main() {
	// 设置日志文件 (如果提供目录)
	if len(os.Args) > 1 {
		logPath := fmt.Sprintf("%s/fakessh-%s.log", os.Args[1], time.Now().Format("2006-01-02-15-04-05-000"))
		logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			log.Println("Failed to open log file:", logPath, err)
			return
		}
		defer logFile.Close()
		log.SetOutput(logFile)
	}

	log.SetFlags(log.LstdFlags | log.Lmicroseconds) // 设置日志时间格式

	// 配置 SSH 服务器
	serverConfig := &ssh.ServerConfig{
		MaxAuthTries:     6, // 最大认证尝试次数
		PasswordCallback: passwordCallback, // 密码认证回调
		ServerVersion:    serverVersions[0], // 使用 Ubuntu24.04
	}

	// 生成 RSA 密钥并添加主机密钥
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	signer, _ := ssh.NewSignerFromSigner(privateKey)
	serverConfig.AddHostKey(signer)

	listener, err := net.Listen("tcp", ":22")
	if err != nil {
		log.Println("Failed to listen:", err)
		return
	}
	defer listener.Close()

	// 循环接受连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept:", err)
			break
		}
		go handleConn(conn, serverConfig) // 处理连接
	}
}

// 记录认证尝试并拒绝
func passwordCallback(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
	log.Println(conn.RemoteAddr(), string(conn.ClientVersion()), conn.User(), string(password))
	time.Sleep(100 * time.Millisecond) // 模拟延迟
	return nil, errBadPassword
}

// 处理客户端连接
func handleConn(conn net.Conn, serverConfig *ssh.ServerConfig) {
	defer conn.Close()
	log.Println(conn.RemoteAddr())
	ssh.NewServerConn(conn, serverConfig)
}