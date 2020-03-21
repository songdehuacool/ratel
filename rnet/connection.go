package rnet

import "C"
import (
	"errors"
	"fmt"
	"io"
	"net"
	"ratel/riface"
	"ratel/utils"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/3/17 9:01 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

/*
	链接模块
*/
type Connection struct {
	// 当前链接的socket TCP套接字
	Conn *net.TCPConn

	// 链接的ID
	ConnID uint32

	// 当前的链接状态
	isClosed bool

	// 告知当前链接已经退出/停止 channel(由Reader告知退出)
	ExitChan chan bool

	// 无缓冲的管道，用于读、写Groutine之间读消息通信
	msgChan chan []byte

	// 消息的管理MsgID 和对应的处理业务API关系
	MsgHandler riface.IMsgHandler
}

// 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, msgHandler riface.IMsgHandler) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		isClosed:   false,
		MsgHandler: msgHandler,
		ExitChan:   make(chan bool, 1),
		msgChan:    make(chan []byte),
	}
	return c
}

/*
	写消息Goroutine，专门发送给客户端消息模块
*/
func (c *Connection) StartWrite() {
	fmt.Println("[Write Goroutine is running...]")
	defer fmt.Println(c.RemoteAddr().String(), " [conn Writer exit!]")

	// 不断的阻塞的等待channel的消息，进行写给客户端
	for {
		select {
		case data := <-c.msgChan:
			// 有数据写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data error, ", err)
				return
			}
		case <-c.ExitChan:
			// 代表Reader已经退出，此时Write也要退出
			return

		}
	}
}

// 链接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connID = ", c.ConnID, " [Reader is exit], remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中
		//buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		//_, err := c.Conn.Read(buf)
		//if err != nil {
		//	fmt.Println("recv buf err", err)
		//	continue
		//}

		// 创建一个拆包解包对象
		dp := NewDataPack()

		// 读取客户端的Msg Head 二进制 8个字节
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error: ", err)
			break
		}
		// 拆包，得到msgID 和 msgDatalen 放在msg消息中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error: ", err)
			break
		}
		// 根据dataLen 再次读取data 放在msg.Data中
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error: ", err)
				break
			}
		}

		msg.SetData(data)
		// 得到当前conn数据的Request请求数据
		req := Request{
			conn: c,
			msg:  msg,
		}

		if utils.GlobalObject.WorkerPoolSize > 0 {
			// 已经开启了工作池机制，将消息发送给Worker工作池处理即可
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			// 从路由中，找到注册绑定的Conn对应的router调用
			// 根据绑定好的MsgID 找到对应处理api业务 执行
			go c.MsgHandler.DoMsgHandler(&req)
		}
	}
}

// 发送数据，将数据发送给远程的客户端
// 提供一个SendMsg方法 将我们要发送给客户端的数据，先进行封包，再发送
func (c *Connection) SendMsg(msgId uint32, data []byte, int2 int) error {
	if c.isClosed {
		return errors.New("Connection closed when send msg")
	}
	// 将data进行封包 MsgDataLen MsgID Data
	dp := NewDataPack()

	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg")
	}

	// 将数据发送给客户端
	c.msgChan <- binaryMsg
	return nil
}

// 启动链接 让当前的链接准备开始工作
func (c *Connection) Start() {
	fmt.Println("Conn Start() ... ConnID = ", c.ConnID)
	// 启动从当前链接的读数据的业务
	go c.StartReader()
	// TODO 启动从当前链接写数据的业务
	go c.StartWrite()
}

// 停止链接 结束当前链接的工作
func (c *Connection) Stop() {
	fmt.Println("Conn Stop().. ConnID = ", c.ConnID)

	// 如果当前链接已经关闭
	if c.isClosed {
		return
	}
	c.isClosed = true

	// 关闭socket链接
	c.Conn.Close()

	// 告知Write关闭
	c.ExitChan <- true

	// 回收资源
	close(c.ExitChan)
	close(c.msgChan)
}

// 获取当前链接的绑定sockert conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.GetTCPConnection()
}

// 获取当前链接模块的链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的 TCP状态 IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.RemoteAddr()
}
