package riface

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/3/18 8:26 上午
 * @description：
 * @modified By：
 * @version    ：$
 */

/*
	IRequest接口：
	实际是将客户端请求链接信息 和 请求的数据 包装到了一个Request中
*/
type IRequest interface {
	// 得到当前链接
	GetConnection() IConnection

	// 得到请求的消息数据
	GetData() []byte

	// 得到当前消息的ID
	GetMsgID() uint32
}
