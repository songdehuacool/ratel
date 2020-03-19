package rnet

import (
	"fmt"
	"net"
	"ratel/riface"
	"ratel/utils"
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
	// 当前的Server添加一个router，server注册的链接对应的处理业务
	Router riface.IRouter
}

func (s *Server) Start() {
	fmt.Printf("[Ratel] Server Name : %s, listenner at IP : %s, Port: %d is starting \n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[Ratel] Version %s, MaxConn: %d, MaxPackageSize: %d \n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxCount, utils.GlobalObject.MaxPackageSize)
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
			dealConn := NewConnection(conn, cid, s.Router)
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

//  路由功能，给当前的服务注册一个路由方法，供客户端链接处理使用
func (s *Server) AddRouter(router riface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Success!!!")
}

/*
	初始化Server模块的方法
*/
func NewServer(name string) riface.IServer {
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}

	return s
}
