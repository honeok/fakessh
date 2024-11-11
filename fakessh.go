package main

import (
	"crypto/rand"        // 用于生成随机数（用于密钥生成）
	"crypto/rsa"         // RSA加密算法库，用来生成服务器的私钥
	"errors"             // 错误处理库
	"fmt"                // 格式化输出库
	"log"                // 日志库
	"net"                // 网络库，用于监听TCP连接
	"os"                 // 操作系统相关库，用于获取命令行参数和文件操作
	"time"               // 时间库，用于时间相关操作

	"golang.org/x/crypto/ssh"  // SSH协议库
)

var (
	errBadPassword = errors.New("permission denied") // 自定义错误信息，表示权限被拒绝
	serverVersions = []string{
		// 不同的SSH服务器版本信息，伪装成常见的SSH版本
		"SSH-2.0-OpenSSH_8.9p1 Ubuntu-3ubuntu0.10",
		"SSH-2.0-OpenSSH_6.6.1p1 Ubuntu-2ubuntu2.3",
		"SSH-2.0-OpenSSH_6.7p1 Debian-5+deb8u3",
		"SSH-2.0-OpenSSH_7.2p2 Ubuntu-4ubuntu2.10",
		"SSH-2.0-OpenSSH_7.4",
		"SSH-2.0-OpenSSH_8.0",
		"SSH-2.0-OpenSSH_8.4p1 Debian-2~bpo10+1",
		"SSH-2.0-OpenSSH_8.4p1 Debian-5+deb11u1",
		"SSH-2.0-OpenSSH_8.9p1 Ubuntu-3ubuntu0.6",
	}
)

func main() {
	// 检查是否有命令行参数，若有参数则创建日志文件
	if len(os.Args) > 1 {
		logPath := fmt.Sprintf("%s/fakessh-%s.log", os.Args[1], time.Now().Format("2006-01-02-15-04-05-000"))
		// 创建日志文件并设置文件权限为644
		logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			log.Println("Failed to open log file:", logPath, err)
			return
		}
		defer logFile.Close()     // 程序退出时关闭日志文件
		log.SetOutput(logFile)    // 设置日志输出到文件
	}

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)  // 设置日志标记格式，包括日期和微秒时间

	serverConfig := &ssh.ServerConfig{
		MaxAuthTries:     5,                     // 设置最大尝试次数为5次
		PasswordCallback: passwordCallback,      // 设置密码回调函数
		ServerVersion:    serverVersions[0],     // 设置SSH服务器版本信息
	}

	// 生成2048位的RSA密钥
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	// 创建一个签名者对象作为服务器的主机密钥
	signer, _ := ssh.NewSignerFromSigner(privateKey)
	serverConfig.AddHostKey(signer)            // 将生成的密钥添加到服务器配置中

	// 监听TCP连接，端口为22（SSH的默认端口）
	listener, err := net.Listen("tcp", ":22")
	if err != nil {
		log.Println("Failed to listen:", err)
		return
	}
	defer listener.Close()                      // 程序退出时关闭监听器

	for {
		// 接受新的连接
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept:", err)
			break
		}
		// 为每个连接启动一个新协程（goroutine）来处理
		go handleConn(conn, serverConfig)
	}
}

// passwordCallback 用于处理密码验证尝试的回调函数
func passwordCallback(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
	// 记录连接的IP地址、客户端SSH版本、用户名和尝试的密码
	log.Println(conn.RemoteAddr(), string(conn.ClientVersion()), conn.User(), string(password))
	time.Sleep(100 * time.Millisecond)        // 延时100毫秒，防止暴力破解
	return nil, errBadPassword                // 返回错误，拒绝密码验证
}

// handleConn 处理每个新连接
func handleConn(conn net.Conn, serverConfig *ssh.ServerConfig) {
	defer conn.Close()                        // 连接结束时关闭
	log.Println(conn.RemoteAddr())            // 记录连接的IP地址
	// 尝试创建SSH服务器连接（由于passwordCallback始终返回错误，实际会拒绝连接）
	ssh.NewServerConn(conn, serverConfig)
}
