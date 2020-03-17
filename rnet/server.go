package rnet

import (
	"errors"
	"fmt"
	"net"
	"ratel/riface"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/3/16 11:22 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

// iserver的接口实现，定义一个Server的服务器模块
type Server struct {
	// 服务器的名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的IP
	IP string
	// 服务器监听的端口
	Port int
}

// 定义当前客户端链接的所绑定handle api(暂时写死)
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// 回显的业务
	fmt.Println("[Conn Handle] CallbackToClient ... ")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP :%s, Post: %d, is starting \n", s.IP, s.Port)
	go func() {
		// 1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}

		// 2 监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, " err", err)
			return
		}
		fmt.Println("start Ratel server succ, ", s.Name, " succ, Listening...")
		var cid uint32
		cid = 0
		// 3 阻塞的等待客户端链接，处理客户端链接业务(读写)
		for {
			// 如果有客户端链接过来，阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			//  将处理新链接的业务方法 和 conn 进行绑定 得到我们的链接模块
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++

			// 启动当前的链接业务模块
			go dealConn.Start()
		}
	}()
}

// 停止服务
func (s *Server) Stop() {

}

// 运行服务器
func (s *Server) Server() {
	// 启动server的服务
	s.Start()

	// 阻塞状态
	select {}
}

/*
	初始化Server模块的方法
*/
func NewServer(name string) riface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}