package rnet

import (
	"fmt"
	"net"
	"github.com/mrsongindezhou/ratel/riface"
	"github.com/mrsongindezhou/ratel/utils"
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
	// 当前的Server的消息管理模块，用来绑定MsgID和对应的处理业务API关系
	MsgHandler riface.IMsgHandler
	// 该Server的连接管理器
	ConnMgr riface.IConneManager
	// 该Server创建链接之后自动调用Hook函数--OnConnStart
	OnConnStart func(conn riface.IConnection)
	// 该Server销毁链接之前自动调用的Hook函数--OnConnStop
	OnConnStop func(conn riface.IConnection)
}

func (s *Server) Start() {
	fmt.Printf("[Ratel] Server Name : %s, listenner at IP : %s, Port: %d is starting \n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[Ratel] Version %s, MaxConn: %d, MaxPackageSize: %d \n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxCount, utils.GlobalObject.MaxPackageSize)
	go func() {
		// 0 开启消息队列及Worker工作池
		s.MsgHandler.StartWorkerPool()

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

			// 设置最大连接个数的判断，如果超过最大连接，那么则关闭此次的连接
			if s.ConnMgr.Len() >= utils.GlobalObject.MaxCount {
				// TODO 给客户端响应一个超出最大连接的错误包
				fmt.Println("Too Many Connection MaxConn = ", utils.GlobalObject.MaxCount)
				conn.Close()
				continue
			}

			//  将处理新链接的业务方法 和 conn 进行绑定 得到我们的链接模块
			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
			cid++

			// 启动当前的链接业务模块
			go dealConn.Start()
		}
	}()
}

// 停止服务
func (s *Server) Stop() {
	// 将一些服务器的资源、状态或者一些已经开辟的链接信息 进行停止或者回收
	fmt.Println("[STOP] Ratel server name", s.Name)
	s.ConnMgr.ClearConn()
}

// 运行服务器
func (s *Server) Server() {
	// 启动server的服务
	s.Start()

	// TODO 做一些启动服务器之后的额外业务
	// 阻塞状态
	select {}
}

//  路由功能，给当前的服务注册一个路由方法，供客户端链接处理使用
func (s *Server) AddRouter(msgID uint32, router riface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("Add Router Success!!!")
}

//
func (s *Server) GetConnMgr() riface.IConneManager {
	return s.ConnMgr
}

/*
	初始化Server模块的方法
*/
func NewServer(name string) riface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandler(),
		ConnMgr:    NewConnManager(),
	}

	return s
}

// 注册OnConnStart 钩子函数的方法
func (s *Server) SetOnConnStart(hookFunc func(connection riface.IConnection)) {
	s.OnConnStart = hookFunc
}

// 调用OnConnStop 钩子函数的方法
func (s *Server) SetOnConnStop(hookFunc func(connection riface.IConnection)) {
	s.OnConnStop = hookFunc
}

// 调用OnConnStart钩子函数的方法
func (s *Server) CallOnConnStart(conn riface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("-----> Call OnConnStart() ... ")
		s.OnConnStart(conn)
	}
}

// 调用OnConnStop钩子函数的方法
func (s *Server) CallOnConnStop(conn riface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("-----> Call OnConnStop() ... ")
		s.OnConnStop(conn)
	}
}
