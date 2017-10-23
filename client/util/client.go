package client

import (
	"github.com/fari-proxy/service"
	"github.com/fari-proxy/encryption"
	"net"
	"log"
)

type client struct {
	*service.Service
}

// 新建一个本地端
// 本地端的职责是:
// 0.监听来自本地浏览器的代理请求
// 1.转发前加密数据
// 2.转发socket数据到服务端
// 3.把服务端返回的数据转发给用户的浏览器
func NewClient(remote, listen, password string) *client {
	c := encryption.NewCipher([]byte(password))
	listenAddr, _ := net.ResolveTCPAddr("tcp", listen)
	remoteAddr, _ := net.ResolveTCPAddr("tcp", remote)
	return &client{
		&service.Service{
			Cipher:	c,
			ListenAddr:	listenAddr,
			RemoteAddr: remoteAddr,
		},
	}
}

// 本地端启动监听给用户的浏览器调用
func (c *client) Listen() error {
	listener, err := net.ListenTCP("tcp", c.ListenAddr)
	if err != nil {
		return err
	}
	log.Printf("启动成功,监听在 %s:%d, 密码: %s", c.ListenAddr.IP, c.ListenAddr.Port, c.Cipher.Password)
	defer listener.Close()

	for {
		userConn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		// userConn被关闭时直接清除所有数据 不管没有发送的数据
		userConn.SetLinger(0)
		go c.handleConn(userConn)
	}
	return nil
}

func (c *client) handleConn(userConn *net.TCPConn) {
	defer userConn.Close()

	proxyServer, err := c.DialRemote()
	if err != nil {
		log.Println(err)
		return
	}
	defer proxyServer.Close()
	// Conn被关闭时直接清除所有数据 不管没有发送的数据
	proxyServer.SetLinger(0)

	// 进行转发
	// 从 proxyServer 读取数据发送到 localUser
	go func() {
		err := c.DecodeCopy(userConn, proxyServer)
		if err != nil {
			// 在 copy 的过程中可能会存在网络超时等 error 被 return，只要有一个发生了错误就退出本次工作
			userConn.Close()
			proxyServer.Close()
		}
	}()
	// 从 localUser 发送数据发送到 proxyServer，这里因为处在翻墙阶段出现网络错误的概率更大
	c.EncodeCopy(proxyServer, userConn)
}

