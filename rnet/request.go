package rnet

import "github.com/mrsongindezhou/ratel/riface"

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/3/18 8:26 上午
 * @description：
 * @modified By：
 * @version    ：$
 */
type Request struct {
	// 已经和客户端建立好的链接
	conn riface.IConnection

	// 客户端请求的数据
	msg riface.IMessage
}

// 得到当前链接
func (r *Request) GetConnection() riface.IConnection {
	return r.conn
}

// 得到请求的消息数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}
