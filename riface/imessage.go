package riface

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/3/19 5:04 下午
 * @description：将请求的消息封装到Message中，定义一个抽象模块
 * @modified By：
 * @version    ：$
 */
type IMessage interface {
	// 获取消息的ID
	GetMsgID() uint32
	// 获取消息的长度
	GetMsgLen() uint32
	// 获取消息的内容
	GetData() []byte

	// 设置消息的ID
	SetMsgID(uint32)
	// 设置消息的内容
	SetData([]byte)
	// 设置消息的长度
	SetDataLen(uint32)
}
